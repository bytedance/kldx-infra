package oss

import (
	cException "github.com/bytedance/kldx-common/exceptions"
	"io/ioutil"
)

type File struct {}

func NewFile() *File {
	return &File{}
}

func (f *File) UploadWithContent(name string, content []byte, option *Option) (*UploadResult, error) {
	return uploadWithContent(name, content, option)
}

func (f *File) UploadWithURL(name string, targetUrl string, option *Option) (*UploadResult, error) {
	data, err := readFromURL(targetUrl)
	if err != nil {
		return nil, cException.InvalidParamError("fetch data from targetUrl error: %v", err)
	}
	return uploadWithContent(name, data, option)
}

func (f *File) UploadWithPath(name string, filePath string, option *Option) (*UploadResult, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, cException.InvalidParamError("read data from filePath error: %v", err)
	}
	return uploadWithContent(name, data, option)
}
