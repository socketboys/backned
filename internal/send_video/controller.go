package send_video

import (
	"github.com/gin-gonic/gin"
	"net/http"
	TaskPool "project-x/internal/task_pool"
	"project-x/internal/utils"
)

func getLanguageEnum(string) string {
	return ""
}

// DubAudio Create Dub Request
// @Summary Create Dub Request
// @Description Places a dub request on the machine, if the task is successfully created then you will get the uuid, that you can use to poll and get the task status
// @Tags Create DubRequest
// @Accept json
// @Produce json
// @Param request body AudioRequest true "request body"
//
// @Success 200 {object} Response
// @Router /audio/dub [post]
func DubAudio(c *gin.Context) {
	var req AudioRequest

	if err := c.ShouldBind(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "You made a wrong request, please try later"+err.Error())
		return
	}

	utils.Logger.Info(req.FileLink)

	euid, err := TaskPool.CreateTask(req.FileLink, req.Language, req.EmailId, req.AudioLength)
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
