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
