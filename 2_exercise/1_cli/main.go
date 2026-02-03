package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Define Flag
	fileFlag := flag.String("file", "", "File route")
	ascFlag := flag.Bool("asc", true, "whether print out in ascending order")

	// Parsing Flag
	flag.Parse()

	var fileName string

	if *fileFlag != "" {
		fileName = *fileFlag
	} else {
		// no flag input
		fileName = scanShortText("Which file will you read? ")
	}

	ascVal := *ascFlag

	// Trimming Space in file name text
	fileName = strings.TrimSpace(fileName)

	// Read file
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	if !ascVal {
		runes := []rune(string(data))
		n := len(runes)
		for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		data = []byte(string(runes))
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
