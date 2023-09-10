package main

import (
	"github.com/gin-gonic/gin"
	server2 "project-x/internal/gRPC/server"
	"project-x/internal/http_server"
	"project-x/internal/http_server/utils"
)

func main() {
	gin.SetMode(gin.TestMode)

	utils.InitLogger()

	app := gin.Default()
	server.Run(app)

	go server2.RungRPC()

	err := app.Run(":9010")
	if err != nil {
		return
	}
}
