package file

import (
	"FrontEndOJGolang/pkg/setting"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

type localFile struct {
	fileBase
}

func (l *localFile) SetUserId(userId uint64) {
	l.UserId = userId
}

func (l *localFile) Put(file *multipart.FileHeader) (string, error) {
	// gen new hashed filename
	h := sha1.New()
	h.Write([]byte(file.Filename))
	fileName := hex.EncodeToString(h.Sum(nil)) + path.Ext(file.Filename)

	dst := strconv.FormatUint(l.UserId, 10)
	os.MkdirAll(filepath.Join(setting.ToolSetting.FileBaseDir, dst), os.ModePerm)
	dst += "/" + fileName

	// put file
	src, err := file.Open()
	if err != nil {
		return dst, err
	}
	defer src.Close()

	out, err := os.Create(filepath.Join(setting.ToolSetting.FileBaseDir, dst))
	if err != nil {
		return dst, err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return dst, err
}

func (l *localFile) Get(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

func (l *localFile) Delete(filePath string) error {
	panic("implement me")
}

