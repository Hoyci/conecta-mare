package storage

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type StorageClient struct {
	client     *minio.Client
	bucketName string
}

func NewStorageClient(endpoint, accessKey, secretKey, bucketName string) *StorageClient {
	storageClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln("failed to create minio client: ", err)
	}

	return &StorageClient{client: storageClient, bucketName: bucketName}
}

func (c *StorageClient) UploadFile(objectName string, fileHeader *multipart.FileHeader) (string, error) {
	ctx := context.Background()
	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer file.Close()

	_, err = c.client.PutObject(ctx, c.bucketName, fmt.Sprintf("%s/%s", "", objectName), file, fileHeader.Size, minio.PutObjectOptions{
		ContentType: fileHeader.Header.Get("Content-Type"),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload object %w", err)
	}

	avatarURL := fmt.Sprintf("%s/%s/%s", c.client.EndpointURL(), c.bucketName, objectName)

	return avatarURL, nil
}
