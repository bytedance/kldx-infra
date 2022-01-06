package impl

import (
	"code.byted.org/apaas/goapi_infra/common/utils"
	cond "code.byted.org/apaas/goapi_infra/mongodb/condition"
	"code.byted.org/apaas/goapi_infra/mongodb/structs"
	"testing"
)

type Goods struct {
	ID        string     `json:"_id"`
	Item      string     `json:"item"`
	Qty       int64      `json:"qty"`
	Info      *GoodsInfo `json:"info,omitempty"`
	CreatedAt string     `json:"createdAt,omitempty"`
	UpdatedAt string     `json:"updatedAt,omitempty"`
}

type GoodsInfo struct {
	City string   `json:"city,omitempty"`
	Tag  []string `json:"tag,omitempty"`
}

func TestQuery_Find_AllGoods(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")
	var result []Goods
	err := T.Where(nil).Find(&result)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

func TestQuery_FindOne_OneGoods(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")
	var result Goods
	err := T.Where(nil).FindOne(&result)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

func TestQuery_Where_Eq(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")
	var result []*Goods
	err := T.Where(
		cond.M{
			"_id": "61cd58f0791435c7bf31453b",
		},
	).Find(&result)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

func TestQuery_Where_In(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")
	var result []Goods
	err := T.Where(
		cond.M{
			"info.city": cond.In([]string{"beijing", "shanghai"}),
		},
	).Find(&result)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

func TestQuery_Where_Eployee_In(t *testing.T) {
	db := NewMongodb()
	T := db.Table("emp")
	var result interface{}
	err := T.Where(
		cond.M{
			"_id": cond.In([]string{"61d3f7b088e069bd971f5552", "61d3f7b5ccc793268ce1da72", "61d3f7b5ccc793268ce1da73"}),
		},
	).Find(&result)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

func TestQuery_Where_Eployee(t *testing.T) {
	db := NewMongodb()
	T := db.Table("emp")
	var result interface{}
	err := T.Where(nil).Find(&result)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

func TestQuery_Where_Or(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")
	var result []Goods
	err := T.Where(
		cond.Or(
			cond.M{"item": cond.Eq("iphone 7")},
			cond.M{"item": cond.Eq("iphone 6")},
		),
	).Find(&result)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

func TestQuery_Where_Where(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")
	var result []Goods
	err := T.Where(cond.M{"item": cond.Eq("iphone 7")}).Where(cond.M{"info.city": cond.Eq("shanghai")}).Find(&result)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

// Limit Offset
func TestQuery_LimitOffset(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")
	var result []Goods
	err := T.Where(nil).Offset(1).Limit(1).Find(&result)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

// TODO 无法表达排序字段的优先级
func TestQuery_OrderBy(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")
	var result interface{}
	err := T.Where(nil).OrderByDesc("qty").OrderBy("item").Find(&result)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

func TestQuery_Count(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")
	count, err := T.Where(nil).Count()
	if err != nil {
		panic(err)
	}

	utils.PrintLog(count)
}

func TestQuery_Where_Gte_Count(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")
	count, err := T.Where(cond.M{"qty": cond.Gte(100)}).Count()
	if err != nil {
		panic(err)
	}

	utils.PrintLog(count)
}

// Distinct
func TestQuery_Distinct(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")

	var cities []string
	err := T.Where(nil).Distinct("info.city", &cities)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(cities)
}

func TestQuery_Project(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")

	var results []*Goods
	err := T.Where(nil).Project(cond.M{"createdAt": 0, "updatedAt": 0, "info": 0}).Find(&results)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(results)
}

// Update
func TestQuery_Update(t *testing.T) {
	db := NewMongodb()

	T := db.Table("student")
	err := T.Where(cond.M{"_id": cond.Eq("61c99b7a96414a5793012868")}).Update(cond.M{"age": "22"})
	if err != nil {
		panic(err)
	}

	var result interface{}
	err = T.Where(cond.M{"_id": cond.Eq("61c99b7a96414a5793012868")}).FindOne(&result)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

// Upsert
func TestQuery_Upsert(t *testing.T) {
	db := NewMongodb()
	T := db.Table("student")
	err := T.Where(cond.M{"_id": "61c99b7a96414a5793012868"}).Upsert(cond.M{"age": "18"})
	if err != nil {
		panic(err)
	}

	var result interface{}
	err = T.Where(cond.M{"_id": "61cac55683420b07931a0190"}).FindOne(&result)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

// BatchUpdate
func TestQuery_BatchUpdate(t *testing.T) {
	db := NewMongodb()
	T := db.Table("student")

	var results []*structs.RecordOnlyId
	err := T.Where(nil).Find(&results)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(results)

	var ids []interface{}
	for _, r := range results {
		ids = append(ids, r.Id)
	}

	err = T.Where(cond.M{"_id": cond.In(ids)}).Find(&results)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(results)

	err = T.Where(cond.M{"_id": cond.In(ids)}).BatchDelete()
	if err != nil {
		panic(err)
	}

	err = T.Where(cond.M{"_id": cond.In([]interface{}{"61c99b7a96414a5793012868", "61c92930a8125721fb44f257"})}).BatchUpdate(cond.M{"age": "32"})
	if err != nil {
		panic(err)
	}

	var result1 interface{}
	err = T.Where(cond.M{"_id": "61c99b7a96414a5793012868"}).Find(&result1)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(result1)

	var result2 interface{}
	err = T.Where(cond.M{"_id": "61c92930a8125721fb44f257"}).Find(&result2)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(result2)
}

// Delete
func TestQuery_Delete(t *testing.T) {
	db := NewMongodb()
	T := db.Table("student")
	result, err := T.Create(map[string]interface{}{"name": "小刚", "age": 18})
	if err != nil {
		panic(err)
	}
	utils.PrintLog(result)

	err = T.Where(cond.M{"_id": result.Id}).Delete()
	if err != nil {
		panic(err)
	}
}

// Delete
func TestQuery_BatchDelete(t *testing.T) {
	db := NewMongodb()
	T := db.Table("student")
	results, err := T.BatchCreate([]map[string]interface{}{
		{"name": "小明", "age": 19},
		{"name": "小花", "age": 20},
	})
	if err != nil {
		panic(err)
	}
	utils.PrintLog(results)

	err = T.Where(cond.M{"_id": cond.In(results)}).BatchDelete()
	if err != nil {
		panic(err)
	}
}

func TestTable_BatchDelete_Goods(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")

	var res1 interface{}
	err := T.Where(nil).Find(&res1)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(res1)

	var res2 interface{}
	err = T.Where(cond.M{"qty": cond.Gt(0)}).Find(&res2)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(res2)
}
