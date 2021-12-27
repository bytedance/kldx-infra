package impl

import (
	cExceptions "code.byted.org/apaas/goapi_common/exceptions"
)

type OpType int

const (
	OpType_Insert OpType = iota + 1
	OpType_Find
	OpType_FindOne
	OpType_Delete
	OpType_Count
	OpType_Update
	OpType_Upsert
	OpType_Aggregate
)

var opTypeString = map[OpType]string{
	OpType_Insert:    "insert",
	OpType_Find:      "find",
	OpType_FindOne:   "findOne",
	OpType_Delete:    "delete",
	OpType_Count:     "count",
	OpType_Update:    "update",
	OpType_Upsert:    "replace",
	OpType_Aggregate: "aggregate",
}

type MongodbParam struct {
	TableName string       `json:"tableName"`
	Args      *MongodbArgs `json:"args"`
	Err       error        `json:"-"`
}

type MongodbArgs struct {
	Op           string           `json:"op"`
	Docs         interface{}      `json:"docs,omitempty"`
	Query        interface{}      `json:"query,omitempty"`
	Collection   string           `json:"collection,omitempty"`
	Sort         map[string]int64 `json:"sort,omitempty"`
	Projection   interface{}      `json:"projection,omitempty"`
	Project      interface{}      `json:"project,omitempty"`
	Hint         interface{}      `json:"hint,omitempty"`
	Skip         int64            `json:"skip,omitempty"`
	Limit        int64            `json:"limit,omitempty"`
	ArrayFilters interface{}      `json:"arrayFilters,omitempty"`
	Upsert       bool             `json:"upsert,omitempty"`
	Distinct     string           `json:"distinct,omitempty"`
	Pipeline     interface{}      `json:"pipeline,omitempty"`
}

func NewMongodbArgs() *MongodbArgs {
	return &MongodbArgs{
		Sort: make(map[string]int64),
	}
}

func NewMongodbParam(tableName string) *MongodbParam {
	return &MongodbParam{TableName: tableName, Args: NewMongodbArgs()}
}

func (p *MongodbParam) SetTableName(tableName string) {
	p.TableName = tableName
}

func (p *MongodbParam) SetOp(op OpType) {
	p.Args.Op = opTypeString[op]
}

func (a *MongodbParam) GetOp() string {
	return a.Args.Op
}

func (a *MongodbParam) SetDocs(docs interface{}) {
	a.Args.Docs = docs
}

func (p *MongodbParam) SetLimit(limit int64) {
	if limit < 1 || limit > 1000 {
		p.Err = cExceptions.InvalidParamError("Limit received invalid value (%d), should be 1~1000", limit)
	}

	p.Args.Limit = limit
}

func (p *MongodbParam) SetOffset(offset int64) {
	if offset < 0 {
		p.Err = cExceptions.InvalidParamError("Offset received invalid value (%d), should be >= 0", offset)
	}

	p.Args.Skip = offset
}

func (p *MongodbParam) SetQuery(condition interface{}) {
	p.Args.Query = condition
}

func (p *MongodbParam) AddSort(field string, direct int64) {
	p.Args.Sort[field] = direct
}
