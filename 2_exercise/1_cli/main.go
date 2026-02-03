package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// Get user input for filename
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Which file will you read?")

	scanner.Scan()
	link := scanner.Text()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Read file
	data, err := os.ReadFile(link)
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(data)
	fmt.Println() // Prevent '%' letter end of output
}
