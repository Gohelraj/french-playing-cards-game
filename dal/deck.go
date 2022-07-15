package dal

import (
	"database/sql"
	"encoding/json"
	"github.com/Gohelraj/french-playing-cards-game/restapiv1/models"
)

func CreateDeck(db *sql.DB, deck models.DeckModel) error {
	cards, err := json.Marshal(deck.Cards)
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO deck (id, is_shuffled, cards) VALUES($1, $2, $3);", deck.DeckID, deck.Shuffled, cards)
	if err != nil {
		return err
	}
	return nil
}

func FetchDeck(db *sql.DB, deckID string) (deck models.DeckModel, err error) {
	var cardsJSONString string
	row := db.QueryRow("SELECT id, is_shuffled, cards FROM deck WHERE id = $1;", deckID)
	err = row.Scan(&deck.DeckID, &deck.Shuffled, &cardsJSONString)
	if err != nil {
		return models.DeckModel{}, err
	}
	err = json.Unmarshal([]byte(cardsJSONString), &deck.Cards)
	return
}

func UpdateDeckCards(db *sql.DB, deck models.DeckModel) error {
	cards, err := json.Marshal(deck.Cards)
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE deck SET cards = $1 WHERE id = $2;", cards, deck.DeckID)
	if err != nil {
		return err
	}
	return nil
}
