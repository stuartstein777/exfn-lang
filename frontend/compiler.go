package frontend

import (
	"fmt"
	"strconv"

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

const (
	PREC_NONE       = iota
	PREC_ASSIGNMENT // =
	PREC_OR         // or
	PREC_AND        // and
	PREC_EQUALITY   // == !=
	PREC_COMPARISON // < > <= >=
	PREC_TERM       // + -
	PREC_FACTOR     // * /
	PREC_UNARY      // ! -
	PREC_CALL       // . () []
	PREC_PRIMARY    // literals, identifiers, this, super
)

type ParseFn func()

type ParseRule interface {
	Prefix() ParseFn
	Infix() ParseFn
	Precedence() int
}

var parser Parser = Parser{}
var compilingChunk *t.Chunk = &t.Chunk{}

func currentChunk() *t.Chunk {
	return compilingChunk
}

func Unary() {
	fmt.Printf("In compiler.unary()\n")
	operatorType := parser.Previous.Type

	// Compile the operand.
	parsePrecedence(PREC_UNARY)

	// Emit the operator instruction.
	switch operatorType {
	case TOKEN_MINUS:
		emitByte(h.OP_NEGATE)
	default:
		return // Unreachable.
	}
}

func Binary() {
	fmt.Printf("In compiler.binary()\n")
	operatorType := parser.Previous.Type
	rule := GetRule(operatorType)
	parsePrecedence(rule.Precedence() + 1)

	switch operatorType {
	case TOKEN_PLUS:
		emitByte(h.OP_ADD)
	case TOKEN_MINUS:
		emitByte(h.OP_SUBTRACT)
	case TOKEN_STAR:
		emitByte(h.OP_MULTIPLY)
	case TOKEN_SLASH:
		emitByte(h.OP_DIVIDE)
	default:
		return // Unreachable.
	}
}

var rules = []ParseRule{
	TOKEN_LEFT_PAREN:    LeftParen{},
	TOKEN_RIGHT_PAREN:   RightParen{},
	TOKEN_LEFT_BRACE:    LeftBrace{},
	TOKEN_RIGHT_BRACE:   RightBrace{},
	TOKEN_COMMA:         Comma{},
	TOKEN_DOT:           Dot{},
	TOKEN_MINUS:         Minus{},
	TOKEN_PLUS:          Plus{},
	TOKEN_SEMICOLON:     Semicolon{},
	TOKEN_SLASH:         Slash{},
	TOKEN_STAR:          Star{},
	TOKEN_BANG:          Bang{},
	TOKEN_BANG_EQUAL:    BangEqual{},
	TOKEN_EQUAL:         Equal{},
	TOKEN_EQUAL_EQUAL:   EqualEqual{},
	TOKEN_GREATER:       Greater{},
	TOKEN_GREATER_EQUAL: GreaterEqual{},
	TOKEN_LESS:          Less{},
	TOKEN_LESS_EQUAL:    LessEqual{},
	TOKEN_IDENTIFIER:    Identifier{},
	TOKEN_STRING:        String{},
	TOKEN_NUMBER:        NumberPrec{},
	TOKEN_AND:           And{},
	TOKEN_CLASS:         Class{},
	TOKEN_ELSE:          Else{},
	TOKEN_FALSE:         False{},
	TOKEN_FOR:           For{},
	TOKEN_FUN:           Fun{},
	TOKEN_IF:            If{},
	TOKEN_NIL:           Nil{},
	TOKEN_OR:            Or{},
	TOKEN_PRINT:         Print{},
	TOKEN_RETURN:        Return{},
	TOKEN_SUPER:         Super{},
	TOKEN_TRUE:          True{},
	TOKEN_VAR:           Var{},
	TOKEN_WHILE:         While{},
	TOKEN_ERROR:         Error{},
	TOKEN_EOF:           EOF{},
}

func Number() {
	fmt.Printf("In compiler.number()\n")
	token := string(scanner.Source[parser.Previous.Start : parser.Previous.Start+parser.Previous.Length])
	value, _ := strconv.ParseFloat(token, 32)
	fmt.Printf("number:: value = %f\n", value)
	// what to do on error here ?
	emitConstant(value)
}

func GetRule(tokenType TokenType) ParseRule {
	fmt.Printf("In compiler.getRule()\n")
	fmt.Printf("compiler.getRule :: precedence = %v\n", rules[tokenType].Precedence())
	return rules[tokenType]
}

func parsePrecedence(precedence int) {
	fmt.Printf("In compiler.parsePrecedence()\n")
	fmt.Printf("compiler.parsePrecedence :: precedence = %d\n", precedence)
	advanceCompiler()
	fmt.Printf("compiler.parsePrecedence :: parser.previous.type = %d\n", parser.Previous.Type)
	prefixRule := GetRule(parser.Previous.Type).Prefix
	if prefixRule() == nil {
		errorAtCurrent("Expect expression.")
		return
	}

	prefixRule()()

	for {
		currentPrecedence := GetRule(parser.Current.Type).Precedence()
		if precedence > currentPrecedence {
			break
		}
		fmt.Printf("compiler.parsePrecedence :: Higher precedence: precedence = %d, currentPrecedence = %d\n", precedence, currentPrecedence)
		advanceCompiler()
		infixRule := GetRule(parser.Previous.Type).Infix
		infixRule()()
	}
}

func expression() {
	fmt.Printf("In compiler.expression()\n")
	parsePrecedence(PREC_ASSIGNMENT)
}

func interpret(source []rune) int {
	fmt.Printf("In compiler.interpret()\n")
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
	fmt.Printf("In compiler.compile()\n")
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
	fmt.Printf("In compiler.consume()\n")
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
	fmt.Printf("In compiler.emitReturn()\n")
	emitByte(h.OP_RETURN)
}

func emitByte(byte byte) {
	fmt.Printf("In compiler.emitByte()\n")
	fmt.Printf("Emitting byte: %d\n", byte)
	t.WriteToChunk(compilingChunk, byte, parser.Previous.Line)
}

func emitBytes(byte1 byte, byte2 byte) {
	fmt.Printf("In compiler.emitBytes()\n")
	emitByte(byte1)
	emitByte(byte2)
}

func makeConstant(value t.Value) byte {
	fmt.Printf("In compiler.makeConstant()\n")

	constant := t.AddConstant(compilingChunk, value, parser.Current.Line)
	if constant > 255 { //TODO: Allow this to be higher.
		errorAtCurrent("Too many constants in one chunk.")
		return 0
	}

	return byte(constant)
}

func emitConstant(value t.Value) {
	fmt.Printf("In compiler.emitConstant()\n")
	emitBytes(h.OP_CONSTANT, makeConstant(value))
	//TODO: Call WriteCOnstantToChunk
	//TODO: How do I know the index ?
	//t.WriteConstantToChunk(compilingChunk, h.OP_CONSTANT, int(len(compilingChunk.Constants)-1), parser.Current.Line)
}

func advanceCompiler() {
	fmt.Printf("In compiler.advance()\n")
	//fmt.Printf("advanceCompiler:: parser.Current:: %v\n", parser.Current)
	parser.Previous = parser.Current

	for {
		errorToken := ErrorToken{}
		errorToken, parser.Current = ScanToken()

		if errorToken == (ErrorToken{}) {
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
