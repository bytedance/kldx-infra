package impl

import (
	"github/kldx/infra/common/utils"
	cond "github/kldx/infra/mongodb/condition"
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

// Count
func TestQuery_Count(t *testing.T) {
	db := NewMongodb()
	T := db.Table("student")
	count, err := T.Where(nil).Limit(100).Count()
	if err != nil {
		panic(err)
	}

	utils.PrintLog(count)
}

// Update
func TestQuery_Update(t *testing.T) {
	db := NewMongodb()
	T := db.Table("student")
	err := T.Where(cond.M{"_id": "61c99b7a96414a5793012868"}).Update(cond.M{"age": "22"})
	if err != nil {
		panic(err)
	}

	var result interface{}
	err = T.Where(cond.M{"_id": "61c99b7a96414a5793012868"}).FindOne(&result)
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

	var results interface{}
	err := T.Where(nil).Find(&results)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(results)

	err = T.Where(cond.M{"_id": cond.In([]string{"61c99b7a96414a5793012868", "61cac55683420b07931a0190"})}).BatchUpdate(cond.M{"age": "30"})
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
	err = T.Where(cond.M{"_id": "61cac9253431e5bda4a43ce3"}).Find(&result2)
	if err != nil {
		panic(err)
	}
	utils.PrintLog(result2)

	// [{"_id":"61c92930a8125721fb44f257","age":18,"createdAt":"2021-12-27T10:47:12.771+08:00","name":"小刚","updatedAt":"2021-12-27T10:47:12.771+08:00"},{"_id":"61c9297ca8125721fb44f258","age":19,"createdAt":"2021-12-27T10:48:28.96+08:00","name":"小明","updatedAt":"2021-12-27T10:48:28.96+08:00"},{"_id":"61c9297ca8125721fb44f259","age":20,"createdAt":"2021-12-27T10:48:28.96+08:00","name":"小花","updatedAt":"2021-12-27T10:48:28.96+08:00"},{"_id":"61c92ade1737b258ce2e7a2d","age":20,"createdAt":"2021-12-27T10:54:22.899+08:00","name":"小花","updatedAt":"2021-12-27T10:54:22.899+08:00"},{"_id":"61c92ade1737b258ce2e7a2c","age":19,"createdAt":"2021-12-27T10:54:22.899+08:00","name":"小明","updatedAt":"2021-12-27T10:54:22.899+08:00"},{"_id":"61c99b7a83c498de0ce5c23b","age":18,"createdAt":"2021-12-27T18:54:50.578+08:00","name":"小刚","updatedAt":"2021-12-27T18:54:50.578+08:00"},{"_id":"61c99b7a96414a5793012867","age":19,"createdAt":"2021-12-27T18:54:50.997+08:00","name":"小明","updatedAt":"2021-12-27T18:54:50.997+08:00"},{"_id":"61c99b7a96414a5793012868","age":20,"createdAt":"2021-12-27T18:54:50.997+08:00","name":"小花","updatedAt":"2021-12-27T18:54:50.997+08:00"},{"_id":"61c9bc1d8d932bb8fd2cdbdd","age":18,"createdAt":"2021-12-27T21:14:05.496+08:00","name":"小刚","updatedAt":"2021-12-27T21:14:05.496+08:00"},{"_id":"61c9bc428d932bb8fd2cdbde","age":18,"createdAt":"2021-12-27T21:14:42.122+08:00","name":"小刚","updatedAt":"2021-12-27T21:14:42.122+08:00"},{"_id":"61c9bcdf8d932bb8fd2cdbdf","age":19,"createdAt":"2021-12-27T21:17:19.259+08:00","name":"小明","updatedAt":"2021-12-27T21:17:19.259+08:00"},{"_id":"61c9bcdf8d932bb8fd2cdbe0","age":20,"createdAt":"2021-12-27T21:17:19.259+08:00","name":"小花","updatedAt":"2021-12-27T21:17:19.259+08:00"},{"_id":"61cac550ecd08e029bfa1906","age":18,"createdAt":"2021-12-28T16:05:36.052+08:00","name":"小刚","updatedAt":"2021-12-28T16:05:36.052+08:00"},{"_id":"61cac55083420b07931a018e","age":20,"createdAt":"2021-12-28T16:05:36.496+08:00","name":"小花","updatedAt":"2021-12-28T16:05:36.496+08:00"},{"_id":"61cac55083420b07931a018d","age":19,"createdAt":"2021-12-28T16:05:36.496+08:00","name":"小明","updatedAt":"2021-12-28T16:05:36.496+08:00"},{"_id":"61cac555ecd08e029bfa1907","age":18,"createdAt":"2021-12-28T16:05:41.513+08:00","name":"小刚","updatedAt":"2021-12-28T16:05:41.513+08:00"},{"_id":"61cac55683420b07931a018f","age":19,"createdAt":"2021-12-28T16:05:42.182+08:00","name":"小明","updatedAt":"2021-12-28T16:05:42.182+08:00"},{"_id":"61cac55683420b07931a0190","age":20,"createdAt":"2021-12-28T16:05:42.182+08:00","name":"小花","updatedAt":"2021-12-28T16:05:42.182+08:00"},{"_id":"61c99b7a96414a5793012868","age":"30","updatedAt":"2021-12-28T16:15:21.804+08:00"}]

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
