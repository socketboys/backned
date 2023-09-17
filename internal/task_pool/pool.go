package TaskPool

import (
	"github.com/google/uuid"
)

var tp taskPool

func init() {
	tp.Task = make(map[string]*TaskStatus)
}

func CreateTask(link, language string) (string, error) {
	euid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	go startProcessing(euid, link, language)

	tp.Task[euid.String()] = &TaskStatus{
		AudioProcessingComplete: false,
		Err:                     nil,
		Links:                   map[string]map[string]string{},
	}

	return euid.String(), nil
}

func DeleteTask(euid string) {
	delete(tp.Task, euid)
	// TODO: ask ishar to create a delete request in case the browser is closed during processing, to remove unnecessary memory allocation and abort gRPC stream for that request
}

func UpdateTaskStatus(euid string, completion bool, links map[string]map[string]string, err error) {
	tp.Task[euid] = &TaskStatus{
		AudioProcessingComplete: completion,
		Err:                     err,
		Links:                   links,
	}
}

func UpdateTaskLink(euid, language, key, value string) {
	if tp.Task[euid].Links[language] == nil {
		tp.Task[euid].Links[language] = make(map[string]string)
	}
	tp.Task[euid].Links[language][key] = value
}

func GetTaskStatus(euid string) *TaskStatus {
	if tp.Task[euid].Err != nil {
		defer DeleteTask(euid)
	}
	return tp.Task[euid]
}
