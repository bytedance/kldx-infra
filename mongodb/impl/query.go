package impl

import (
	"code.byted.org/apaas/goapi_infra/http/faasinfra"
	"code.byted.org/apaas/goapi_infra/mongodb"
)

type Query struct {
	*MongodbParam
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
	panic("implement me")
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
	panic("implement me")
}

func (q *Query) Offset(offset int64) mongodb.IQuery {
	panic("implement me")
}

func (q *Query) Sort(v interface{}) mongodb.IQuery {
	panic("implement me")
}

func (q *Query) Count() (int64, error) {
	panic("implement me")
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
