package TaskPool

import (
	"github.com/google/uuid"
	"os"
	"os/exec"
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

	go func() {
		// for streamed download
		//DownloadFile(euid.String(), "external/input/", link)

		extension := ".wav"
		path := "external/input/"
		DirectDownloadFile(euid.String(), path, link, extension)

		defer func() {
			os.Remove(path + euid.String() + extension)
		}()

		cmd := exec.Command("python3", "main.py", language, "--audioname", euid.String()+extension)
		cmd.Dir = "./pipeline-cli/langline"
		err = cmd.Run()
		if err != nil {
			// error handling
		}
		err = cmd.Wait()

		// if there is any error during the process then
		// add the error to the corresponding euid in pool
		// and on polling remove it from the task pool
	}()

	tp.Task[euid.String()] = &TaskStatus{
		AudioProcessingComplete: false,
		Err:                     nil,
		ProgressPct:             0,
	}

	return euid.String(), nil
}

func DeleteTask(euid string) {
	delete(tp.Task, euid)
	// TODO: ask ishar to create a delete request in case the browser is closed during processing, to remove unnecessary memory allocation and abort gRPC stream for that request
}

func UpdateTaskStatus(euid string, completion bool, percent float32, err error) {

	tp.Task[euid] = &TaskStatus{
		AudioProcessingComplete: completion,
		Err:                     err,
		ProgressPct:             percent,
	}
}

func GetTaskStatus(euid string) *TaskStatus {
	return tp.Task[euid]
}
