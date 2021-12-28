package impl

import (
	"github/kldx/infra/mongodb"
)

type Mongodb struct {
}

func NewMongodb() *Mongodb {
	return &Mongodb{}
}

func (m *Mongodb) Table(tableName string) mongodb.ITable {
	return NewTable(tableName)
}

