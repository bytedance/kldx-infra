package http

import (
	cConstants "github/kldx/common/constants"
	cUtils "github/kldx/common/utils"
	"net/http"
)

func FaaSInfraMiddleware(req *http.Request) error {
	req.Header.Add(cConstants.HttpHeaderKey_Tenant, cUtils.GetTenant().Name)
	req.Header.Add(cConstants.HttpHeaderKey_User, "-1")
	req.Header.Add(cConstants.HttpHeaderKey_MicroserviceId, cUtils.GetMicroserviceId())
	return nil
}
