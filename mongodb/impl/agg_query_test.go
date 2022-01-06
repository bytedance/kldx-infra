package impl

import (
	"github.com/bytedance/kldx-infra/common/utils"
	cond "github.com/bytedance/kldx-infra/mongodb/condition"
	"testing"
)

// GroupBy
func TestQuery_GroupBy_Push(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")

	var results interface{}
	err := T.GroupBy("info.city", "城市").Push(map[string]string{"商品": "item", "数量": "qty"}, "列表").Find(&results)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(results)
}

func TestQuery_GroupBy_NoPush(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")

	var results interface{}
	err := T.GroupBy("info.city", "城市").Find(&results)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(results)
}

func TestQuery_GroupBy_NoAlias(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")

	var results interface{}
	err := T.GroupBy("info.city").Push(map[string]string{"商品": "item", "数量": "qty"}, "列表").Find(&results)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(results)
}

// GroupBy
func TestQuery_GroupBy_MulFields(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")

	var results interface{}
	err := T.GroupBy([]string{"item", "qty"}, "item-qty").Find(&results)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(results)
}

// GroupBy First
func TestQuery_GroupBy_First_SingleField(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")

	var results interface{}
	err := T.GroupBy("info.city", "city").First("qty", "count").Find(&results)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(results)
}

func TestQuery_GroupBy_First_MulFields(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")

	var results interface{}
	err := T.GroupBy("info.city", "city").First([]string{"item", "qty"}, "first-item").Find(&results)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(results)
}

// GroupBy Sum
func TestQuery_GroupBy_Sum(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")

	var results []interface{}
	err := T.GroupBy("info.city", "city").Sum("qty", "total").Find(&results)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(results)
}

// GroupBy Avg
func TestQuery_GroupByAvg(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")

	var results []interface{}
	err := T.GroupBy("info.city", "city").Avg("qty", "avg").Find(&results)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(results)
}

// GroupBy StdDevPop
func TestQuery_StdDevPop(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")

	var results []interface{}
	err := T.GroupBy("info.city", "city").StdDevPop("qty", "stdDevPop").Find(&results)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(results)
}

// GroupBy StdDevSamp
func TestQuery_GroupByStdDevSamp(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")

	var results []interface{}
	err := T.GroupBy("info.city", "city").StdDevSamp("qty", "stdDevSamp").Find(&results)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(results)
}

// GroupBy AddToSet
func TestQuery_GroupByAddToSet(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")

	var results []interface{}
	err := T.GroupBy("info.city", "city").AddToSet("item", "items").Find(&results)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(results)
}

func TestQuery_GroupBy_Having(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")

	var results interface{}
	err := T.GroupBy("info.city", "city").Push([]string{"item", "qty"}, "list").Having(cond.M{"qty": cond.Gt(40)}).Find(&results)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(results)
}
