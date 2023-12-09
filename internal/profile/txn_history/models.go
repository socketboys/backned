package txn_history

import (
	"database/sql"
	"time"
)

type Txn struct {
	Txn_ID   string         `json:"txn_id" field:"txn_id"`
	Credits  float32        `json:"final_credits" field:"remaining_credits"`
	Amount   float32        `json:"txn_amount" field:"txn_amount"`
	Time     sql.NullTime   `json:"time" field:"time"`
	Subtitle sql.NullString `json:"subtitle" field:"subtitle"`
	Video    sql.NullString `json:"video" field:"video"`
}

type TxnResponse struct {
	Txn_ID   string    `json:"txn_id" field:"txn_id"`
	Credits  float32   `json:"final_credits" field:"remaining_credits"`
	Amount   float32   `json:"txn_amount" field:"txn_amount"`
	Time     time.Time `json:"time" field:"time"`
	Subtitle string    `json:"subtitle" field:"subtitle"`
	Video    string    `json:"video" field:"video"`
}
