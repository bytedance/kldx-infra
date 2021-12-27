package impl

import (
	cExceptions "code.byted.org/apaas/goapi_common/exceptions"
	"code.byted.org/apaas/goapi_infra/http/faasinfra"
	"code.byted.org/apaas/goapi_infra/mongodb"
	cond "code.byted.org/apaas/goapi_infra/mongodb/condition"
	op "code.byted.org/apaas/goapi_infra/mongodb/operator"
	"reflect"
)

type AggQuery struct {
	*MongodbParam
	conditions   []interface{}
	group        map[string]interface{}
	groupIDAlias string
}

func NewAggQuery(tableName string) *AggQuery {
	return &AggQuery{
		MongodbParam: &MongodbParam{
			TableName: tableName,
			Args: &MongodbArgs{
				Aggregate: true,
			},
		},
	}
}

func (a *AggQuery) Find(records interface{}) error {
	if a.Err != nil {
		return a.Err
	}
	a.SetOp(OpType_Aggregate)
	a.buildPipeline()
	return faasinfra.Find(a.MongodbParam, records)
}

func (a *AggQuery) FindOne(record interface{}) error {
	if a.Err != nil {
		return a.Err
	}
	a.SetOp(OpType_Aggregate)
	a.SetOne(true)
	a.buildPipeline()
	return faasinfra.FindOne(a.MongodbParam, record)
}

func (a *AggQuery) GroupBy(field interface{}, alias ...interface{}) mongodb.IAggQuery {
	_, value := a.parseKeyValue(field, alias...)
	if a.Err != nil || value == nil {
		return a
	}
	a.addGroup("_id", value)
	if len(alias) > 0 {
		if k, ok := alias[0].(string); ok {
			a.groupIDAlias = k
		}
	}
	return a
}

func (a *AggQuery) Push(field interface{}, alias ...interface{}) mongodb.IAggQuery {
	if a.Err != nil {
		return a
	}

	if field == nil {
		a.Err = cExceptions.InvalidParamError("field cannot be empty")
		return a
	}

	key := ""
	if len(alias) > 0 {
		if k, ok := alias[0].(string); ok {
			key = k
		}
	}
	if key == "" {
		a.Err = cExceptions.InvalidParamError("field cannot empty alias must be string and cannot be empty")
		return a
	}

	_, value := a.parseKeyValue(field, alias...)
	a.addGroup(key, cond.M{op.Push: value})
	return a
}

func (a *AggQuery) Having(condition interface{}) mongodb.IAggQuery {
	if a.Err != nil {
		return a
	}

	if condition == nil {
		return a
	}

	typ := reflect.TypeOf(condition)
	val := reflect.ValueOf(condition)
	if typ.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	switch typ.Kind() {
	case reflect.Slice:
		a.conditions = append(a.conditions, cond.M{op.And: condition})
	case reflect.Struct, reflect.Map:
		a.conditions = append(a.conditions, condition)
	default:
		a.Err = cExceptions.InvalidParamError("Having received invalid type, should be slice, struct or map, but received %s ", typ)
	}
	return a
}

func (a *AggQuery) First(field interface{}, alias ...interface{}) mongodb.IAggQuery {
	return a.appendGroup(op.First, field, alias...)
}

func (a *AggQuery) Last(field interface{}, alias ...interface{}) mongodb.IAggQuery {
	return a.appendGroup(op.Last, field, alias...)
}

func (a *AggQuery) Sum(field string, alias ...interface{}) mongodb.IAggQuery {
	return a.appendGroup(op.Sum, field, alias...)
}

func (a *AggQuery) Num(field string) mongodb.IAggQuery {
	if len(field) == 0 {
		field = "count"
	}
	a.addGroup(field, cond.M{op.Sum: 1})
	return a
}

func (a *AggQuery) Avg(field string, alias ...interface{}) mongodb.IAggQuery {
	return a.appendGroup(op.Avg, field, alias...)
}

func (a *AggQuery) StdDevPop(field string, alias ...interface{}) mongodb.IAggQuery {
	return a.appendGroup(op.StdDevPop, field, alias...)
}

func (a *AggQuery) StdDevSamp(field string, alias ...interface{}) mongodb.IAggQuery {
	return a.appendGroup(op.StdDevSamp, field, alias...)
}

func (a *AggQuery) AddToSet(field string, alias ...interface{}) mongodb.IAggQuery {
	return a.appendGroup(op.AddToSet, field, alias...)
}

func (a *AggQuery) appendGroup(op string, field interface{}, alias ...interface{}) mongodb.IAggQuery {
	key, value := a.parseKeyValue(field, alias...)
	if a.Err != nil || key == "" || value == nil {
		return a
	}
	a.addGroup(key, cond.M{op: value})
	return a
}

func (a *AggQuery) parseKeyValue(field interface{}, alias ...interface{}) (key string, value interface{}) {
	if a.Err != nil || field == nil {
		return
	}

	if len(alias) > 0 {
		if k, ok := alias[0].(string); ok {
			key = k
		}
	}

	switch field.(type) {
	case string:
		newField := field.(string)
		if len(key) == 0 {
			key = newField
		}
		value = "$" + newField
	case []string:
		newFields := field.([]string)
		if len(newFields) == 0 {
			return
		}

		if len(key) == 0 {
			a.Err = cExceptions.InvalidParamError("The first element of alias must be string type and not empty")
			return
		}

		m := cond.M{}
		for _, v := range newFields {
			m[v] = "$" + v
		}
		value = m
	case map[string]string:
		if len(key) == 0 {
			a.Err = cExceptions.InvalidParamError("The first element of alias must be string type and not empty")
			return
		}

		m := cond.M{}
		for k, v := range field.(map[string]string) {
			m[k] = "$" + v
		}
		value = m
	case map[string]interface{}:
		if len(key) == 0 {
			a.Err = cExceptions.InvalidParamError("The first element of alias must be string type and not empty")
			return
		}

		m := cond.M{}
		for k, v := range field.(map[string]interface{}) {
			if val, ok := v.(string); ok {
				m[k] = "$" + val
			} else {
				a.Err = cExceptions.InvalidParamError("The first element of alias must be string type and not empty")
				return
			}
		}
		value = m
	default:
		// struct or others
		if len(key) == 0 {
			a.Err = cExceptions.InvalidParamError("The first element of alias must be string type and not empty")
			return
		}
		value = field
	}
	return
}

func (p *AggQuery) addGroup(key string, value interface{}) {
	if p.group == nil {
		p.group = make(map[string]interface{})
	}
	p.group[key] = value
}

func (p *AggQuery) buildPipeline() {
	p.conditionAddPipeline()
	p.groupAddPipeLine()
}

func (p *AggQuery) groupAddPipeLine() {
	if p.group == nil {
		return
	}

	g := cond.M{
		"type":  "group",
		"group": p.group,
	}
	if len(p.groupIDAlias) > 0 {
		g["idAlias"] = p.groupIDAlias
	}
	p.Args.Pipeline = append(p.Args.Pipeline, g)
}

func (p *AggQuery) conditionAddPipeline() {
	if p.group == nil {
		return
	}

	g := cond.M{
		"type":  "matchGeneral",
	}

	if len(p.conditions) == 1 {
		g["match"] = p.conditions[0]
	} else if len(p.conditions) > 1 {
		g["match"] = cond.M{op.And: p.conditions}
	}
	p.Args.Pipeline = append(p.Args.Pipeline, g)
}
