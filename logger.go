package log

import (
	"log"
	"os"
)

// InitLogger инициализирует логирование
func InitLogger() *log.Logger {
	return log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
}
