package main

import "fmt"

// TODO: Take in col.
// Count back Source, to get col on that line.
func ReportError(line, col int, message, source string) {
	fmt.Printf("[line %d] %s\n", line, message)
}

type StringReadingError struct {
	Line                 int
	Column               int
	Message              string
	UnclosedStringLength int
}

type NumberReadingError struct {
	Line       int
	Column     int
	Message    string
	TokensRead int
}
