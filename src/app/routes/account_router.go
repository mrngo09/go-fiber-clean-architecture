package routes

import (
	"clean-architecture-go-fiber/src/app/middlewares"
	component "clean-architecture-go-fiber/src/components"
	accounttrpt "clean-architecture-go-fiber/src/module/account/transport"

	driver "clean-architecture-go-fiber/src/platform/driver/postgresql"
	"fmt"

	"github.com/gin-gonic/gin"
)

func accountRoutes(superRoute *gin.RouterGroup, appCtx component.AppContext) {

	var db = driver.ConnectToPostgreSQL()

	fmt.Println("Opened connection to Postgresql")

	accountRouter := superRoute.Group("/accounts")
	{
		accountRouter.POST("/login", accounttrpt.HandleLoginWithEmail(db.SQL, appCtx))
		accountRouter.POST("/", accounttrpt.HandleCreateAccount(db.SQL))

		accountRouter.GET("/:id", accounttrpt.HandlerFindAnAccount(db.SQL))
		accountRouter.GET("/", middlewares.RequireAuth(appCtx), accounttrpt.HandleRetrieveAccounts(db.SQL))

		accountRouter.PATCH("/:id", accounttrpt.HandleUpdateAccount(db.SQL))
		accountRouter.DELETE("/:id", accounttrpt.HandleDeleteAccount(db.SQL))
	}
}
