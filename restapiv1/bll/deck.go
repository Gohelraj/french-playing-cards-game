// This package can be used to handle business logic layer.
package bll

import (
	"github.com/Gohelraj/french-playing-cards-game/restapiv1/constants"
	"github.com/Gohelraj/french-playing-cards-game/restapiv1/models"
	"strconv"
	"strings"
)

func DefaultDeckGenerator() models.Deck {
	newDeck := make(models.Deck, constants.TotalValidCardsCount)
	idx := 0
	for i := models.Spades; i <= models.Hearts; i++ {
		for j := models.Ace; j <= models.King; j++ {
			suitString := strings.ToUpper(i.String())
			valueString := strings.ToUpper(j.String())
			newDeck[idx] = models.Card{
				Suit:  suitString,
				Value: valueString,
				Code:  string([]byte{valueString[0], suitString[0]}),
			}
			idx++
		}
	}
	return newDeck
}

func DecodeValueAndSuit(code string) (string, string, error) {
	var (
		value    models.CardValue
		suit     models.CardSuit
		parseErr error
	)
	valueCode := string(code[0])
	suitCode := string(code[1])

	valueInt, err := strconv.Atoi(valueCode)
	// if error is not null that means the provided value code is alphabet.
	if err != nil {
		switch valueCode {
		case "A":
			value = models.Ace
		case "J":
			value = models.Jack
		case "Q":
			value = models.Queen
		case "K":
			value = models.King
		default:
			return "", "", models.GenerateInvalidCardCodeError(code)
		}
	} else if valueInt == 1 {
		// 1 value code is used for "10" number card
		value = models.Ten
	} else {
		value = models.CardValue(valueInt)
	}

	switch suitCode {
	case "S":
		suit = models.Spades
	case "D":
		suit = models.Diamonds
	case "C":
		suit = models.Clubs
	case "H":
		suit = models.Hearts
	default:
		return "", "", models.GenerateInvalidCardCodeError(code)
	}
	return strings.ToUpper(value.String()), strings.ToUpper(suit.String()), parseErr
}

func CustomDeckGenerator(codesList []string, totalCodes int) (models.Deck, error) {
	newDeck := make(models.Deck, totalCodes)
	for i, cardCode := range codesList {
		value, suit, err := DecodeValueAndSuit(cardCode)
		if err != nil {
			return models.Deck{}, err
		}
		newDeck[i] = models.Card{
			Suit:  suit,
			Value: value,
			Code:  string([]byte{value[0], suit[0]}),
		}
	}
	return newDeck, nil
}

func DrawCards(d models.Deck, n int) (drawnCards models.Deck, remainingCards models.Deck) {
	drawnCards, remainingCards = d[:n], d[n:]
	return
}
