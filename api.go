package infra

import (
	"code.byted.org/apaas/goapi_infra/file"
	mongodbImpl "code.byted.org/apaas/goapi_infra/mongodb/impl"
	"code.byted.org/apaas/goapi_infra/redis"
)

var (
	MongoDB = mongodbImpl.NewMongodb()
	Redis = redis.NewRedis()
	File = file.NewFile()
)
