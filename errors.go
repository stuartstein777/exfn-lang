package main

import "fmt"

func ReportError(line int, message string) {
	fmt.Printf("[line %d] %s\n", line, message)
}

type StringReadingError struct {
	Line                 int
	Column               int
	Message              string
	UnclosedStringLength int
}
