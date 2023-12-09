package add_money

type AddReqModel struct {
	EmailId string `json:"email" example:"rajatn@gmail.com"`
	Credits int32  `json:"credits" example:"1000"`
}
