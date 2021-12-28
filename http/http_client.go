package http

import (
	cHttp "code.byted.org/apaas/goapi_common/http"
	"code.byted.org/apaas/goapi_infra/common/utils"
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
