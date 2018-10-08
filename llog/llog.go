package llog

import (
	"fmt"
	"time"
)

// Good prints a "good" (green) message with a timestamp and the given message
func Good(text string) {
	fmt.Printf("%s\033[36m [+] \033[0m%-60s\t", currentTime(), text)
}

// Bad prints a "good" (green) message with a timestamp and the given message
func Bad(text string) {
	fmt.Printf("%s\033[31m [+] \033[0m%-60s\t", currentTime(), text)
}

// current_time returns the current time as a string in the HH:MM:SS format
func currentTime() string {
	t := time.Now()
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
}
