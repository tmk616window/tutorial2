package minio

import (
	minio_go "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const endpoint = "minio:9000"
const accessKeyID = "minio"
const secretAccessKey = "password"
const useSSL = false
// const expiry time.Duration = time.Second * 24 * 60 * 60 // 1 day TODO:後で使う

type client struct {
	*minio_go.Client
}

type ClientInterface interface {

}

func NewClient() (client, error) {
	minioClient, err := minio_go.New(endpoint, &minio_go.Options{
	Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	Secure: useSSL,
	})
	if err != nil {
		return client{nil}, err
	}

	return client{minioClient}, nil
} 