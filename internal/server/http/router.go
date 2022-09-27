package http

import (
	"net/http"
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

	frontendGroup := router.Group("/").Use(middleware.Cors(app))
	frontendGroup.GET("/", func(c *gin.Context) { c.File("web/index.html") })
	frontendGroup.StaticFS("/static", http.Dir("web/static"))

	authGroup := router.Group("/api/v1").Use(middleware.Cors(app))
	authGroup.POST("/auth", func(c *gin.Context) { user.AuthHandler(c, app) })
	authGroup.POST("/register", func(c *gin.Context) { user.RegisterHandler(c, app) })

	apiGroup := router.Group("/api/v1").
		Use(middleware.Cors(app)).
		Use(middleware.Auth(app))
	apiGroup.GET("/users", func(c *gin.Context) { user.GetUsersHandler(c, app) })
	apiGroup.GET("/user/:id", func(c *gin.Context) { user.GetUserHandler(c, app) })
	apiGroup.PUT("/user", func(c *gin.Context) { user.UpdateUserHandler(c, app) })
	apiGroup.POST("/user", func(c *gin.Context) { user.CreateUserHandler(c, app) })
	apiGroup.DELETE("/user/:id", func(c *gin.Context) { user.DeleteUserHandler(c, app) })

	return router
}
