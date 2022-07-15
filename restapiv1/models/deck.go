package models

import (
	"github.com/google/uuid"
	"math/rand"
	"time"
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
	r := rand.New(rand.NewSource(time.Now().Unix()))
	// We start at the end of the slice, inserting our random
	// values one at a time.
	for idx := len(d); idx > 0; idx-- {
		randomIndex := r.Intn(idx)
		// We swap the value at index idx-1 and the random index
		// to move our randomly chosen value to the end of the
		// slice, and to move the value that was at idx-1 into our
		// unshuffled portion of the slice.
		d[idx-1], d[randomIndex] = d[randomIndex], d[idx-1]
	}
}
