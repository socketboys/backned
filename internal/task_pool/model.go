package TaskPool

type taskPool struct {
	Task map[string]*TaskStatus // "uuid": {ML job completed?, "error string if there was any error in process?"}
}

type TaskStatus struct {
	AudioProcessingComplete bool                         `json:"processing_complete" example:"false"`
	Err                     string                       `json:"error" example:"There was an error"`
	Links                   map[string]map[string]string `json:"links"`
}

type DeductRequest struct {
	Cost     float32 `json:"cost"`
	EmailId  string  `json:"email_id"`
	Subtitle string  `json:"subtitle"`
	Video    string  `json:"video"`
	Euid     string  `json:"euid"`
}

type PublishTranslationRequest struct {
	Euid        string   `json:"euid"`
	Link        string   `json:"link"`
	Language    []string `json:"language"`
	EmailId     string   `json:"email_id"`
	AudioLength float32  `json:"audio_length"`
}
