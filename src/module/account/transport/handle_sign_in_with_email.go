package accounttrpt

import (
	accountmodel "clean-architecture-go-fiber/src/module/account/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleSignInWithEmail(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		type typeDataLogin struct {
			Email    string `json:"email" `
			Password string `json:"password"`
		}

		var dataLogin typeDataLogin

		var dataAccount accountmodel.Account

		if err := ctx.ShouldBind(&dataLogin); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("email = ?", dataLogin.Email).First(&dataAccount).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Wrong email, try again.",
			})
			return
		}

		if dataLogin.Password != dataAccount.Password {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Wrong password, try again.",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"token": "Generating....",
		})
	}
}
