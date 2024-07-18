package pkg

import (
	"context"
	"log"
	"os"

	"github.com/FelipeMCassiano/Apostoli/cg"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

// TODO: download from s3

func UploadFile(key string) error {
	fil, err := os.Open(key)
	if err != nil {
		log.Println("open err:", err.Error())
		return err
	}
	out, err := cg.S3Uploader.Upload(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Key:    aws.String(key),
		Body:   fil,
	})
	if err != nil {
		log.Println("s3 error:", err.Error())
		return err
	}
	log.Println(out)
	return nil
}
