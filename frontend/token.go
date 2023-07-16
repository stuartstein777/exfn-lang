package frontend

type TokenType int

const (
	// Single-character tokens.
	TOKEN_LEFT_PAREN TokenType = iota
	TOKEN_RIGHT_PAREN
	TOKEN_LEFT_BRACE
	TOKEN_RIGHT_BRACE
	TOKEN_LEFT_SQUARE_BRACKET
	TOKEN_RIGHT_SQUARE_BRACKET
	TOKEN_COMMA
	TOKEN_DOT
	TOKEN_MINUS
	TOKEN_PLUS
	TOKEN_SEMICOLON
	TOKEN_SLASH
	TOKEN_STAR
	TOKEN_PERCENT

	// One or two character tokens.
	TOKEN_BANG
	TOKEN_BANG_EQUAL
	TOKEN_EQUAL
	TOKEN_EQUAL_EQUAL
	TOKEN_GREATER
	TOKEN_GREATER_EQUAL
	TOKEN_LESS
	TOKEN_LESS_EQUAL

	// Literals.
	TOKEN_IDENTIFIER
	TOKEN_STRING
	TOKEN_NUMBER

	// Keywords.
	TOKEN_AND
	TOKEN_CLASS
	TOKEN_ELSE
	TOKEN_FALSE
	TOKEN_FUN
	TOKEN_FOR
	TOKEN_IF
	TOKEN_NIL
	TOKEN_OR
	TOKEN_PRINT
	TOKEN_RETURN
	TOKEN_SUPER
	TOKEN_THIS
	TOKEN_TRUE
	TOKEN_VAR
	TOKEN_WHILE

	TOKEN_NEWLINE

	TOKEN_EOF
	TOKEN_ERROR
)

type Token struct {
	Type   TokenType
	Start  int
	Length int
	Line   int
}

type ErrorToken struct {
	Type    TokenType
	Message string
	Line    int
}

// func (t Token) ToString() string {
// 	return fmt.Sprint(t.Type, " ", t.Lexeme, " ", t.Literal)
// }
