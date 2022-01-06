package file

import (
	"errors"
	"fmt"
)

// file

type Option struct {
	Type   string `json:"type"`   // http content type
	Region string `json:"region"` // region of storage
}

type UploadResult struct {
	ID  string `json:"id,omitempty"`
	URL string `json:"url,omitempty"`
}

type uploadError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e uploadError) error() error {
	if e.Code != 0 {
		if len(e.Message) == 0 {
			e.Message = "upload file fail"
		}
		return errors.New(fmt.Sprintf(`code: %d, message:"%s"`, e.Code, e.Message))
	}
	return nil
}

type fileUploadResult struct {
	ID  string `json:"id,omitempty"`
	URL string `json:"url,omitempty"`
	*uploadError
}