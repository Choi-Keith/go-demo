package uploader

import (
	"demo01/util"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-resty/resty/v2"
)

var (
	local = NewLocal()
)

type uploader interface {
	PutImage(data []byte) (string, error)
	PutObject(key string, data []byte) (string, error)
	CopyImage(originUrl string) (string, error)
}

func PutImage(data []byte) (string, error) {
	return local.PutImage(data)
}

func PutObject(key string, data []byte) (string, error) {
	return local.PutObject(key, data)
}

func CopyImage(originUrl string) (string, error) {
	return local.CopyImage(originUrl)
}

// 本地文件系统
type localUploader struct{}

func NewLocal() *localUploader {
	return &localUploader{}
}

func generatetImageKey(data []byte) string {
	md5 := util.MD5Bytes(data)
	return filepath.Join("public", "image", md5+".jpg")
}

func (local *localUploader) PutImage(data []byte) (string, error) {
	key := generatetImageKey(data)
	return local.PutObject(key, data)
}

func (local *localUploader) PutObject(key string, data []byte) (string, error) {
	//if err := os.MkdirAll("/", os.ModeDir); err != nil {
	//	return "", err
	//}

	filename := filepath.Join(key)
	if err := os.MkdirAll(filepath.Dir(filename), os.ModePerm); err != nil {
		return "", err
	}
	if err := ioutil.WriteFile(filename, data, os.ModePerm); err != nil {
		return "", err
	}
	return util.UrlJoin(key), nil
}

func (local *localUploader) CopyImage(originUrl string) (string, error) {
	data, err := download(originUrl)
	if err != nil {
		return "", err
	}
	return local.PutImage(data)
}

func download(url string) ([]byte, error) {
	rsp, err := resty.New().R().Get(url)
	if err != nil {
		return nil, err
	}
	return rsp.Body(), nil
}
