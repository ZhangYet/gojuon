package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

var roma = map[string][]string{
	"a":  []string{"a", "i", "u", "e", "o"},
	"k":  []string{"ka", "ki", "ku", "ke", "ko"},
	"s":  []string{"sa", "shi", "su", "se", "so"},
	"t":  []string{"ta", "chi", "tsu", "te", "to"},
	"n":  []string{"na", "ni", "nu", "ne", "no"},
	"h":  []string{"ha", "hi", "fu", "he", "ho"},
	"m":  []string{"ma", "mi", "mu", "me", "mo"},
	"y":  []string{"ya", "yu", "yo"},
	"r":  []string{"ra", "ri", "ru", "re", "ro"},
	"w":  []string{"wa", "wi", "wo"},
	"nn": []string{"n"},
}

var hiragana = map[string][]string{
	"a":  []string{"あ", "い", "う", "え", "お"},
	"k":  []string{"か", "き", "く", "け", "こ"},
	"s":  []string{"さ", "し", "す", "せ", "そ"},
	"t":  []string{"た", "ち", "つ", "て", "と"},
	"n":  []string{"な", "に", "ぬ", "ね", "の"},
	"h":  []string{"は", "ひ", "ふ", "へ", "ほ"},
	"m":  []string{"ま", "み", "む", "め", "も"},
	"y":  []string{"や", "ゆ", "よ"},
	"r":  []string{"ら", "り", "る", "れ", "ろ"},
	"w":  []string{"わ", "ゐ", "を"},
	"nn": []string{"ん"},
}

var katakana = map[string][]string{
	"a":  []string{"ア", "イ", "ウ", "エ", "オ"},
	"k":  []string{"カ", "キ", "ク", "ケ", "コ"},
	"s":  []string{"サ", "シ", "ス", "セ", "ソ"},
	"t":  []string{"タ", "チ", "ツ", "テ", "ト"},
	"n":  []string{"ナ", "ニ", "ヌ", "ネ", "ノ"},
	"h":  []string{"ハ", "ヒ", "フ", "ヘ", "ホ"},
	"m":  []string{"マ", "ミ", "ム", "メ", "モ"},
	"y":  []string{"ヤ", "ユ", "ヨ"},
	"r":  []string{"ラ", "リ", "ル", "レ", "ロ"},
	"w":  []string{"ワ", "ヰ", "ヲ"},
	"nn": []string{"ン"},
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
	return fmt.Sprintf("%s: %s %s", g.roma, g.hira, g.hira)
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

func main() {
	app := cli.NewApp()
	app.Name = "gojuon"
	app.Usage = "help japanese amateur learn gojuon."
	app.Commands = []cli.Command{
		{
			Name:    "reference",
			Aliases: []string{"r"},
			Usage:   "print gojuon list",
			Action: func(c *cli.Context) {
				reference(c.Args())
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
