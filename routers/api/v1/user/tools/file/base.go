package file

import (
	"errors"
	"mime/multipart"
)

type Ifile interface {
	Put(*multipart.FileHeader) (string, error)
	Get(filePath string) ([]byte, error)
	Delete(filePath string) error
	SetUserId(userId uint64)
}

type fileBase struct {
	UserId uint64
}

const LOCALFILE = "LOCALFILE"
const UCLOUD = "UCLOUD"

var fileTools = map[string]Ifile{
	LOCALFILE: &localFile{},
	UCLOUD: &ucloud{},
}

func GetFileTool(toolType string) (Ifile, error) {
	if _, ok := fileTools[toolType]; !ok {
		return nil, errors.New("file tool type not exist")
	}
	return fileTools[toolType], nil
}

func GetFileToolWithUser(toolType string, userId uint64) (Ifile, error) {
	tool, err := GetFileTool(toolType)
	if err == nil {
		tool.SetUserId(userId)
	}
	return tool, err
}