package impl

import (
	cExceptions "github.com/bytedance/kldx-common/exceptions"
)

type OpType int

const (
	OpType_Insert OpType = iota + 1
	OpType_Find
	OpType_FindOne
	OpType_Distinct
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
	OpType_Distinct:  "distinct",
	OpType_Delete:    "delete",
	OpType_Count:     "count",
	OpType_Update:    "update",
	OpType_Upsert:    "replace",
	OpType_Aggregate: "aggregate",
}

type MongodbParam struct {
	TableName string       `bson:"tableName" json:"tableName"`
	Args      *MongodbArgs `bson:"args" json:"args"`
	Err       error        `bson:"-" json:"-"`
}

type MongodbArgs struct {
	Op           string                   `bson:"op" json:"op"`
	Docs         interface{}              `bson:"docs,omitempty" json:"docs,omitempty"`
	Query        interface{}              `bson:"query,omitempty" json:"query,omitempty"`
	Collection   string                   `bson:"collection,omitempty" json:"collection,omitempty"`
	Sort         map[string]int64         `bson:"sort,omitempty" json:"sort,omitempty"`
	Projection   interface{}              `bson:"projection,omitempty" json:"projection,omitempty"`
	Hint         interface{}              `bson:"hint,omitempty" json:"hint,omitempty"`
	Skip         int64                    `bson:"skip,omitempty" json:"skip,omitempty"`
	Limit        int64                    `bson:"limit,omitempty" json:"limit,omitempty"`
	ArrayFilters interface{}              `bson:"arrayFilters,omitempty" json:"arrayFilters,omitempty"`
	Upsert       *bool                    `bson:"upsert,omitempty" json:"upsert,omitempty"`
	Pipeline     []map[string]interface{} `bson:"pipeline,omitempty" json:"pipeline,omitempty"`
	Update       interface{}              `bson:"update,omitempty" json:"update,omitempty"`
	One          *bool                    `bson:"one,omitempty" json:"one,omitempty"`
	// Distinct
	Key string `bson:"key,omitempty" json:"key,omitempty"`
	// aggregate
	Aggregate bool `bson:"aggregate,omitempty" json:"aggregate,omitempty"`
}

func NewMongodbArgs() *MongodbArgs {
	return &MongodbArgs{}
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

func (p *MongodbParam) SetKey(key string) {
	p.Args.Key = key
}

func (p *MongodbParam) SetProjection(projection interface{}) {
	p.Args.Projection = projection
}

func (p *MongodbParam) SetOne(one bool) {
	p.Args.One = &one
}

func (p *MongodbParam) SetAggregate(agg bool) {
	p.Args.Aggregate = agg
}

func (p *MongodbParam) SetUpsert(upsert bool) {
	p.Args.Upsert = &upsert
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
	if p.Args.Sort == nil {
		p.Args.Sort = make(map[string]int64)
	}
	p.Args.Sort[field] = direct
}

func (p *MongodbParam) SetUpdate(field2value interface{}) {
	p.Args.Update = field2value
}
