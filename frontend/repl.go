package frontend

import (
	"bufio"
	"fmt"
	"os"
)

func Repl() {
	for {
		fmt.Printf("> ")

		reader := bufio.NewReader(os.Stdin)

		text, error := reader.ReadString('\n')

		if error == nil {
			line := []rune(text)
			interpret(line)
		} else {
			fmt.Println(error)
			fmt.Println("Exiting repl!")
			break
		}
	}
}
