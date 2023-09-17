package TaskPool

import (
	"github.com/google/uuid"
	"os"
	"os/exec"
	"sync"
)

func startProcessing(euid uuid.UUID, link, language string) {
	// TODO for streamed download for files bigger than 5 MiB
	//DownloadFile(euid.String(), "external/input/", link)

	extension := ".wav"
	path := "external/input/"
	DirectDownloadFile(euid.String(), path, link, extension)

	cmd := exec.Command("python3", "main.py", language, "--audioname", euid.String()+extension)
	cmd.Dir = "./pipeline-cli/langline"
	err := cmd.Run()
	if err != nil {
		UpdateTaskStatus(euid.String(), false, map[string]map[string]string{}, err)
		err = os.Remove(path + euid.String() + extension)
		if err != nil {
			// logger error handling
		}
		return
	}
	err = cmd.Wait()

	err = os.Remove(path + euid.String() + extension)
	if err != nil {
		// logger error handling
	}
	err = os.Remove(path + euid.String() + ".srt")
	if err != nil {
		// logger error handling
	}

	var wg sync.WaitGroup

	wg.Add(2)
	go UploadAudio(euid, euid.String()+getFilePrefix(language)+".wav", "external/audio/"+euid.String()+getFilePrefix(language)+".wav", language, &wg)
	go UploadSub(euid, euid.String()+getFilePrefix(language)+".srt", "external/subtitle/"+euid.String()+getFilePrefix(language)+".srt", language, &wg)
	wg.Wait()

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
