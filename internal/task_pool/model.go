package TaskPool

type taskPool struct {
	Task map[string]*TaskStatus // "uuid": {ML job completed?, "error string if there was any error in process?"}
}

type TaskStatus struct {
	AudioProcessingComplete bool    `json:"audio_processing_complete"`
	Err                     error   `json:"error"`
	ProgressPct             float32 `json:"progress_pct"`
}
