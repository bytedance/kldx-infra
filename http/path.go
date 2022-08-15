package http

import (
	"strings"

	cConstants "github.com/bytedance/kldx-common/constants"
	cUtils "github.com/bytedance/kldx-common/utils"
)

const (
	FaaSInfraPathMongodb = "/resource/v3/namespaces/:namespace/db"
	FaaSInfraPathRedis   = "/resource/v2/namespaces/:namespace/cache"
	FaaSInfraPathFile  = "/resource/v2/namespaces/:namespace/file"
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