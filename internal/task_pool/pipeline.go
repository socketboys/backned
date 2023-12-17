package TaskPool

import (
	"github.com/google/uuid"
	"os"
	"os/exec"
	"project-x/internal/utils"
	"sync"
)

func startProcessing(euid uuid.UUID, link string, languages []string, emailId string, audioLength float32) {
	// TODO for streamed download for files bigger than 5 MiB
	//DownloadFile(euid.String(), "external/input/", link)

	// TODO get audio length from FE
	// TODO add check for credit amount less than processing cost

	extension := ".wav" // TODO create a way for this
	path := "external/input/"
	utils.Logger.Info("Audio Download started")
	DirectDownloadFile(euid.String(), path, link, extension)
	utils.Logger.Info("File Download complete")

	utils.Logger.Info("Transformers starting")
	executeTransformers(languages, euid.String(), path, extension)
	utils.Logger.Info("Transformers done")

	cleanResidualFiles(path, euid.String(), extension)

	utils.Logger.Info("Starting upload")
	var wg sync.WaitGroup

	wg.Add(2 * len(languages))
	for _, langs := range languages {
		go UploadAudio(euid, euid.String()+getFilePrefix(langs)+".wav", "external/audio/"+euid.String()+getFilePrefix(langs)+".wav", langs, &wg)
		go UploadSub(euid, euid.String()+getFilePrefix(langs)+".srt", "external/subtitle/"+euid.String()+getFilePrefix(langs)+".srt", langs, &wg)
	}
	wg.Wait()

	for _, language := range languages {
		DeductMoney(audioLength*7.083, emailId, "subtitle/"+euid.String()+getFilePrefix(language)+".srt", "audio/"+euid.String()+getFilePrefix(language)+".wav", euid) // cost as per $5/hr
	} // TODO fix this
}

func executeTransformers(languages []string, euid, path, extension string) {
	utils.Logger.Info("Starting python command execution")

	langs := ""
	for i, lang := range languages {
		if i < len(languages)-1 {
			langs += lang + " "
		}
	}
	if len(languages) > 0 {
		langs += languages[len(languages)-1]
	}

	cmd := exec.Command("python3", "inference.py", "--lang", langs, "--audioname", euid+extension)

	cmd.Dir = "./Vaani-ML"

	utils.Logger.Info(cmd.String())
	err := cmd.Run()

	if err != nil {
		utils.Logger.Error(err.Error())
		UpdateTaskStatus(euid, false, map[string]map[string]string{}, err)
		err = os.Remove(path + euid + extension)
		if err != nil {
			utils.Logger.Error(err.Error())
		}
		// TODO mail user
		return
	}
	err = cmd.Wait()
	s, err := cmd.Output()
	utils.Logger.Info(string(s))
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
