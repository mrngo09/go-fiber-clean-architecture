package accounttrpt

import (
	accountbusiness "clean-architecture-go-fiber/src/module/account/business"
	accountstorage "clean-architecture-go-fiber/src/module/account/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandlerFindAnAccount(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		storage := accountstorage.NewMySQLStorage(db)
		biz := accountbusiness.NewFindAccountBiz(storage)
		data, err := biz.FindAnAccount(ctx.Request.Context(), map[string]interface{}{"id": id})

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": data})

	}
}
