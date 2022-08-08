package main

import (
	"clean-architecture-go-fiber/src/app/routes"
	accountmodel "clean-architecture-go-fiber/src/module/account/model"
	database "clean-architecture-go-fiber/src/platform/database/mysql"

	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	db, err := database.OpenConnectToMySQL()
	fmt.Println("Opened connection to MySQL")

	routes.InitRouter()

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	db.AutoMigrate(&accountmodel.Account{})

}
