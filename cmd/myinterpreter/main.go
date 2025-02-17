package main

import (
	"fmt"
	"os"
	"unicode"
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
	line := 1

	for i := 0; i < len(fileContents); i++ {
		lex := fileContents[i]

		if lex == '\n' {
			line++
			continue
		}

		if unicode.IsSpace(rune(lex)) {
			continue
		}

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
		case '<':
			if i+1 < len(fileContents) && fileContents[i+1] == '=' {
				fmt.Println("LESS_EQUAL <= null")
				i++
			} else {
				fmt.Println("LESS < null")
			}
		case '>':
			if i+1 < len(fileContents) && fileContents[i+1] == '=' {
				fmt.Println("GREATER_EQUAL >= null")
				i++
			} else {
				fmt.Println("GREATER > null")
			}
		case '/':
			if i+1 < len(fileContents) && fileContents[i+1] == '/' {
				for i < len(fileContents) && fileContents[i] != '\n' {
					i++
				}
				i--
				continue
			} else {
				fmt.Println("SLASH / null")
			}
		case '"':
			newIndex, errOccurred := scanString(fileContents, i, line)
			if errOccurred {
				hasError = true
				goto end
			}
			i = newIndex
		default:
			fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %c\n", line, lex)
			hasError = true
		}
	}

end:
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

func scanString(contents []byte, start int, line int) (int, bool) {
	i := start + 1
	for i < len(contents) && contents[i] != '"' {
		if contents[i] == '\n' {
			fmt.Fprintf(os.Stderr, "[line %d] Error: Unterminated string.\n", line)
			return i, true
		}
		i++
	}

	if i >= len(contents) {
		fmt.Fprintf(os.Stderr, "[line %d] Error: Unterminated string.\n", line)
		return i, true
	}

	literal := string(contents[start+1 : i])
	lexeme := string(contents[start : i+1])

	fmt.Printf("STRING %s %s\n", lexeme, literal)

	return i, false
}
