package logger

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

// Here will be implemented injected interface with ELK system or Prometheus
// We might create custom errors for any case and system
// Also we need to define always where exactly error happened so we can use runtime caller and our custom interface params to add service, etc

func LogError(payload ...interface{}) {
	errorString := strings.Trim(fmt.Sprint(payload), "[]")

	_, file, line, ok := runtime.Caller(2)
	if ok {
		log.Printf("%s (%s:%d)", errorString, file, line)
	} else {
		log.Print(errorString)
	}

	stringSlice := strings.Split(file, "/")
	fileName := stringSlice[len(stringSlice)-1]

	log.Print(fmt.Sprintf("[Error] File: %s, line: %v, ERROR: %s", fileName, line, errorString))
}
