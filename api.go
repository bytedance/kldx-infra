package infra

import (
	"github.com/bytedance/kldx-infra/file"
	mongodbImpl "github.com/bytedance/kldx-infra/mongodb/impl"
	"github.com/bytedance/kldx-infra/redis"
)

var (
	MongoDB = mongodbImpl.NewMongodb()
	Redis = redis.NewRedis()
	File = file.NewFile()
)
