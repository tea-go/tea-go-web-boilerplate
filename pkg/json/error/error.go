package json

// Error create a error response body basic structure
func Error(message string, statusCode int) map[string]interface{} {
	h := make(map[string]interface{})

	h["status"] = "error"
	h["message"] = message

	// define a custom status code
	if statusCode != 0 {
		h["statusCode"] = statusCode
	}

	return h
}

// InternalServerError return a internal server error map
func InternalServerError(message string, statusCode int) map[string]interface{} {
	h := Error(message, statusCode)
	return h
}
