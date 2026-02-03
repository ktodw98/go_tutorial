package main

import (
	"flag"
	"log"
	"strings"
)

func main() {
	// Define Flag
	fileFlag := flag.String("file", "example-big.txt", "File route")
	ascFlag := flag.Bool("asc", true, "whether print out in ascending order")
	modeFlag := flag.String("mode", "all", "all, line")

	// Parsing Flag
	flag.Parse()

	// Scanner
	handler := NewInputHandler()

	var filePath string
	var mode string

	if *fileFlag != "" {
		filePath = *fileFlag
	} else {
		// no flag input
		filePath = handler.Ask("Which file will you read? ")
	}

	if *modeFlag != "" {
		mode = *modeFlag
	} else {
		// no flag input
		mode = handler.Ask("Which mode will you read? (all / line) ")
	}

	ascVal := *ascFlag

	// Trimming Space in file name text
	filePath = strings.TrimSpace(filePath)

	if filePath == "" {
		log.Fatal("No Input")
	}

	processFileData(filePath, ascVal, mode)
}
