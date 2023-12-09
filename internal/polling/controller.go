package polling

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project-x/internal/task_pool"
	"project-x/internal/utils"
)

// GetProcessStatus godoc
// @Summary poll ongoing task status.
// @Description get the status of ongoing task request of processing the audio.
// @Tags Poll Get
// @Accept json
// @Produce json
//
// @Param uuid path string true "uuid"
//
// @Success 200 {object} TaskPool.TaskStatus
// @Router /poll/:uuid [get]
func GetProcessStatus(c *gin.Context) {
	uuid := c.Param("uuid")

	if len(uuid) == 0 || uuid == "/" {
		utils.SendError(c, http.StatusBadRequest, "You made a wrong request, please try later")
		return
	}

	c.JSON(http.StatusOK, TaskPool.GetTaskStatus(uuid))
}
