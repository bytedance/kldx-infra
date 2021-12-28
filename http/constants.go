package http

import (
	cConstants "code.byted.org/apaas/goapi_common/constants"
	cUtils "code.byted.org/apaas/goapi_common/utils"
	"strings"
)

const (
	FaasinfraPath_Mongodb  = "/resource/v1/namespaces/:namespace/db"
	FaasinfraPath_File  = "/resource/v1/namespaces/:namespace/file"
)

func GetFaasinfraPath_Mongodb() string {
	return strings.ReplaceAll(FaasinfraPath_Mongodb, cConstants.ReplaceNamespace, cUtils.GetTenant().Namespace)
}

func GetFaasinfraPath_File() string {
	return strings.ReplaceAll(FaasinfraPath_File, cConstants.ReplaceNamespace, cUtils.GetTenant().Namespace)
}