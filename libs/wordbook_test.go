package libs

import (
	"os"
	"testing"
)

func TestWordBook(t *testing.T) {
	testPath := "./testWordBook"
	defer func() {
		os.Remove(testPath)
	}()
	t.Run("search and record", func(t *testing.T) {
		w := NewFileWorkBook(testPath)
		record, err := w.Search("世界")
		if err != nil {
			t.Fatal(err)
		}
		t.Log(record.String())
		if err := w.Record(record); err != nil {
			t.Fatal()
		}
	})
}
