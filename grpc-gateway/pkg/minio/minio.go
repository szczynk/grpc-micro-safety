package minio

import (
	"context"
	"gateway/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinioClient(ctx context.Context, cfg *config.Config) (*minio.Client, error) {
	endpoint := cfg.Minio.Endpoint //examples: localhost:9000
	accessKeyId := cfg.Minio.AccessKeyId
	secretAccessKey := cfg.Minio.SecretAccessKey
	useSSL := cfg.Minio.UseSSL //examples: false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyId, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err // fatal
	}

	if cfg.Minio.NewBucket {
		bucketName := cfg.Minio.Bucket //examples: safety
		location := cfg.Minio.Location //examples: us-east-1

		err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{
			Region: location,
		})
		if err != nil {
			// Check to see if we already own this bucket (which happens if you run this twice)
			exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
			if errBucketExists == nil && exists {
				return minioClient, err //errorf
			} else {
				return nil, err //fatal
			}
		}
	}

	return minioClient, err
}
