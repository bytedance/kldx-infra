package http

import (
	"fmt"
	"testing"
)

func TestGetAppToken(t *testing.T) {
	token, err := GetAppToken()
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}
