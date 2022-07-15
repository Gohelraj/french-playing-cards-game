package bll

import (
	"github.com/Gohelraj/french-playing-cards-game/restapiv1/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultDeckGenerator(t *testing.T) {
	tests := []struct {
		name string
		want models.Deck
	}{
		{
			name: "success",
			want: []models.Card{
				{Value: "ACE", Suit: "SPADES", Code: "AS"},
				{Value: "2", Suit: "SPADES", Code: "2S"},
				{Value: "3", Suit: "SPADES", Code: "3S"},
				{Value: "4", Suit: "SPADES", Code: "4S"},
				{Value: "5", Suit: "SPADES", Code: "5S"},
				{Value: "6", Suit: "SPADES", Code: "6S"},
				{Value: "7", Suit: "SPADES", Code: "7S"},
				{Value: "8", Suit: "SPADES", Code: "8S"},
				{Value: "9", Suit: "SPADES", Code: "9S"},
				{Value: "10", Suit: "SPADES", Code: "1S"},
				{Value: "JACK", Suit: "SPADES", Code: "JS"},
				{Value: "QUEEN", Suit: "SPADES", Code: "QS"},
				{Value: "KING", Suit: "SPADES", Code: "KS"},
				{Value: "ACE", Suit: "DIAMONDS", Code: "AD"},
				{Value: "2", Suit: "DIAMONDS", Code: "2D"},
				{Value: "3", Suit: "DIAMONDS", Code: "3D"},
				{Value: "4", Suit: "DIAMONDS", Code: "4D"},
				{Value: "5", Suit: "DIAMONDS", Code: "5D"},
				{Value: "6", Suit: "DIAMONDS", Code: "6D"},
				{Value: "7", Suit: "DIAMONDS", Code: "7D"},
				{Value: "8", Suit: "DIAMONDS", Code: "8D"},
				{Value: "9", Suit: "DIAMONDS", Code: "9D"},
				{Value: "10", Suit: "DIAMONDS", Code: "1D"},
				{Value: "JACK", Suit: "DIAMONDS", Code: "JD"},
				{Value: "QUEEN", Suit: "DIAMONDS", Code: "QD"},
				{Value: "KING", Suit: "DIAMONDS", Code: "KD"},
				{Value: "ACE", Suit: "CLUBS", Code: "AC"},
				{Value: "2", Suit: "CLUBS", Code: "2C"},
				{Value: "3", Suit: "CLUBS", Code: "3C"},
				{Value: "4", Suit: "CLUBS", Code: "4C"},
				{Value: "5", Suit: "CLUBS", Code: "5C"},
				{Value: "6", Suit: "CLUBS", Code: "6C"},
				{Value: "7", Suit: "CLUBS", Code: "7C"},
				{Value: "8", Suit: "CLUBS", Code: "8C"},
				{Value: "9", Suit: "CLUBS", Code: "9C"},
				{Value: "10", Suit: "CLUBS", Code: "1C"},
				{Value: "JACK", Suit: "CLUBS", Code: "JC"},
				{Value: "QUEEN", Suit: "CLUBS", Code: "QC"},
				{Value: "KING", Suit: "CLUBS", Code: "KC"},
				{Value: "ACE", Suit: "HEARTS", Code: "AH"},
				{Value: "2", Suit: "HEARTS", Code: "2H"},
				{Value: "3", Suit: "HEARTS", Code: "3H"},
				{Value: "4", Suit: "HEARTS", Code: "4H"},
				{Value: "5", Suit: "HEARTS", Code: "5H"},
				{Value: "6", Suit: "HEARTS", Code: "6H"},
				{Value: "7", Suit: "HEARTS", Code: "7H"},
				{Value: "8", Suit: "HEARTS", Code: "8H"},
				{Value: "9", Suit: "HEARTS", Code: "9H"},
				{Value: "10", Suit: "HEARTS", Code: "1H"},
				{Value: "JACK", Suit: "HEARTS", Code: "JH"},
				{Value: "QUEEN", Suit: "HEARTS", Code: "QH"},
				{Value: "KING", Suit: "HEARTS", Code: "KH"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DefaultDeckGenerator()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDecodeValueAndSuit(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name        string
		args        args
		valueString string
		suitString  string
		wantErr     bool
	}{
		{
			name:        "success - Code:AS",
			args:        args{"AS"},
			valueString: "ACE",
			suitString:  "SPADES",
			wantErr:     false,
		},
		{
			name:        "success - Code:2S",
			args:        args{"2S"},
			valueString: "2",
			suitString:  "SPADES",
			wantErr:     false,
		},
		{
			name:        "success - Code:3S",
			args:        args{"3S"},
			valueString: "3",
			suitString:  "SPADES",
			wantErr:     false,
		},
		{
			name:        "success - Code:4S",
			args:        args{"4S"},
			valueString: "4",
			suitString:  "SPADES",
			wantErr:     false,
		},
		{
			name:        "success - Code:5S",
			args:        args{"5S"},
			valueString: "5",
			suitString:  "SPADES",
			wantErr:     false,
		},
		{
			name:        "success - Code:6S",
			args:        args{"6S"},
			valueString: "6",
			suitString:  "SPADES",
			wantErr:     false,
		},
		{
			name:        "success - Code:7S",
			args:        args{"7S"},
			valueString: "7",
			suitString:  "SPADES",
			wantErr:     false,
		},
		{
			name:        "success - Code:8S",
			args:        args{"8S"},
			valueString: "8",
			suitString:  "SPADES",
			wantErr:     false,
		},
		{
			name:        "success - Code:9S",
			args:        args{"9S"},
			valueString: "9",
			suitString:  "SPADES",
			wantErr:     false,
		},
		{
			name:        "success - Code:1S",
			args:        args{"1S"},
			valueString: "10",
			suitString:  "SPADES",
			wantErr:     false,
		},
		{
			name:        "success - Code:JS",
			args:        args{"JS"},
			valueString: "JACK",
			suitString:  "SPADES",
			wantErr:     false,
		},
		{
			name:        "success - Code:QS",
			args:        args{"QS"},
			valueString: "QUEEN",
			suitString:  "SPADES",
			wantErr:     false,
		},
		{
			name:        "success - Code:KS",
			args:        args{"KS"},
			valueString: "KING",
			suitString:  "SPADES",
			wantErr:     false,
		},
		{
			name:        "success - Code:AD",
			args:        args{"AD"},
			valueString: "ACE",
			suitString:  "DIAMONDS",
			wantErr:     false,
		},
		{
			name:        "success - Code:2D",
			args:        args{"2D"},
			valueString: "2",
			suitString:  "DIAMONDS",
			wantErr:     false,
		},
		{
			name:        "success - Code:3D",
			args:        args{"3D"},
			valueString: "3",
			suitString:  "DIAMONDS",
			wantErr:     false,
		},
		{
			name:        "success - Code:4D",
			args:        args{"4D"},
			valueString: "4",
			suitString:  "DIAMONDS",
			wantErr:     false,
		},
		{
			name:        "success - Code:5D",
			args:        args{"5D"},
			valueString: "5",
			suitString:  "DIAMONDS",
			wantErr:     false,
		},
		{
			name:        "success - Code:6D",
			args:        args{"6D"},
			valueString: "6",
			suitString:  "DIAMONDS",
			wantErr:     false,
		},
		{
			name:        "success - Code:7D",
			args:        args{"7D"},
			valueString: "7",
			suitString:  "DIAMONDS",
			wantErr:     false,
		},
		{
			name:        "success - Code:8D",
			args:        args{"8D"},
			valueString: "8",
			suitString:  "DIAMONDS",
			wantErr:     false,
		},
		{
			name:        "success - Code:9D",
			args:        args{"9D"},
			valueString: "9",
			suitString:  "DIAMONDS",
			wantErr:     false,
		},
		{
			name:        "success - Code:1D",
			args:        args{"1D"},
			valueString: "10",
			suitString:  "DIAMONDS",
			wantErr:     false,
		},
		{
			name:        "success - Code:JD",
			args:        args{"JD"},
			valueString: "JACK",
			suitString:  "DIAMONDS",
			wantErr:     false,
		},
		{
			name:        "success - Code:QD",
			args:        args{"QD"},
			valueString: "QUEEN",
			suitString:  "DIAMONDS",
			wantErr:     false,
		},
		{
			name:        "success - Code:KD",
			args:        args{"KD"},
			valueString: "KING",
			suitString:  "DIAMONDS",
			wantErr:     false,
		},
		{
			name:        "success - Code:AC",
			args:        args{"AC"},
			valueString: "ACE",
			suitString:  "CLUBS",
			wantErr:     false,
		},
		{
			name:        "success - Code:2C",
			args:        args{"2C"},
			valueString: "2",
			suitString:  "CLUBS",
			wantErr:     false,
		},
		{
			name:        "success - Code:3C",
			args:        args{"3C"},
			valueString: "3",
			suitString:  "CLUBS",
			wantErr:     false,
		},
		{
			name:        "success - Code:4C",
			args:        args{"4C"},
			valueString: "4",
			suitString:  "CLUBS",
			wantErr:     false,
		},
		{
			name:        "success - Code:5C",
			args:        args{"5C"},
			valueString: "5",
			suitString:  "CLUBS",
			wantErr:     false,
		},
		{
			name:        "success - Code:6C",
			args:        args{"6C"},
			valueString: "6",
			suitString:  "CLUBS",
			wantErr:     false,
		},
		{
			name:        "success - Code:7C",
			args:        args{"7C"},
			valueString: "7",
			suitString:  "CLUBS",
			wantErr:     false,
		},
		{
			name:        "success - Code:8C",
			args:        args{"8C"},
			valueString: "8",
			suitString:  "CLUBS",
			wantErr:     false,
		},
		{
			name:        "success - Code:9C",
			args:        args{"9C"},
			valueString: "9",
			suitString:  "CLUBS",
			wantErr:     false,
		},
		{
			name:        "success - Code:1C",
			args:        args{"1C"},
			valueString: "10",
			suitString:  "CLUBS",
			wantErr:     false,
		},
		{
			name:        "success - Code:JC",
			args:        args{"JC"},
			valueString: "JACK",
			suitString:  "CLUBS",
			wantErr:     false,
		},
		{
			name:        "success - Code:QC",
			args:        args{"QC"},
			valueString: "QUEEN",
			suitString:  "CLUBS",
			wantErr:     false,
		},
		{
			name:        "success - Code:KC",
			args:        args{"KC"},
			valueString: "KING",
			suitString:  "CLUBS",
			wantErr:     false,
		},
		{
			name:        "success - Code:AH",
			args:        args{"AH"},
			valueString: "ACE",
			suitString:  "HEARTS",
			wantErr:     false,
		},
		{
			name:        "success - Code:2H",
			args:        args{"2H"},
			valueString: "2",
			suitString:  "HEARTS",
			wantErr:     false,
		},
		{
			name:        "success - Code:3H",
			args:        args{"3H"},
			valueString: "3",
			suitString:  "HEARTS",
			wantErr:     false,
		},
		{
			name:        "success - Code:4H",
			args:        args{"4H"},
			valueString: "4",
			suitString:  "HEARTS",
			wantErr:     false,
		},
		{
			name:        "success - Code:5H",
			args:        args{"5H"},
			valueString: "5",
			suitString:  "HEARTS",
			wantErr:     false,
		},
		{
			name:        "success - Code:6H",
			args:        args{"6H"},
			valueString: "6",
			suitString:  "HEARTS",
			wantErr:     false,
		},
		{
			name:        "success - Code:7H",
			args:        args{"7H"},
			valueString: "7",
			suitString:  "HEARTS",
			wantErr:     false,
		},
		{
			name:        "success - Code:8H",
			args:        args{"8H"},
			valueString: "8",
			suitString:  "HEARTS",
			wantErr:     false,
		},
		{
			name:        "success - Code:9H",
			args:        args{"9H"},
			valueString: "9",
			suitString:  "HEARTS",
			wantErr:     false,
		},
		{
			name:        "success - Code:1H",
			args:        args{"1H"},
			valueString: "10",
			suitString:  "HEARTS",
			wantErr:     false,
		},
		{
			name:        "success - Code:JH",
			args:        args{"JH"},
			valueString: "JACK",
			suitString:  "HEARTS",
			wantErr:     false,
		},
		{
			name:        "success - Code:QH",
			args:        args{"QH"},
			valueString: "QUEEN",
			suitString:  "HEARTS",
			wantErr:     false,
		},
		{
			name:        "success - Code:KH",
			args:        args{"KH"},
			valueString: "KING",
			suitString:  "HEARTS",
			wantErr:     false,
		},
		{
			name:        "failure - invalid value code in argument",
			args:        args{code: "VS"},
			valueString: "",
			suitString:  "",
			wantErr:     true,
		},
		{
			name:        "failure - invalid suit code in argument",
			args:        args{code: "AX"},
			valueString: "",
			suitString:  "",
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotSuit, err := DecodeValueAndSuit(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeValueAndSuit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.valueString, gotValue)
			assert.Equal(t, tt.suitString, gotSuit)
		})
	}
}

func TestCustomDeckGenerator(t *testing.T) {
	type args struct {
		codesList  []string
		totalCodes int
	}
	tests := []struct {
		name    string
		args    args
		want    models.Deck
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				codesList:  []string{"AS", "KD", "AC", "2C", "KH", "9H", "1D"},
				totalCodes: 7,
			},
			want: []models.Card{
				{Value: "ACE", Suit: "SPADES", Code: "AS"},
				{Value: "KING", Suit: "DIAMONDS", Code: "KD"},
				{Value: "ACE", Suit: "CLUBS", Code: "AC"},
				{Value: "2", Suit: "CLUBS", Code: "2C"},
				{Value: "KING", Suit: "HEARTS", Code: "KH"},
				{Value: "9", Suit: "HEARTS", Code: "9H"},
				{Value: "10", Suit: "DIAMONDS", Code: "1D"},
			},
			wantErr: false,
		},
		{
			name: "failure - invalid value in codes",
			args: args{
				codesList:  []string{"AS", "KZ", "AC", "2X", "KH", "9H", "1D"},
				totalCodes: 7,
			},
			want:    []models.Card{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CustomDeckGenerator(tt.args.codesList, tt.args.totalCodes)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomDeckGenerator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDrawCards(t *testing.T) {
	type args struct {
		deck models.Deck
		n    int
	}
	tests := []struct {
		name           string
		args           args
		drawnCards     models.Deck
		remainingCards models.Deck
		wantErr        bool
	}{
		{
			name: "success",
			args: args{
				deck: []models.Card{
					{Value: "ACE", Suit: "SPADES", Code: "AS"},
					{Value: "KING", Suit: "DIAMONDS", Code: "KD"},
					{Value: "ACE", Suit: "CLUBS", Code: "AC"},
					{Value: "2", Suit: "CLUBS", Code: "2C"},
					{Value: "KING", Suit: "HEARTS", Code: "KH"},
					{Value: "9", Suit: "HEARTS", Code: "9H"},
					{Value: "10", Suit: "DIAMONDS", Code: "1D"},
				},
				n: 2,
			},
			drawnCards: []models.Card{
				{Value: "ACE", Suit: "SPADES", Code: "AS"},
				{Value: "KING", Suit: "DIAMONDS", Code: "KD"},
			},
			remainingCards: []models.Card{
				{Value: "ACE", Suit: "CLUBS", Code: "AC"},
				{Value: "2", Suit: "CLUBS", Code: "2C"},
				{Value: "KING", Suit: "HEARTS", Code: "KH"},
				{Value: "9", Suit: "HEARTS", Code: "9H"},
				{Value: "10", Suit: "DIAMONDS", Code: "1D"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			drawnCards, remainingCards := DrawCards(tt.args.deck, tt.args.n)
			assert.Equal(t, tt.drawnCards, drawnCards)
			assert.Equal(t, tt.remainingCards, remainingCards)
		})
	}
}
