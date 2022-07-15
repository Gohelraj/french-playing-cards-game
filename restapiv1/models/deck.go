package models

import (
	"github.com/google/uuid"
	"math/rand"
)

type Deck []Card

type DeckModel struct {
	DeckID   uuid.UUID
	Cards    Deck
	Shuffled bool
}

type CreateDeckResponseBody struct {
	DeckId    uuid.UUID `json:"deck_id"`
	Shuffled  bool      `json:"shuffled"`
	Remaining int       `json:"remaining"`
}

type GetDeckResponseBody struct {
	DeckId    uuid.UUID `json:"deck_id"`
	Shuffled  bool      `json:"shuffled"`
	Remaining int       `json:"remaining"`
	Cards     Deck      `json:"cards"`
}

type DrawCardsResponseBody struct {
	Cards Deck `json:"cards"`
}

func (d Deck) Shuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}
