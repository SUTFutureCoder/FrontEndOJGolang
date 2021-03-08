package file

import "mime/multipart"

type FileTools interface {
	Put(*multipart.FileHeader) (string, error)
	Get(filePath string) error
	Delete(filePath string) error
}
