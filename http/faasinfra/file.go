package faasinfra

import (
	"bytes"
	cException "github/kldx/common/exceptions"
	"github/kldx/infra/common/constants"
	"github/kldx/infra/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"sync"
)

var (
	httpClientOnce sync.Once
	httpClient     *http.Client
)

func getHttpClient() *http.Client {
	httpClientOnce.Do(func() {
		httpClient = &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 10,
			},
		}
	})
	return httpClient
}

func ReadFromURL(targetURL string) ([]byte, error) {

	u, err := url.Parse(targetURL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	rsp, err := getHttpClient().Do(req)
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

func UploadWithContent(name string, content []byte, option *structs.FileOption) (*structs.FileUploadResult, error) {
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

	out, err := doRequestFile(writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	res := &structs.FileUploadResult{}
	if err := json.Unmarshal(out, res); err != nil {
		return nil, err
	}
	return res, nil
}
