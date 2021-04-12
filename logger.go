package logger

import (
	"fmt"
	"time"
)

func Info(message string) {
	fmt.Printf("[Info] %s: %s\n", timeFormat(), message)
}

func timeFormat() string {
	return time.Now().UTC().Format(time.RFC3339)
}
