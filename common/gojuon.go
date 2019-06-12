package common

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type gojuon struct {
	roma string
	hira string
	kata string
}

var (
	romaCol  = []string{"a", "ka", "sa", "ta", "na", "ha", "ma", "ya", "ra", "wa", "n"}
	hiraCol  = []string{"あ", "か", "さ", "た", "な", "は", "ま", "や", "ら", "わ", "ん"}
	kataCol  = []string{"ア", "カ", "サ", "タ", "ナ", "ハ", "マ", "ヤ", "ラ", "ワ", "ン"}
	colIndex = []string{"a", "k", "s", "t", "n", "h", "m", "y", "r", "w", "nn"}

	romaColIndex = make(map[string]string)
	hiraColIndex = make(map[string]string)
	kataColIndex = make(map[string]string)
	indexMap     = make(map[string]string)

	romaGojuonDict map[string][]gojuon
)

func init() {
	for idx, index := range colIndex {
		romaColIndex[romaCol[idx]] = index
		hiraColIndex[hiraCol[idx]] = index
		kataColIndex[kataCol[idx]] = index
		indexMap[index] = index
	}
	romaGojuonDict = make(map[string][]gojuon)
	for k, roma := range roma {
		romaGojuonDict[k] = make([]gojuon, 0)
		for idx, r := range roma {
			g := gojuon{
				roma: r,
				hira: hiragana[k][idx],
				kata: katakana[k][idx],
			}
			romaGojuonDict[k] = append(romaGojuonDict[k], g)
		}
	}
}

var roma = map[string][]string{
	"a":  {"a", "i", "u", "e", "o"},
	"k":  {"ka", "ki", "ku", "ke", "ko"},
	"s":  {"sa", "shi", "su", "se", "so"},
	"t":  {"ta", "chi", "tsu", "te", "to"},
	"n":  {"na", "ni", "nu", "ne", "no"},
	"h":  {"ha", "hi", "fu", "he", "ho"},
	"m":  {"ma", "mi", "mu", "me", "mo"},
	"y":  {"ya", "yu", "yo"},
	"r":  {"ra", "ri", "ru", "re", "ro"},
	"w":  {"wa", "wi", "wo"},
	"nn": {"n"},
}

var hiragana = map[string][]string{
	"a":  {"あ", "い", "う", "え", "お"},
	"k":  {"か", "き", "く", "け", "こ"},
	"s":  {"さ", "し", "す", "せ", "そ"},
	"t":  {"た", "ち", "つ", "て", "と"},
	"n":  {"な", "に", "ぬ", "ね", "の"},
	"h":  {"は", "ひ", "ふ", "へ", "ほ"},
	"m":  {"ま", "み", "む", "め", "も"},
	"y":  {"や", "ゆ", "よ"},
	"r":  {"ら", "り", "る", "れ", "ろ"},
	"w":  {"わ", "ゐ", "を"},
	"nn": {"ん"},
}

var katakana = map[string][]string{
	"a":  {"ア", "イ", "ウ", "エ", "オ"},
	"k":  {"カ", "キ", "ク", "ケ", "コ"},
	"s":  {"サ", "シ", "ス", "セ", "ソ"},
	"t":  {"タ", "チ", "ツ", "テ", "ト"},
	"n":  {"ナ", "ニ", "ヌ", "ネ", "ノ"},
	"h":  {"ハ", "ヒ", "フ", "ヘ", "ホ"},
	"m":  {"マ", "ミ", "ム", "メ", "モ"},
	"y":  {"ヤ", "ユ", "ヨ"},
	"r":  {"ラ", "リ", "ル", "レ", "ロ"},
	"w":  {"ワ", "ヰ", "ヲ"},
	"nn": {"ン"},
}

func QueryKanaRows(typ string, cols ...string) (randomRows []string) {
	var dict map[string][]string
	switch typ {
	case "hira":
		dict = hiragana
	case "kata":
		dict = katakana
	default:
		dict = roma
	}
	temp := make([]string, 0)
	index := getIndex(cols)
	for _, i := range index {
		temp = append(temp, dict[i]...)
	}
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(temp), func(i, j int) {
		temp[i], temp[j] = temp[j], temp[i]
	})
	return temp
}

func getIndex(cols []string) (index []string) {
	for _, col := range cols {
		var dictIndex string
		if index, ok := indexMap[col]; ok {
			dictIndex = index
		}
		if index, ok := romaColIndex[col]; ok {
			dictIndex = index
		}
		if index, ok := hiraColIndex[col]; ok {
			dictIndex = index
		}
		if index, ok := kataColIndex[col]; ok {
			dictIndex = index
		}
		index = append(index, dictIndex)
	}
	return index
}

func printGojuon(g gojuon) string {
	return fmt.Sprintf("%s: %s %s", g.roma, g.hira, g.kata)
}

func Reference(cols []string) {
	index := getIndex(cols)
	for _, k := range index {
		gojuonList, _ := romaGojuonDict[k]
		output := make([]string, len(gojuonList))
		for idx, g := range gojuonList {
			output[idx] = printGojuon(g)
		}
		fmt.Println(strings.Join(output, "\t"))
	}
}
