package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Gohelraj/french-playing-cards-game/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/Gohelraj/french-playing-cards-game/repository"
	"github.com/Gohelraj/french-playing-cards-game/restapiv1/bll"
	"github.com/Gohelraj/french-playing-cards-game/restapiv1/constants"
	"github.com/Gohelraj/french-playing-cards-game/restapiv1/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateDeck(c *gin.Context) {
	shuffleQueryParamString := c.DefaultQuery("shuffle", "false")
	shuffle, err := strconv.ParseBool(shuffleQueryParamString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shuffle parameter"})
		return
	}

	cardsQuery := c.DefaultQuery("cards", "")
	var deckCards models.Deck
	if cardsQuery != "" {
		deckCards, err = generateDeckCardsFromCustomQuery(cardsQuery)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		deckCards = bll.DefaultDeckGenerator()
	}
	if shuffle {
		deckCards.Shuffle()
	}
	db := c.Value(constants.DBConnCtxKey).(*sql.DB)
	deckRepo := repository.NewDeckRepo(db)
	deckID := uuid.New()
	deck := models.DeckModel{
		DeckID:   deckID,
		Cards:    deckCards,
		Shuffled: shuffle,
	}
	err = deckRepo.CreateDeck(deck)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.CreateDeckResponseBody{
		DeckId:    deckID,
		Shuffled:  shuffle,
		Remaining: len(deckCards),
	})
}

func generateDeckCardsFromCustomQuery(cardsQuery string) (models.Deck, error) {
	codesList := strings.Split(cardsQuery, ",")
	codesCount := len(codesList)
	if codesCount > constants.TotalValidCardsCount {
		return nil, errors.New(fmt.Sprintf("cards can no be more than %d", constants.TotalValidCardsCount))
	}
	duplicateValues := utils.FindDuplicatesFromStringSlice(codesList)
	if len(duplicateValues) > 0 {
		return nil, errors.New(fmt.Sprintf("Duplicate values are not allowed: %s", duplicateValues))
	}
	deckCards, err := bll.CustomDeckGenerator(codesList, codesCount)
	if err != nil {
		return nil, err
	}
	return deckCards, nil
}

func OpenDeck(c *gin.Context) {
	db := c.Value(constants.DBConnCtxKey).(*sql.DB)
	deckId := c.Params.ByName("deckId")
	deckRepo := repository.NewDeckRepo(db)
	deck, err := deckRepo.FetchDeck(deckId)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "deck not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.GetDeckResponseBody{
		DeckId:    deck.DeckID,
		Shuffled:  deck.Shuffled,
		Remaining: len(deck.Cards),
		Cards:     deck.Cards,
	})
}

func Draw(c *gin.Context) {
	count, err := strconv.Atoi(c.DefaultQuery("count", ""))
	if err != nil || count < 1 || count > constants.TotalValidCardsCount {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("count must be specified and must be between 1 and %d", constants.TotalValidCardsCount)})
		return
	}

	db := c.Value(constants.DBConnCtxKey).(*sql.DB)
	deckRepo := repository.NewDeckRepo(db)
	deckId := c.Params.ByName("deckId")
	deck, err := deckRepo.FetchDeck(deckId)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "deck not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	availableCardsCount := len(deck.Cards)
	if availableCardsCount == 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "deck is empty"})
		return
	}

	if count > availableCardsCount {
		c.JSON(http.StatusBadRequest, gin.H{"error": "requested number of cards is greater than the cards remaining in the deck"})
		return
	}

	drawnCards, remainingCards := bll.DrawCards(deck.Cards, count)
	err = deckRepo.UpdateDeckCards(models.DeckModel{
		DeckID: deck.DeckID,
		Cards:  remainingCards,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.DrawCardsResponseBody{
		Cards: drawnCards,
	})
}
