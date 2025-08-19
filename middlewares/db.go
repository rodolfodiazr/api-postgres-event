package middlewares

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Database(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
