package send_video

type AudioRequest struct {
	FileLink  string   `json:"file_link" binding:"required,url"`
	Languages []string `json:"languages" binding:"required"` // TODO: Add support for dubbing multiple languages at the same time
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
