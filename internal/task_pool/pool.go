package TaskPool

import (
	"github.com/google/uuid"
	"project-x/internal/utils"
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

	utils.Logger.Info("Starting pipeline")

	go startProcessing(euid, link, language)

	utils.Logger.Info("Task Created")
	tp.Task[euid.String()] = &TaskStatus{
		AudioProcessingComplete: false,
		Err:                     "",
		Links:                   map[string]map[string]string{},
	}

	return euid.String(), nil
}

func DeleteTask(euid string) {
	delete(tp.Task, euid)
	utils.Logger.Debug(euid + " task deleted")
	// TODO: ask ishar to create a delete request in case the browser is closed during processing, to remove unnecessary memory allocation and abort gRPC stream for that request
}

func UpdateTaskStatus(euid string, completion bool, links map[string]map[string]string, err error) {
	utils.Logger.Info("Task status updated")
	tp.Task[euid] = &TaskStatus{
		AudioProcessingComplete: completion,
		Err:                     err.Error(),
		Links:                   links,
	}
}

func UpdateTaskLink(euid, language, key, value string) {
	utils.Logger.Debug("Task link added ")
	if tp.Task[euid].Links[language] == nil {
		tp.Task[euid].Links[language] = make(map[string]string)
	}
	tp.Task[euid].Links[language][key] = value
}

func GetTaskStatus(euid string) *TaskStatus {
	utils.Logger.Info("Task status requested")
	if tp.Task[euid].Err != "" {
		defer DeleteTask(euid)
	}
	return tp.Task[euid]
}
