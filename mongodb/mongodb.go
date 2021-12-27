package mongodb

import (
	"code.byted.org/apaas/goapi_infra/structs"
)

type IMongodb interface {
	Table(tableName string) ITable
}

// 表
type ITable interface {
	// 创建
	Create(record interface{}) (*structs.RecordOnlyId, error)
	BatchCreate(records interface{}) ([]string, error)

	// 更新、删除、查询都是通过条件过滤后操作
	Where(condition interface{}, args ...interface{}) IQuery
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
	Find(records interface{}) error
	FindOne(record interface{}) error

	Where(condition interface{}, args ...interface{}) IQuery
	// @Deprecated
	And(elems ...interface{}) IQuery
	// @Deprecated
	Or(elems ...interface{}) IQuery
	// @Deprecated
	Nor(elems ...interface{}) IQuery
	Limit(limit int64) IQuery
	Offset(offset int64) IQuery
	// @Deprecated
	Sort(v interface{}) IQuery
	OrderBy(fields ...string) IQuery
	OrderByDesc(fields ...string) IQuery
	Count() (int64, error)
	Distinct(field string, args ...interface{}) IQuery
	Project(v interface{}) IQuery

	// 聚合操作
	GroupBy(field string, args ...interface{}) IAggQuery
}

// 聚合查询
type IAggQuery interface {
	Find(records interface{}) error
	FindOne(record interface{}) error

	// 分组
	GroupBy(field string, args ...interface{}) IAggQuery
	Having(condition interface{}) IAggQuery

	// 分组中的第 1 个
	First(field string, args ...interface{}) IAggQuery
	// 分组中的最后 1 个
	Last(field string, args ...interface{}) IAggQuery

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
