package faasinfra

import (
	cExceptions "code.byted.org/apaas/goapi_common/exceptions"
	"code.byted.org/apaas/goapi_infra/structs"
	"encoding/json"
	"fmt"
)

func BatchCreate(param interface{}) ([]string, error) {
	data, err := doRequestMongodb(param)
	if err != nil {
		return nil, err
	}

	var result structs.BatchCreateResult
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, cExceptions.InternalError("BatchCreate failed, err: %v", err)
	}

	var ids []string
	for _, id := range result.Ids {
		ids = append(ids, id)
	}

	return ids, nil
}

func Create(param interface{}) (*structs.RecordOnlyId, error) {
	ids, err := BatchCreate(param)
	if err != nil {
		return nil, err
	}

	if len(ids) > 0 {
		return &structs.RecordOnlyId{Id: ids[0]}, nil
	}

	return nil, nil
}

func Find(param, records interface{}) error {
	data, err := doRequestMongodb(param)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &records)
	if err != nil {
		return cExceptions.InternalError("Find failed, err: %v", err)
	}
	return nil
}

func FindOne(param, record interface{}) error {
	data, err := doRequestMongodb(param)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &record)
	if err != nil {
		return cExceptions.InternalError("FindOne failed, err: %v", err)
	}
	return nil
}

func Count(param interface{}) (int64, error) {
	data, err := doRequestMongodb(param)
	if err != nil {
		return 0, err
	}

	fmt.Println(string(data))

	return 0, nil
}
