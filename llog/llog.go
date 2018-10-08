package llog

import (
	"fmt"
)

// Good prints a "good" (green) message with a timestamp and the given message
func Good(text string) {
	fmt.Printf("\033[36m [+] \033[0m%-60s\t", text)
}

// Bad prints a "good" (green) message with a timestamp and the given message
func Bad(text string) {
	fmt.Printf("\033[31m [+] \033[0m%-60s\t", text)
}
