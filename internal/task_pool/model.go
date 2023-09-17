package TaskPool

type taskPool struct {
	Task map[string]*TaskStatus // "uuid": {ML job completed?, "error string if there was any error in process?"}
}

type TaskStatus struct {
	AudioProcessingComplete bool                         `json:"processing_complete"`
	Err                     error                        `json:"error"`
	Links                   map[string]map[string]string `json:"links"`
}
