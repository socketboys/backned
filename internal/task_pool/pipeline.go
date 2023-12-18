package TaskPool

import (
	"github.com/google/uuid"
	"github.com/wagslane/go-rabbitmq"
	"log"
)

func startProcessing(euid uuid.UUID, link string, languages []string, emailId string, audioLength float32) {
	// TODO for streamed download for files bigger than 5 MiB
	//DownloadFile(euid.String(), "external/input/", link)

	// TODO get audio length from FE
	// TODO add check for credit amount less than processing cost

	// TODO shift to ML
	//extension := ".wav" // TODO create a way for this
	//path := "external/input/"
	//utils.Logger.Info("Audio Download started")
	//DirectDownloadFile(euid.String(), path, link, extension)
	//utils.Logger.Info("File Download complete")

	//utils.Logger.Info("Transformers starting")
	//executeTransformers(languages, euid.String(), path, extension)
	//utils.Logger.Info("Transformers done")

	//cleanResidualFiles(path, euid.String(), extension)

	//utils.Logger.Info("Starting upload")
	//var wg sync.WaitGroup
	//
	//wg.Add(2 * len(languages))
	//for _, langs := range languages {
	//go UploadAudio(euid, euid.String()+getFilePrefix(langs)+".wav", "external/audio/"+euid.String()+getFilePrefix(langs)+".wav", langs, &wg)
	//go UploadSub(euid, euid.String()+getFilePrefix(langs)+".srt", "external/subtitle/"+euid.String()+getFilePrefix(langs)+".srt", langs, &wg)
	//}
	//wg.Wait()
}

func getFilePrefix(language string) string {
	if language == "hindi" {
		return "_hi"
	} else if language == "telugu" {
		return "_te"
	} else if language == "marathi" {
		return "_ma"
	} else if language == "bengali" {
		return "_be"
	} else if language == "tamil" {
		return "_ta"
	} else {
		return ""
	}
}

func StartTaskPoolConsumer(msg rabbitmq.Delivery) (act rabbitmq.Action) {
	log.Printf("consumed: %v", string(msg.Body))
	// rabbitmq.Ack, rabbitmq.NackDiscard, rabbitmq.NackRequeue

	// Deduct money on successful processing
	//TODO fix on consuming message
	//for _, language := range languages {
	//	DeductMoney(audioLength*7.083, emailId, "subtitle/"+euid.String()+getFilePrefix(language)+".srt", "audio/"+euid.String()+getFilePrefix(language)+".wav", euid) // cost as per $5/hr
	//}
	return rabbitmq.Ack
}
