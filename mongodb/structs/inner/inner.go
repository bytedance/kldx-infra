package inner

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
)

type BatchCreateResult struct {
	IDs []primitive.ObjectID `bson:"data"`
}

type CountResult struct {
	Data struct {
		Count int64 `bson:"count"`
	} `bson:"data"`
}

type Result interface {
	Bind(interface{})
	Err() error
}

type RawResult struct {
	Data interface{} `json:"data" bson:"data"`
	E    error       `json:"-" bson:"-"`
}

func (r *RawResult) Bind(v interface{}) {
	r.Data = v
}

func (r *RawResult) Err() error {
	return nil
}

func (r *RawResult) UnmarshalBSON(b []byte) error {
	elems, err := bson.Raw(b).Elements()
	if err != nil {
		return err
	}

	for _, elem := range elems {
		switch elem.Key() {
		case "data":
			if r.Data == nil {
				r.Data = bson.M{}
				if err := elem.Value().Unmarshal(&r.Data); err != nil {
					return err
				}
			} else {
				typ := reflect.TypeOf(r.Data)
				if typ.Kind() != reflect.Ptr {
					r.E = fmt.Errorf("received invalid type for result, should be ptr, received %s", typ)
					return r.E
				}
				if err := elem.Value().Unmarshal(r.Data); err != nil {
					r.E = err
					return r.E
				}
			}
		}
	}
	return nil
}
