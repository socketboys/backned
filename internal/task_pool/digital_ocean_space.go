package TaskPool

import (
	"context"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"os"
	"sync"
)

var client *minio.Client

const host = "https://backned.blr1.cdn.digitaloceanspaces.com/"
const spaceName = "backned"

func InitSpace() {
	accessEndpoint := "blr1.digitaloceanspaces.com"
	ssl := true
	accessKey := "DO00Q89RLRRGNK7AZAUH"
	secretAccessKey := "oaVwJJOlMlWwVTDJArVrahWsAVFbtmTxFriF7DNTLUY"
	//token := "dop_v1_6805dff11c48a4422e126bf19bacdd95e4249aab651a97f54243caf9fe38af7e"

	var err error

	if client, err = minio.New(accessEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretAccessKey, ""),
		Region: "blr1",
		Secure: ssl,
	}); err != nil {
		log.Fatalf("%v", err.Error())
	}
}

func UploadAudio(euid uuid.UUID, objectName, filePath, language string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			// logger error handling
		}
	}(filePath)
	ctx := context.Background()

	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		// logger
		UpdateTaskStatus(euid.String(), false, map[string]map[string]string{}, err)
		return
	}
	fi, err := f.Stat()
	if err != nil {
		// logger
		UpdateTaskStatus(euid.String(), false, map[string]map[string]string{}, err)
		return
	}

	var uploadOptions minio.PutObjectOptions

	if fi.Size() >= 5*1024*1024 {
		uploadOptions = minio.PutObjectOptions{
			ContentType: "audio/wav",
			//ContentEncoding:         "",
			ContentLanguage:       language,
			CacheControl:          "public, max-age=31536000",
			NumThreads:            4,
			PartSize:              1024 * 256,
			ConcurrentStreamParts: true,
			Internal:              minio.AdvancedPutOptions{},
		}
	} else {
		uploadOptions = minio.PutObjectOptions{
			ContentType: "audio/wav",
			//ContentEncoding:         "",
			ContentLanguage: language,
			CacheControl:    "public, max-age=31536000",
		}
	}

	_, err = client.FPutObject(ctx, spaceName, "audio/"+objectName, filePath, uploadOptions)
	if err != nil {
		log.Println(err.Error()) // logger
		_, err = client.FPutObject(ctx, spaceName, "audio/"+objectName, filePath, uploadOptions)
		if err != nil {
			log.Println(err.Error()) // logger
			UpdateTaskStatus(euid.String(), false, map[string]map[string]string{}, err)
			return
		} else {
			UpdateTaskLink(euid.String(), language, "audio", host+"audio/"+objectName)
		}
	} else {
		UpdateTaskLink(euid.String(), language, "audio", host+"audio/"+objectName)
	}
}

func UploadSub(euid uuid.UUID, objectName, filePath, language string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			// logger error handling
		}
	}(filePath)
	ctx := context.Background()

	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		// logger
		UpdateTaskStatus(euid.String(), false, map[string]map[string]string{}, err)
	}
	fi, err := f.Stat()
	if err != nil {
		// logger
		UpdateTaskStatus(euid.String(), false, map[string]map[string]string{}, err)
	}

	var uploadOptions minio.PutObjectOptions

	if fi.Size() >= 5*1024*1024 {
		uploadOptions = minio.PutObjectOptions{
			ContentType: "text/plain",
			//ContentEncoding:         "",
			ContentLanguage:       language,
			CacheControl:          "public, max-age=31536000",
			NumThreads:            4,
			PartSize:              1024 * 256,
			ConcurrentStreamParts: true,
			Internal:              minio.AdvancedPutOptions{},
		}
	} else {
		uploadOptions = minio.PutObjectOptions{
			ContentType: "text/plain",
			//ContentEncoding:         "",
			ContentLanguage: language,
			CacheControl:    "public, max-age=31536000",
		}
	}

	_, err = client.FPutObject(ctx, spaceName, "subtitle/"+objectName, filePath, uploadOptions)
	if err != nil {
		log.Println(err.Error()) // logger
		_, err = client.FPutObject(ctx, spaceName, "subtitle/"+objectName, filePath, uploadOptions)
		if err != nil {
			log.Println(err.Error()) // logger
			UpdateTaskStatus(euid.String(), false, map[string]map[string]string{}, err)
		} else {
			UpdateTaskLink(euid.String(), language, "subtitle", host+"subtitle/"+objectName)
		}
	} else {
		UpdateTaskLink(euid.String(), language, "subtitle", host+"subtitle/"+objectName)
	}
}
