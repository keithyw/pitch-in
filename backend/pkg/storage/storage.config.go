package storage

import (
	"os"
	"strconv"
)

type StorageConfig struct {
	Endpoint string
	AccessKeyID string
	SecretAccessKey string
	UseSSL bool
	Bucket string // may change in the future. will manage via a folder for now
	Region string
}

func NewStorageConfig() *StorageConfig {
	// todo: get away from env vars
	useSSL, _ := strconv.ParseBool(os.Getenv("STORAGE_USE_SSL"))	
	return &StorageConfig{
		Endpoint: os.Getenv("S3_ENDPOINT"),
		AccessKeyID: os.Getenv("S3_ACCESS_KEY"),
		SecretAccessKey: os.Getenv("S3_SECRET_KEY"),
		UseSSL: useSSL,
		Bucket: os.Getenv("S3_BUCKET"),
		Region: os.Getenv("STORAGE_REGION"),
	}
}