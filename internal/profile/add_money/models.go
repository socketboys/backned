package add_money

type AddReqModel struct {
	EmailId string `json:"email"`
	Credits int32  `json:"credits"`
}
