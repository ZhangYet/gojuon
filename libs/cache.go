package libs

type WordBookCache interface {
	query(index WordBookIndex) WordBookEntry
	create(entry WordBookEntry)
}

type TrivalCache struct {
	data map[WordBookIndex]WordBookEntry
}

func NewTrivalCache() *TrivalCache {
	return &TrivalCache{data: make(map[WordBookIndex]WordBookEntry)}
}

func (c *TrivalCache) query(index WordBookIndex) WordBookEntry {
	return c.data[index]
}

func (c *TrivalCache) create(entry WordBookEntry) {
	index := WordBookIndex{
		Japanese: entry.Japanese,
		Furigana: entry.Furigana,
	}
	c.data[index] = WordBookEntry{}
}
