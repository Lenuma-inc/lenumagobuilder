package utils

import (
	"fmt"
)

// FormatOutput форматирует вывод для логов
func FormatOutput(msg string) string {
	return fmt.Sprintf("INFO: %s", msg)
}
