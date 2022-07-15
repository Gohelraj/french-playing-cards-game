package models

import (
	"errors"
	"fmt"
)

type CardValue int8
type CardSuit int8

/*
 - This will generate cardvalue_string.go file
 - "-linecomment" flag used to modify String() function, It will use names provided in the comment on the same line as the object.
*/
//go:generate stringer -type=CardValue,CardSuit -linecomment

const (
	Spades CardSuit = iota
	Diamonds
	Clubs
	Hearts
)

const (
	Ace   CardValue = iota + 1 // Ace
	Two                        // 2
	Three                      // 3
	Four                       // 4
	Five                       // 5
	Six                        // 6
	Seven                      // 7
	Eight                      // 8
	Nine                       // 9
	Ten                        // 10
	Jack                       // Jack
	Queen                      // Queen
	King                       // King
)

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

func GenerateInvalidCardCodeError(cardCode string) error {
	return errors.New(fmt.Sprintf("Card code %s is invalid", cardCode))
}
