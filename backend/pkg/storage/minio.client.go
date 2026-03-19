package storage

import (
	"context"
	"io"
	"net/url"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type minioClientImpl struct {
	config *StorageConfig
	ctx context.Context
	Client *minio.Client
}
	

func NewMinioClient(ctx context.Context, config *StorageConfig) (StorageClient, error) {
	client, err := minio.New(config.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: config.UseSSL,
	})
	if err != nil {
		return nil, err
	}
	return &minioClientImpl{
		config: config,
		ctx: ctx,
		Client: client,
	}, nil
}

func (c *minioClientImpl) BucketExist(name string) (bool, error) {
	exist, err := c.Client.BucketExists(c.ctx, name)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (c *minioClientImpl) Get(name string) (*GetResponse, error) {
	obj, err := c.StatObject(name)
	if err != nil {
		return nil, err
	}
	expiry := time.Hour * 24
	reqParams := make(url.Values)
	url, err := c.Client.PresignedGetObject(c.ctx, c.config.Bucket, name, expiry, reqParams)
	if err != nil {
		return nil, err
	}
	
	return &GetResponse{
		URL: url.String(),
		ContentType: obj.ContentType,
		Expiry: time.Now().Add(expiry),
		Size: obj.Size,
	}, nil
}

func (c *minioClientImpl) Put(reader io.Reader, name, contentType string) (*UploadResponse, error) {
	var size int64 = -1 // MinIO uses -1 to indicate unknown size for streaming

	// Attempt to determine size if the reader is a "Sizer" or a File
	if sizer, ok := reader.(interface{ Size() int64 }); ok {
		size = sizer.Size()
	} else if file, ok := reader.(*os.File); ok {
		if stat, err := file.Stat(); err == nil {
			size = stat.Size()
		}
	}

	obj, err := c.Client.PutObject(c.ctx, c.config.Bucket, name, reader, size, minio.PutObjectOptions{ ContentType: contentType })
	if err != nil {
		return nil, err
	}

	expiry := time.Hour * 24
	reqParams := make(url.Values)
	url, err := c.Client.PresignedGetObject(c.ctx, c.config.Bucket, name, expiry, reqParams)
	if err != nil {
		return nil, err
	}

	return &UploadResponse{
		ETag: obj.ETag,		
		Size: obj.Size,
		URL: url.String(),
	}, nil
	
	
}

func (c *minioClientImpl) Remove(name string) error {
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
	}

	err := c.Client.RemoveObject(c.ctx, c.config.Bucket, name, opts)
	if err != nil {
		return err
	}

	return nil
}

func (c *minioClientImpl) StatObject(name string) (*StorageObject, error) {
	obj, err := c.Client.StatObject(c.ctx, c.config.Bucket, name, minio.StatObjectOptions{})
	if err != nil {
		return nil, err
	}
	return &StorageObject{
		Key: obj.Key,
		ETag: obj.ETag,
		Size: obj.Size,
		ContentType: obj.ContentType,
		Expires: obj.Expires,
	}, nil
}