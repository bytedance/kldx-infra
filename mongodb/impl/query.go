package impl

import (
	cExceptions "code.byted.org/apaas/goapi_common/exceptions"
	"code.byted.org/apaas/goapi_infra/http/faasinfra"
	"code.byted.org/apaas/goapi_infra/mongodb"
	cond "code.byted.org/apaas/goapi_infra/mongodb/condition"
	op "code.byted.org/apaas/goapi_infra/mongodb/operator"
	"reflect"
)

const (
	Asc = 1
	Desc = -1
)

type Query struct {
	*MongodbParam
	conditions []interface{}
}

func NewQuery(param *MongodbParam) *Query {
	q := &Query{MongodbParam: param}
	return q
}

func (q *Query) Update() error {
	if q.Err != nil {
		return q.Err
	}
	panic("implement me")
}

func (q *Query) Upsert() error {
	if q.Err != nil {
		return q.Err
	}
	panic("implement me")
}

func (q *Query) BatchUpdate() error {
	if q.Err != nil {
		return q.Err
	}
	panic("implement me")
}

func (q *Query) Delete() error {
	if q.Err != nil {
		return q.Err
	}
	panic("implement me")
}

func (q *Query) BatchDelete() error {
	if q.Err != nil {
		return q.Err
	}
	panic("implement me")
}

func (q *Query) Find(records interface{}) error {
	if q.Err != nil {
		return q.Err
	}
	q.SetOp(OpType_Find)

	if len(q.conditions) == 1 {
		q.SetQuery(q.conditions[0])
	} else if len(q.conditions) > 1 {
		q.SetQuery(cond.M{op.And: q.conditions})
	}

	return faasinfra.Find(q.MongodbParam, records)
}

func (q *Query) FindOne(record interface{}) error {
	if q.Err != nil {
		return q.Err
	}
	q.SetOp(OpType_FindOne)
	q.SetLimit(1)
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

func (q *Query) And(elems ...interface{}) mongodb.IQuery {
	panic("implement me")
}

func (q *Query) Or(elems ...interface{}) mongodb.IQuery {
	panic("implement me")
}

func (q *Query) Nor(elems ...interface{}) mongodb.IQuery {
	panic("implement me")
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

func (q *Query) Sort(v interface{}) mongodb.IQuery {
	panic("implement me")
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
	return faasinfra.Count(q.MongodbParam)
}

func (q *Query) Distinct(field string, args ...interface{}) mongodb.IQuery {
	panic("implement me")
}

func (q *Query) Project(v interface{}) mongodb.IQuery {
	panic("implement me")
}

func (q *Query) GroupBy(field string, args ...interface{}) mongodb.IAggQuery {
	panic("implement me")
}
