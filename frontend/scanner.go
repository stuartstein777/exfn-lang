package frontend

type Scanner struct {
	Source       []rune
	SourceLength int
	Current      int
	Start        int
	Line         int
}

var (
	scanner Scanner = Scanner{[]rune{}, 0, 0, 0, 1}
	S               = scanner
)

func GetToken() string {
	return string(scanner.Source[scanner.Start:scanner.Current])
}
func peek() rune {
	return scanner.Source[scanner.Current]
}

// Read the next token and advance the scanner.
func advance() rune {
	scanner.Current++
	return rune(scanner.Source[scanner.Current-1])

}

func IsDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func IsAlpha(char rune) bool {
	return (char >= 'a' && char <= 'z') ||
		(char >= 'A' && char <= 'Z') ||
		char == '_'
}

func MakeToken(tokenType TokenType) Token {
	return Token{
		tokenType,
		scanner.Start,
		scanner.Current - scanner.Start,
		scanner.Line,
	}
}

func SkipWhitespace() {
	for {
		c := peek()

		switch c {
		case ' ', '\r', '\t':
			advance()
		case '\n':
			scanner.Line++
			advance()
		case '/':
			if peekNext() == '/' {
				for peek() != '\n' && !isAtEnd() {
					advance()
				}
			}
		default:
			return
		}
	}
}

func isAtEnd() bool {
	return scanner.Current >= scanner.SourceLength
}

func readString() (ErrorToken, Token) {
	for {
		if isAtEnd() {
			return errorToken("Unterminated string."), Token{}
		}

		if peek() == '"' {
			break
		}

		if peek() == '\n' {
			scanner.Line++
		}
		advance()
	}

	if isAtEnd() {
		return errorToken("Unterminated string."), Token{}
	}
	token := MakeToken(TOKEN_STRING)

	advance() // to advance past the closing "
	return ErrorToken{}, token
}

func peekNext() rune {
	if scanner.Current+1 >= scanner.SourceLength {
		return '\000'
	}

	return scanner.Source[scanner.Current+1]
}

func readNumber() Token {
	for {
		if isAtEnd() || !IsDigit(peek()) {
			break
		}
		advance()
	}

	if isAtEnd() {
		return MakeToken(TOKEN_NUMBER)
	}
	if peek() == '.' && IsDigit(peekNext()) {
		advance()

		for {
			if isAtEnd() || !IsDigit(peek()) {
				break
			}
			advance()
		}
	}

	return MakeToken(TOKEN_NUMBER)
}

func errorToken(message string) ErrorToken {
	return ErrorToken{
		TOKEN_ERROR,
		message,
		scanner.Line,
	}
}

func InitScanner(source string) {
	scanner.Source = []rune(source)
	scanner.SourceLength = len(source)
	scanner.Current = 0
	scanner.Start = 0
	scanner.Line = 1
}

func ScanToken() (ErrorToken, Token) {
	SkipWhitespace()
	scanner.Start = scanner.Current

	if scanner.Current >= scanner.SourceLength {
		return ErrorToken{}, MakeToken(TOKEN_EOF)
	}

	c := advance()

	if IsDigit(c) {
		return ErrorToken{}, readNumber()
	}

	switch c {
	case '(':
		return ErrorToken{}, MakeToken(TOKEN_LEFT_PAREN)
	case ')':
		return ErrorToken{}, MakeToken(TOKEN_RIGHT_PAREN)
	case '{':
		return ErrorToken{}, MakeToken(TOKEN_LEFT_BRACE)
	case '}':
		return ErrorToken{}, MakeToken(TOKEN_RIGHT_BRACE)
	case ',':
		return ErrorToken{}, MakeToken(TOKEN_COMMA)
	case '.':
		return ErrorToken{}, MakeToken(TOKEN_DOT)
	case '-':
		return ErrorToken{}, MakeToken(TOKEN_MINUS)
	case '+':
		return ErrorToken{}, MakeToken(TOKEN_PLUS)
	case '*':
		return ErrorToken{}, MakeToken(TOKEN_STAR)
	case ';':
		return ErrorToken{}, MakeToken(TOKEN_SEMICOLON)
	// = can be == or =
	case '=':
		if peek() == '=' {
			advance()
			return ErrorToken{}, MakeToken(TOKEN_EQUAL_EQUAL)
		} else {
			return ErrorToken{}, MakeToken(TOKEN_EQUAL)
		}
	// ! can be != or !
	case '!':
		if peek() == '=' {
			advance()
			return ErrorToken{}, MakeToken(TOKEN_BANG_EQUAL)
		} else {
			return ErrorToken{}, MakeToken(TOKEN_BANG)
		}
	case '/':
		return ErrorToken{}, MakeToken(TOKEN_SLASH)
	}

	return ErrorToken{}, Token{}
}
