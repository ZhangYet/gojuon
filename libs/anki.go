package libs

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/levigross/grequests"
)

type Anki struct {
	ankiUrl string
}

func NewAnki(url string) *Anki {
	return &Anki{
		ankiUrl: url,
	}
}

const (
	defaultAnkiConnectVersion = 6
	AnkiUrl                   = "http://localhost:8765"
)

var (
	DefaultFields    = []string{"FrontField", "BackField1", "BackField2"}
	DefaultTemplates = []map[string]string{
		{
			"Front": DefaultFront,
			"Back":  DefaultBack,
		},
	}
	DefaultTags = []string{"gojuon"}
)

const (
	DefaultFront = `
<div class="front">
<span class="large japanese">{{FrontField}}</span>
<br/
</div>
	`
	DefaultBack = `
<div class="back">
<span class="large">{{BackField1}}</span>
<hr/>
{{BackField2}}
<br/>
</span>
</div>
	`
	DefaultCSS = `
div.front, div.back {
	text-align:center;
	font-family: Courier;
	font-size: 30px;
}

span.small {font-size: 15px;}
span.normal {font-size: 30px;}
span.large {font-size: 60px;}
span.italic {font-style:italic;}

.left {float:left;}


span.courier {font-family: Courier;}

.win .japanese {font-family: "MS Mincho";}
.mac .japanese {font-family: "Hiragino Mincho Pro";}
.linux .japanese {font-family: "Kochi Mincho";}
.mobile .japanese {font-family: "Hiragino Mincho ProN";}
	`
)

var defaultHeader = map[string]string{
	"Content-Type": "application/json",
}

func (this *Anki) request(ro *grequests.RequestOptions) (*grequests.Response, error) {
	return grequests.Post(this.ankiUrl, ro)
}

func (this *Anki) wrapRequest(ro *grequests.RequestOptions) (interface{}, error) {
	resp, err := this.request(ro)
	if err != nil {
		return nil, err
	}
	content, err := parseBytesToMap(resp.Bytes())
	if err != nil {
		return nil, err
	}
	result, respErr, err := getErrorAndContent(content)
	if respErr != nil {
		return nil, respErr
	}
	if err != nil {
		return nil, respErr
	}
	return result, nil
}

func (this *Anki) DeckNamesAndIds() (map[string]Meta, error) {
	ro := &grequests.RequestOptions{
		JSON:          map[string]interface{}{"action": "modelNamesAndIds", "version": defaultAnkiConnectVersion},
		Headers:       defaultHeader,
		BeforeRequest: logReqFunc,
	}
	result, err := this.wrapRequest(ro)
	if err != nil {
		return nil, err
	}
	ret := mapToDeckMap(result.(map[string]interface{}))
	return ret, nil
}

func (this *Anki) CreateDeck(deckName string) (int64, error) {
	params := map[string]interface{}{
		"deck": deckName,
	}
	ro := &grequests.RequestOptions{
		JSON: map[string]interface{}{
			"action":  "createDeck",
			"version": defaultAnkiConnectVersion,
			"params":  params,
		},
		Headers:       defaultHeader,
		BeforeRequest: logReqFunc,
	}
	result, err := this.wrapRequest(ro)
	if err != nil {
		return 0, err
	}
	return int64(result.(float64)), nil
}

func (this *Anki) DeleteDeck(deckName string, cardsToo bool) error {
	params := map[string]interface{}{
		"decks":    []string{deckName},
		"cardsToo": cardsToo,
	}
	ro := &grequests.RequestOptions{
		JSON: map[string]interface{}{
			"action":  "deleteDecks",
			"version": defaultAnkiConnectVersion,
			"params":  params,
		},
		Headers:       defaultHeader,
		BeforeRequest: logReqFunc,
	}
	_, err := this.wrapRequest(ro)
	return err
}

func (this *Anki) ModelNamesAndIds() (map[string]Meta, error) {
	ro := &grequests.RequestOptions{
		JSON: map[string]interface{}{
			"action":  "modelNamesAndIds",
			"version": defaultAnkiConnectVersion,
		},
		Headers:       defaultHeader,
		BeforeRequest: logReqFunc,
	}
	result, err := this.wrapRequest(ro)
	if err != nil {
		return nil, err
	}
	ret := mapToDeckMap(result.(map[string]interface{}))
	return ret, nil
}

func (this *Anki) CreateModel(modelName string, inOrderFields []string, css string,
	cardTemplates []map[string]string) error {
	params := map[string]interface{}{
		"modelName":     modelName,
		"inOrderFields": inOrderFields,
		"css":           css,
		"cardTemplates": cardTemplates,
	}
	ro := wrapReqOption("createModel", params)
	_, err := this.wrapRequest(ro)
	return err
}

func (this *Anki) AddNote(deckName, modelName string, fields map[string]string, tags []string) (int64, error) {
	params := map[string]interface{}{
		"note": map[string]interface{}{
			"deckName":  deckName,
			"modelName": modelName,
			"fields":    fields,
			"options": map[string]interface{}{
				"allowDuplicate": false,
			},
			"tags": tags,
		},
	}
	ro := wrapReqOption("addNote", params)
	result, err := this.wrapRequest(ro)
	if err != nil {
		return 0, err
	}
	return int64(result.(float64)), nil
}

type Model struct {
	Name string
	Id   int64
}

type Meta struct {
	Name string
	Id   int64
}

func wrapReqOption(action string, params map[string]interface{}) *grequests.RequestOptions {
	jsonData := map[string]interface{}{
		"action":  action,
		"version": defaultAnkiConnectVersion,
	}
	if params != nil {
		jsonData["params"] = params
	}
	return &grequests.RequestOptions{
		JSON:          jsonData,
		Headers:       defaultHeader,
		BeforeRequest: logReqFunc,
	}
}

func logReqFunc(req *http.Request) error {
	requestDump, _ := httputil.DumpRequest(req, true)
	fmt.Println(string(requestDump))
	return nil
}

func mapToDeckMap(content map[string]interface{}) map[string]Meta {
	ret := make(map[string]Meta)
	for key, value := range content {
		id := int64(value.(float64))
		ret[key] = Meta{Name: key, Id: id}
	}
	return ret
}

func parseBytesToMap(content []byte) (map[string]interface{}, error) {
	ret := make(map[string]interface{})
	err := json.Unmarshal(content, &ret)
	return ret, err
}

func getErrorAndContent(content map[string]interface{}) (result interface{}, respErr error, err error) {
	result, ok := content["result"]
	if !ok {
		return result, respErr, errors.New("no result")
	}
	errInterface, ok := content["error"]
	if !ok {
		return result, respErr, nil
	}
	if errInterface != nil {
		msg := errInterface.(string)
		respErr = errors.New(msg)
	}
	return result, respErr, err
}
