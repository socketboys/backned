package TaskPool

import (
	"github.com/google/uuid"
	"os"
	"os/exec"
	"project-x/internal/utils"
	"sync"
)

func startProcessing(euid uuid.UUID, link, language string) {
	// TODO for streamed download for files bigger than 5 MiB
	//DownloadFile(euid.String(), "external/input/", link)

	extension := ".wav"
	path := "external/input/"
	utils.Logger.Info("Audio Download started")
	DirectDownloadFile(euid.String(), path, link, extension)

	convertTo16kHz(path, euid.String(), extension)
	executeWhisper(path, euid.String(), extension)

	executeTransformers(language, euid.String(), path, extension)
	cleanResidualFiles(path, euid.String(), extension)

	utils.Logger.Info("Starting upload")
	var wg sync.WaitGroup

	wg.Add(2)
	go UploadAudio(euid, euid.String()+getFilePrefix(language)+".wav", "external/audio/"+euid.String()+getFilePrefix(language)+".wav", language, &wg)
	go UploadSub(euid, euid.String()+getFilePrefix(language)+".srt", "external/subtitle/"+euid.String()+getFilePrefix(language)+".srt", language, &wg)
	wg.Wait()

	utils.Logger.Info("Uploading executed")
	tp.Task[euid.String()].AudioProcessingComplete = true
}

func convertTo16kHz(path, euid, extension string) {
	utils.Logger.Info("Converting to 16kHz using FFmPEG")
	cmd := exec.Command("ffmpeg", "-i", "../../external/input/"+euid+extension, "-ar", "16000", "../../external/input/"+euid+"_"+extension)
	cmd.Dir = "./pipeline-cli/whisper"
	err := cmd.Run()
	if err != nil {
		utils.Logger.Error(err.Error())
		UpdateTaskStatus(euid, false, map[string]map[string]string{}, err)
		err = os.Remove(path + euid + extension)
		if err != nil {
			utils.Logger.Error(err.Error())
		}
		return
	}
	err = cmd.Wait()
	s, err := cmd.Output()
	err = os.Remove("./external/input/" + euid + extension)
	if err != nil {
		utils.Logger.Error(err.Error())
		UpdateTaskStatus(euid, false, map[string]map[string]string{}, err)
		err = os.Remove(path + euid + extension)
		if err != nil {
			utils.Logger.Error(err.Error())
		}
		return
	}
	err = os.Rename("./external/input/"+euid+"_"+extension, "./external/input/"+euid+extension)
	if err != nil {
		utils.Logger.Error(err.Error())
		UpdateTaskStatus(euid, false, map[string]map[string]string{}, err)
		err = os.Remove(path + euid + extension)
		if err != nil {
			utils.Logger.Error(err.Error())
		}
		return
	}
	utils.Logger.Info(string(s))
}

func executeWhisper(path, euid, extension string) {
	utils.Logger.Info("Starting whisper.cpp command execution")
	cmd := exec.Command("./main", "--output-srt", "true", "-of", "../../external/input/"+euid+extension, "-f", "../../external/input/"+euid+extension)
	cmd.Dir = "./pipeline-cli/whisper"
	utils.Logger.Info(cmd.String())
	err := cmd.Run()
	if err != nil {
		utils.Logger.Error(err.Error())
		UpdateTaskStatus(euid, false, map[string]map[string]string{}, err)
		err = os.Remove(path + euid + extension)
		if err != nil {
			utils.Logger.Error(err.Error())
		}
		return
	}
	err = cmd.Wait()
	s, err := cmd.Output()
	utils.Logger.Info(string(s))
}

func executeTransformers(language, euid, path, extension string) {
	utils.Logger.Info("Starting python command execution")
	cmd := exec.Command("python3", "main.py", language, "--audioname", euid+extension+".srt")
	cmd.Dir = "./pipeline-cli/langline"
	utils.Logger.Info(cmd.String())
	err := cmd.Run()
	if err != nil {
		utils.Logger.Error(err.Error())
		UpdateTaskStatus(euid, false, map[string]map[string]string{}, err)
		err = os.Remove(path + euid + extension)
		if err != nil {
			utils.Logger.Error(err.Error())
		}
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
