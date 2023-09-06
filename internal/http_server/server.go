package server

import (
	"github.com/gin-gonic/gin"
	"project-x/internal/http_server/polling"
	"project-x/internal/http_server/send_video"
)

func Run(app *gin.Engine) {

	app.Use(CORSMiddleware())
	app.Use(gin.Logger())
	app.Use(gin.CustomRecovery(RecoveryAPI))

	//app.Static("/", "../views/layouts")

	app.POST("/audio/dub", send_video.DubAudio)

	app.GET("/poll/:uuid", polling.GetProcessStatus)
}
