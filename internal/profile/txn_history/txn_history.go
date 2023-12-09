package txn_history

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project-x/internal/postgres"
	"project-x/internal/utils"
)

func GetTxnHistory(ctx *gin.Context) {
	email := ctx.Param("email")

	rows, err := postgres.PostgresConn.Query(`select txn_id, remaining_credits, txn_amount, txn_time, subtitle, video from txns where email=$1`, email)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var txns []TxnResponse

	var txn Txn
	for rows.Next() {
		if err := rows.Scan(&txn.Txn_ID, &txn.Credits, &txn.Amount, &txn.Time, &txn.Subtitle, &txn.Video); err != nil {
			utils.SendError(ctx, http.StatusInternalServerError, err.Error())
			return
		} else {
			var tx = TxnResponse{
				Txn_ID:   txn.Txn_ID,
				Credits:  txn.Credits,
				Amount:   txn.Amount,
				Time:     txn.Time.Time,
				Subtitle: txn.Subtitle.String,
				Video:    txn.Video.String,
			}
			txns = append(txns, tx)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"txns": txns,
	})
}
