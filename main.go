package main

import (
	"fmt"
	"io"
	"os"
)

// TODO: add error handling, return error or nil if it ran fine.
func run(source string) {
	// get tokens from scanner
	tokens := ScanTokens(source)

	curLine := 0
	for _, token := range tokens {
		if token.Line > curLine {
			curLine = token.Line
			fmt.Println()
			fmt.Printf("%d:", token.Line)
		}
		if token.Type == STRING {
			fmt.Printf("\"%s\"", token.Literal)
		} else {
			fmt.Printf("`%s` ", token.Lexeme)
		}
	}
	fmt.Println()

}

func runFile(path string) {
	source, err := os.ReadFile(path)

	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}
	run(string(source))
}

func runPrompt() {
	for {
		fmt.Print("> ")
		var line string
		_, err := fmt.Scanln(&line)

		if err != nil {
			if err == io.EOF {
				os.Exit(1)
			} else if err.Error() == "unexpected newline" {
				continue
			}
			fmt.Println("Error reading input")
			os.Exit(1)
		}

		if line == "" {
			continue
		}
		if line == ":quit" {
			return
		}

		run(line)
	}
}

func main() {
	// args := os.Args[1:] // 1: to exclude the program name
	// numberOfArgs := len(args)
	// if numberOfArgs > 1 {
	// 	fmt.Println("Usage: jlox [script]")
	// } else if numberOfArgs == 1 {
	// 	runFile(args[0])
	// } else {
	// 	runPrompt()
	// }

	runFile("test.xfn")
}
