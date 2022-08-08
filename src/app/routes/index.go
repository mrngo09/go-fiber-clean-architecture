package routes

import (
	"clean-architecture-go-fiber/src/app/middlewares"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func AddRouter(r *gin.RouterGroup) {
	accountRoutes(r)
	authRoutes(r)
}

func InitRouter() {
	app := gin.New()
	app.Use(middlewares.LoggingMiddleware())

	store := cookie.NewStore([]byte("secret"))
	app.Use(sessions.Sessions("mysession", store))

	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello world")
	})

	router := app.Group("/api/v1")
	{
		AddRouter(router)
	}

	app.Run(":8080")
}
