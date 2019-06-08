package data

import (
	"bytes"
	"compress/gzip"

	"github.com/gojp/nihongo/lib/dictionary"

	"github.com/gobuffalo/packr/v2"
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
