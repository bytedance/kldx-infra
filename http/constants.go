package http

import (
	"code.byted.org/apaas/goapi_infra/common/constants"
	"code.byted.org/apaas/goapi_infra/common/utils"
	"strings"
)

const (
	FaasinfraPath_GetToken = "/auth/v1/app/token"
	FaasinfraPath_Mongodb  = "/resource/v1/namespaces/:namespace/db"
)

func GetFaasinfraPath_Mongodb() string {
	return strings.ReplaceAll(FaasinfraPath_Mongodb, constants.ReplaceNamespace, utils.GetTenant().Namespace)
}