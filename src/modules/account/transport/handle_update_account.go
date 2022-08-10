package accounttrpt

import (
	accountbiz "clean-architecture-go-fiber/src/module/account/business"
	accountmodel "clean-architecture-go-fiber/src/module/account/model"
	accountstorage "clean-architecture-go-fiber/src/module/account/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleUpdateAccount(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var id, err = strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		var dataUpdate accountmodel.Account
		if err := ctx.ShouldBind(&dataUpdate); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := accountstorage.NewPostgresStorage(db)
		biz := accountbiz.NewUpdateAccountBiz(storage)

		if err := biz.UpdateAccount(ctx.Request.Context(), map[string]interface{}{"id": id}, &dataUpdate); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}
