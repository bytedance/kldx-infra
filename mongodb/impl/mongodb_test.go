package impl

import (
	"code.byted.org/apaas/goapi_infra/common/utils"
	cond "code.byted.org/apaas/goapi_infra/mongodb/condition"
	"testing"
)

func TestTable_Create(t *testing.T) {
	db := NewMongodb()
	T := db.Table("student")
	result, err := T.Create(map[string]interface{}{"name": "小刚", "age": 18})
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

func TestTable_BatchCreate(t *testing.T) {
	db := NewMongodb()
	T := db.Table("student")
	result, err := T.BatchCreate([]map[string]interface{}{
		{"name": "小明", "age": 19},
		{"name": "小花", "age": 20},
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

func TestQuery_Where_Case01(t *testing.T) {
	db := NewMongodb()
	T := db.Table("student")
	var result interface{}
	err := T.Where(
		cond.M{
			"name": cond.Eq("小明"),
		},
	).Find(&result)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

func TestQuery_Where_Case02(t *testing.T) {
	db := NewMongodb()
	T := db.Table("student")
	var result interface{}
	err := T.Where(
		cond.Or(
			cond.M{"name": cond.Eq("小明")},
			cond.M{"name": cond.Eq("小花")},
		),
	).Find(&result)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

func TestQuery_Where_Case03(t *testing.T) {
	db := NewMongodb()
	T := db.Table("student")
	var result interface{}
	err := T.Where(cond.M{"name": cond.Eq("小花")}).Where(cond.M{"age": cond.Eq(20)}).Find(&result)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

// Limit Offset
func TestQuery_LimitOffset(t *testing.T) {
	db := NewMongodb()
	T := db.Table("student")
	var result interface{}
	err := T.Where(nil).Offset(1).Limit(1).Find(&result)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

// Order
func TestQuery_OrderBy(t *testing.T) {
	db := NewMongodb()
	T := db.Table("student")
	var result interface{}
	err := T.Where(nil).OrderByDesc("age", "updatedAt").Find(&result)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}
