package impl

import (
	"code.byted.org/apaas/goapi_infra/mongodb"
)

type AggQuery struct {
	
}

func (a *AggQuery) Find(records interface{}) error {
	panic("implement me")
}

func (a *AggQuery) FindOne(record interface{}) error {
	panic("implement me")
}

func (a *AggQuery) GroupBy(field string, args ...interface{}) mongodb.IAggQuery {
	panic("implement me")
}

func (a *AggQuery) Having(condition interface{}) mongodb.IAggQuery {
	panic("implement me")
}

func (a *AggQuery) First(field string, args ...interface{}) mongodb.IAggQuery {
	panic("implement me")
}

func (a *AggQuery) Last(field string, args ...interface{}) mongodb.IAggQuery {
	panic("implement me")
}

func (a *AggQuery) Sum(field string, args ...interface{}) mongodb.IAggQuery {
	panic("implement me")
}

func (a *AggQuery) Num(field string) mongodb.IAggQuery {
	panic("implement me")
}

func (a *AggQuery) Avg(field string, args ...interface{}) mongodb.IAggQuery {
	panic("implement me")
}

func (a *AggQuery) StdDevPop(field string, args ...interface{}) mongodb.IAggQuery {
	panic("implement me")
}

func (a *AggQuery) StdDevSamp(field string, args ...interface{}) mongodb.IAggQuery {
	panic("implement me")
}

