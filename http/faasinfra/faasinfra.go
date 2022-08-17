package faasinfra

import (
	"context"
	"encoding/base64"

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

func doRequestMongodb(ctx context.Context, param interface{}) ([]byte, error) {
	data, err := errorWrapper(http.GetFaaSInfraClient().PostBson(ctx, http.GetFaaSInfraPathMongodb(), nil, param, cHttp.AppTokenMiddleware, cHttp.TenantAndUserMiddleware, cHttp.ServiceIDMiddleware))
	if err != nil {
		return data, err
	}
	return base64.StdEncoding.DecodeString(string(data))
}
