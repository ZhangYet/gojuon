package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli"
)

var Version string

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

type gojuon struct {
	roma string
	hira string
	kata string
}

var romaGojuonDict map[string][]gojuon

func initGojuon() {
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

func init() {
	initGojuon()
}

func printGojuon(g gojuon) string {
	return fmt.Sprintf("%s: %s %s", g.roma, g.hira, g.kata)
}

func reference(lines []string) {
	for _, k := range lines {
		gojuonList, _ := romaGojuonDict[k]
		output := make([]string, len(gojuonList))
		for idx, g := range gojuonList {
			output[idx] = printGojuon(g)
		}
		fmt.Println(strings.Join(output, "\t"))
	}
}

func genTest(typ string, lines []string) {
	var dict map[string][]string
	switch typ {
	case "hira":
		dict = hiragana
	case "kata":
		dict = katakana
	default:
		dict = roma
	}
	t := []string{}
	for _, k := range lines {
		data, ok := dict[k]
		if ok {
			t = append(t, data...)
		}
	}
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(t), func(i, j int) {
		t[i], t[j] = t[j], t[i]
	})
	fmt.Println(strings.Join(t, ", "))
}

func main() {
	app := cli.NewApp()
	app.Name = "gojuon"
	app.Usage = "help japanese amateur learn gojuon."
	app.Version = Version
	app.Commands = []cli.Command{
		{
			Name:    "reference",
			Aliases: []string{"r"},
			Usage:   "print gojuon list",
			Action: func(c *cli.Context) {
				reference(c.Args())
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
				genTest(typ, c.Args())
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
