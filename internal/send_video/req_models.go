package send_video

type AudioRequest struct {
	FileLink    string   `json:"file_link" example:"https://www.emaple.com/file"`
	Languages   []string `json:"languages" example:"hindi, telugu, marathi, bengali, tamil"`
	EmailId     string   `json:"email" example:"rajatn@gmail.com"`
	AudioLength float32  `json:"audio_length" example:"10"`
}

type UploadHeader struct {
	FileName string
	Size     int
}

type UploadStatus struct {
	Code   int    `json:"code,omitempty"`
	Status string `json:"status,omitempty"`
	Pct    *int   `json:"pct,omitempty"`
}
