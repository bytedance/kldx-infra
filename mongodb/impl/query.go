package impl

import (
	cExceptions "github.com/bytedance/kldx-common/exceptions"
	"github.com/bytedance/kldx-infra/http/faasinfra"
	"github.com/bytedance/kldx-infra/mongodb"
	cond "github.com/bytedance/kldx-infra/mongodb/condition"
	op "github.com/bytedance/kldx-infra/mongodb/operator"
	"reflect"
)

const (
	Asc  = 1
	Desc = -1
)

type Query struct {
	*MongodbParam
	conditions []interface{}
}

func NewQuery(tableName string) *Query {
	q := &Query{MongodbParam: &MongodbParam{
		TableName: tableName,
		Args: NewMongodbArgs(),
	}}
	return q
}

func (q *Query) Update(record interface{}) error {
	if q.Err != nil {
		return q.Err
	}

	typ := reflect.TypeOf(record)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct && typ.Kind() != reflect.Map {
		return cExceptions.InvalidParamError("Update failed: record should be map or struct, but %s", typ)
	}

	q.SetOp(OpType_Update)
	q.SetUpdate(cond.M{op.Set: record})
	q.SetOne(true)
	q.SetUpsert(false)
	q.buildQuery()
	return faasinfra.Update(q.MongodbParam)
}

func (q *Query) Upsert(record interface{}) error {
	if q.Err != nil {
		return q.Err
	}

	typ := reflect.TypeOf(record)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct && typ.Kind() != reflect.Map {
		return cExceptions.InvalidParamError("Update failed: record should be map or struct, but %s", typ)
	}

	q.SetOp(OpType_Update)
	q.SetUpdate(cond.M{op.Set: record})
	q.SetOne(true)
	q.SetUpsert(true)
	q.buildQuery()
	return faasinfra.Update(q.MongodbParam)
}

func (q *Query) BatchUpdate(record interface{}) error {
	if q.Err != nil {
		return q.Err
	}

	typ := reflect.TypeOf(record)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct && typ.Kind() != reflect.Map {
		return cExceptions.InvalidParamError("Update failed: record should be map or struct, but %s", typ)
	}

	q.SetOp(OpType_Update)
	q.SetUpdate(cond.M{op.Set: record})
	q.SetOne(false)
	q.SetUpsert(false)
	q.buildQuery()
	return faasinfra.Update(q.MongodbParam)
}

func (q *Query) Delete() error {
	if q.Err != nil {
		return q.Err
	}
	q.SetOp(OpType_Delete)
	q.SetOne(true)
	q.buildQuery()
	return faasinfra.Delete(q.MongodbParam)
}

func (q *Query) BatchDelete() error {
	if q.Err != nil {
		return q.Err
	}
	q.SetOp(OpType_Delete)
	q.SetOne(false)
	q.buildQuery()
	return faasinfra.Delete(q.MongodbParam)
}

func (q *Query) Find(records interface{}) error {
	if q.Err != nil {
		return q.Err
	}
	q.SetOp(OpType_Find)
	q.buildQuery()
	return faasinfra.Find(q.MongodbParam, records)
}

func (q *Query) FindOne(record interface{}) error {
	if q.Err != nil {
		return q.Err
	}
	q.SetOp(OpType_FindOne)
	q.SetLimit(1)
	q.buildQuery()
	return faasinfra.FindOne(q.MongodbParam, record)
}

func (q *Query) Where(condition interface{}, args ...interface{}) mongodb.IQuery {
	if q.Err != nil {
		return q
	}

	if condition == nil {
		return q
	}

	typ := reflect.TypeOf(condition)
	val := reflect.ValueOf(condition)
	if typ.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	switch typ.Kind() {
	case reflect.Slice:
		q.conditions = append(q.conditions, cond.M{op.And: condition})
	case reflect.Struct, reflect.Map:
		q.conditions = append(q.conditions, condition)
	default:
		q.Err = cExceptions.InvalidParamError("Query.Where received invalid type, should be slice, struct or map, but received %s ", typ)
	}
	return q
}

func (q *Query) Limit(limit int64) mongodb.IQuery {
	if q.Err != nil {
		return q
	}
	q.SetLimit(limit)
	return q
}

func (q *Query) Offset(offset int64) mongodb.IQuery {
	if q.Err != nil {
		return q
	}
	q.SetOffset(offset)
	return q
}

func (q *Query) OrderBy(fields ...string) mongodb.IQuery {
	if q.Err != nil {
		return q
	}
	for _, field := range fields {
		q.AddSort(field, Asc)
	}
	return q
}

func (q *Query) OrderByDesc(fields ...string) mongodb.IQuery {
	if q.Err != nil {
		return q
	}
	for _, field := range fields {
		q.AddSort(field, Desc)
	}
	return q
}

func (q *Query) Count() (int64, error) {
	if q.Err != nil {
		return 0, q.Err
	}

	q.SetOp(OpType_Count)
	q.buildQuery()
	return faasinfra.Count(q.MongodbParam)
}

func (q *Query) Distinct(field string, v interface{}) error {
	if q.Err != nil {
		return q.Err
	}

	q.SetOp(OpType_Distinct)
	q.SetKey(field)
	return faasinfra.Distinct(q.MongodbParam, v)
}

func (q *Query) Project(projection interface{}) mongodb.IQuery {
	if q.Err != nil {
		return q
	}

	q.SetProjection(projection)
	return q
}

func (q *Query) buildQuery() {
	if len(q.conditions) == 1 {
		q.SetQuery(q.conditions[0])
	} else if len(q.conditions) > 1 {
		q.SetQuery(cond.M{op.And: q.conditions})
	}
}
