package faasinfra

import (
	"bytes"
	cConstants "github.com/bytedance/kldx-common/constants"
	cExceptions "github.com/bytedance/kldx-common/exceptions"
	cHttp "github.com/bytedance/kldx-common/http"
	"github.com/bytedance/kldx-infra/http"
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
		return nil, cExceptions.InternalError("request faaSInfra failed, code: %s, msg: %s", code, msg)
	} else {
		return nil, cExceptions.InvalidParamError("request faaSInfra failed, code: %s, msg: %s", code, msg)
	}
}

func doRequestMongodb(param interface{}) ([]byte, error) {
	return errorWrapper(http.GetFaaSInfraClient().PostBson(http.GetFaaSInfraPathMongodb(), nil, param, cHttp.AppTokenMiddleware, cHttp.TenantAndUserMiddleware, cHttp.ServiceIDMiddleware))
}

func DoRequestFile(contentType string, body *bytes.Buffer) ([]byte, error) {
	return errorWrapper(http.GetFaaSInfraClient().PostFormData(http.GetFaaSInfraPathFile(), map[string][]string{
		cConstants.HttpHeaderKey_ContentType: {contentType},
	}, body, cHttp.AppTokenMiddleware, cHttp.TenantAndUserMiddleware, cHttp.ServiceIDMiddleware))
}
