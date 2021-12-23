package utils

import (
	"code.byted.org/apaas/goapi_infra/common/constants"
	"code.byted.org/apaas/goapi_infra/conf"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type httpConfig struct {
	Url                 string
	MaxIdleConns        int
	MaxIdleConnsPerHost int
	IdleConnTimeout     time.Duration
}

type Tenant struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Type      int64  `json:"type"`
}

func GetMicroserviceId() string {
	// TODO 暂时 mock，后面从 env 中获取
	return "ujpa81"
}

func GetFaasinfraAppidAndSecret() (string, string) {
	// TODO 暂时 mock，后面从 env 中获取
	return "c_fb79be28fae349ca90c0", "cd8fc6cb3c0a423d985e918e8019ec77"
}

func GetEnv() string {
	// TODO 暂时 mock，后面从 env 中获取
	return "development"
}

func GetBoeTag() string {
	// TODO 暂时 mock，后面从 env 中获取
	return constants.BoeTag
}

func GetInExtranetTag() string {
	// TODO 暂时 mock，后面从 env 中获取
	return constants.IntranetNetTag
}

func GetTenant() Tenant {
	// TODO 暂时 mock，后面从 env 中获取
	return Tenant{
		Id:        6187,
		Name:      "zwx_01",
		Namespace: "microService__c",
		Type:      1,
	}
}

func GetOpenapiUrl() string {
	key := fmt.Sprintf("%s:%s:%s", GetEnv(), GetBoeTag(), GetInExtranetTag())
	url, _ := conf.Env2url[key]
	return url
}

func GetFaasinfraClientConf() *httpConfig {
	return &httpConfig{
		Url:                 GetOpenapiUrl(),
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     10 * time.Second,
	}
}

// 单测用
func PrintLog(contents ...interface{}) {
	isPrint := false
	for _, content := range contents {
		if content == nil {
			fmt.Println(content)
			isPrint = true
			continue
		}

		typ := reflect.TypeOf(content)
		val := reflect.ValueOf(content)
		if typ.Kind() == reflect.Ptr {
			val = val.Elem()
			typ = typ.Elem()
		}

		switch typ.Kind() {
		case reflect.String:
			fmt.Println(content)
			isPrint = true
		default:
			content, err := json.Marshal(content)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(content))
			isPrint = true
		}
	}

	if isPrint {
		fmt.Println()
	}
}
