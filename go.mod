module github.com/bytedance/kldx-infra

go 1.16

require (
	github.com/bytedance/kldx-common v0.0.5-0.20220815133320-8099d99cd3f5
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/tidwall/gjson v1.14.0
	go.mongodb.org/mongo-driver v1.8.3
)

replace go.mongodb.org/mongo-driver v1.8.3 => go.mongodb.org/mongo-driver v1.4.1
