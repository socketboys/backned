package polling

import (
	"github.com/gin-gonic/gin"
	"net/http"
	TaskPool "project-x/internal/task_pool"
	"project-x/internal/utils"
)

func GetProcessStatus(c *gin.Context) {
	uuid := c.Param("uuid")

	if uuid == "" || uuid == "/" {
		utils.SendError(c, http.StatusBadRequest, "You made a wrong request, please try later")
		return
	}

	c.JSON(http.StatusOK, TaskPool.GetTaskStatus(uuid))
}
