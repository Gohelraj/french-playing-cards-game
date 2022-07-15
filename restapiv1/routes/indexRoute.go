package routes

import (
	"database/sql"
	"net/http"

	"github.com/Gohelraj/french-playing-cards-game/middleware"
	"github.com/Gohelraj/french-playing-cards-game/restapiv1/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.InjectDBConnection(db))

	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello There!!")
		return
	})

	r.POST("/deck", controllers.CreateDeck)
	r.GET("/deck/:deckId", controllers.OpenDeck)
	r.PATCH("/deck/:deckId/draw", controllers.Draw)
	return r
}
