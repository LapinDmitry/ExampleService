package utils

import (
	"os"
	"strconv"
)

// GetEnv - для удобного чтения переменных окружения
func GetEnv(key string, defaultVal interface{}) interface{} {
	if value, exists := os.LookupEnv(key); exists {

		var result = defaultVal

		switch defaultVal.(type) {
		case int:
			if res, err := strconv.Atoi(value); err == nil {
				result = res
			}
		case int64:
			if res, err := strconv.ParseInt(value, 10, 64); err == nil {
				result = res
			}
		case float64:
			if res, err := strconv.ParseFloat(value, 64); err == nil {
				result = res
			}
		case string:
			result = value
		default:
			return defaultVal
		}

		return result
	}

	return defaultVal
}
