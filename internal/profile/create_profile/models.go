package create_profile

type CreateReqModel struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Phone          int64  `json:"phone,omitempty"`
	InitialCredits int32  `json:"initial_credits,omitempty"` // default value
}
