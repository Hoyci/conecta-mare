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
  endpoint string
  useSSL bool
}

func NewStorageClient(endpoint, accessKey, secretKey, bucketName, environment string) *StorageClient {
  useSSL := environment == "prod" 
	storageClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln("failed to create minio client: ", err)
	}

  ctx := context.Background()
  exists, err := storageClient.BucketExists(ctx, bucketName)
  if err != nil {
    log.Fatalf("failed to check if bucket %s exists: %v", bucketName, err)
  }
  if !exists {
    err = storageClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
    if err != nil {
      log.Fatalf("failed to create bucket %s: %v", bucketName, err)
    }
  }

  return &StorageClient{
    client: storageClient, 
    bucketName: bucketName, 
    endpoint: endpoint,
    useSSL: useSSL,
  }
}

func (c *StorageClient) UploadFile(
  objectName string, 
  fileHeader *multipart.FileHeader,
) (string, error) {
  ctx := context.Background()
  file, err := fileHeader.Open()
  if err != nil {
    return "", fmt.Errorf("failed to open uploade file: %w", err)
  }
  defer file.Close()

  _, err = c.client.PutObject(
    ctx,
    c.bucketName,
    objectName,
    file,
    fileHeader.Size,
    minio.PutObjectOptions{
      ContentType: fileHeader.Header.Get("Content-Type"),
    },
    )
  if err != nil {
    return "", fmt.Errorf("failed to upload object: %w", err)
  } 
  scheme := "http"
  if c.useSSL {
    scheme = "https"
  }
  objectURL := fmt.Sprintf("%s://%s/%s/%s", scheme, c.endpoint, c.bucketName, objectName)
  return objectURL, nil
}
