package impl

import (
	"code.byted.org/apaas/goapi_infra/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mongodb struct {
}

func NewMongodb() *Mongodb {
	return &Mongodb{}
}

func (m *Mongodb) Table(tableName string) mongodb.ITable {
	return NewTable(tableName)
}

func (m *Mongodb) NewObjectId() primitive.ObjectID {
	return primitive.NewObjectID()
}

func (m *Mongodb) ParseObjectId(_id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(_id)
}

