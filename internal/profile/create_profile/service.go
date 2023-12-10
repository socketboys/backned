package create_profile

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project-x/internal/postgres"
	"project-x/internal/utils"
	"strings"
)

// CreateProfile Create a profile on Vaaani
// @Summary create a profile for the user using google auth
// @Description Adds a new user where the primary key will be the gmail id and profile id will be given in response
// @Tags Create Profile
// @Accept json
// @Produce json
// @Param request body CreateReqModel true "request body"
//
// @Success 200
// @Router /profile/create [post]
func CreateProfile(ctx *gin.Context) {
	var req CreateReqModel

	if err := ctx.ShouldBind(&req); err != nil {
		utils.SendError(ctx, http.StatusBadRequest, "You made a wrong request, please try later"+err.Error())
		return
	}

	tx, err := (postgres.PostgresConn).Begin()
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "profile creation txn failed "+err.Error())
		return
	}

	var profileId string
	err = tx.QueryRow(`insert into users(name, phone, email) values ($1, $2, $3) returning profile_id`, req.Name, req.Phone, req.Email).Scan(&profileId)
	if strings.Trim(err.Error(), "/n") == `ERROR: duplicate key value violates unique constraint "users_pk" (SQLSTATE 23505)` {
		utils.SendError(ctx, http.StatusBadRequest, "A profile exists for the same email id")
		return
	} else if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "profile creation txn failed "+err.Error())
		return
	}

	_, err = tx.Exec(`insert into transactions(remaining_credits, txn_amount, email) values ($1, $2, $3)`, req.InitialCredits, req.InitialCredits, req.Email)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "profile creation txn failed "+err.Error())
		return
	}
	if err = tx.Commit(); err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "profile creation txn failed "+err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"final_credits": req.InitialCredits,
		"profile_id":    profileId,
	})
}

// if we want to onboard users without credits
//if req.InitialCredits > 0 {
// 	existing code
//} else {
//	var profileId string
//	err := postgres.PostgresConn.QueryRow(`insert into users(name, phone, email) values ($1, $2, $3) returning profile_id`, req.Name, req.Phone, req.Email).Scan(&profileId)
//	if err != nil {
//		utils.SendError(ctx, http.StatusInternalServerError, "profile creation txn failed "+err.Error())
//		return
//	}
//}
