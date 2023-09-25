package TaskPool

import (
	"errors"
	"io"
	"net/http"
	"os"
	"project-x/internal/utils"
)

func DirectDownloadFile(uuid, path, url, extension string) {
	file, err := os.Create(path + uuid + extension)
	if err != nil {
		utils.Logger.Error(err.Error())
		UpdateTaskStatus(uuid, false, map[string]map[string]string{}, err)
		os.Remove(path + uuid + extension)
		return
	}

	if err != nil {
		utils.Logger.Error(err.Error())
		UpdateTaskStatus(uuid, false, map[string]map[string]string{}, err)
		os.Remove(path + uuid + extension)
		return
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		utils.Logger.Error(err.Error())
		UpdateTaskStatus(uuid, false, map[string]map[string]string{}, err)
		os.Remove(path + uuid + extension)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if n, err := io.Copy(file, resp.Body); err != nil || n == 0 {
		if n == 0 {
			err = errors.New("Empty/invalid media file provided")
		}

		utils.Logger.Error(err.Error())
		UpdateTaskStatus(uuid, false, map[string]map[string]string{}, err)

		os.Remove(path + uuid + extension)
		return
	}

	if check := checkMimeType(resp.Header.Get("Content-Type")); check == false {
		utils.Logger.Error(err.Error())
		UpdateTaskStatus(uuid, false, map[string]map[string]string{}, errors.New("Invalid media link provided"))
		os.Remove(path + uuid + extension)
		return
	}

	return
}

func checkMimeType(mime string) bool {
	if mime == "audio/mpeg" || mime == "audio/wav" || mime == "audio/webm" || mime == "audio/ogg" || mime == "audio/flac" || mime == "audio/aac" || mime == "audio/mp4" || mime == "audio/x-ms-wma" || mime == "audio/x-wav" || mime == "audio/x-aiff" || mime == "audio/x-matroska" || mime == "audio/x-pn-realaudio" {
		return true
	} else {
		return false
	}
}
