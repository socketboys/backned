package main

import (
	"github.com/gin-gonic/gin"
	"project-x/internal"
	"project-x/internal/task_pool"
	"project-x/internal/utils"
)

func main() {
	gin.SetMode(gin.TestMode)

	// TODO shift to os.Getenv()
	TaskPool.InitSpace()
	utils.InitLogger()

	app := gin.Default()
	internal.Run(app)

	err := app.Run(":9010")
	if err != nil {
		return
	}
}
