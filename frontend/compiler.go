package frontend

import (
	"fmt"

	h "github.com/stuartstein777/exfnlang/helpers"
)

func interpret(source []rune) int {
	return h.INTERPRET_OK
}

func RunFile(path string) int {
	source, error := ReadSource(path)

	if error != nil {
		fmt.Printf("An error occurred reading file:\n%s\n", error)
		return h.ERROR_READING_FILE
	} else {
		return interpret([]rune(source))
	}
}
