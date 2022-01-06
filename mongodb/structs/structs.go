package structs

type RecordOnlyId struct {
	Id string `json:"_id"`
}

type BatchCreateResult struct {
	Ids []string `json:"_ids"`
}

type CountResult struct {
	Count int64 `json:"count"`
}
