package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Gwen")
	fmt.Println(message)
}
