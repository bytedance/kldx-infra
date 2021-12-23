package impl

import (
	"code.byted.org/apaas/goapi_infra/common/utils"
	"testing"
)

func TestTable_Create(t *testing.T) {
	db := NewMongodb()
	T := db.Table("student")
	result, err := T.Create(map[string]interface{}{"name": "小刚", "age":18})
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

func TestTable_BatchCreate(t *testing.T) {
	db := NewMongodb()
	T := db.Table("student")
	result, err := T.BatchCreate([]map[string]interface{}{
		{"name": "小明", "age":19},
		{"name": "小花", "age":20},
	})
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

func TestQuery_Find(t *testing.T) {
	db := NewMongodb()
	T := db.Table("student")
	var result interface{}
	err := T.Where(nil).Find(&result)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

func TestQuery_FindOne(t *testing.T) {
	db := NewMongodb()
	T := db.Table("student")
	var result interface{}
	err := T.Where(nil).FindOne(&result)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}
