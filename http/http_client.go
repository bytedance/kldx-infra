package http

import (
	cHttp "github/kldx/common/http"
	"github/kldx/infra/common/utils"
	"net/http"
	"sync"
)

var (
	faasinfraClientOnce sync.Once
	faasinfraClient     *cHttp.HttpClient
)

func GetFaasinfraClient() *cHttp.HttpClient {
	conf := utils.GetFaasinfraClientConf()
	faasinfraClientOnce.Do(func() {
		faasinfraClient = &cHttp.HttpClient{
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
	return faasinfraClient
}
