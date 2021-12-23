package structs

type TokenResult struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		AccessToken string `json:"accessToken"`
		ExpireTime  int64  `json:"expireTime"`
	} `json:"data"`
}

type RecordOnlyId struct {
	Id string `json:"_id"`
}

type BatchCreateResult struct {
	Ids []string `json:"_ids"`
}
