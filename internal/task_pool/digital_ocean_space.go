package TaskPool

import (
	"context"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"os"
	"project-x/internal/utils"
	"sync"
)

var client *minio.Client

var host string
var spaceName string

func InitSpace() {
	host = os.Getenv("DO_CDN_HOST")
	spaceName = os.Getenv("DO_SPACE_NAME")
	accessEndpoint := os.Getenv("DO_ACCESS_ENDPOINT")
	ssl := true
	accessKey := os.Getenv("DO_ACCESS_KEY")
	secretAccessKey := os.Getenv("DO_SECRET_ACCESS_KEY")
	//token := os.GetEnv("DO_TOKEN")

	var err error

	if client, err = minio.New(accessEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretAccessKey, ""),
		Region: os.Getenv("DO_REGION"),
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
			utils.Logger.Error(err.Error())
		}
	}(filePath)
	ctx := context.Background()

	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		// logger
		utils.Logger.Error(err.Error())
		UpdateTaskStatus(euid.String(), false, map[string]map[string]string{}, err)
		// TODO mail user
		return
	}
	fi, err := f.Stat()
	if err != nil {
		utils.Logger.Error(err.Error())
		UpdateTaskStatus(euid.String(), false, map[string]map[string]string{}, err)
		// TODO mail user
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
		utils.Logger.Error(err.Error())
		_, err = client.FPutObject(ctx, spaceName, "audio/"+objectName, filePath, uploadOptions)
		if err != nil {
			utils.Logger.Error(err.Error())
			UpdateTaskStatus(euid.String(), false, map[string]map[string]string{}, err)
			return
		} else {
			utils.Logger.Debug("Successful audio upload")
			UpdateTaskLink(euid.String(), language, "audio", host+"audio/"+objectName)
		}
		// TODO mail user
	} else {
		utils.Logger.Debug("Successful audio upload")
		UpdateTaskLink(euid.String(), language, "audio", host+"audio/"+objectName)
	}
}

func UploadSub(euid uuid.UUID, objectName, filePath, language string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			utils.Logger.Error(err.Error())
		}
	}(filePath)
	ctx := context.Background()

	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		utils.Logger.Error(err.Error())
		UpdateTaskStatus(euid.String(), false, map[string]map[string]string{}, err)
		// TODO mail user
	}
	fi, err := f.Stat()
	if err != nil {
		utils.Logger.Error(err.Error())
		UpdateTaskStatus(euid.String(), false, map[string]map[string]string{}, err)
		// TODO mail user
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
		utils.Logger.Error(err.Error())
		_, err = client.FPutObject(ctx, spaceName, "subtitle/"+objectName, filePath, uploadOptions)
		if err != nil {
			utils.Logger.Error(err.Error())
			UpdateTaskStatus(euid.String(), false, map[string]map[string]string{}, err)
		} else {
			utils.Logger.Debug("Successful upload of subtitle")
			UpdateTaskLink(euid.String(), language, "subtitle", host+"subtitle/"+objectName)
		}
		// TODO mail user
	} else {
		utils.Logger.Debug("Successful upload of subtitle")
		UpdateTaskLink(euid.String(), language, "subtitle", host+"subtitle/"+objectName)
	}
}
