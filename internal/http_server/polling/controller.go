package polling

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project-x/internal/http_server/task_pool"
	"project-x/internal/http_server/utils"
)

func GetProcessStatus(c *gin.Context) {
	uuid := c.Param("uuid")

	if uuid == "" || uuid == "/" {
		utils.SendError(c, http.StatusBadRequest, "You made a wrong request, please try later")
		return
	}

	c.JSON(http.StatusOK, TaskPool.GetTaskStatus(uuid))
}
