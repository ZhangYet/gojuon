package libs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type WordBookStorage interface {
	save(entry WordBookEntry) error
	load() ([]WordBookEntry, error)
}

type FileStorage struct {
	file string
}

func NewFileStorage(file string) *FileStorage {
	return &FileStorage{file: file}
}

func (s *FileStorage) save(entry WordBookEntry) error {
	line := fmt.Sprintf("%d|%s|%s|%s\n", entry.CreateTime.Unix(), entry.Japanese, entry.Furigana, entry.English)
	f, err := os.OpenFile(s.file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("%v : %s", err, s.file)
	}
	defer func() { _ = f.Close() }()
	_, err = f.Write([]byte(line))
	return err
}

func (s *FileStorage) load() ([]WordBookEntry, error) {
	f, err := os.Open(s.file)
	if err != nil {
		return nil, err
	}
	ret := make([]WordBookEntry, 0)
	reader := bufio.NewReader(f)
	for {
		content, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		line := string(content)
		sline := strings.Split(line, "|")
		if len(sline) < 4 {
			return nil, fmt.Errorf("file format error")
		}
		createTs, err := strconv.Atoi(sline[0])
		if err != nil {
			return nil, err
		}
		entry := WordBookEntry{
			CreateTime: time.Unix(int64(createTs), 0),
			Japanese:   sline[1],
			Furigana:   sline[2],
			English:    sline[3],
		}
		ret = append(ret, entry)
	}
	return ret, nil
}
