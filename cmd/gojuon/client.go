package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
	"google.golang.org/grpc"

	"github.com/ZhangYet/gojuon/cmd"
	"github.com/ZhangYet/gojuon/common"

	gojuon_dict "github.com/ZhangYet/gojuon/rpc"
)

var (
	Version   string
	rpcConn   *grpc.ClientConn
	rpcClient gojuon_dict.DictServiceClient
)

func shell() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		ctx := context.Background()
		in := gojuon_dict.SearchRequest{
			Keyword: strings.Trim(input, "\n"),
		}
		rep, err := rpcClient.Search(ctx, &in)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
		output := fmt.Sprintf("Japanese:\t%s\nFurigana:\t%s\nEnglish:\t%s\n",
			rep.Record.Japanese, rep.Record.Furigana, rep.Record.English)
		fmt.Fprintf(os.Stdout, output)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "gojuon"
	app.Usage = "help japanese amateur learn gojuon."
	app.Version = Version
	app.Before = func(context *cli.Context) error {
		cmd.SetupConfig()
		rpcConn, err := grpc.Dial(cmd.RpcAddr, grpc.WithInsecure())
		if err != nil {
			return err
		}
		rpcClient = gojuon_dict.NewDictServiceClient(rpcConn)
		return nil

	}
	app.After = func(context *cli.Context) error {
		if rpcConn != nil {
			return rpcConn.Close()
		}
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:    "shell",
			Aliases: []string{"sh"},
			Action: func(c *cli.Context) {
				shell()
			},
		},
		{
			Name:    "recordWord",
			Aliases: []string{"rw"},
			Usage:   "search the meaning and record a japanese new word",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "word",
					Usage: "input a new word",
				},
			},
			Action: func(c *cli.Context) {
				keyword := c.String("word")
				ctx := context.Background()
				in := gojuon_dict.SearchRequest{
					Keyword: keyword,
				}
				rep, err := rpcClient.Search(ctx, &in)
				if err != nil {
					panic(err)
				}
				fmt.Printf("Japanese:\t%s\nFurigana:\t%s\nEnglish:\t%s\n",
					rep.Record.Japanese, rep.Record.Furigana, rep.Record.English)
				recordIn := gojuon_dict.RecordRequest{
					Record: rep.Record,
				}
				if _, err := rpcClient.Record(ctx, &recordIn); err != nil {
					panic(err)
				}
			},
		},
		{
			Name:    "reference",
			Aliases: []string{"r"},
			Usage:   "print gojuon list",
			Action: func(c *cli.Context) {
				common.Reference(c.Args())
			},
		},
		{
			Name:    "gen",
			Aliases: []string{"g"},
			Usage:   "print gojuon test",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "type",
					Value: "roma",
					Usage: "roma|hira(gana)|kata(gana)",
				},
			},
			Action: func(c *cli.Context) {
				typ := c.String("type")
				ret := common.QueryKanaRows(typ, c.Args()...)
				fmt.Println(strings.Join(ret, ", "))
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
