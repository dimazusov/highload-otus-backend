package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"social/internal/app"
)

func Auth(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, ok := c.Get("X-Auth-Token")
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{})
			return
		}
		if !app.Domain.AuthToken.Service.CheckAuthToken(token) {
			c.JSON(http.StatusForbidden, gin.H{})
			return
		}
		c.Next()
	}
}