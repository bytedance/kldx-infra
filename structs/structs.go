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

// file

type FileOption struct {
	Type   string `json:"type"`   // http content type
	Region string `json:"region"` // region of storage
}

type FileUploadResult struct {
	ID  string `json:"id,omitempty"`
	URL string `json:"url,omitempty"`
}