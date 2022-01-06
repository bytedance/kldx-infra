package impl

import (
	"github.com/bytedance/kldx-infra/common/utils"
	"encoding/json"
	"testing"
)

func getInitData() interface{} {
	originData := `[{"_id":"61cc295251d84d3687cd405f","createdAt":"2021-12-29T17:24:34.677+08:00","info":{"city":"shanghai","tag":["hot"]},"item":"iphone 7","qty":150,"updatedAt":"2021-12-29T17:24:34.677+08:00"},{"_id":"61cc295251d84d3687cd4060","createdAt":"2021-12-29T17:24:34.677+08:00","info":{"city":"beijing","tag":["new"]},"item":"iphone X","qty":100,"updatedAt":"2021-12-29T17:24:34.677+08:00"},{"_id":"61cc295251d84d3687cd4061","createdAt":"2021-12-29T17:24:34.677+08:00","info":{"city":"beijing","tag":["new","hot"]},"item":"Mac Pro","qty":75,"updatedAt":"2021-12-29T17:24:34.677+08:00"},{"_id":"61cc295251d84d3687cd4062","createdAt":"2021-12-29T17:24:34.677+08:00","info":{"city":"beijing","tag":[]},"item":"Mac Air","qty":45,"updatedAt":"2021-12-29T17:24:34.677+08:00"},{"_id":"61cc295251d84d3687cd405e","createdAt":"2021-12-29T17:24:34.677+08:00","info":{"city":"shanghai","tag":[]},"item":"iphone 6","qty":35,"updatedAt":"2021-12-29T17:24:34.677+08:00"}]`
	var data interface{}
	err := json.Unmarshal([]byte(originData), &data)
	if err != nil {
		panic(err)
	}

	return data
}

func TestTable_BatchCreate_Goods(t *testing.T) {
	db := NewMongodb()
	T := db.Table("goods")
	result, err := T.BatchCreate(getInitData())
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

func TestTable_Create_Student(t *testing.T) {
	db := NewMongodb()
	T := db.Table("student")
	result, err := T.Create(map[string]interface{}{"name": "小刚", "age": 19})
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

func TestTable_Create_Employee(t *testing.T) {
	db := NewMongodb()
	T := db.Table("emp")
	result, err := T.Create(map[string]interface{}{"name": "小刚", "age": 19})
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}

func TestTable_BatchCreate_Employee(t *testing.T) {
	db := NewMongodb()
	T := db.Table("emp")
	result, err := T.BatchCreate([]*map[string]interface{}{
		{"name": "小花", "age": 20},
		{"name": "小明", "age": 18},
	})
	if err != nil {
		panic(err)
	}

	utils.PrintLog(result)
}
