package TaskPool

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func DownloadFile(uuid, path, url string) {
	fileDetail, err := (&http.Client{Timeout: time.Second * 120}).Head(url)

	if err != nil {
		log.Println(err.Error())
		UpdateTaskStatus(uuid, false, -1, err)
		return
	}

	file, err := os.Create(path + uuid + ".wav")

	defer file.Close()

	if err != nil {
		log.Println(err.Error())
		UpdateTaskStatus(uuid, false, -1, err)
		return
	}

	chunkSize := 1024 // TODO find the optimal chunk size instead of 1024 every time
	chunks := fileDetail.ContentLength / int64(chunkSize)

	if ((fileDetail.ContentLength) % chunks) != 0 {
		chunks++
	}

	var wg sync.WaitGroup
	wg.Add(int(chunks))
	bytes := make([][]byte, chunks)

	for i := int64(0); i < chunks; i++ {
		go StreamChunks(i, chunkSize, uuid, url, &bytes, fileDetail, &wg)
	}
	wg.Wait()

	for i := 0; i < int(chunks); i++ {
		if n, err := file.Write(bytes[i]); n == 0 || err != nil {
			UpdateTaskStatus(uuid, false, -1, err)
			return
		}
	}

	return
}

func StreamChunks(i int64, chunkSize int, uuid, url string, bytes *[][]byte, fileDetail *http.Response, wg *sync.WaitGroup) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
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

	resp, err := (&http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}).Do(req)
	if err != nil {
		UpdateTaskStatus(uuid, false, -1, err)
		return
	} else if resp.StatusCode != http.StatusPartialContent {
		UpdateTaskStatus(uuid, false, -1, fmt.Errorf("expected HTTP 206, but the file was not partially chunked and status was %v", resp.StatusCode))
		return
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		UpdateTaskStatus(uuid, false, -1, err)
		return
	}
	(*bytes)[i] = data
}
