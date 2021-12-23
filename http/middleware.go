package http

import (
	"code.byted.org/apaas/goapi_infra/common/constants"
	"code.byted.org/apaas/goapi_infra/common/exceptions"
	"code.byted.org/apaas/goapi_infra/common/structs"
	"code.byted.org/apaas/goapi_infra/common/utils"
	"encoding/json"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type ReqMiddleWare func(req *http.Request) error

func AppTokenMiddleware(req *http.Request) error {
	if req == nil || req.Header == nil {
		return nil
	}
	token, err := GetAppToken()
	if err != nil {
		return err
	}
	req.Header.Add(constants.HttpHeaderKey_Authorization, token)
	req.Header.Add(constants.HttpHeaderKey_Tenant, utils.GetTenant().Name)
	req.Header.Add(constants.HttpHeaderKey_User, "-1")
	req.Header.Add(constants.HttpHeaderKey_MicroserviceId, utils.GetMicroserviceId())
	return nil
}

var (
	openapiToken           atomic.Value
	openapiTokenExpireTime atomic.Value
	tokenRemainingTime     int64 = 600 // 10min
)

func getAppTokenFromMem() string {
	expireTime, ok := openapiTokenExpireTime.Load().(int64)
	if !ok {
		return ""
	}

	token, ok := openapiToken.Load().(string)
	if !ok {
		return ""
	}

	// token 为空 或 10分钟内过期，不再使用
	if expireTime-time.Now().Unix() < tokenRemainingTime || token == "" {
		return ""
	}

	return token
}

func GetAppToken() (string, error) {
	// 1.取内存值
	token := getAppTokenFromMem()
	if token != "" {
		return token, nil
	}

	// 2.取远程值
	token, err := refreshAppToken()
	if err != nil {
		return "", err
	}
	return token, nil
}

func refreshAppToken() (string, error) {
	// 1.获取锁
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()

	// 2.锁获取成功后，确认 token 是否已经被刷新
	token := getAppTokenFromMem()
	if token != "" {
		return token, nil
	}

	// 3.需要刷新 token
	appid, secret := utils.GetFaasinfraAppidAndSecret()
	data := map[string]string{
		"app_id":     appid,
		"app_secret": secret,
	}

	body, err := GetFaasinfraClient().PostJson(FaasinfraPath_GetToken, nil, data)
	if err != nil {
		return "", err
	}

	tokenResult := structs.TokenResult{}
	err = json.Unmarshal(body, &tokenResult)
	if err != nil {
		return "", exceptions.InternalError("unmarshal OpenapiTokenResult failed, err: %v", err)
	}

	if tokenResult.Data.AccessToken == "" {
		return "", exceptions.InternalError("openapi accessToken is empty")
	}

	openapiToken.Store(tokenResult.Data.AccessToken)
	openapiTokenExpireTime.Store(tokenResult.Data.ExpireTime)
	return tokenResult.Data.AccessToken, nil
}
