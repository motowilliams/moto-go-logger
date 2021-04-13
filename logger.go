package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func Info(message string) {
	fmt.Printf("[Info] %s: %s\n", timeFormat(), message)
}

func InfoJson(message string) {

	type LogEntry struct {
		Level   string
		Time    string
		Message string
	}

	logEntry := LogEntry{"Info", timeFormat(), message}

	prettyJSON, err := json.MarshalIndent(logEntry, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", string(prettyJSON))
}

func timeFormat() string {
	return time.Now().UTC().Format(time.RFC3339)
}
