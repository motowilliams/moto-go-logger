package logger

import "fmt"

func Info(message string) {
	fmt.Printf("[Info]: %s\n", message)
}
