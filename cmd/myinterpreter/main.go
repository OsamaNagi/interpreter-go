package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	hasError := false

	for i := 0; i < len(fileContents); i++ {
		lex := fileContents[i]

		switch lex {
		case '(', ')', '{', '}', ',', '.', '-', '+', ';', '*':
			fmt.Printf("%s %c null\n", getTokenName(lex), lex)
		case '=':
			if i+1 < len(fileContents) && fileContents[i+1] == '=' {
				fmt.Println("EQUAL_EQUAL == null")
				i++
			} else {
				fmt.Println("EQUAL = null")
			}
		case '!':
			if i+1 < len(fileContents) && fileContents[i+1] == '=' {
				fmt.Println("BANG_EQUAL != null")
				i++
			} else {
				fmt.Println("BANG ! null")
			}
		default:
			fmt.Fprintf(os.Stderr, "[line 1] Error: Unexpected character: %c\n", lex)
			hasError = true
		}
	}

	fmt.Println("EOF  null")

	if hasError {
		os.Exit(65)
	} else {
		os.Exit(0)
	}
}

func getTokenName(ch byte) string {
	switch ch {
	case '(':
		return "LEFT_PAREN"
	case ')':
		return "RIGHT_PAREN"
	case '{':
		return "LEFT_BRACE"
	case '}':
		return "RIGHT_BRACE"
	case ',':
		return "COMMA"
	case '.':
		return "DOT"
	case '-':
		return "MINUS"
	case '+':
		return "PLUS"
	case ';':
		return "SEMICOLON"
	case '*':
		return "STAR"
	}
	return "UNKNOWN"
}
