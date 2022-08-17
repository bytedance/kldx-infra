package faasinfra

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"

	cExceptions "github.com/bytedance/kldx-common/exceptions"
	"github.com/bytedance/kldx-infra/mongodb/structs"
	"github.com/bytedance/kldx-infra/mongodb/structs/inner"
)

func BatchCreate(ctx context.Context, param interface{}) ([]primitive.ObjectID, error) {
	data, err := doRequestMongodb(ctx, param)
	if err != nil {
		return nil, err
	}

	var result inner.BatchCreateResult
	err = bson.Unmarshal(data, &result)
	if err != nil {
		return nil, cExceptions.InternalError("BatchCreate failed, err: %v", err)
	}

	return result.IDs, nil
}

func Create(ctx context.Context, param interface{}) (*structs.RecordOnlyId, error) {
	ids, err := BatchCreate(ctx, param)
	if err != nil {
		return nil, err
	}

	if len(ids) > 0 {
		return &structs.RecordOnlyId{ID: ids[0]}, nil
	}

	return nil, nil
}

func Find(ctx context.Context, param, results interface{}) error {
	resultsVal := reflect.ValueOf(results)
	if resultsVal.Kind() != reflect.Ptr {
		return fmt.Errorf("[Find] results argument must be a pointer to a slice, but was a %s", resultsVal.Kind())
	}

	data, err := doRequestMongodb(ctx, param)
	if err != nil {
		return err
	}

	res := &inner.RawResult{}
	res.Bind(results)

	err = bson.Unmarshal(data, res)
	if err != nil {
		return cExceptions.InternalError("[Find] Unmarshal failed, err: %v", err)
	}

	return err
}

func FindOne(ctx context.Context, param, result interface{}) error {

	resultsVal := reflect.ValueOf(result)
	if resultsVal.Kind() != reflect.Ptr {
		return fmt.Errorf("[FindOne] results argument must be a pointer to a slice, but was a %s", resultsVal.Kind())
	}

	data, err := doRequestMongodb(ctx, param)
	if err != nil {
		return err
	}

	res := &inner.RawResult{}
	res.Bind(result)

	err = bson.Unmarshal(data, res)
	if err != nil {
		return cExceptions.InternalError("[FindOne] Unmarshal failed, err: %v", err)
	}
	return nil
}

func Count(ctx context.Context, param interface{}) (int64, error) {
	data, err := doRequestMongodb(ctx, param)
	if err != nil {
		return 0, err
	}

	result := &inner.CountResult{}

	err = bson.Unmarshal(data, &result)
	if err != nil {
		return 0, cExceptions.InternalError("[Count] Unmarshal failed, err: %v", err)
	}

	return result.Data.Count, nil
}

func Distinct(ctx context.Context, param interface{}, results interface{}) error {
	resultsVal := reflect.ValueOf(results)
	if resultsVal.Kind() != reflect.Ptr {
		return fmt.Errorf("[Distinct] results argument must be a pointer to a slice, but was a %s", resultsVal.Kind())
	}

	data, err := doRequestMongodb(ctx, param)
	if err != nil {
		return err
	}

	res := &inner.RawResult{}
	res.Bind(results)

	err = bson.Unmarshal(data, &res)
	if err != nil {
		return cExceptions.InternalError("[Distinct] Unmarshal failed, err: %v", err)
	}

	return nil
}

func Update(ctx context.Context, param interface{}) error {
	_, err := doRequestMongodb(ctx, param)
	if err != nil {
		return err
	}

	return nil
}

func Delete(ctx context.Context, param interface{}) error {
	_, err := doRequestMongodb(ctx, param)
	if err != nil {
		return err
	}

	return nil
}
