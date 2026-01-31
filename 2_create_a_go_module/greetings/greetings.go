package greetings

import "fmt"

func Hello(name string) string {
	// Return a greeting that embeds the name in message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
