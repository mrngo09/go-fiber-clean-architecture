package routes

import (
	"clean-architecture-go-fiber/src/app/middlewares"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func AddRouter(r *gin.RouterGroup) {
	accountRoutes(r)
	authRoutes(r)
}

func InitRouter() {
	app := gin.New()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	app.Use(cors.New(corsConfig))

	app.Use(middlewares.LoggingMiddleware())
	app.Use(gin.Recovery())

	// store := cookie.NewStore([]byte("secret"))
	// app.Use(sessions.Sessions("mysession", store))

	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello world")
	})

	router := app.Group("/api/v1")
	{
		AddRouter(router)
	}

	app.Run(":8080")
}
