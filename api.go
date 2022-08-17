package infra

import (
	mongodbImpl "github.com/bytedance/kldx-infra/mongodb/impl"
	"github.com/bytedance/kldx-infra/redis"
)

var (
	MongoDB = mongodbImpl.NewMongodb()
	Redis   = redis.NewRedis()
)
