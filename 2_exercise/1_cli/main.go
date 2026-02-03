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
	var ascVal bool

	if *fileFlag != "" {
		fileName = *fileFlag
	} else {
		// no flag input
		fileName = scanShortText("Which file will you read? ")
	}

	if *ascFlag {
		ascVal = true
	} else {
		ascVal = false
	}

	// Trimming Space in file name text
	fileName = strings.TrimSpace(fileName)

	// Read file
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	if !ascVal {
		reversed := make([]byte, 0)
		for i := len(data) - 1; i >= 0; i = i - 1 {
			reversed = append(reversed, data[i])
		}
		data = reversed
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
