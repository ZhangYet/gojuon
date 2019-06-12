package data

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"strings"

	"github.com/gobuffalo/packr/v2"
	"github.com/gojp/nihongo/lib/dictionary"
)

var Box *packr.Box
var Dict dictionary.Dictionary

func init() {
	Box = packr.New("edict", "./")
	dict, err := Box.Find("edict2.json.gz")
	if err != nil {
		panic(err)
	}
	reader := bytes.NewBuffer(dict)
	gReader, err := gzip.NewReader(reader)
	if err != nil {
		panic(err)
	}
	Dict, err = dictionary.Load(gReader)
	if err != nil {
		panic(err)
	}
}

func GetEnglish(entry dictionary.Entry) string {
	englishList := make([]string, len(entry.Glosses))
	for idx, gloss := range entry.Glosses {
		englishList[idx] = fmt.Sprintf("%d: %s", idx, gloss.English)
	}
	return strings.Join(englishList, "; ")
}
