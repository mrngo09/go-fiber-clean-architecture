package accounttrpt

import (
	accountbiz "clean-architecture-go-fiber/src/module/account/business"
	accountstorage "clean-architecture-go-fiber/src/module/account/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleDeleteAccount(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		storage := accountstorage.NewPostgresStorage(db)
		biz := accountbiz.NewDeleteAccountBiz(storage)
		if err := biz.DeleteAccount(ctx.Request.Context(), map[string]interface{}{"id": id}); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}
