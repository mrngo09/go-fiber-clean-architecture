package accounttrpt

import (
	component "clean-architecture-go-fiber/src/components"
	accountbiz "clean-architecture-go-fiber/src/module/account/business"
	accountmodel "clean-architecture-go-fiber/src/module/account/model"
	accountstorage "clean-architecture-go-fiber/src/module/account/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleLoginWithEmail(db *gorm.DB, appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var data accountmodel.UserLogin

		if err := ctx.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := accountstorage.NewPostgresStorage(db)
		biz := accountbiz.NewLoginBiz(store, appCtx.GetTokenProvider(), 60*60*24*30)

		account, err := biz.UserLogin(ctx.Request.Context(), &data)
		if err != nil {
			panic(err)
		}
		ctx.JSON(200, gin.H{"data": account})
	}
}

// func HandleSignInWithEmail(db *gorm.DB) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {

// 		type typeDataLogin struct {
// 			Email    string `json:"email" `
// 			Password string `json:"password"`
// 		}

// 		var dataLogin typeDataLogin

// 		var dataAccount accountmodel.Account

// 		if err := ctx.ShouldBind(&dataLogin); err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if err := db.Where("email = ?", dataLogin.Email).First(&dataAccount).Error; err != nil {
// 			ctx.JSON(http.StatusNotFound, gin.H{
// 				"message": "Wrong email, try again.",
// 			})
// 			return
// 		}

// 		if dataLogin.Password != dataAccount.Password {
// 			ctx.JSON(http.StatusNotFound, gin.H{
// 				"message": "Wrong password, try again.",
// 			})
// 			return
// 		}

// 		ctx.JSON(http.StatusOK, gin.H{
// 			"token": "Generating....",
// 		})
// 	}
// }
