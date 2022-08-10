package accounttrpt

import (
	accountbiz "clean-architecture-go-fiber/src/module/account/business"
	accountmodel "clean-architecture-go-fiber/src/module/account/model"
	accountstorage "clean-architecture-go-fiber/src/module/account/storage"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleCreateAccount(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dataAccount accountmodel.Account

		if err := ctx.ShouldBind(&dataAccount); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		dataAccount.Email = strings.TrimSpace(dataAccount.Email)

		storage := accountstorage.NewPostgresStorage(db)
		business := accountbiz.NewCreateAccountbiz(storage)

		if err := business.CreateNewAccount(ctx.Request.Context(), &dataAccount); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": dataAccount.Id})

	}
}
