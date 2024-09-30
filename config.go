package config

import (
	"os"
)

// GetEnv получает переменные окружения с указанием значения по умолчанию
func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
