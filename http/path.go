package http

import (
	cConstants "code.byted.org/apaas/goapi_common/constants"
	cUtils "code.byted.org/apaas/goapi_common/utils"
	"strings"
)

const (
	FaaSInfraPathMongodb = "/resource/v1/namespaces/:namespace/db"
	FaaSInfraPathRedis   = "/resource/v1/namespaces/:namespace/cache"
	FaaSInfraPathFile  = "/resource/v1/namespaces/:namespace/file"
)

func GetFaaSInfraPathMongodb() string {
	return strings.ReplaceAll(FaaSInfraPathMongodb, cConstants.ReplaceNamespace, cUtils.GetTenant().Namespace)
}

func GetFaaSInfraPathRedis() string {
	return strings.ReplaceAll(FaaSInfraPathRedis, cConstants.ReplaceNamespace, cUtils.GetTenant().Namespace)
}

func GetFaaSInfraPathFile() string {
	return strings.ReplaceAll(FaaSInfraPathFile, cConstants.ReplaceNamespace, cUtils.GetTenant().Namespace)
}