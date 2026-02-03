package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	// Get user input for filename
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Which file will you read?")

	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		log.Fatal("No input")
	}

	link := strings.TrimSpace(scanner.Text())

	if link == "" {
		log.Fatal("No input")
	}

	// Read file
	data, err := os.ReadFile(link)
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(data)
	fmt.Println() // Prevent '%' letter end of output
}
