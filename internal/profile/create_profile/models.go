package create_profile

type CreateReqModel struct {
	Name           string `json:"name" example:"Rajat Kumar"`
	Email          string `json:"email" example:"rajatn@gmail.com"`
	Phone          int64  `json:"phone" example:"8010201921"`
	InitialCredits int32  `json:"initial_credits" example:"2100"` // default value
}
