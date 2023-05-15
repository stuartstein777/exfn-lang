package main

import (
	"fmt"
	"strconv"
)

func peek(source string, current int, sourceLen int, expected byte) bool {
	if current >= sourceLen {
		return false
	}
	return source[current+1] == expected
}

/*
Valid: `123`, `123.456`
Invalid: `123.`, `.123`, `123.456.789`, `123x` (where x is anything not . or digit)

If we hit 2 decimal points, we should return an error and the length i so far, then
the . can be read as a separate token
*/
func ReadNumericLiteral(source string, current, line, sourceLen int) (Token, NumberReadingError) {
	// keep reading as long as we have digits or .
	// but only allow one .
	literal := ""
	i := current
	hasDecimalPlace := false
	for i < sourceLen {
		if IsDigit(source[i]) {
			literal += string(source[i])
			i++
		} else if source[i] == '.' {
			// We've read a decimal place, but we can't have two
			if hasDecimalPlace {
				return Token{},
					NumberReadingError{
						line,
						current,
						fmt.Sprintf("Invalid number starting at line: %d, col: %d", line, i-1),
						i - current}
			} else {
				hasDecimalPlace = true
				literal += string(source[i])
				i++
			}
		} else {
			break
		}
	}
	nl, _ := strconv.ParseFloat(literal, 64)
	return Token{NUMBER, literal, nl, line, i - current}, NumberReadingError{}
}

func ReadStringLiteral(source string, current, line, sourceLen int) (Token, StringReadingError) {
	literal := ""
	i := current + 1
	closed := false
	for i < sourceLen {
		if source[i] == '"' {
			closed = true
			break
		} else if source[i] == '\n' {
			break
		} else {
			literal += string(source[i])
			i++
		}
	}

	if !closed {
		return Token{},
			StringReadingError{
				line,
				current,
				fmt.Sprintf("Unterminated string starting at line: %d, col: %d", line, current),
				len(literal)}
	}

	return Token{STRING, literal, literal, line, i - current + 1}, StringReadingError{}
}

func ReadIdentifier(source string, current, line, sourceLen int) (Token, int) {

}

func IsDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func IsAlpha(char byte) bool {
	return (char >= 'a' && char <= 'z') ||
		(char >= 'A' && char <= 'Z') ||
		char == '_'
}

func ScanToken(source string, current, line int) (Token, int, string) {
	l := len(source)
	switch source[current] {
	case '(':
		return Token{LEFT_PAREN, "(", nil, line, 1}, 1, ""
	case ')':
		return Token{RIGHT_PAREN, ")", nil, line, 1}, 1, ""
	case '{':
		return Token{LEFT_BRACE, "{", nil, line, 1}, 1, ""
	case '}':
		return Token{RIGHT_BRACE, "}", nil, line, 1}, 1, ""
	case '[':
		return Token{LEFT_SQUARE_BRACKET, "[", nil, line, 1}, 1, ""
	case ']':
		return Token{RIGHT_SQUARE_BRACKET, "]", nil, line, 1}, 1, ""
	case ',':
		return Token{COMMA, ",", nil, line, 1}, 1, ""
	case '.':
		return Token{DOT, ".", nil, line, 1}, 1, ""
	case '=':
		// need to look ahead to see if it's ==
		if peek(source, current, l, '=') {
			return Token{EQUAL_EQUAL, "==", nil, line, 2}, 2, ""
		} else {
			return Token{EQUAL, "=", nil, line, 1}, 1, ""
		}
	case '-':
		return Token{MINUS, "-", nil, line, 1}, 1, ""
	case '+':
		return Token{PLUS, "+", nil, line, 1}, 1, ""
	case ';':
		return Token{SEMICOLON, ";", nil, line, 1}, 1, ""
	case '*':
		return Token{STAR, "*", nil, line, 1}, 1, ""
	case '!':
		// need to look ahead to see if it's !=
		if peek(source, current, l, '=') {
			return Token{BANG_EQUAL, "!=", nil, line, 2}, 2, ""
		} else {
			return Token{BANG, "!", nil, line, 1}, 1, ""
		}
	case '<':
		// need to look ahead to see if it's <=
		if peek(source, current, l, '=') {
			return Token{LESS_EQUAL, "<=", nil, line, 2}, 2, ""
		} else {
			return Token{LESS, "<", nil, line, 1}, 1, ""
		}
	case '>':
		// need to look ahead to see if it's >=
		if peek(source, current, l, '=') {
			return Token{GREATER_EQUAL, ">=", nil, line, 2}, 2, ""
		} else {
			return Token{GREATER, ">", nil, line, 1}, 1, ""
		}
	case '/':
		// need to look ahead to see if it's a comment
		if peek(source, current, l, '/') {
			// a comment goes until the end of the line
			n := current
			for source[n] != '\n' && n < l {
				n += 1
			}
			return Token{}, n - current, ""
		} else {
			return Token{SLASH, "/", nil, line, 1}, 1, ""
		}
	// string literals
	case '"':
		token, err := ReadStringLiteral(source, current, line, l)
		if err != (StringReadingError{}) {
			ReportError(line, current, err.Message, source)
			return Token{}, err.UnclosedStringLength + 1, ""
		}
		return token, token.Length, ""
	case '\r':
		fallthrough
	case '\n':
		return Token{NEWLINE, "", nil, line, 1}, 1, ""
	default:
		if IsDigit(source[current]) {
			token, err := ReadNumericLiteral(source, current, line, l)
			if err != (NumberReadingError{}) {
				ReportError(line, current, err.Message, source)
				return Token{}, err.TokensRead, ""
			}
			return token, token.Length, ""
		} else if IsAlpha(source[current]) {
			token, err := ReadIdentifier(source, current, line, l)
		} else {
			return Token{}, 1, "" //TODO: Error handling for invalid lexemes! e.g. rogue &^%$ etc
		}
	}
}

func ScanTokens(source string) []Token {
	var tokens []Token
	sourceLen := len(source)
	current := 0
	line := 1

	// current is the actual token we are looking at.
	for current < sourceLen {
		// We are at the beginning of the next lexeme.
		token, l, _ := ScanToken(source, current, line)

		// ignore new lines, but increment line number for error reporting
		if token != (Token{}) {
			if token.Type != NEWLINE {
				tokens = append(tokens, token)
			} else {
				line += 1
			}
		}

		current += l
		//fmt.Printf("Read %s with length %d\n. Current [%d]", token.Lexeme, l, current)

		if token.Type == EOF {
			break
		}
	}

	return tokens
}
