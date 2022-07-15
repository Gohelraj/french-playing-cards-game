package repository

import (
	"database/sql"

	"github.com/Gohelraj/french-playing-cards-game/dal"
	"github.com/Gohelraj/french-playing-cards-game/restapiv1/models"
)

type DeckRepoInterface interface {
	CreateDeck(deck models.DeckModel) error
	FetchDeck(deckID string) (deck models.DeckModel, err error)
	UpdateDeckCards(cards models.DeckModel) error
}

type DeckRepo struct {
	DBConn *sql.DB
}

func NewDeckRepo(db *sql.DB) DeckRepoInterface {
	deckRepo := &DeckRepo{
		DBConn: db,
	}
	return deckRepo
}

func (deckRepo *DeckRepo) CreateDeck(deck models.DeckModel) error {
	return dal.CreateDeck(deckRepo.DBConn, deck)
}

func (deckRepo *DeckRepo) FetchDeck(deckID string) (deck models.DeckModel, err error) {
	return dal.FetchDeck(deckRepo.DBConn, deckID)
}

func (deckRepo *DeckRepo) UpdateDeckCards(deck models.DeckModel) error {
	return dal.UpdateDeckCards(deckRepo.DBConn, deck)
}
