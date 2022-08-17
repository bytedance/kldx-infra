package impl

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"testing"
	"time"

	"github.com/bytedance/kldx-infra/common/utils"
	cond "github.com/bytedance/kldx-infra/mongodb/condition"
	"github.com/bytedance/kldx-infra/mongodb/structs"
)

type Goods struct {
	ID        primitive.ObjectID `bson:"_id"`
	Item      string             `bson:"item"`
	Qty       int64              `bson:"qty"`
	Info      *GoodsInfo         `bson:"info,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty"`

	Age       int64              `bson:"age,omitempty"`
}

type GoodsInfo struct {
	City string   `bson:"city,omitempty"`
	Tag  []string `bson:"tag,omitempty"`
}

var (
	ctx = context.Background()
)

func Init() {
	// TODO setEnv
	_ = os.Setenv("KTenantName", "zhouzilong")
	_ = os.Setenv("KNamespace", "test_case__c")
	_ = os.Setenv("KSvcID", "ttboes3cb")
	_ = os.Setenv("KClientID", "bb4bd82a33e14df854f610156b067e76f86be115131478cd37eadac6a68fe9ad")
	_ = os.Setenv("KClientSecret", "2a87d861e18735da7c7e0b2e3a0ff71a8134fc51efd37b2e16f17ce81c8bca9cbb2fde343a1ff26a0ffff4bbd41e1ba7")
	_ = os.Setenv("KOpenApiDomain", "http://boe-apaas-oapi-dev.byted.org")
	_ = os.Setenv("KFaaSInfraDomain", "http://apaas-faasinfra-dev.byted.org")
	//_ = os.Setenv("KFaaSInfraDomain", "http://10.94.93.221:6789")
	//_ = os.Setenv("KFaaSInfraDomain", "http://127.0.0.1:6789")
}

func Before() {
	goods := NewMongodb().Table("goods")
	student := NewMongodb().Table("student")
	emp := NewMongodb().Table("emp")

	var err error
	err = goods.Where(nilMap).Delete(ctx)
	err = student.Where(nilMap).Delete(ctx)
	err = emp.Where(nilMap).Delete(ctx)

	db := NewMongodb()
	T := db.Table("goods")
	create, err := T.BatchCreate(ctx, getInitData())
	if err != nil {
		panic(err)
	} else {
		fmt.Println(create)
	}

	T = db.Table("student")
	stuRes, err := T.Create(ctx, map[string]interface{}{"name": "小刚", "age": 19})
	if err != nil {
		panic(err)
	} else {
		fmt.Println(stuRes)
	}

	T = db.Table("emp")
	empRes, err := T.BatchCreate(ctx, []*map[string]interface{}{
		{"name": "小花", "age": 20},
		{"name": "小明", "age": 18},
	})
	if err != nil {
		panic(empRes)
	}

	utils.PrintLog(emp)
}

func TestMain(m *testing.M) {
	Init()
	Before()
	m.Run()
}

func TestQuery_Find_AllGoods(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")
	var result []Goods
	err := T.Where(nil).Find(ctx, &result)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(result)
}

func TestQuery_FindOne_OneGoods(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")
	var result Goods
	err := T.Where(nil).FindOne(ctx, &result)
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
			"_id": objID,
		},
	).Find(ctx, &result)
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
	).Find(ctx, &result)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

