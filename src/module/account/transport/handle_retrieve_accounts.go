package accounttrpt

import (
	accountbiz "clean-architecture-go-fiber/src/module/account/business"
	accountmodel "clean-architecture-go-fiber/src/module/account/model"
	accountstorage "clean-architecture-go-fiber/src/module/account/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleRetrieveAccounts(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var pagging accountmodel.DataPaging

		if err := ctx.ShouldBind(&pagging); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		pagging.Process()

		store := accountstorage.NewMySQLStorage(db)
		biz := accountbiz.NewRetrieveAccountsBiz(store)

		data, err := biz.RetrieveAccounts(ctx.Request.Context(), nil, &pagging)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": data, "pagging": pagging})
	}
}
