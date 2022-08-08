package routes

import (
	accounttrpt "clean-architecture-go-fiber/src/module/account/transport"

	"github.com/gin-gonic/gin"
)

func authRoutes(superRoute *gin.RouterGroup) {
	authRouter := superRoute.Group("/oauth2")
	{
		authRouter.GET("/google", accounttrpt.HandlerSignInWithGoogle())
		authRouter.GET("/google/callback", accounttrpt.HandleCallbackGoogle())
	}
}

//localhost:8080/api/v1/oauth2/google
