package impl

import (
	"code.byted.org/apaas/goapi_infra/mongodb"
)

type Mongodb struct {
}

func NewMongodb() *Mongodb {
	return &Mongodb{}
}

func (m *Mongodb) Table(tableName string) mongodb.ITable {
	return NewTable(tableName)
}

