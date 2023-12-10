package send_video

type AudioRequest struct {
	FileLink    string   `json:"file_link" binding:"required,url" example:"https://www.emaple.com/file"`
	Languages   []string `json:"languages" binding:"required" example:"hindi|telugu|marathi|bengali|tamil"` // TODO: Add support for dubbing multiple languages at the same time
	EmailId     string   `json:"email" binding:"required" example:"rajatn@gmail.com"`
	AudioLength float32  `json:"audio_length" binding:"required" example:"10"`
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
