package component

import (
	"clean-architecture-go-fiber/src/components/tokenprovider"
)

type AppContext interface {
	GetTokenProvider() tokenprovider.TokenProvider
}

type appCtx struct {
	tokenProvider tokenprovider.TokenProvider
}

func NewAppContext(tokenProvider tokenprovider.TokenProvider) *appCtx {
	return &appCtx{
		tokenProvider: tokenProvider,
	}
}

func (ctx *appCtx) GetTokenProvider() tokenprovider.TokenProvider {
	return ctx.tokenProvider
}
