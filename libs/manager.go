package libs

import (
	"fmt"
	"time"
)

type WordBookIndex struct {
	Japanese string
	Furigana string
}

type WordBookEntry struct {
	Japanese   string
	Furigana   string
	English    string
	CreateTime time.Time
}

type Manager interface {
	Query(index WordBookIndex) WordBookEntry
	Save(entry WordBookEntry) error
	Init() error
}

type WordBookManager struct {
	storage WordBookStorage
	cache   WordBookCache
}

func (m *WordBookManager) Query(index WordBookIndex) WordBookEntry {
	return m.cache.query(index)
}

func (m *WordBookManager) Save(entry WordBookEntry) error {
	return m.storage.save(entry)
}

func (m *WordBookManager) Init() error {
	entryList, err := m.storage.load()
	if err != nil {
		return err
	}
	for _, entry := range entryList {
		m.cache.create(entry)
	}
	return nil
}

func NewManager(typ string, args ...interface{}) (m Manager, err error) {
	switch typ {
	case "file":
		fileName, ok := args[0].(string)
		if !ok {
			return m, fmt.Errorf("FileManager needs a string parameter")
		}
		storage := NewFileStorage(fileName)
		cache := NewTrivalCache()
		return &WordBookManager{
			storage: storage,
			cache:   cache,
		}, nil
	default:
		return m, fmt.Errorf("not supported manager type")
	}
}
