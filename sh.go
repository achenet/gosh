package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type shell struct {
	variables map[string]string
}

func (s *shell) setVar(k, v string) {
	s.variables[k] = v
}

func (s *shell) processInput(input string) {
	// split line with spaces
	args := strings.Split(input, " ")
	
	// Swap this out for a switch statement,
	// or maybe a map of works to functions
	switch args[0] {
	case "set":
		s.setVar(args[1], args[2])
		fmt.Printf("set variable %s to value %s\n", args[1], args[2])
	case "show":
		fmt.Println(s.variables)
	default:
		fmt.Println("unrecognized command")
	}

}

func main() {
	fmt.Println("gosh!")
	env := shell{
		variables: make(map[string]string),
	}
	for {
		// Read std input
		r := bufio.NewReader(os.Stdin)
		// ReadLine will NOT include the \n character at the end of the line.
		l, _, err := r.ReadLine()
		handle(err)
		s := string(l)
		fmt.Println(">", s)
		// Process input
		env.processInput(s)

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
	return s == "exit"
}