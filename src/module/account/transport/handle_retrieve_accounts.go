package accounttrpt

import (
	accountmodel "clean-architecture-go-fiber/src/module/account/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleRetrieveAccounts(db *gorm.DB) gin.HandlerFunc {

	dataAccount := []accountmodel.Account{}
	return func(ctx *gin.Context) {
		if err := db.Find(&dataAccount).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": dataAccount,
		})
	}
}
