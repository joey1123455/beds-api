package logger

import (
	"log"
	"os"
)

type Fields map[string]interface{}

type Logger interface {
	Info(message string, properties Fields)
	Error(err error, properties Fields)
	Fatal(err error, properties Fields)
	Debug(message string, properties Fields)
	SetLevel(level Level)
}

type Level int8

const (
	LevelInfo Level = iota
	LevelError
	LevelFatal
	LevelOff
	LevelDebug
)

func (l Level) String() string {
	switch l {
	case LevelInfo:
		return "INFO"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	case LevelDebug:
		return "DEBUG"
	default:
		return ""
	}
}

func openLogFile(path string) (*os.File, error) {
	// Open the file in append mode, creating it if it doesn't exist
	logFile, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

func ErrorLogger(e error) {
	logFile, err := openLogFile("./errors.log")
	if err != nil {
		log.Println(err)
	}

	// Set the output destination for log messages
	log.SetOutput(logFile)

	// Configure log flags
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)

	// Example log message
	log.Println("Error:", e.Error())

	// Close the log file after use
	logFile.Close()
}
