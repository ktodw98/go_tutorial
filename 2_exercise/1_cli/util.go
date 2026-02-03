package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func processFileData(filePath string, ascVal bool, mode string) {

	// Read file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	switch mode {
	case "all":
		fileScanner.Split(bufio.ScanRunes)
	case "line":
		fileScanner.Split(bufio.ScanLines)
	default:
		fileScanner.Split(bufio.ScanLines)
	}

	for fileScanner.Scan() {
		data := fileScanner.Text()
		if !ascVal {
			data = reverseString(data)
		}

		fmt.Printf(data)
	}

	if err := fileScanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error while Scanning:", err)
	}

	fmt.Println() // Prevent '%' letter end of output
}

func reverseString(data string) string {
	runes := []rune(data)
	n := len(runes)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
