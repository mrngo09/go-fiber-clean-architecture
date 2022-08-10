package routes

import (
	"clean-architecture-go-fiber/src/app/middlewares"
	component "clean-architecture-go-fiber/src/components"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func AddRouter(r *gin.RouterGroup, appCtx component.AppContext) {
	accountRoutes(r, appCtx)
	authRoutes(r)
}

func InitRouter() {
	app := gin.New()

	appCtx := SetupAppContext()

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
		AddRouter(router, appCtx)
	}

	app.Run(":8080")
}
