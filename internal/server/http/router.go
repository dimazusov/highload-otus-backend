package http

import (
	"social/internal/app"
	"social/internal/server/http/handlers/user"
	"social/internal/server/http/middleware"

	"github.com/gin-gonic/gin"
)

// @title Swagger API
// @version 1.0
// @description social api
func NewGinRouter(app *app.App) *gin.Engine {
	router := gin.Default()
	authMiddleware := middleware.Auth(app)

	frontendGroup := router.Group("/")
	frontendGroup.GET("/")

	apiGroup := router.Group("/api/v1")

	apiGroup.GET("/auth", func(c *gin.Context) { user.AuthHandler(c, app) })
	apiGroup.GET("/register", func(c *gin.Context) { user.RegisterHandler(c, app) })

	apiGroup.Use(authMiddleware).GET("/users", func(c *gin.Context) { user.GetUsersHandler(c, app) })
	apiGroup.Use(authMiddleware).GET("/user/:id", func(c *gin.Context) { user.GetUserHandler(c, app) })
	apiGroup.Use(authMiddleware).PUT("/user", func(c *gin.Context) { user.UpdateUserHandler(c, app) })
	apiGroup.Use(authMiddleware).POST("/user", func(c *gin.Context) { user.CreateUserHandler(c, app) })
	apiGroup.Use(authMiddleware).DELETE("/user/:id", func(c *gin.Context) { user.DeleteUserHandler(c, app) })

	return router
}
