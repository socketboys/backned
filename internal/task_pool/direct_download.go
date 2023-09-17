package TaskPool

import (
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
		return
	}

	if err != nil {
		utils.Logger.Error(err.Error())
		UpdateTaskStatus(uuid, false, map[string]map[string]string{}, err)
		return
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		utils.Logger.Error(err.Error())
		UpdateTaskStatus(uuid, false, map[string]map[string]string{}, err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if n, err := io.Copy(file, resp.Body); err != nil || n == 0 {
		utils.Logger.Error(err.Error())
		UpdateTaskStatus(uuid, false, map[string]map[string]string{}, err)
		return
	}

	return
}
