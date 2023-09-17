package main

import (
	"github.com/gin-gonic/gin"
	"project-x/internal"
	"project-x/internal/digital_ocean"
	"project-x/internal/utils"
)

func main() {
	gin.SetMode(gin.TestMode)

	digital_ocean.InitSpace()
	utils.InitLogger()

	app := gin.Default()
	internal.Run(app)

	err := app.Run(":9010")
	if err != nil {
		return
	}
}
