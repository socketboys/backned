package internal

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"project-x/internal/polling"
	"project-x/internal/profile/add_money"
	"project-x/internal/profile/create_profile"
	"project-x/internal/profile/txn_history"
	"project-x/internal/send_video"
	"project-x/internal/task_pool"
)

func Run(app *gin.Engine) {

	app.Use(CORSMiddleware())
	app.Use(gin.Logger())
	app.Use(gin.CustomRecovery(RecoveryAPI))

	//app.Static("/", "../views/layouts")

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.POST("/audio/dub", send_video.DubAudio)
	app.GET("/poll/:uuid", polling.GetProcessStatus)

	profile := app.Group("/profile")
	{
		profile.POST("/create", create_profile.CreateProfile)
		profile.POST("/add_money", add_money.AddMoney)
		profile.POST("/deduct_money", TaskPool.CutMoneyService)
		profile.GET("/txn_history/:email", txn_history.GetTxnHistory)
	}
}
