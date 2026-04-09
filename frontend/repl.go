package frontend

func Repl() {
	//interpret([]rune("6+7"))
	//interpret([]rune("((6+7*4)+2-1)"))
	//interpret([]rune("(6 / 2) * 3 + 4 / 2"))

	//interpret([]rune("-(5*2+3) - 6 / 3"))
	interpret([]rune("5+nil"))

	// for {
	// 	fmt.Printf("> ")

	// 	reader := bufio.NewReader(os.Stdin)

	// 	text, error := reader.ReadString('\n')

	// 	if error == nil {
	// 		line := []rune(text)
	// 		interpret(line)
	// 	} else {
	// 		fmt.Println(error)
	// 		fmt.Println("Exiting repl!")
	// 		break
	// 	}
	// }
}
