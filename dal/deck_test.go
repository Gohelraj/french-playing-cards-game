package dal

import (
	"database/sql"
	"github.com/Gohelraj/french-playing-cards-game/restapiv1/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateDeck(t *testing.T) {
	type args struct {
		db   *sql.DB
		deck models.DeckModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - with shuffled false",
			args: args{
				db: testDB,
				deck: models.DeckModel{
					DeckID: uuid.New(),
					Cards: []models.Card{
						{Value: "ACE", Suit: "SPADES", Code: "AS"},
						{Value: "KING", Suit: "DIAMONDS", Code: "KD"},
						{Value: "ACE", Suit: "CLUBS", Code: "AC"},
						{Value: "2", Suit: "CLUBS", Code: "2C"},
						{Value: "KING", Suit: "HEARTS", Code: "KH"},
						{Value: "9", Suit: "HEARTS", Code: "9H"},
						{Value: "10", Suit: "DIAMONDS", Code: "1D"},
					},
					Shuffled: false,
				},
			},
			wantErr: false,
		},
		{
			name: "success - with shuffled true",
			args: args{
				db: testDB,
				deck: models.DeckModel{
					DeckID: uuid.New(),
					Cards: []models.Card{
						{Value: "KING", Suit: "DIAMONDS", Code: "KD"},
						{Value: "ACE", Suit: "CLUBS", Code: "AC"},
						{Value: "ACE", Suit: "SPADES", Code: "AS"},
						{Value: "9", Suit: "HEARTS", Code: "9H"},
						{Value: "2", Suit: "CLUBS", Code: "2C"},
						{Value: "10", Suit: "DIAMONDS", Code: "1D"},
						{Value: "KING", Suit: "HEARTS", Code: "KH"},
					},
					Shuffled: true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateDeck(tt.args.db, tt.args.deck); (err != nil) != tt.wantErr {
				t.Errorf("CreateDeck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFetchDeck(t *testing.T) {
	type args struct {
		db     *sql.DB
		deckID string
	}
	tests := []struct {
		name     string
		args     args
		wantDeck models.DeckModel
		wantErr  bool
	}{
		{
			name: "success",
			args: args{
				db:     testDB,
				deckID: "6818b769-c508-43de-910a-4c0a63ad8dad",
			},
			wantDeck: models.DeckModel{
				DeckID: uuid.MustParse("6818b769-c508-43de-910a-4c0a63ad8dad"),
				Cards: []models.Card{
					{Value: "KING", Suit: "CLUBS", Code: "KC"},
					{Value: "ACE", Suit: "CLUBS", Code: "AC"},
					{Value: "5", Suit: "HEARTS", Code: "5H"},
					{Value: "KING", Suit: "HEARTS", Code: "KH"},
					{Value: "ACE", Suit: "DIAMONDS", Code: "AD"},
					{Value: "10", Suit: "DIAMONDS", Code: "1D"},
					{Value: "8", Suit: "HEARTS", Code: "8H"},
					{Value: "QUEEN", Suit: "CLUBS", Code: "QC"},
					{Value: "6", Suit: "DIAMONDS", Code: "6D"},
					{Value: "9", Suit: "CLUBS", Code: "9C"},
					{Value: "QUEEN", Suit: "DIAMONDS", Code: "QD"},
					{Value: "3", Suit: "SPADES", Code: "3S"},
					{Value: "3", Suit: "CLUBS", Code: "3C"},
					{Value: "3", Suit: "DIAMONDS", Code: "3D"},
					{Value: "2", Suit: "HEARTS", Code: "2H"},
					{Value: "10", Suit: "SPADES", Code: "1S"},
					{Value: "2", Suit: "SPADES", Code: "2S"},
					{Value: "4", Suit: "CLUBS", Code: "4C"},
					{Value: "3", Suit: "HEARTS", Code: "3H"},
					{Value: "7", Suit: "HEARTS", Code: "7H"},
					{Value: "ACE", Suit: "SPADES", Code: "AS"},
					{Value: "5", Suit: "CLUBS", Code: "5C"},
					{Value: "4", Suit: "DIAMONDS", Code: "4D"},
					{Value: "QUEEN", Suit: "HEARTS", Code: "QH"},
					{Value: "6", Suit: "SPADES", Code: "6S"},
					{Value: "7", Suit: "DIAMONDS", Code: "7D"},
					{Value: "2", Suit: "CLUBS", Code: "2C"},
					{Value: "ACE", Suit: "HEARTS", Code: "AH"},
					{Value: "KING", Suit: "DIAMONDS", Code: "KD"},
					{Value: "JACK", Suit: "CLUBS", Code: "JC"},
					{Value: "QUEEN", Suit: "SPADES", Code: "QS"},
					{Value: "7", Suit: "SPADES", Code: "7S"},
					{Value: "6", Suit: "HEARTS", Code: "6H"},
					{Value: "JACK", Suit: "DIAMONDS", Code: "JD"},
					{Value: "10", Suit: "CLUBS", Code: "1C"},
					{Value: "JACK", Suit: "SPADES", Code: "JS"},
					{Value: "5", Suit: "DIAMONDS", Code: "5D"},
					{Value: "4", Suit: "HEARTS", Code: "4H"},
					{Value: "2", Suit: "DIAMONDS", Code: "2D"},
					{Value: "9", Suit: "SPADES", Code: "9S"},
					{Value: "JACK", Suit: "HEARTS", Code: "JH"},
					{Value: "10", Suit: "HEARTS", Code: "1H"},
					{Value: "KING", Suit: "SPADES", Code: "KS"},
					{Value: "5", Suit: "SPADES", Code: "5S"},
					{Value: "8", Suit: "SPADES", Code: "8S"},
					{Value: "4", Suit: "SPADES", Code: "4S"},
					{Value: "7", Suit: "CLUBS", Code: "7C"},
					{Value: "8", Suit: "DIAMONDS", Code: "8D"},
					{Value: "9", Suit: "DIAMONDS", Code: "9D"},
					{Value: "8", Suit: "CLUBS", Code: "8C"},
					{Value: "9", Suit: "HEARTS", Code: "9H"},
					{Value: "6", Suit: "CLUBS", Code: "6C"},
				},
				Shuffled: false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDeck, err := FetchDeck(tt.args.db, tt.args.deckID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchDeck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantDeck, gotDeck)
		})
	}
}

func TestUpdateDeckCards(t *testing.T) {
	type args struct {
		db   *sql.DB
		deck models.DeckModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				db: testDB,
				deck: models.DeckModel{
					DeckID: uuid.MustParse("2e904357-ae8e-401b-9071-0a0a3f4b5770"),
					Cards: []models.Card{
						{Value: "2", Suit: "SPADES", Code: "2S"},
						{Value: "4", Suit: "CLUBS", Code: "4C"},
						{Value: "3", Suit: "HEARTS", Code: "3H"},
						{Value: "7", Suit: "HEARTS", Code: "7H"},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateDeckCards(tt.args.db, tt.args.deck); (err != nil) != tt.wantErr {
				t.Errorf("UpdateDeckCards() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
