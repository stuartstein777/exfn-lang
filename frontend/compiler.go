package frontend

import (
	"fmt"

	h "github.com/stuartstein777/exfnlang/helpers"
	t "github.com/stuartstein777/exfnlang/types"
	vmn "github.com/stuartstein777/exfnlang/vm"
)

type Parser struct {
	Current   Token
	Previous  Token
	HadError  bool
	PanicMode bool
}

var parser Parser = Parser{}
var compilingChunk *t.Chunk = &t.Chunk{}

func interpret(source []rune) int {
	chunk := t.Chunk{}
	initChunk(&chunk)

	if !compile(source, &chunk) {
		return h.INTERPRET_COMPILE_ERROR
	}
	vm := vmn.VM{}
	vm.Chunk = &chunk
	vm.IP = 0
	vm.ResetStack()

	result := vmn.Run(&vm)

	return result
}

func compile(source []rune, chunk *t.Chunk) bool {
	InitScanner(source)
	compilingChunk = chunk
	parser.HadError = false
	parser.PanicMode = false
	advanceCompiler()
	expression()
	consume(TOKEN_EOF, "Expect end of expression.")
	endCompiler()
	return !parser.HadError
}

// TODO: To build this out, it will be the basis for syntax checking and reporting syntax errors.
func consume(tokenType TokenType, message string) {
	if parser.Current.Type == tokenType {
		advanceCompiler()
		return
	}

	errorAtCurrent(message)
}

func endCompiler() {
	emitReturn()
}

func emitReturn() {
	emitByte(h.OP_RETURN)
}

func emitByte(byte byte) {
	t.WriteToChunk(compilingChunk, byte, parser.Previous.Line)
}

func emitBytes(byte1 byte, byte2 byte) {
	emitByte(byte1)
	emitByte(byte2)
}

func advanceCompiler() {
	parser.Previous = parser.Current

	for {
		errorToken := ErrorToken{}
		errorToken, parser.Current = ScanToken()

		if errorToken != (ErrorToken{}) {
			break
		}

		errorAtCurrent(errorToken.Message)
	}
}

func errorAtCurrent(message string) {
	errorAt(&parser.Current, message)
}

func errorAt(token *Token, message string) {
	if parser.PanicMode {
		return
	}

	parser.PanicMode = true

	fmt.Printf("[line %d] Error", token.Line)

	if token.Type == TOKEN_EOF {
		fmt.Printf(" at end")
	} else if token.Type == TOKEN_ERROR {
		// Nothing.
	} else {
		fmt.Printf(" at '%s'", GetToken())
	}

	fmt.Printf(": %s\n", message)
	parser.HadError = true
}

func initChunk(chunk *t.Chunk) {
	chunk.Code = []byte{}
	chunk.Constants = []t.Value{}
	chunk.LineNumbers = []int{}
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
