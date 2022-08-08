package accounttrpt

import (
	accountmodel "clean-architecture-go-fiber/src/module/account/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandlerFindAnAccount(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var dataAccount accountmodel.Account

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		if err := db.Where("id =?", id).First(&dataAccount).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": dataAccount,
		})
	}
}
