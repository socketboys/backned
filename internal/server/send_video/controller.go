package send_video

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project-x/internal/task_pool"
	"project-x/internal/utils"
)

func getLanguageEnum(string) string {
	return ""
}

func DubAudio(c *gin.Context) {
	var req AudioRequest

	if err := c.ShouldBind(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "You made a wrong request, please try later"+err.Error())
		return
	}

	euid, err := TaskPool.CreateTask(req.FileLink, req.Language)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "We couldn't complete this step: "+err.Error())
		return
	}

	c.JSON(http.StatusAccepted, Response{
		Msg:  "Your request has been successfully submitted",
		EUID: euid,
	})

	return
}
