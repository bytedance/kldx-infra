package infra

import (
	"github.com/bytedance/kldx-infra/common/utils"
	cond "github.com/bytedance/kldx-infra/mongodb/condition"
	"testing"
)

func TestMongodb(t *testing.T) {
	var result interface{}
	err := MongoDB.Table("goods").Where(cond.M{
		"qty": cond.Gt(0),
	}).Find(&result)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(result)
}

func TestMongodb_Delete(t *testing.T) {
	err := MongoDB.Table("goods").Where(cond.M{
		"qty": cond.Gt(0),
	}).BatchDelete()
	if err != nil {
		panic(err)
	}
}
