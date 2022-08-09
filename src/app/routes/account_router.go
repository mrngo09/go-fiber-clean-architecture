package routes

import (
	// accountmodel "clean-architecture-go-fiber/src/module/account/model"
	accounttrpt "clean-architecture-go-fiber/src/module/account/transport"
	driver "clean-architecture-go-fiber/src/platform/driver/postgresql"
	"fmt"

	"github.com/gin-gonic/gin"
)

func accountRoutes(superRoute *gin.RouterGroup) {

	var db = driver.ConnectToPostgreSQL()

	// db.SQL.AutoMigrate(&accountmodel.Account{})

	fmt.Println("Opened connection to Postgresql")

	accountRouter := superRoute.Group("/accounts")
	{
		accountRouter.GET("/:id", accounttrpt.HandlerFindAnAccount(db.SQL))

		accountRouter.POST("/", accounttrpt.HandleCreateAccount(db.SQL))

		accountRouter.POST("/login", accounttrpt.HandleSignInWithEmail(db.SQL))

		accountRouter.GET("/", accounttrpt.HandleRetrieveAccounts(db.SQL))

	}
}
