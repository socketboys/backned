package main

import (
	"github.com/gin-gonic/gin"
	"project-x/internal"
	"project-x/internal/task_pool"
	"project-x/internal/utils"
)

// @title SocketBoys/Backned APIs
// @version 1.0
// @description Testing Swagger APIs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8081
// @schemes http

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
