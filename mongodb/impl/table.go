package impl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"

	cExceptions "github.com/bytedance/kldx-common/exceptions"
	"github.com/bytedance/kldx-infra/http/faasinfra"
	"github.com/bytedance/kldx-infra/mongodb"
	"github.com/bytedance/kldx-infra/mongodb/structs"
)

type Table struct {
	*MongodbParam
}

func NewTable(tableName string) *Table {
	t := &Table{MongodbParam: NewMongodbParam(tableName)}
	if len(tableName) == 0 {
		t.Err = cExceptions.InvalidParamError("tableName is empty")
	}
	return t
}

func (t *Table) Create(ctx context.Context, record interface{}) (*structs.RecordOnlyId, error) {
	if t.Err != nil {
		return nil, t.Err
	}

	t.SetOp(OpType_Insert)
	t.SetDocs([]interface{}{record})
	return faasinfra.Create(ctx, t.MongodbParam)
}

func (t *Table) BatchCreate(ctx context.Context, records interface{}) ([]primitive.ObjectID, error) {
	if t.Err != nil {
		return nil, t.Err
	}

	t.SetOp(OpType_Insert)
	t.SetDocs(records)
	return faasinfra.BatchCreate(ctx, t.MongodbParam)
}

func (t *Table) Where(condition interface{}, args ...interface{}) mongodb.IQuery {
	return NewQuery(t.MongodbParam.TableName).Where(condition, args)
}

func (q *Table) GroupBy(field interface{}, alias ...interface{}) mongodb.IAggQuery {
	return NewAggQuery(q.TableName).GroupBy(field, alias...)
}
