package logger

import (
	"io"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func Init() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	// create logs on current machine/environment if it does not exist
	if err := os.MkdirAll("/app/logs", 0755); err != nil {
		log.Printf("Failed to create logs directory: %v", err)
	}

	// set file logging for current machine/environment
	file, err := os.OpenFile("/app/logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Failed to open log file: %v", err)
		// stdout if file fails, not in the mood to relive debugging a tool that lacks
		// a clear documentaton and relies on blogs and mostly the community to do
		// the docs
		logger.SetOutput(os.Stdout)
	} else {
		multiWriter := io.MultiWriter(os.Stderr, file)
		logger.SetOutput(multiWriter)
	}
}

// returns the logger instance
func Get() *logrus.Logger {
	return logger
}

// TODO: refine!
