package middlewares

import (
	"clean-architecture-go-fiber/src/common"
	component "clean-architecture-go-fiber/src/components"
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewFullErrorResponse(401,
		err,
		fmt.Sprintf("Wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
		"Unauthorization",
	)
}

func extracTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

func RequireAuth(appCtx component.AppContext) func(c *gin.Context) {
	tokenProvider := appCtx.GetTokenProvider()
	return func(c *gin.Context) {
		token, err := extracTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		c.Set(common.KeyUserHeader, payload.UserId)
		c.Next()
	}
}

func RequireAdminAuth(appCtx component.AppContext) func(c *gin.Context) {
	tokenProvider := appCtx.GetTokenProvider()
	return func(c *gin.Context) {
		token, err := extracTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		if payload.Role != common.RoleAdmin {
			panic(common.ErrPermissionDenied)
		}

		c.Set(common.KeyUserHeader, payload.UserId)
		c.Next()
	}
}

var (
	ErrNotFound = common.NewFullErrorResponse(401,
		errors.New("token not found"),
		"token not found",
		"ErrNotFound",
		"ErrNotFound",
	)
	ErrInvalidToken = common.NewFullErrorResponse(401, errors.New("invalid token provided"),
		"invalid token provided",
		"ErrInvalidToken",
		"ErrInvalidToken",
	)
)
