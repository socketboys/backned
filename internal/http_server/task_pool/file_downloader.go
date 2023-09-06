package TaskPool

import (
	"fmt"
	"net/http"
	"sync"
)

func DownloadFile(uuid, path, url string) {
	fileDetail, err := http.Head(url)
	if err != nil {
		UpdateTaskStatus(uuid, false, -1, err)
		return
	}

	chunkSize := 1024 // TODO find the optimal chunk size instead of 1024 every time
	chunks := fileDetail.ContentLength / int64(chunkSize)
	if fileDetail.ContentLength%chunks != 0 {
		chunks++
	}

	var wg sync.WaitGroup
	wg.Add(int(chunks))

	for i := int64(0); i < chunks; i++ {
		go StreamChunks(i, chunkSize, uuid, url, fileDetail, &wg)
	}

	return
}

func StreamChunks(i int64, chunkSize int, uuid, url string, fileDetail *http.Response, wg *sync.WaitGroup) {
	defer (*wg).Done()

	start := int(i) * chunkSize     // chunkSize
	end := int64(start + chunkSize) // chunkSize
	if end > fileDetail.ContentLength {
		end = fileDetail.ContentLength
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		UpdateTaskStatus(uuid, false, -1, err)
		return
	} else {
		req.Header.Add("Range", fmt.Sprintf("bytes=%v-%v", i, end))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		UpdateTaskStatus(uuid, false, -1, err)
		return
	} else if resp.StatusCode != http.StatusPartialContent {
		UpdateTaskStatus(uuid, false, -1, fmt.Errorf("expected HTTP 206, but the file was not partially chunked and status was %v", resp.StatusCode))
		return
	}

	// create a heap implementation to stream these chunks according to their order to ML BE
}
