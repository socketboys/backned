package TaskPool

import (
	"github.com/google/uuid"
	"os"
	"os/exec"
	"project-x/internal/utils"
	"sync"
)

func startProcessing(euid uuid.UUID, link string, languages []string) {
	// TODO for streamed download for files bigger than 5 MiB
	//DownloadFile(euid.String(), "external/input/", link)

	extension := ".wav"
	path := "external/input/"
	utils.Logger.Info("Audio Download started")
	DirectDownloadFile(euid.String(), path, link, extension)

	utils.Logger.Info("Starting command execution")

	language := ""
	for i, l := range languages {
		if i < len(languages)-1 {
			language += l + " "
		}
	}
	language += languages[len(languages)-1]

	cmd := exec.Command("python3", "inference.py", "--lang", language, "--audioname", euid.String()+extension)
	print("python3 inference.py --lang " + language + " --audioname " + euid.String() + extension)
	cmd.Dir = "../../Vaani-ML"
	err := cmd.Run()
	if err != nil {
		utils.Logger.Error(err.Error())
		UpdateTaskStatus(euid.String(), false, map[string]map[string]string{}, err)
		err = os.Remove(path + euid.String() + extension)
		if err != nil {
			utils.Logger.Error(err.Error())
		}
		return
	}
	err = cmd.Wait()
	s, err := cmd.Output()
	utils.Logger.Info(string(s))

	err = os.Remove(path + euid.String() + extension)
	if err != nil {
		utils.Logger.Error(err.Error())
	}
	err = os.Remove(path + euid.String() + ".srt")
	if err != nil {
		utils.Logger.Error(err.Error())
	}

	utils.Logger.Info("Starting upload")
	var wg sync.WaitGroup

	wg.Add(2)
	go UploadAudio(euid, euid.String()+getFilePrefix(language)+".wav", "external/audio/"+euid.String()+getFilePrefix(language)+".wav", language, &wg)
	go UploadSub(euid, euid.String()+getFilePrefix(language)+".srt", "external/subtitle/"+euid.String()+getFilePrefix(language)+".srt", language, &wg)
	wg.Wait()

	utils.Logger.Info("Uploading executed")
	tp.Task[euid.String()].AudioProcessingComplete = true
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
