package file

import (
	cException "code.byted.org/apaas/goapi_common/exceptions"
	"io/ioutil"
)

func UploadWithContent(name string, content []byte, option *Option) (*UploadResult, error) {
	return uploadWithContent(name, content, option)
}

func UploadWithURL(name string, targetUrl string, option *Option) (*UploadResult, error) {
	data, err := readFromURL(targetUrl)
	if err != nil {
		return nil, cException.InvalidParamError("fetch data from targetUrl error: %v", err)
	}
	return uploadWithContent(name, data, option)
}

func UploadWithPath(name string, filePath string, option *Option) (*UploadResult, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, cException.InvalidParamError("read data from filePath error: %v", err)
	}
	return uploadWithContent(name, data, option)
}
