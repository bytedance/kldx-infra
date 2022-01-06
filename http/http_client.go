package http

import (
	cHttp "github/kldx/common/http"
	"github/kldx/infra/common/utils"
	"net/http"
	"sync"
)

var (
	fiOnce   sync.Once
	fiClient *cHttp.HttpClient
)

func GetFaaSInfraClient() *cHttp.HttpClient {
	conf := utils.GetFaasinfraClientConf()
	fiOnce.Do(func() {
		fiClient = &cHttp.HttpClient{
			Client: http.Client{
				Transport: &http.Transport{
					MaxIdleConns:        conf.MaxIdleConns,
					MaxIdleConnsPerHost: conf.MaxIdleConnsPerHost,
					IdleConnTimeout:     conf.IdleConnTimeout,
				},
			},
			Url: conf.Url,
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

