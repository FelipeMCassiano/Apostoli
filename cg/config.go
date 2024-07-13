package cg

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var (
	S3Uploader  *manager.Uploader
	S3Client    *s3.Client
	RedisClient *redis.Client
)

func LoadConfigs() error {
	// The only way i found to
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		return err
	}
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	cfg, err := loadAWSConfig()
	if err != nil {
		return err
	}

	S3Client = createNewS3Client(cfg)
	S3Uploader = createNewS3Uploader(S3Client)
	RedisClient = createNewRedisClient()
	return nil
}

func loadAWSConfig() (*aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func createNewS3Client(cfg *aws.Config) *s3.Client {
	return s3.NewFromConfig(*cfg)
}

func createNewS3Uploader(*s3.Client) *manager.Uploader {
	return manager.NewUploader(S3Client)
}

func createNewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDRESS"),
		DB:   0,
	})

	return client
}
