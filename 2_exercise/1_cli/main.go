package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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

	// Read file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	info, _ := file.Stat()
	totalSize := info.Size()

	progressChan := make(chan int64, 10)

	pReader := &progressReader{
		Reader: file,
		Total:  totalSize,
		pChan:  progressChan,
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for current := range progressChan {
			percent := float64(current) / float64(totalSize) * 100

			if percent > 100 {
				percent = 100
			}

			fmt.Fprintf(os.Stderr, "\rprocessing: %.2f%% (%d/%d)", percent, current, pReader.Total)

		}
	}()

	processFileData(pReader, ascVal, mode)
	close(progressChan)

	wg.Wait()
	fmt.Println("\nFinished")
}
