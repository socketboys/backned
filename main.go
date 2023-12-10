package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"project-x/internal"
	"project-x/internal/postgres"
	"project-x/internal/task_pool"
	"project-x/internal/utils"
)

// @title SocketBoys/Backned APIs
// @version 1.0
// @description Testing Swagger APIs.
// @termsOfService https://vaaani.live
// @contact.name API Support
// @contact.url https://www.vaaani.live
// @contact.email support@vaaani.live
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host api.vaaani.live
// @schemes http

func main() {
	gin.SetMode(gin.TestMode)

	err := godotenv.Load(".env")
	if err != nil {
		utils.Logger.Error(err.Error())
		panic(err.Error())
	}

	// TODO shift to os.Getenv()
	postgres.InitPostgres()
	TaskPool.InitSpace()
	utils.InitLogger()

	app := gin.Default()
	internal.Run(app)

	utils.Logger.Info("Running Vidyavani on :9010")

	err = app.Run(":9010")
	if err != nil {
		return
	}
}
