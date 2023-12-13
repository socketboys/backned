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
	gin.SetMode(gin.ReleaseMode)
	//
	//os.Setenv("SUPPORT_EMAIL", "support@vaaani.live")
	//os.Setenv("PASSWORD", "R@nd71kg")
	//os.Setenv("SMTP_HOST", "smtp.titan.email")
	//os.Setenv("SMTP_PORT", "587")
	//os.Setenv("DO_ACCESS_ENDPOINT", "blr1.digitaloceanspaces.com")
	//os.Setenv("DO_ACCESS_KEY", "DO00Q89RLRRGNK7AZAUH")
	//os.Setenv("DO_SECRET_ACCESS_KEY", "oaVwJJOlMlWwVTDJArVrahWsAVFbtmTxFriF7DNTLUY")
	//os.Setenv("DO_REGION", "blr1")
	//os.Setenv("DO_CDN_HOST", "https://backned.blr1.cdn.digitaloceanspaces.com/")
	//os.Setenv("DO_SPACE_NAME", "backned")
	//os.Setenv("DO_TOKEN", "dop_v1_0a2e144b571f8a682103283c60743f102584c7e7b3188565217927daa58c66fe")
	//
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
