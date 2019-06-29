package libs

import (
	"testing"
)

const ankiUrl = "http://localhost:8765"

func TestAnki_DeckNamesAndIds(t *testing.T) {
	type fields struct {
		ankiUrl string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Get Deck Names And Ids",
			fields: fields{
				ankiUrl: ankiUrl,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Anki{
				ankiUrl: tt.fields.ankiUrl,
			}
			got, err := this.DeckNamesAndIds()
			if (err != nil) != tt.wantErr {
				t.Errorf("Anki.DeckNamesAndIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("result: %v\n", got)
		})
	}
}

func TestAnki_CreateAndDeleteDeck(t *testing.T) {
	type fields struct {
		ankiUrl string
	}
	type args struct {
		deckName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Create a new deck",
			fields: fields{
				ankiUrl: ankiUrl,
			},
			args: args{
				deckName: "testDeck",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Anki{
				ankiUrl: tt.fields.ankiUrl,
			}
			got, err := this.CreateDeck(tt.args.deckName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Anki.CreateDeck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("the id of the new deck: %d\n", got)
			if err := this.DeleteDeck(tt.args.deckName, true); err != nil {
				t.Errorf("Anki.DeleteDeck() error = %v", err)
			}
		})
	}
}

func TestAnki_ModelNamesAndIds(t *testing.T) {
	type fields struct {
		ankiUrl string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Get Model Names And Ids",
			fields: fields{
				ankiUrl: ankiUrl,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Anki{
				ankiUrl: tt.fields.ankiUrl,
			}
			got, err := this.ModelNamesAndIds()
			if (err != nil) != tt.wantErr {
				t.Errorf("Anki.ModelNamesAndIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("result: %v\n", got)
		})
	}
}

func TestAnki_CreateModel(t *testing.T) {
	type fields struct {
		ankiUrl string
	}
	type args struct {
		modelName     string
		inOrderFields []string
		css           string
		cardTemplates []map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Create models",
			fields: fields{
				ankiUrl: ankiUrl,
			},
			args: args{
				modelName:     "gojuon_model_test",
				inOrderFields: []string{"FrontField", "BackField1", "BackField2"},
				css:           defaultCSS,
				cardTemplates: []map[string]string{
					{
						"Front": defaultFront,
						"Back":  defaultBack,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Anki{
				ankiUrl: tt.fields.ankiUrl,
			}
			if err := this.CreateModel(tt.args.modelName, tt.args.inOrderFields, tt.args.css, tt.args.cardTemplates); (err != nil) != tt.wantErr {
				t.Errorf("Anki.CreateModel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAnki_AddNote(t *testing.T) {
	type fields struct {
		ankiUrl string
	}
	type args struct {
		deckName  string
		modelName string
		fields    map[string]string
		tags      []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Add note",
			fields: fields{
				ankiUrl: ankiUrl,
			},
			args: args{
				deckName:  "testDeck",
				modelName: "gojuon_model_test",
				fields: map[string]string{
					"FrontField": "FrontFieldTest",
					"BackField1": "BackField1Test",
					"BackField2": "BackField2Test",
				},
				tags: []string{"gojuon"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Anki{
				ankiUrl: tt.fields.ankiUrl,
			}
			got, err := this.AddNote(tt.args.deckName, tt.args.modelName, tt.args.fields, tt.args.tags)
			if (err != nil) != tt.wantErr {
				t.Errorf("Anki.AddNote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("the id of new card is %d\n", got)
		})
	}
}
