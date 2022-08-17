package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// unittest to use
func PrintLog(contents ...interface{}) {
	isPrint := false
	for _, content := range contents {
		if content == nil {
			fmt.Println(content)
			isPrint = true
			continue
		}

		typ := reflect.TypeOf(content)
		val := reflect.ValueOf(content)
		if typ.Kind() == reflect.Ptr {
			val = val.Elem()
			typ = typ.Elem()
		}

		switch typ.Kind() {
		case reflect.String:
			fmt.Println(content)
			isPrint = true
		default:
			content, err := json.Marshal(content)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(content))
			isPrint = true
		}
	}

	if isPrint {
		fmt.Println()
	}
}
