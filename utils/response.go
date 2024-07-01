package utils

func Response(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
		"data":    data,
	}
}
