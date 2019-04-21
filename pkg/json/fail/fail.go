package json

import (
	"errors"
)

// ErrNotFound is used when a beer could not be found.
var ErrNotFound = errors.New("beer not found")

// ErrDuplicate is used when a beer already existed.
var ErrDuplicate = errors.New("beer already exists")

// Faild create a failed response body basic structure
func Faild(message string, statusCode int) map[string]interface{} {
	h := make(map[string]interface{})

	h["status"] = "failed"
	h["message"] = message

	// define a custom status code
	if statusCode != 0 {
		h["statusCode"] = statusCode
	}

	return h
}

// NotFound return a not found failed map
func NotFound(message string, statusCode int) map[string]interface{} {
	h := Faild(message, statusCode)
	return h
}

// Forbidden return a forbidden failed map
func Forbidden(message string, statusCode int) map[string]interface{} {
	h := Faild(message, statusCode)
	return h
}

// InvalidParameter return a invalid parameter failed map
func InvalidParameter(message string, statusCode int) map[string]interface{} {
	h := Faild(message, statusCode)
	return h
}

// BindJSONFail return a bind json failed map
func BindJSONFail(message string, statusCode int) map[string]interface{} {
	h := Faild(message, statusCode)
	return h
}

// InvalidID return a invalid id failed map
func InvalidID(message string, statusCode int) map[string]interface{} {
	h := Faild(message, statusCode)
	return h
}
