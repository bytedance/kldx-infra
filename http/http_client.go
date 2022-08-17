package http

import (
	"net/http"
	"sync"
	"time"

	cHttp "github.com/bytedance/kldx-common/http"
	cUtils "github.com/bytedance/kldx-common/utils"
)

var (
	fiOnce   sync.Once
	fiClient *cHttp.HttpClient
)

func GetFaaSInfraClient() *cHttp.HttpClient {
	fiOnce.Do(func() {
		fiClient = &cHttp.HttpClient{
			Client: http.Client{
				Transport: &http.Transport{
					MaxIdleConns:        100,
					MaxIdleConnsPerHost: 10,
					IdleConnTimeout:     10 * time.Second,
				},
			},
			Url: cUtils.GetFaasinfraUrl(),
		}
	})
	return fiClient
}

var (
	httpClientOnce sync.Once
	httpClient     *http.Client
)

func GetCommonHttpClient() *http.Client {
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
