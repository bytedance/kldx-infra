package file

import (
	cException "code.byted.org/apaas/goapi_common/exceptions"
	"code.byted.org/apaas/goapi_infra/http/faasinfra"
	"code.byted.org/apaas/goapi_infra/structs"
	"io/ioutil"
)

func UploadWithContent(name string, content []byte, option *structs.FileOption) (*structs.FileUploadResult, error) {
	return faasinfra.UploadWithContent(name, content, option)
}

func UploadWithURL(name string, targetUrl string, option *structs.FileOption) (*structs.FileUploadResult, error) {
	data, err := faasinfra.ReadFromURL(targetUrl)
	if err != nil {
		return nil, cException.InvalidParamError("fetch data from targetUrl error: %v", err)
	}
	return faasinfra.UploadWithContent(name, data, option)
}

func UploadWithPath(name string, filePath string, option *structs.FileOption) (*structs.FileUploadResult, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, cException.InvalidParamError("read data from filePath error: %v", err)
	}
	return faasinfra.UploadWithContent(name, data, option)
}