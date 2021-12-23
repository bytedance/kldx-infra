package exceptions

import (
	"fmt"
	"reflect"
)

const (
	ErrorTypes_InternalError     = "InternalError"
	ErrorTypes_InvalidParamError = "InvalidParamError"
)

type systemError struct {
	errorType string
	msg       string
}

func (e *systemError) Error() string {
	return fmt.Sprintf("%s: %s", e.errorType, e.msg)
}

func InternalError(format string, args ...interface{}) *systemError {
	msg := fmt.Sprintf(format, args...)
	err := &systemError{
		errorType: ErrorTypes_InternalError,
		msg:       msg,
	}

	// TODO 打 error 日志
	fmt.Println(err)
	return err
}

type businessError struct {
	errorType string
	msg       string
}

func (e *businessError) Error() string {
	return fmt.Sprintf("%s: %s", e.errorType, e.msg)
}

func InvalidParamError(format string, args ...interface{}) *businessError {
	msg := fmt.Sprintf(format, args...)
	err := &businessError{
		errorType: ErrorTypes_InvalidParamError,
		msg:       msg,
	}

	// TODO 打 warning 日志
	fmt.Println(err)
	return err
}

func ErrorWrap(err error) error {
	if err == nil {
		return err
	}

	typ := reflect.TypeOf(err)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	// 已知类型，直接返回
	if typ.Name() == "systemError" || typ.Name() == "businessError" {
		return err
	}

	// 未知类型当成系统错误
	return InternalError(err.Error())
}
