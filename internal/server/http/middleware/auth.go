package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"social/internal/app"
	"social/internal/domain/auth_token"
)

func Auth(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, ok := c.Get("X-Auth-Token")
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{})
			return
		}
		_, err := app.Domain.AuthToken.Service.Parse(context.Background(), token.(string))
		if err != nil {
			if errors.Is(err, auth_token.ErrWrongToken) {
				c.JSON(http.StatusForbidden, gin.H{})
				return
			}
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		c.Next()
	}
}

func Cors(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("test")
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}
	}
}