func TestQuery_Where_Employee_In(t *testing.T) {
	db := NewMongodb()
	T := db.Table("emp")
	var result interface{}
	err := T.Where(
		cond.M{
			"_id": cond.In([]string{"61d3f7b088e069bd971f5552", "61d3f7b5ccc793268ce1da72", "61d3f7b5ccc793268ce1da73"}),
		},
	).Find(ctx, &result)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

func TestQuery_Where_Employee(t *testing.T) {
	db := NewMongodb()
	T := db.Table("emp")
	var result interface{}
	err := T.Where(nil).Find(ctx, &result)
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
	).Find(ctx, &result)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

func TestQuery_Where_Where(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")
	var result []Goods
	err := T.Where(cond.M{"item": cond.Eq("iphone 7")}).Where(cond.M{"info.city": cond.Eq("shanghai")}).Find(ctx, &result)
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
	err := T.Where(nil).Offset(1).Limit(1).Find(ctx, &result)
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
	err := T.Where(nil).OrderByDesc("qty").OrderBy("item").Find(ctx, &result)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

func TestQuery_Count(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")
	count, err := T.Where(nil).Count(ctx)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(count)
}

func TestQuery_Where_Gte_Count(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")
	count, err := T.Where(cond.M{"qty": cond.Gte(100)}).Count(ctx)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(count)
}

// Distinct
//func TestQuery_Distinct(t *testing.T) {
//	db := NewMongodb()
//	T := db.Table("goods")
//
//	var cities []string
//	err := T.Where(nil).Distinct(ctx, "info.city", &cities)
//	if err != nil {
//		panic(err)
//	}
//	utils.PrintLog(cities)
//}

func TestQuery_Project(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")

	var results []*Goods
	err := T.Where(nil).Project(cond.M{"createdAt": 0, "updatedAt": 0, "info": 0}).Find(ctx, &results)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(results)
}

// Update
func TestQuery_Update(t *testing.T) {
	db := NewMongodb()

	T := db.Table("student")
	err := T.Where(cond.M{"_id": cond.Eq(objID)}).Update(ctx, cond.M{"age": "22"})
	if err != nil {
		panic(err)
	}

	var result interface{}
	err = T.Where(cond.M{"_id": cond.Eq(objID1)}).FindOne(ctx, &result)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

// Upsert
func TestQuery_Upsert(t *testing.T) {
	db := NewMongodb()
	T := db.Table("student")
	err := T.Where(cond.M{"_id": "61c99b7a96414a5793012868"}).Upsert(ctx, cond.M{"age": "18"})
	if err != nil {
		panic(err)
	}

	var result interface{}
	err = T.Where(cond.M{"_id": "61cac55683420b07931a0190"}).FindOne(ctx, &result)
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

// BatchUpdate
func TestQuery_BatchUpdate(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")

	var results []*structs.RecordOnlyId
	err := T.Where(nil).Find(ctx, &results)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(results)

	var ids []interface{}
	for _, r := range results {
		if r.ID.Hex() == objID.Hex() || r.ID.Hex() == objID1.Hex() {
			continue
		}
		ids = append(ids, r.ID)
	}

	var list []bson.M
	err = T.Where(cond.M{"_id": cond.In(ids)}).Find(ctx, &list)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(list)

	err = T.Where(cond.M{"_id": cond.In(ids)}).BatchDelete(ctx)
	if err != nil {
		panic(err)
	}

	err = T.Where(cond.M{"_id": cond.In([]interface{}{objID, objID1})}).BatchUpdate(ctx, cond.M{"qty": 66})
	if err != nil {
		panic(err)
	}

	var result1 []bson.M
	err = T.Where(cond.M{"_id": objID}).Find(ctx, &result1)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(result1)

	var result2 []bson.M
	err = T.Where(cond.M{"_id": objID1}).Find(ctx, &result2)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(result2)

	err = T.Where(cond.M{}).Find(ctx, &list)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(list)
}

// Delete
func TestQuery_Delete(t *testing.T) {
	db := NewMongodb()
	T := db.Table("student")
	result, err := T.Create(ctx, map[string]interface{}{"name": "小刚", "age": 18})
	if err != nil {
		panic(err)
	}
	utils.PrintLog(result)

	err = T.Where(cond.M{"_id": result.ID}).Delete(ctx)
	if err != nil {
		panic(err)
	}
}

// Delete
func TestQuery_BatchDelete(t *testing.T) {
	db := NewMongodb()
	T := db.Table("student")
	results, err := T.BatchCreate(ctx, []map[string]interface{}{
		{"name": "小明", "age": 19},
		{"name": "小花", "age": 20},
	})
	if err != nil {
		panic(err)
	}
	utils.PrintLog(results)

	err = T.Where(cond.M{"_id": cond.In(results)}).BatchDelete(ctx)
	if err != nil {
		panic(err)
	}
}

func TestTable_BatchDelete_Goods(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")

	var res1 interface{}
	err := T.Where(nil).Find(ctx, &res1)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(res1)

	var res2 interface{}
	err = T.Where(cond.M{"qty": cond.Gt(0)}).Find(ctx, &res2)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(res2)
}
