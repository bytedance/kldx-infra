package http

import (
	"bytes"
	"code.byted.org/apaas/goapi_infra/common/constants"
	"code.byted.org/apaas/goapi_infra/common/exceptions"
	"code.byted.org/apaas/goapi_infra/common/utils"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type HttpClient struct {
	http.Client
	url string
}

var (
	// 内网版 管理接口 Faasinfra Client
	faasinfraClientOnce sync.Once
	faasinfraClient     *HttpClient
)

func GetFaasinfraClient() *HttpClient {
	conf := utils.GetFaasinfraClientConf()
	faasinfraClientOnce.Do(func() {
		faasinfraClient = &HttpClient{
			Client: http.Client{
				Transport: &http.Transport{
					MaxIdleConns:        conf.MaxIdleConns,
					MaxIdleConnsPerHost: conf.MaxIdleConnsPerHost,
					IdleConnTimeout:     conf.IdleConnTimeout,
				},
			},
			url: conf.Url,
		}
	})
	return faasinfraClient
}

func (c *HttpClient) doRequest(req *http.Request, headers map[string][]string, mids []ReqMiddleWare) ([]byte, error) {
	for _, mid := range mids {
		err := mid(req)
		if err != nil {
			return nil, err
		}
	}

	for key, values := range headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// 配置超时时间
	ctx, cancel := getTimeoutCtx()
	defer cancel()

	resp, err := c.Do(req.WithContext(ctx))
	if err != nil {
		return nil, exceptions.InternalError("doRequest failed, err: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, exceptions.InternalError("doRequest failed: statusCode is %d", resp.StatusCode)
	}

	// http resp body
	datas, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, exceptions.InternalError("doRequest failed, err: %v", err)
	}

	return datas, nil
}

func (c *HttpClient) Get(path string, headers map[string][]string, mids ...ReqMiddleWare) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, c.url+path, nil)
	if err != nil {
		return nil, exceptions.InternalError("HttpClient.Get failed, err: %v", err)
	}

	return c.doRequest(req, headers, mids)
}

func (c *HttpClient) PostJson(path string, headers map[string][]string, data interface{}, mids ...ReqMiddleWare) ([]byte, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, exceptions.InternalError("HttpClient.PostJson failed, err: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, c.url+path, bytes.NewReader(body))
	if err != nil {
		return nil, exceptions.InternalError("HttpClient.PostJson failed, err: %v", err)
	}

	if headers == nil {
		headers = map[string][]string{}
	}
	headers[constants.HttpHeaderKey_ContentType] = []string{constants.HttpHeaderValue_Json}
	return c.doRequest(req, headers, mids)
}

func (c *HttpClient) PostFormData(path string, headers map[string][]string, body *bytes.Buffer, mids ...ReqMiddleWare) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, c.url+path, body)
	if err != nil {
		return nil, exceptions.InternalError("HttpClient.PostFormData failed, err: %v", err)
	}
	return c.doRequest(req, headers, mids)
}

func (c *HttpClient) PatchJson(path string, headers map[string][]string, data interface{}, mids ...ReqMiddleWare) ([]byte, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, exceptions.InternalError("HttpClient.PatchJson failed, err: %v", err)
	}

	req, err := http.NewRequest(http.MethodPatch, c.url+path, bytes.NewReader(body))
	if err != nil {
		return nil, exceptions.InternalError("HttpClient.PatchJson failed, err: %v", err)
	}

	if headers == nil {
		headers = map[string][]string{}
	}
	headers[constants.HttpHeaderKey_ContentType] = []string{constants.HttpHeaderValue_Json}
	return c.doRequest(req, headers, mids)
}

func (c *HttpClient) DeleteJson(path string, headers map[string][]string, data interface{}, mids ...ReqMiddleWare) ([]byte, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, exceptions.InternalError("HttpClient.DeleteJson failed, err: %v", err)
	}

	req, err := http.NewRequest(http.MethodDelete, c.url+path, bytes.NewReader(body))
	if err != nil {
		return nil, exceptions.InternalError("HttpClient.DeleteJson failed, err: %v", err)
	}

	if headers == nil {
		headers = map[string][]string{}
	}
	headers[constants.HttpHeaderKey_ContentType] = []string{constants.HttpHeaderValue_Json}
	return c.doRequest(req, headers, mids)
}

func getTimeoutCtx() (context.Context, context.CancelFunc) {
	// 暂时定死为 10s
	return context.WithTimeout(context.Background(), 10 * time.Second)
}
