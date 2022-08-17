package impl

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"

	"github.com/bytedance/kldx-infra/common/utils"
)

var (
	nilMap = bson.M{}
	objID  primitive.ObjectID
	objID1 primitive.ObjectID
)

func getInitData() interface{} {
	objID, _ = primitive.ObjectIDFromHex("61cc295251d84d3687cd405f")
	objID1, _ = primitive.ObjectIDFromHex("61cc295251d84d3687cd40601")
	return []*Goods{{
		ID:   objID,
		Item: "iphone 7",
		Qty:  150,
		Info: &GoodsInfo{
			City: "shanghai",
			Tag:  []string{"hot"},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, {
		ID:   objID1,
		Item: "iphone X",
		Qty:  100,
		Info: &GoodsInfo{
			City: "beijing",
			Tag:  []string{"new"},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, {
		ID:   primitive.NewObjectID(),
		Item: "Mac Pro",
		Qty:  75,
		Info: &GoodsInfo{
			City: "shanghai",
			Tag:  []string{"new", "hot"},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, {
		ID:   primitive.NewObjectID(),
		Item: "Mac Air",
		Qty:  45,
		Info: &GoodsInfo{
			City: "beijing",
			Tag:  []string{},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, {
		ID:   primitive.NewObjectID(),
		Item: "iphone 6",
		Qty:  35,
		Info: &GoodsInfo{
			City: "shanghai",
			Tag:  []string{},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}}
}

func TestTable_BatchCreate_Goods(t *testing.T) {
	//db := NewMongodb()
	//T := db.Table("goods")
	//result, err := T.BatchCreate(ctx, getInitData())
	//if err != nil {
	//	panic(err)
	//}
	//
	//utils.PrintLog(result)
}

func TestTable_Create_Student(t *testing.T) {
	db := NewMongodb()
	T := db.Table("student")
	fmt.Println(T.Where(bson.M{"name": "小刚"}).Delete(ctx))
	result, err := T.Create(ctx, map[string]interface{}{"name": "小刚", "age": 19})
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

func TestTable_Create_Employee(t *testing.T) {
	db := NewMongodb()
	T := db.Table("emp")
	result, err := T.Create(ctx, map[string]interface{}{"name": "小刚", "age": 19})
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

func TestTable_BatchCreate_Employee(t *testing.T) {
	db := NewMongodb()
	T := db.Table("emp")
	result, err := T.BatchCreate(ctx, []*map[string]interface{}{
		{"name": "小花", "age": 20},
		{"name": "小明", "age": 18},
	})
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}