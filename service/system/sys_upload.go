package system

import (
	upload2 "go-shop-api/utils/upload"
	"mime/multipart"
)

type UploadService struct {
}

// UploadFile 上传图片
func (u *UploadService) UploadFile(header *multipart.FileHeader) (err error, filePath string) {
	upload := upload2.NewFileStore()
	filePath, _, err = upload.UploadFile(header)
	if err != nil {
		panic(err)
	}
	return

}
