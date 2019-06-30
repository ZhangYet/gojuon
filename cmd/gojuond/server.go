package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gojp/nihongo/lib/dictionary"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"github.com/ZhangYet/gojuon/cmd"
	"github.com/ZhangYet/gojuon/data"
	"github.com/ZhangYet/gojuon/libs"

	gojuon_dict "github.com/ZhangYet/gojuon/rpc"
)

var _ gojuon_dict.DictServiceServer = &GojuonService{}

var anki *libs.Anki

const (
	ankiDeck  = "gojuon_deck"
	ankiModel = "gojuon_model"
)

type GojuonService struct {
	manager libs.Manager
	dict    dictionary.Dictionary
}

func (g *GojuonService) Search(_ context.Context, req *gojuon_dict.SearchRequest) (rep *gojuon_dict.SearchResponse, err error) {
	keyword := req.Keyword
	if keyword == "" {
		return nil, status.Errorf(codes.InvalidArgument, "keyword is empty")
	}

	ret := g.dict.Search(keyword, 1)
	if len(ret) <= 0 {
		grpclog.Warningf("%s not found in dict", keyword)
		return rep, nil
	}
	index := libs.WordBookIndex{
		Japanese: ret[0].Japanese,
		Furigana: ret[0].Furigana,
	}
	workRecord := gojuon_dict.WordRecord{
		Japanese: ret[0].Japanese,
		Furigana: ret[0].Furigana,
		English:  data.GetEnglish(ret[0]),
	}

	r := g.manager.Query(index)
	if r.English != "" {
		workRecord.CreateTime = &timestamp.Timestamp{
			Seconds: r.CreateTime.Unix(),
			Nanos:   0,
		}
	} else {
		workRecord.CreateTime = &timestamp.Timestamp{
			Seconds: time.Now().Unix(),
			Nanos:   0,
		}
	}
	go func(japanese, furigana, english string) {
		fields0 := map[string]string{
			"FrontField": japanese,
			"BackField1": furigana,
			"BackField2": english,
		}
		fields1 := map[string]string{
			"FrontField": furigana,
			"BackField1": english,
			"BackField2": japanese,
		}
		fields2 := map[string]string{
			"FrontField": english,
			"BackField1": japanese,
			"BackField2": furigana,
		}
		fieldsList := []map[string]string{
			fields0,
			fields1,
			fields2,
		}
		var wg sync.WaitGroup
		for _, fields := range fieldsList {
			wg.Add(1)
			go func(f map[string]string) {
				id, err := anki.AddNote(ankiDeck, ankiModel, f, libs.DefaultTags)
				grpclog.Infof("add note to %s, id: %, err: %v\n", ankiDeck, id, err)
				wg.Done()
			}(fields)
		}
		wg.Wait()
	}(workRecord.Japanese, workRecord.Furigana, workRecord.English)

	rep = &gojuon_dict.SearchResponse{}
	rep.Record = &workRecord
	return rep, nil
}

func (g *GojuonService) Record(_ context.Context, req *gojuon_dict.RecordRequest) (rep *gojuon_dict.RecordResponse, err error) {
	if req.Record == nil {
		return nil, status.Errorf(codes.InvalidArgument, "no records")
	}
	rep = &gojuon_dict.RecordResponse{}
	index := libs.WordBookIndex{
		Japanese: req.Record.Japanese,
		Furigana: req.Record.Furigana,
	}
	r := g.manager.Query(index)
	if r.English != "" {
		rep.CreateTime = &timestamp.Timestamp{
			Seconds: r.CreateTime.Unix(),
			Nanos:   0,
		}

	}
	createTs := time.Now()
	entry := libs.WordBookEntry{
		Japanese:   req.Record.Japanese,
		Furigana:   req.Record.Furigana,
		English:    req.Record.English,
		CreateTime: createTs,
	}
	if err := g.manager.Save(entry); err != nil {
		grpclog.Errorf("error in saving %v", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	rep.CreateTime = &timestamp.Timestamp{Seconds: createTs.Unix(), Nanos: 0}
	return rep, nil
}

func NewGojuonService(file string) (*GojuonService, error) {
	m, err := libs.NewManager("file", file)
	if err != nil {
		return nil, err
	}
	return &GojuonService{
		dict:    data.Dict,
		manager: m,
	}, nil
}

func setupLog(f *os.File) {
	logger := grpclog.NewLoggerV2(f, f, f)
	grpclog.SetLoggerV2(logger)
}

func init() {
	cmd.SetupConfig()
	f, err := os.Create(cmd.LogFile)
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		panic(err)
	}
	setupLog(f)

	anki = libs.NewAnki(libs.AnkiUrl)
	decks, err := anki.DeckNamesAndIds()
	if err != nil {
		grpclog.Error(err)
		return
	}
	if _, ok := decks[ankiDeck]; !ok {
		id, err := anki.CreateDeck(ankiDeck)
		if err != nil {
			grpclog.Error(err)
			return
		}
		grpclog.Infof("the id of anki deck is %d\n", id)
	}
	models, err := anki.ModelNamesAndIds()
	if err != nil {
		grpclog.Error(err)
		return
	}
	if _, ok := models[ankiModel]; !ok {
		err := anki.CreateModel(ankiModel, libs.DefaultFields, libs.DefaultCSS, libs.DefaultTemplates)
		if err != nil {
			grpclog.Error(err)
			return
		}
		grpclog.Infof("created model sucessfully")
	}
}

func main() {
	lis, err := net.Listen("tcp", cmd.RpcAddr)
	grpclog.Infof("listening on %s", cmd.RpcAddr)
	if err != nil {
		panic(err)
	}
	var grpcServer *grpc.Server
	go func() {
		grpcServer = grpc.NewServer()
		rpcService, err := NewGojuonService(cmd.SavingData)
		if err != nil {
			panic(err)
		}
		gojuon_dict.RegisterDictServiceServer(grpcServer, rpcService)
		reflection.Register(grpcServer)
		err = grpcServer.Serve(lis)
		if err != nil {
			grpclog.Error(err)
		}
	}()
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, os.Kill)
	s := <-interrupt
	grpclog.Infof("process interrupt by signal: %v", s)
	grpcServer.GracefulStop()
}
