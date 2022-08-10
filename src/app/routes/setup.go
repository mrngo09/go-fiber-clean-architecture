package routes

import (
	component "clean-architecture-go-fiber/src/components"
	"clean-architecture-go-fiber/src/components/tokenprovider/jwt"
)

func SetupAppContext() component.AppContext {
	tokenProvider := jwt.NewJwtProvider("SCRESETLOLS")

	appCtx := component.NewAppContext(tokenProvider)

	return appCtx
}
