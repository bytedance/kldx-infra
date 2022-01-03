package file

import (
	"bytes"
	cException "code.byted.org/apaas/goapi_common/exceptions"
	"code.byted.org/apaas/goapi_infra/common/constants"
	http2 "code.byted.org/apaas/goapi_infra/http"
	"code.byted.org/apaas/goapi_infra/http/faasinfra"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
)

func readFromURL(targetURL string) ([]byte, error) {

	u, err := url.Parse(targetURL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	rsp, err := http2.GetCommonHttpClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("statusCode: %d", rsp.StatusCode)
	}
	return b, err
}

func uploadWithContent(name string, content []byte, option *Option) (*UploadResult, error) {
	if len(content) > constants.MaxFileSize {
		return nil, cException.InvalidParamError("file too large, exceed %v", constants.MaxFileSize)
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(constants.FileFieldName, name)
	if err != nil {
		return nil, err
	}

	if _, err := part.Write(content); err != nil {
		return nil, err
	}

	if option != nil {
		data, err := json.Marshal(option)
		if err != nil {
			return nil, err
		}

		if err := writer.WriteField(constants.FileOptionFieldName, string(data)); err != nil {
			return nil, err
		}

	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	out, err := faasinfra.DoRequestFile(writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	res := &fileUploadResult{}
	if err := json.Unmarshal(out, res); err != nil {
		return nil, err
	}
	if res.uploadError == nil {
		return &UploadResult{
			ID: res.ID,
			URL: res.URL,
		}, nil
	}
	return nil, res.uploadError.error()
}
