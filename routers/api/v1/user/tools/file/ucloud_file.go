package file

import (
	"mime/multipart"
)

type ucloud struct {
	fileBase
}

func (u ucloud) Put(header *multipart.FileHeader) (string, error) {
	panic("implement me")
}

func (u ucloud) Get(filePath string) ([]byte, error) {
	panic("implement me")
}

func (u ucloud) Delete(filePath string) error {
	panic("implement me")
}

func (u ucloud) SetUserId(userId uint64) {
	u.UserId = userId
}
