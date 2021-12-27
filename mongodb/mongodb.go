package mongodb

import (
	"code.byted.org/apaas/goapi_infra/structs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IMongodb interface {
	Table(tableName string) ITable
	NewObjectId() primitive.ObjectID
	ParseObjectId(_id string) (primitive.ObjectID, error)
}

// 表
type ITable interface {
	// 创建
	Create(record interface{}) (*structs.RecordOnlyId, error)
	BatchCreate(records interface{}) ([]string, error)

	// 条件查询、条件更新、条件删除
	Where(condition interface{}, args ...interface{}) IQuery

	// 聚合查询
	GroupBy(field interface{}, alias ...interface{}) IAggQuery
}

// 查询
type IQuery interface {
	// 更新
	Update(record interface{}) error
	Upsert(record interface{}) error
	BatchUpdate(record interface{}) error

	// 删除
	Delete() error
	BatchDelete() error

	// 查询
	Find(v interface{}) error
	FindOne(v interface{}) error

	Where(condition interface{}, args ...interface{}) IQuery
	Limit(limit int64) IQuery
	Offset(offset int64) IQuery
	OrderBy(fields ...string) IQuery
	OrderByDesc(fields ...string) IQuery
	Count() (int64, error)
	Distinct(field string, v interface{}) error
	Project(v interface{}) IQuery
}

// 聚合查询
type IAggQuery interface {
	Find(records interface{}) error
	FindOne(record interface{}) error

	// 分组
	GroupBy(field interface{}, alias ...interface{}) IAggQuery
	Having(condition interface{}) IAggQuery

	// 显示的字段
	Push(field interface{}, alias ...interface{}) IAggQuery
	// 分组中的第 1 个
	First(field interface{}, alias ...interface{}) IAggQuery
	// 分组中的最后 1 个
	Last(field interface{}, alias ...interface{}) IAggQuery
	// 分组中去重
	AddToSet(field string, alias ...interface{}) IAggQuery

	// 计算
	// 分组求和
	Sum(field string, args ...interface{}) IAggQuery
	// 分组求计数
	Num(field string) IAggQuery
	// 分组求平均
	Avg(field string, args ...interface{}) IAggQuery
	// 分组求总体标准差
	StdDevPop(field string, args ...interface{}) IAggQuery
	// 分组求样本标准差
	StdDevSamp(field string, args ...interface{}) IAggQuery
}
