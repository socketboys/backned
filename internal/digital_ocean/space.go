package digital_ocean

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"os"
)

var client *minio.Client

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

func UploadAudio(objectName, filePath, language string) {
	defer os.Remove(filePath)
	ctx := context.Background()

	f, err := os.Open(filePath)
	fi, err := f.Stat()
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
		log.Println(err.Error())
		_, err = client.FPutObject(ctx, spaceName, "audio/"+objectName, filePath, uploadOptions)
		if err != nil {
			log.Println(err.Error())
			// error handling
		}
	}
}

func UploadSub(objectName, filePath, language string) {
	defer os.Remove(filePath)
	ctx := context.Background()

	f, err := os.Open(filePath)
	fi, err := f.Stat()

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
		log.Println(err.Error())
		_, err = client.FPutObject(ctx, spaceName, "subtitle/"+objectName, filePath, uploadOptions)
		if err != nil {
			log.Println(err.Error())
			// error handling
		}
	}
}
