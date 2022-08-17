package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type RecordOnlyId struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
}
