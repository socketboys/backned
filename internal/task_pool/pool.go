package TaskPool

import (
	"github.com/google/uuid"
	"project-x/internal/utils"
)

var tp taskPool

func init() {
	tp.Task = make(map[string]*TaskStatus)
}

func CreateTask(link, language, emailId string, audioLength float32) (string, error) {
	euid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	utils.Logger.Info("Starting pipeline")

	go startProcessing(euid, link, language, emailId, audioLength)

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
	// TODO: ask ishar to create a delete request in case the browser is closed during processing, to remove unnecessary memory allocation
}

func UpdateTaskStatus(euid string, completion bool, links map[string]map[string]string, err error) {
	utils.Logger.Info("Task status updated")
	tp.Task[euid] = &TaskStatus{
		AudioProcessingComplete: completion,
		Err:                     err.Error(),
		Links:                   links,
	}
}

func UpdateTaskCompletionStatus(euid string, completion bool, err error) {
	utils.Logger.Info("Task status updated")
	tp.Task[euid].AudioProcessingComplete = completion
	if err != nil {
		tp.Task[euid].Err = err.Error()
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
	if len(tp.Task[euid].Err) > 0 {
		defer DeleteTask(euid)
	}
	return tp.Task[euid]
}
