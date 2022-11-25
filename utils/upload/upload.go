package upload

import (
	"go-shop-api/global"
	"mime/multipart"
)

type FileStore interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

func NewFileStore() FileStore {
	switch global.GA_CONFIG.ApplicationConfig.UploadType {
	case "local":
		return &Local{}
	default:
		return &Local{}
	}
}
