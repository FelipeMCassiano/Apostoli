package pkg

import (
	"context"
	"os"

	"github.com/FelipeMCassiano/Apostoli/cg"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

// TODO: download from s3

func UploadFile(ctx context.Context, key string) error {
	_, err := cg.S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Key:    aws.String(key),
		Body:   os.Stdin,
	})
	if err != nil {
		return err
	}
	return nil
}
