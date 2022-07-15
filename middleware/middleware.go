package middleware

import (
	"database/sql"

	"github.com/Gohelraj/french-playing-cards-game/restapiv1/constants"
	"github.com/gin-gonic/gin"
)

func InjectDBConnection(dbConn *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Set(constants.DBConnCtxKey, dbConn)
		c.Next()
	}
}
