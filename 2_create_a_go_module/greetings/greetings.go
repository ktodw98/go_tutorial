package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	// If name is not given, return error
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. nice to meet you",
		"How are you doing? %v",
		"I'm goooood %v.",
	}

	return formats[rand.Intn(len(formats))]
}
