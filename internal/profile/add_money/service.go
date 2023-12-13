package add_money

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project-x/internal/postgres"
	"project-x/internal/utils"
)

// AddMoney Add money to user's account
// @Summary Adds a new transaction record to the transaction table
// @Description Adds a new transaction record to the transaction table, and you will get the final credits of user's account
// @Tags Create AddMoney
// @Accept json
// @Produce json
// @Param request body AddReqModel true "request body"
//
// @Success 200
// @Router /profile/add_money [post]
func AddMoney(ctx *gin.Context) {
	var req AddReqModel

	if err := ctx.BindJSON(&req); err != nil {
		utils.SendError(ctx, http.StatusBadRequest, "You made a wrong request, please try later"+err.Error())
		return
	}

	if req.Credits > 0 {
		var money float32
		err := postgres.PostgresConn.QueryRow(`insert into transactions(remaining_credits, txn_amount, email)
		select tc.credits, $1, $2 from (select (remaining_credits + $1) as credits from txns where email = $2 limit(1)) as tc returning remaining_credits;`, req.Credits, req.EmailId).Scan(&money)
		if err != nil {
			utils.SendError(ctx, http.StatusInternalServerError, "profile creation txn failed "+err.Error())
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"final_credits": money,
			})
			// TODO mail user
		}
	} else {
		utils.SendError(ctx, http.StatusInternalServerError, "incorrect credit amount")
		return
	}
}
