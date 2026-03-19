package storage

import (
	"io"
	"time"
)

type GetResponse struct {	
	ContentType string
	Expiry time.Time
	Size int64
	URL string
}

type UploadResponse struct {
	ETag string
	Size int64
	URL string
}

type StorageObject struct {
	ETag string
	Key string
	Size int64
	ContentType string
	Expires time.Time
}

type StorageClient interface {
	BucketExist (name string) (bool, error)
	Get(name string) (*GetResponse, error)
	Put(reader io.Reader, name, contentType string) (*UploadResponse, error)
	Remove(name string) error
	StatObject(name string) (*StorageObject, error)
}