package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"sync"
)

func main() {
	// Define Flag
	fileFlag := flag.String("file", "example-big.txt", "File route")
	ascFlag := flag.Bool("asc", true, "whether print out in ascending order")
	modeFlag := flag.String("mode", "line", "all, line")

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

	progressChan := make(chan int64)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		processFileData(filePath, ascVal, mode, progressChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		var totalRead int64
		for n := range progressChan {
			totalRead += n
			fmt.Printf("Total Read: %v", totalRead)
			fmt.Println()
		}

	}()

	wg.Wait()
}
