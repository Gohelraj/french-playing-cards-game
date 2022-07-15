package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Gohelraj/french-playing-cards-game/middleware"
	"github.com/Gohelraj/french-playing-cards-game/restapiv1/models"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewCreateDeck(t *testing.T) {
	tests := []struct {
		name              string
		shuffleQueryParam string
		cardsQueryParam   string
		wantStatusCode    int
		wantResponse      *models.CreateDeckResponseBody
	}{
		{
			name:              "success - default deck without shuffle",
			shuffleQueryParam: "false",
			wantStatusCode:    http.StatusOK,
			wantResponse: &models.CreateDeckResponseBody{
				Shuffled:  false,
				Remaining: 52,
			},
		},
		{
			name:              "success - default deck with shuffle",
			shuffleQueryParam: "true",
			wantStatusCode:    http.StatusOK,
			wantResponse: &models.CreateDeckResponseBody{
				Shuffled:  true,
				Remaining: 52,
			},
		},
		{
			name:              "success - deck with custom cards without shuffle",
			cardsQueryParam:   "AS,KD,AC,2C,KH,9H,1D",
			shuffleQueryParam: "false",
			wantStatusCode:    http.StatusOK,
			wantResponse: &models.CreateDeckResponseBody{
				Shuffled:  false,
				Remaining: 7,
			},
		},
		{
			name:              "success - deck with custom cards with shuffle",
			cardsQueryParam:   "AS,KD,AC,2C,KH",
			shuffleQueryParam: "true",
			wantStatusCode:    http.StatusOK,
			wantResponse: &models.CreateDeckResponseBody{
				Shuffled:  true,
				Remaining: 5,
			},
		},
		{
			name:              "failure - invalid value in shuffle query param",
			shuffleQueryParam: "gdfdgdfg",
			wantStatusCode:    http.StatusBadRequest,
			wantResponse:      nil,
		},
		{
			name:              "failure - invalid value in cards query param",
			cardsQueryParam:   "AS,KD,AZ,2C,KH2",
			shuffleQueryParam: "false",
			wantStatusCode:    http.StatusBadRequest,
			wantResponse:      nil,
		},
		{
			name:              "failure - duplicate values in cards query param",
			cardsQueryParam:   "AS,KD,AC,2C,KH,KH,KH",
			shuffleQueryParam: "false",
			wantStatusCode:    http.StatusBadRequest,
			wantResponse:      nil,
		},
		{
			name:              "failure - validation error max 52 cards allowed",
			cardsQueryParam:   "AS,2S,3S,4S,5S,6S,7S,8S,9S,1S,JS,QS,KS,AD,2D,3D,4D,5D,6D,7D,8D,9D,1D,JD,QD,KD,AC,2C,3C,4C,5C,6C,7C,8C,9C,1C,JC,QC,KC,AH,2H,3H,4H,5H,6H,7H,8H,9H,1H,JH,QH,KH,1D",
			shuffleQueryParam: "false",
			wantStatusCode:    http.StatusBadRequest,
			wantResponse:      nil,
		},
	}
	router := gin.Default()
	router.Use(middleware.InjectDBConnection(testDB))
	router.POST("/deck", CreateDeck)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/deck", bytes.NewBuffer(nil))

			q := req.URL.Query()
			q.Add("shuffle", tt.shuffleQueryParam)
			q.Add("cards", tt.cardsQueryParam)
			req.URL.RawQuery = q.Encode()

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.wantStatusCode, w.Code)

			if tt.wantResponse != nil {
				var deckResponse models.CreateDeckResponseBody
				json.NewDecoder(w.Body).Decode(&deckResponse)
				assert.Equal(t, tt.wantResponse.Remaining, deckResponse.Remaining)
				assert.Equal(t, tt.wantResponse.Shuffled, deckResponse.Shuffled)
			}

		})
	}
}

