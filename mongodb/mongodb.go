package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/bytedance/kldx-infra/mongodb/structs"
)

type IMongodb interface {
	Table(tableName string) ITable
}

// 表
type ITable interface {
	// 创建
	Create(ctx context.Context, record interface{}) (*structs.RecordOnlyId, error)
	BatchCreate(ctx context.Context, records interface{}) ([]primitive.ObjectID, error)

	// 条件查询、条件更新、条件删除
	Where(condition interface{}, args ...interface{}) IQuery

	// 聚合查询
	GroupBy(field interface{}, alias ...interface{}) IAggQuery
}

// 查询
type IQuery interface {
	// 更新
	Update(ctx context.Context, record interface{}) error
	Upsert(ctx context.Context, record interface{}) error
	BatchUpdate(ctx context.Context, record interface{}) error

	// 删除
	Delete(ctx context.Context, ) error
	BatchDelete(ctx context.Context, ) error

	// 查询
	Find(ctx context.Context, v interface{}) error
	FindOne(ctx context.Context, v interface{}) error

	Count(ctx context.Context, ) (int64, error)
	//Distinct(ctx context.Context, field string, v interface{}) error

	Where(condition interface{}, args ...interface{}) IQuery
	Limit(limit int64) IQuery
	Offset(offset int64) IQuery
	OrderBy(fields ...string) IQuery
	OrderByDesc(fields ...string) IQuery
	Project(v interface{}) IQuery
}

// 聚合查询
type IAggQuery interface {
	Find(ctx context.Context, records interface{}) error
	FindOne(ctx context.Context, record interface{}) error

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
