package http

import (
	cConstants "code.byted.org/apaas/goapi_common/constants"
	cUtils "code.byted.org/apaas/goapi_common/utils"
	"net/http"
)

func FaasinfraMiddleware(req *http.Request) error {
	req.Header.Add(cConstants.HttpHeaderKey_Tenant, cUtils.GetTenant().Name)
	req.Header.Add(cConstants.HttpHeaderKey_User, "-1")
	req.Header.Add(cConstants.HttpHeaderKey_MicroserviceId, cUtils.GetMicroserviceId())
	return nil
}
