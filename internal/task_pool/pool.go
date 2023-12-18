package TaskPool

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/wagslane/go-rabbitmq"
	"project-x/internal/utils"
)

var tp taskPool

func init() {
	tp.Task = make(map[string]*TaskStatus)
}

func CreateTask(link string, language []string, emailId string, audioLength float32) (string, error) {
	euid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	utils.Logger.Info("Publishing message")
	req, err := json.Marshal(PublishTranslationRequest{
		Euid:        euid.String(),
		Link:        link,
		Language:    language,
		EmailId:     emailId,
		AudioLength: audioLength,
	})
	if err != nil {
		utils.Logger.Error("There was an error creating byte msg for publisher")
		return "", errors.Join(errors.New("there was an error creating byte msg for publisher"), err)
	}

	pconf, err := RabbitProducer.PublishWithDeferredConfirmWithContext(
		context.Background(),
		req,
		[]string{"translationkey"},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsMandatory,
		rabbitmq.WithPublishOptionsPersistentDelivery,
		PublisherOptions,
	)

	if err != nil || len(pconf) == 0 {
		utils.Logger.Error("There was an error publishing task to translation queue")
		return "", errors.Join(errors.New("there was an error publishing task to translation queue"), err)
	}

	utils.Logger.Info("Message published")

	tp.Task[euid.String()] = &TaskStatus{
		AudioProcessingComplete: false,
		Err:                     "",
		Links:                   map[string]map[string]string{},
	}
	utils.Logger.Info("Task Created")

	return euid.String(), nil
}

func PublisherOptions(options *rabbitmq.PublishOptions) {
	options.Exchange = "translation"
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
