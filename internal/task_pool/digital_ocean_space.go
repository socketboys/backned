package TaskPool

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"os"
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
