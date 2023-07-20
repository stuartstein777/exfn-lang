package frontend

func Repl() {
	interpret([]rune("6+7"))
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
