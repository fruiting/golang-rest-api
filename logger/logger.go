package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	f, err := os.OpenFile("loggers/testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(f)
}

// Info - function of logging general info
func Info(info string) {
	log.Println(info)
}

// Error - function of logging errors in system
func Error(error string) {
	log.Error(error)
}
