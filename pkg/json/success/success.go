package json

// Success create a success response body basic structure
func Success(message string, statusCode int) map[string]interface{} {
	h := make(map[string]interface{})

	h["status"] = "success"
	h["message"] = message

	// define a custom status code
	if statusCode != 0 {
		h["statusCode"] = statusCode
	}

	return h
}

// OK return a ok response body
func OK(message string, statusCode int) map[string]interface{} {
	h := Success(message, statusCode)
	return h
}

// NoContent return no response body
func NoContent(message string, statusCode int) map[string]interface{} {
	h := Success(message, statusCode)
	return h
}
