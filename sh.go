package main

import (
	"bufio"
	"fmt"
	"os"
)

// Initialize env
var env map[string]interface{}

func main() {
	fmt.Println("gosh!")
	for {
		// Read std input
		r := bufio.NewReader(os.Stdin)
		// ReadLine will NOT include the \n character at the end of the line.
		l, _, err := r.ReadLine()
		handle(err)
		s := string(l)
		fmt.Println(">", s)
		// Process input

		// Write std output

		// Exit
		if exit(s) {
			break
		}

	}
	fmt.Println("Bonne continuation!")
}

func handle(err error) {
	if err != nil {
		fmt.Println("")
		os.Exit(420)
	}
}

func exit(s string) bool {
	if s == "exit" {
		return true
	}
	return false
}
