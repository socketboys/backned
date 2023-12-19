package TaskPool

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/wagslane/go-rabbitmq"
	"log"
	"os"
	"project-x/internal/utils"
	"sync"
)

func startProcessing(euid uuid.UUID, link string, languages []string, emailId string, audioLength float32) {
	// TODO add check for credit amount less than processing cost

	extension := ".wav" // TODO create a way for this
	path := "external/input/"
	utils.Logger.Info("Audio Download started")
	DirectDownloadFile(euid.String(), path, link, extension)
	utils.Logger.Info("File Download complete")

	utils.Logger.Info("Transformers starting")
	// TODO publish message

	utils.Logger.Info("Publishing message")
	req, err := json.Marshal(PublishTranslationRequest{
		Euid:        euid.String() + ".wav",
		Language:    languages,
		EmailId:     emailId,
		AudioLength: audioLength,
	})
	if err != nil {
		utils.Logger.Error("There was an error creating byte msg for publisher", err.Error())
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
		utils.Logger.Error("There was an error publishing task to translation queue", err.Error())
	}

	utils.Logger.Info("Message published")

}

func cleanResidualFiles(path, euid, extension string) {
	err := os.Remove(path + euid + extension)
	if err != nil {
		utils.Logger.Error(err.Error())
	}
	err = os.Remove(path + euid + ".srt")
	if err != nil {
		utils.Logger.Error(err.Error())
	}
}

func getFilePrefix(language string) string {
	if language == "hindi" {
		return "_hi"
	}
	if language == "telugu" {
		return "_tel"
	}
	if language == "bengali" {
		return "_be"
	}
	if language == "assamese" {
		return "_asm"
	}
	if language == "bodo" {
		return "_bod"
	}
	if language == "gujrati" {
		return "_guj"
	}
	if language == "kannada" {
		return "_kan"
	}
	if language == "malyalam" {
		return "_mal"
	}
	if language == "marathi" {
		return "_mar"
	}
	if language == "manipuri" {
		return "_mni"
	}
	if language == "odiya" {
		return "_odi"
	}
	if language == "punjabi" {
		return "_pan"
	}
	if language == "tamil" {
		return "_tam"
	}
	return ""
}

func StartTaskPoolConsumer(msg rabbitmq.Delivery) (act rabbitmq.Action) {
	log.Printf("consumed: %v", string(msg.Body))

	var msgResponse PublishTranslationRequest

	err := json.Unmarshal(msg.Body, &msgResponse)
	if err != nil {
		return 0
	}

	utils.Logger.Info("Py Service completed")

	cleanResidualFiles("./external/input", msgResponse.Euid, ".wav")

	utils.Logger.Info("Starting upload")
	var wg sync.WaitGroup

	wg.Add(2 * len(msgResponse.Language))
	for _, langs := range msgResponse.Language {
		go UploadAudio(msgResponse.Euid, msgResponse.Euid+getFilePrefix(langs)+".wav", "external/audio/"+msgResponse.Euid+getFilePrefix(langs)+".wav", langs, &wg)
		go UploadSub(msgResponse.Euid, msgResponse.Euid+getFilePrefix(langs)+".srt", "external/subtitle/"+msgResponse.Euid+getFilePrefix(langs)+".srt", langs, &wg)
	}
	wg.Wait()

	// Deduct money on successful processing
	//TODO fix on consuming message
	//for _, language := range languages {
	//	DeductMoney(audioLength*7.083, emailId, "subtitle/"+euid.String()+getFilePrefix(language)+".srt", "audio/"+euid.String()+getFilePrefix(language)+".wav", euid) // cost as per $5/hr
	//}
	return rabbitmq.Ack
}
