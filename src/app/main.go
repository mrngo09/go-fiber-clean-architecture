package main

import (
	"clean-architecture-go-fiber/src/app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	routes.InitRouter()
}
