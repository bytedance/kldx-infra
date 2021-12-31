package faasinfra

import (
	"bytes"
	cConstants "code.byted.org/apaas/goapi_common/constants"
	cExceptions "code.byted.org/apaas/goapi_common/exceptions"
	cHttp "code.byted.org/apaas/goapi_common/http"
	"code.byted.org/apaas/goapi_infra/http"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
)

func errorWrapper(body []byte, err error) ([]byte, error) {
	if err != nil {
		return nil, cExceptions.ErrorWrap(err)
	}

	code := gjson.GetBytes(body, "code").String()
	msg := gjson.GetBytes(body, "msg").String()
	if !http.HasError(code) {
		data := gjson.GetBytes(body, "data")
		if data.Type == gjson.String {
			return []byte(data.Str), nil
		}
		return []byte(data.Raw), nil
	} else if http.IsSysError(code) {
		return nil, cExceptions.InternalError("request openapi failed, code: %s, msg: %s", code, msg)
	} else {
		return nil, cExceptions.InvalidParamError("request openapi failed, code: %s, msg: %s", code, msg)
	}
}

func doRequestMongodb(param interface{}) ([]byte, error) {
	pStr, _ := json.Marshal(param)
	fmt.Println(string(pStr))

	return errorWrapper(http.GetFaaSInfraClient().PostJson(http.GetFaaSInfraPathMongodb(), nil, param, cHttp.AppTokenMiddleware, http.FaaSInfraMiddleware))
}

func doRequestFile(contentType string, body *bytes.Buffer) ([]byte, error) {
	return errorWrapper(http.GetFaaSInfraClient().PostFormData(http.GetFaaSInfraPathFile(), map[string][]string{
		cConstants.HttpHeaderKey_ContentType: {contentType},
	}, body, cHttp.AppTokenMiddleware, http.FaaSInfraMiddleware))
}
