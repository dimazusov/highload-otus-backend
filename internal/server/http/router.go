package http

import (
	"github.com/gin-gonic/gin"

	"social/internal/server/http/middleware"
	"social/internal/app"
	"social/internal/server/http/handlers/user"
)

// @title Swagger API
// @version 1.0
// @description social api
func NewGinRouter(app *app.App) *gin.Engine {
	router := gin.Default()


	v1Group := router.Group("/api/v1")

	v1Group.GET("/auth", func(c *gin.Context) { user.GetRegisterHandler(c, app) })
	v1Group.GET("/register", func(c *gin.Context) { user.GetRegisterHandler(c, app) })

	auth := middleware.Auth(app)
	v1Group.Use(auth).GET("/users", func(c *gin.Context) { user.GetUsersHandler(c, app) })
	v1Group.Use(auth).GET("/user/:id", func(c *gin.Context) { user.GetUserHandler(c, app) })
	v1Group.Use(auth).PUT("/user", func(c *gin.Context) { user.UpdateUserHandler(c, app) })
	v1Group.Use(auth).POST("/user", func(c *gin.Context) { user.CreateUserHandler(c, app) })
	v1Group.Use(auth).DELETE("/user/:id", func(c *gin.Context) { user.DeleteUserHandler(c, app) })

	return router
}
