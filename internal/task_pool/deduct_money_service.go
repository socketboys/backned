package TaskPool

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
	"project-x/internal/postgres"
	"project-x/internal/utils"
)

// CutMoneyService deduct money
// @Summary deduct x amount of money from user's account
// @Description -ve credits are not yet handled
// @Tags Delete CutMoneyService
// @Accept json
// @Produce json
// @Param request body DeductRequest true "request body"
//
// @Success 200
// @Router /profile/deduct_money [post]
func CutMoneyService(ctx *gin.Context) {
	var req DeductRequest
	err := ctx.Bind(&req)
	if err != nil {
		utils.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var euid uuid.UUID
	euid.Scan(req.Euid)

	credits, err := DeductMoney(req.Cost, req.EmailId, req.Subtitle, req.Video, euid)
	if err != nil {
		utils.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"final_credits": credits,
	})
}

func DeductMoney(cost float32, emailId, subtitle, video string, euid uuid.UUID) (finalCredits float32, err error) {
	err = postgres.PostgresConn.QueryRow(`insert into transactions(remaining_credits, txn_id, txn_amount, email, subtitle, video) select tc.credits, $1, $2, $3, $4, $5 from (select (remaining_credits + ($2)) as credits from txns where email=$3 limit(1)) as tc returning remaining_credits;`,
		"vaani-"+euid.String(),
		-cost,
		emailId,
		os.Getenv("DO_CDN_HOST")+subtitle,
		os.Getenv("DO_CDN_HOST")+video).Scan(&finalCredits)

	if err != nil {
		utils.Logger.Error(err.Error())
		UpdateTaskStatus(euid.String(), false, map[string]map[string]string{}, err)
		// TODO mail user
		return
	} else {
		utils.Logger.Info("Uploading executed")
		UpdateTaskCompletionStatus(euid.String(), true, err)
		// TODO mail user
		//email_user.MailUser("rajatnd9@gmail.com", "your video processing was successful")
		// /video/[translation_id]
	}
	return
}
