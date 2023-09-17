package TaskPool

import (
	"github.com/google/uuid"
	"os"
	"os/exec"
	"project-x/internal/digital_ocean"
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

func GetTaskStatus(euid string) *TaskStatus {
	return tp.Task[euid]
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

func startProcessing(euid uuid.UUID, link, language string) {
	// for streamed download
	//DownloadFile(euid.String(), "external/input/", link)

	extension := ".wav"
	path := "external/input/"
	DirectDownloadFile(euid.String(), path, link, extension)

	cmd := exec.Command("python3", "main.py", language, "--audioname", euid.String()+extension)
	cmd.Dir = "./pipeline-cli/langline"
	err := cmd.Run()
	if err != nil {
		// error handling
	}
	err = cmd.Wait()

	err = os.Remove(path + euid.String() + extension)
	if err != nil {
		// error handling
	}
	err = os.Remove(path + euid.String() + ".srt")
	if err != nil {
		// error handling
	}

	go digital_ocean.UploadAudio(euid.String()+getFilePrefix(language)+".wav", "external/audio/"+euid.String()+getFilePrefix(language)+".wav", language)
	go digital_ocean.UploadSub(euid.String()+getFilePrefix(language)+".srt", "external/subtitle/"+euid.String()+getFilePrefix(language)+".srt", language)
	// if there is any error during the process then
	// add the error to the corresponding euid in pool
	// and on polling remove it from the task pool
}
