package libs

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ZhangYet/gojuon/data"
)

type Word struct {
	Japanese string
	Furigana string
	English  string
}

func (w Word) String() string {
	return fmt.Sprintf("japanese: %s, furigana: %s, english: %s", w.Japanese, w.Furigana, w.English)
}

type WordRecord struct {
	Word
	Id         int64
	CreateTime time.Time
}

type WordBook interface {
	Search(keyword string) (WordRecord, error) // search a word by japanese or furigana
	Record(WordRecord) error                   // write down a word
}

// FileWordBook implements WordBook by file.
type FileWorkBook struct {
	storePath string
}

func NewFileWorkBook(storePath string) WordBook {
	return &FileWorkBook{storePath: storePath}
}

func (b *FileWorkBook) Search(keyword string) (record WordRecord, err error) {
	searchResult := data.Dict.Search(keyword, 1)
	if len(searchResult) <= 0 {
		return record, fmt.Errorf("can not found this word in dictionary")
	}
	var word Word
	word.Japanese = searchResult[0].Japanese
	word.Furigana = searchResult[0].Furigana
	englishList := make([]string, len(searchResult[0].Glosses))
	for idx, gloss := range searchResult[0].Glosses {
		e := fmt.Sprintf("%d: %s", idx, gloss.English)
		englishList[idx] = e
	}
	word.English = strings.Join(englishList, ", ")
	ts := time.Now()
	record = WordRecord{
		Word:       word,
		Id:         ts.Unix(),
		CreateTime: time.Now(),
	}
	return record, nil
}

func (b *FileWorkBook) Record(record WordRecord) (err error) {
	f, err := os.OpenFile(b.storePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	if err != nil {

		return err

	}
	line := fmt.Sprintf("%d|%s|%s\n", record.Id, record.CreateTime.Format("2006-01-02"), record.String())
	_, err = f.Write([]byte(line))
	return err
}
