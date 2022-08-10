package routes

import (
	accounttrpt "clean-architecture-go-fiber/src/module/account/transport"

	driver "clean-architecture-go-fiber/src/platform/driver/postgresql"
	"fmt"

	"github.com/gin-gonic/gin"
)

func accountRoutes(superRoute *gin.RouterGroup) {

	var db = driver.ConnectToPostgreSQL()

	fmt.Println("Opened connection to Postgresql")

	accountRouter := superRoute.Group("/accounts")
	{
		accountRouter.POST("/login", accounttrpt.HandleSignInWithEmail(db.SQL))
		accountRouter.POST("/", accounttrpt.HandleCreateAccount(db.SQL))

		accountRouter.GET("/:id", accounttrpt.HandlerFindAnAccount(db.SQL))
		accountRouter.GET("/", accounttrpt.HandleRetrieveAccounts(db.SQL))

		accountRouter.PATCH("/:id", accounttrpt.HandleUpdateAccount(db.SQL))
		accountRouter.DELETE("/:id", accounttrpt.HandleDeleteAccount(db.SQL))
	}
}
