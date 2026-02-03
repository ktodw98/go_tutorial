package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// Get user input for filename
	link := scanShortText("Which file will you read?")

	// Read file
	data, err := os.ReadFile(link)
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(data)
	fmt.Println() // Prevent '%' letter end of output
}

func scanShortText(info string) string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("%v > ", info)

	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		log.Fatal("No Input")
	}

	inputText := scanner.Text()

	if inputText == "" {
		log.Fatal("no input")
	}

	return inputText
}