func TestOpenDeck(t *testing.T) {
	tests := []struct {
		name           string
		deckID         string
		wantStatusCode int
		wantResponse   *models.GetDeckResponseBody
	}{
		{
			name:           "success - default deck without shuffle",
			deckID:         "8c3f7db7-45ae-42a8-8dc4-25b0e3745776",
			wantStatusCode: http.StatusOK,
			wantResponse: &models.GetDeckResponseBody{
				Shuffled:  false,
				Remaining: 52,
				Cards: []models.Card{
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
		},
		{
			name:           "failure - deck not found",
			deckID:         "451c86a2-82cd-9037-8b19-a667188006er",
			wantStatusCode: http.StatusNotFound,
			wantResponse:   nil,
		},
	}
	router := gin.Default()
	router.Use(middleware.InjectDBConnection(testDB))
	router.GET("/deck/:deckId", OpenDeck)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", fmt.Sprintf("/deck/%s", tt.deckID), nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.wantStatusCode, w.Code)

			if tt.wantResponse != nil {
				var deckResponse models.GetDeckResponseBody
				json.NewDecoder(w.Body).Decode(&deckResponse)
				assert.Equal(t, tt.wantResponse.Remaining, deckResponse.Remaining)
				assert.Equal(t, tt.wantResponse.Shuffled, deckResponse.Shuffled)
				assert.Equal(t, tt.wantResponse.Cards, deckResponse.Cards)
			}

		})
	}
}

func TestDraw(t *testing.T) {
	tests := []struct {
		name            string
		deckID          string
		countQueryParam string
		wantStatusCode  int
		wantResponse    *models.DrawCardsResponseBody
	}{
		{
			name:            "success",
			deckID:          "351c86a1-52cd-4237-8b12-a667188006ac",
			countQueryParam: "4",
			wantStatusCode:  http.StatusOK,
			wantResponse: &models.DrawCardsResponseBody{
				Cards: []models.Card{
					{Value: "8", Suit: "DIAMONDS", Code: "8D"},
					{Value: "9", Suit: "DIAMONDS", Code: "9D"},
					{Value: "10", Suit: "DIAMONDS", Code: "1D"},
					{Value: "JACK", Suit: "DIAMONDS", Code: "JD"},
				},
			},
		},
		{
			name:            "failure - deck not found",
			deckID:          "sdsc86a1-52cd-4237-8b12-a667188006ac",
			countQueryParam: "3",
			wantStatusCode:  http.StatusNotFound,
			wantResponse:    nil,
		},
		{
			name:            "failure - number of cards must be specified and be between 1 and 52",
			deckID:          "351c86a1-52cd-4237-8b12-a667188006ac",
			countQueryParam: "55",
			wantStatusCode:  http.StatusBadRequest,
			wantResponse:    nil,
		},
		{
			name:            "failure - requested number of cards is greater than the cards remaining in the deck",
			deckID:          "351c86a1-52cd-4237-8b12-a667188006ac",
			countQueryParam: "20",
			wantStatusCode:  http.StatusBadRequest,
			wantResponse:    nil,
		},
		{
			name:            "failure - deck is empty",
			deckID:          "891c86a2-52cd-4237-8b12-a667188006df",
			countQueryParam: "20",
			wantStatusCode:  http.StatusConflict,
			wantResponse:    nil,
		},
	}
	router := gin.Default()
	router.Use(middleware.InjectDBConnection(testDB))
	router.PATCH("/deck/:deckId/draw", Draw)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("PATCH", fmt.Sprintf("/deck/%s/draw", tt.deckID), nil)

			q := req.URL.Query()
			q.Add("count", tt.countQueryParam)
			req.URL.RawQuery = q.Encode()

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.wantStatusCode, w.Code)

			if tt.wantResponse != nil {
				var deckResponse models.GetDeckResponseBody
				json.NewDecoder(w.Body).Decode(&deckResponse)
				assert.Equal(t, tt.wantResponse.Cards, deckResponse.Cards)
			}
		})
	}
}
