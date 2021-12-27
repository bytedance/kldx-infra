package faasinfra

import (
	cExceptions "code.byted.org/apaas/goapi_common/exceptions"
	cHttp "code.byted.org/apaas/goapi_common/http"
	"code.byted.org/apaas/goapi_infra/http"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
)

const (
	FaasinfraSuccessCode_Success = "0"

	FaasinfraFailCode_InternalError  = "k_ec_000001"
	FaasinfraFailCode_TokenExpire    = "k_ident_013000"
	FaasinfraFailCode_IllegalToken   = "k_ident_013001"
	FaasinfraFailCode_MissingToken   = "k_fs_ec_100001"
	FaasinfraFailCode_RateLimitError = "k_fs_ec_000004"
)

func errorWrapper(body []byte, err error) ([]byte, error) {
	if err != nil {
		return nil, cExceptions.ErrorWrap(err)
	}

	code := gjson.GetBytes(body, "code").String()
	msg := gjson.GetBytes(body, "msg").String()
	switch code {
	case FaasinfraSuccessCode_Success:
		data := gjson.GetBytes(body, "data")
		if data.Type == gjson.String {
			return []byte(data.Str), nil
		}
		return []byte(data.Raw), nil
	case FaasinfraFailCode_InternalError, FaasinfraFailCode_TokenExpire, FaasinfraFailCode_IllegalToken,
		FaasinfraFailCode_MissingToken, FaasinfraFailCode_RateLimitError:
		return nil, cExceptions.InternalError("request openapi failed, code: %s, msg: %s", code, msg)
	default:
		return nil, cExceptions.InvalidParamError("request openapi failed, code: %s, msg: %s", code, msg)
	}
}

func doRequestMongodb(param interface{}) ([]byte, error) {
	pStr, _ := json.Marshal(param)
	fmt.Println(string(pStr))

	return errorWrapper(http.GetFaasinfraClient().PostJson(http.GetFaasinfraPath_Mongodb(), nil, param, cHttp.AppTokenMiddleware, http.FaasinfraMiddleware))
}
