package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type InputHandler struct {
	scanner *bufio.Scanner
}

func NewInputHandler() *InputHandler {
	return &InputHandler{
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (h *InputHandler) Ask(prompt string) string {
	fmt.Printf("%v > ", prompt)
	if h.scanner.Scan() {
		return strings.TrimSpace(h.scanner.Text())
	}
	return ""
}
