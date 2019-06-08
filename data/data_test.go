package data

import (
	"compress/gzip"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/gojp/nihongo/lib/dictionary"
)

func TestLoadDict(t *testing.T) {
	t.Run("test load dict", func(t *testing.T) {
		if &Dict == nil {
			t.Fatal("dict didn't init")
		}
	})
	t.Run("test dict load in dev", func(t *testing.T) {
		f, err := os.Open("./edict2.json.gz")
		if err != nil {
			t.Fatal(err)
		}
		gReader, err := gzip.NewReader(f)
		if err != nil {
			t.Fatal(err)
		}
		newDict, err := dictionary.Load(gReader)
		if err != nil {
			t.Fatal(err)
		}
		ret := newDict.Search("愛", 10)
		if len(ret) <= 0 {
			t.Fatal("can't get entry from newDict")
		}

	})
	t.Run("test search", func(t *testing.T) {
		ret := Dict.Search("愛", 10)
		if len(ret) <= 0 {
			t.Fatal("can't get entry from dict")
		}
		t.Logf("word: %s, furigana: %s, english: %v", ret[0].Japanese, ret[0].Furigana, ret[0].Glosses[0].English)
		ret2 := Dict.Search("あい", 10)
		if len(ret2) <= 0 {
			t.Fatal("can't get entry from dict by furigana")

		}
		englishList := make([]string, len(ret2[0].Glosses))
		for idx, gloss := range ret2[0].Glosses {
			englishList[idx] = fmt.Sprintf("%d:%s", idx, gloss.English)
		}
		english := strings.Join(englishList, ";")
		t.Logf("word: %s, furigana: %s, english: %s", ret2[0].Japanese, ret2[0].Furigana, english)
	})
}
