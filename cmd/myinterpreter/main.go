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

	for _, lex := range fileContents {
		switch lex {
		case '(':
			fmt.Printf("LEFT_PAREN %c null\n", lex)
		case ')':
			fmt.Printf("RIGHT_PAREN %c null\n", lex)
		case '{':
			fmt.Printf("LEFT_BRACE %c null\n", lex)
		case '}':
			fmt.Printf("RIGHT_BRACE %c null\n", lex)
		case ',':
			fmt.Printf("COMMA %c null\n", lex)
		case '.':
			fmt.Printf("DOT %c null\n", lex)
		case '-':
			fmt.Printf("MINUS %c null\n", lex)
		case '+':
			fmt.Printf("PLUS %c null\n", lex)
		case ';':
			fmt.Printf("SEMICOLON %c null\n", lex)
		case '*':
			fmt.Printf("STAR %c null\n", lex)
		default:
			fmt.Printf("[line 1] Error: Unexpected character: %c\n", lex)
		}
	}
	fmt.Println("EOF  null")
}
