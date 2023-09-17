package internal

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"project-x/internal/polling"
	"project-x/internal/send_video"
)

func Run(app *gin.Engine) {

	app.Use(CORSMiddleware())
	app.Use(gin.Logger())
	app.Use(gin.CustomRecovery(RecoveryAPI))

	//app.Static("/", "../views/layouts")

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.POST("/audio/dub", send_video.DubAudio)
	app.GET("/poll/:uuid", polling.GetProcessStatus)
}
