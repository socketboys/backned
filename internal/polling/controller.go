package polling

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project-x/internal/task_pool"
	"project-x/internal/utils"
)

func GetProcessStatus(c *gin.Context) {
	uuid := c.Param("uuid")

	if len(uuid) == 0 || uuid == "/" {
		utils.SendError(c, http.StatusBadRequest, "You made a wrong request, please try later")
		return
	}

	//for i := 0; i < 3000; i++ {
	//	go func() {
	//		b, _ := exec.Command("date").Output()
	//		log.Printf("%s", b)
	//	}()
	//}

	c.JSON(http.StatusOK, TaskPool.GetTaskStatus(uuid))
}
