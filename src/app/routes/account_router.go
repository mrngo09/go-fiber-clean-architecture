package routes

import (
	accounttrpt "clean-architecture-go-fiber/src/module/account/transport"
	database "clean-architecture-go-fiber/src/platform/driver/mysql"

	"github.com/gin-gonic/gin"
)

var db, err = database.OpenConnectToMySQL()

func accountRoutes(superRoute *gin.RouterGroup) {
	accountRouter := superRoute.Group("/accounts")
	{
		accountRouter.GET("/:id", accounttrpt.HandlerFindAnAccount(db))

		accountRouter.POST("/", accounttrpt.HandleCreateAccount(db))

		accountRouter.POST("/login", accounttrpt.HandleSignInWithEmail(db))

		accountRouter.GET("/", accounttrpt.HandleRetrieveAccounts(db))

	}
}
