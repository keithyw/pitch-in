package utils

import (
	"fmt"
	"io"
	"net/http"
)

type ImageInfo struct {
	MimeType string
}

func GetImageInfo(reader io.Reader) (*ImageInfo, error) {
	imageData, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed reading image: %w", err)
	}

	buf := 512
	if len(imageData) < buf {
		buf = len(imageData)
	}

	mimeType := http.DetectContentType(imageData[:buf])
	return &ImageInfo{
		MimeType: mimeType,
	}, nil	
}