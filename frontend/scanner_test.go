package frontend

import (
	"fmt"
	"testing"
)

func TestPeek(t *testing.T) {
	InitScanner("hello")

	c := peek()

	if c != 'h' {
		t.Errorf("Expected true, got false")
	}

	scanner.Current = 3

	c = peek()
	if c != 'l' {
		t.Errorf("Expected true, got false")
	}
}

func TestIsDigit(t *testing.T) {

	for _, c := range "0123456789" {
		if !IsDigit(c) {
			t.Errorf(fmt.Sprintf("Expected true, got false. %v should be alpha.", c))
		}
	}

	if IsDigit('a') {
		t.Errorf("Expected false, got true. a is not digit")
	}
}

func TestIsAlpha(t *testing.T) {
	for _, c := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_" {
		if !IsAlpha(c) {
			t.Errorf("Expected true, got false. %v should be alpha.", c)
		}
	}
}

func TestAdvance(t *testing.T) {
	InitScanner("hello")

	c := advance()

	if c != 'h' {
		t.Errorf("Expected 'h', got %v", c)
	}

	c = advance()

	if c != 'e' {
		t.Errorf("Expected 'e', got %v", c)
	}

	c = peek()

	if c != 'l' {
		t.Errorf("Expected 'l', got %v", c)
	}
}

func TestSkipWhiteSpace(t *testing.T) {
	InitScanner(" \t\r\nx")

	SkipWhitespace()

	c := peek()

	if c != 'x' {
		t.Errorf("Expected 'x', got %v", c)
	}
}

func TestReadString(t *testing.T) {
	InitScanner(`func foo(){var x = "hello, world";}`)
	scanner.Current = 20
	scanner.Start = 20
	_, token := readString()

	if token.Type != TOKEN_STRING {
		t.Errorf("Expected %v, got %v", TOKEN_STRING, token.Type)
	}

	if token.Length != 12 {
		t.Errorf("Expected length 12, got %v", token.Length)
	}

	if token.Start != 20 {
		t.Errorf("Expected start 20, got %v", token.Start)
	}

	if token.Line != 1 {
		t.Errorf("Expected line 1, got %v", token.Line)
	}

	str := string(scanner.Source[token.Start : token.Start+token.Length])

	if str != "hello, world" {
		t.Errorf("Expected 'hello, world', got %v", str)
	}

}

func TestReadString2(t *testing.T) {
	InitScanner(`x = "hello, world"`)
	scanner.Start = 5
	scanner.Current = 5
	_, token := readString()

	if token.Type != TOKEN_STRING {
		t.Errorf("Expected %v, got %v", TOKEN_STRING, token.Type)
	}

	if token.Length != 12 {
		t.Errorf("Expected length 12, got %v", token.Length)
	}

	if token.Start != 5 {
		t.Errorf("Expected start 5, got %v", token.Start)
	}

	if token.Line != 1 {
		t.Errorf("Expected line 1, got %v", token.Line)
	}

	str := string(scanner.Source[token.Start : token.Start+token.Length])

	if str != "hello, world" {
		t.Errorf("Expected 'hello, world', got %v", str)
	}
}

func TestReadStringExpectUnterminatedStringError(t *testing.T) {
	InitScanner(`x = "hello, world`)
	scanner.Start = 5
	scanner.Current = 5
	errorToken, _ := readString()

	if errorToken.Type != TOKEN_ERROR {
		t.Errorf("Expected %v, got %v", TOKEN_ERROR, errorToken.Type)
	}

	if errorToken.Message != "Unterminated string." {
		t.Errorf("Expected %v, got %v", "Unterminated string.", errorToken.Message)
	}
}

func TestReadNumber(t *testing.T) {
	InitScanner("x = 1234;")
	scanner.Current = 4
	scanner.Start = 4
	token := readNumber()

	if token.Type != TOKEN_NUMBER {
		t.Errorf("Expected %v, got %v", TOKEN_NUMBER, token.Type)
	}

	if token.Length != 4 {
		t.Errorf("Expected length 4, got %v", token.Length)
	}

	if token.Start != 4 {
		t.Errorf("Expected start 4, got %v", token.Start)
	}

	if token.Line != 1 {
		t.Errorf("Expected line 1, got %v", token.Line)
	}

	num := string(scanner.Source[token.Start : token.Start+token.Length])

	if num != "1234" {
		t.Errorf("Expected '1234', got %v", num)
	}
}

func TestReadNumberNumberEndsAtEndOfSource(t *testing.T) {
	InitScanner("x = 1234")
	scanner.Start = 4
	scanner.Current = 4
	token := readNumber()

	if token.Type != TOKEN_NUMBER {
		t.Errorf("Expected %v, got %v", TOKEN_NUMBER, token.Type)
	}

	if token.Length != 4 {
		t.Errorf("Expected length 4, got %v", token.Length)
	}

	if token.Start != 4 {
		t.Errorf("Expected start 4, got %v", token.Start)
	}

	if token.Line != 1 {
		t.Errorf("Expected line 1, got %v", token.Line)
	}

	num := string(scanner.Source[token.Start : token.Start+token.Length])

	if num != "1234" {
		t.Errorf("Expected '1234', got %v", num)
	}
}

func TestReadNumberWithDecimal(t *testing.T) {
	InitScanner("x = 1234.567;")
	scanner.Start = 4
	scanner.Current = 4
	token := readNumber()

	if token.Type != TOKEN_NUMBER {
		t.Errorf("Expected %v, got %v", TOKEN_NUMBER, token.Type)
	}

	if token.Length != 8 {
		t.Errorf("Expected length 8, got %v", token.Length)
	}

	if token.Start != 4 {
		t.Errorf("Expected start 4, got %v", token.Start)
	}

	if token.Line != 1 {
		t.Errorf("Expected line 1, got %v", token.Line)
	}

	num := string(scanner.Source[token.Start : token.Start+token.Length])

	if num != "1234.567" {
		t.Errorf("Expected '1234.567', got %v", num)
	}
}

func TestScanToken(t *testing.T) {
	InitScanner("1234.567 + 789 = 987654;")
	_, token := ScanToken()

	tk := string(scanner.Source[token.Start : token.Start+token.Length])

	if token.Start != 0 {
		t.Errorf("Expected start 0, got %v", token.Start)
	}

	if token.Length != 8 {
		t.Errorf("Expected length 8, got %v", token.Length)
	}
	if tk != "1234.567" {
		t.Errorf("Expected '1234.567', got _%v_", tk)
	}

	_, token = ScanToken()

	if token.Start != 9 {
		t.Errorf("Expected start 9, got %v", token.Start)
	}

	tk = string(scanner.Source[token.Start : token.Start+token.Length])

	if tk != "+" {
		t.Errorf("Expected '+', got _%v_", tk)
	}

	_, token = ScanToken()
	tk = string(scanner.Source[token.Start : token.Start+token.Length])

	if tk != "789" {
		t.Errorf("Expected '789', got _%v_", tk)
	}

	_, token = ScanToken()
	tk = string(scanner.Source[token.Start : token.Start+token.Length])

	if tk != "=" {
		t.Errorf("Expected '=', got _%v_", tk)
	}

	_, token = ScanToken()
	tk = string(scanner.Source[token.Start : token.Start+token.Length])

	if tk != "987654" {
		t.Errorf("Expected '987654', got _%v_", tk)
	}

	_, token = ScanToken()

	tk = string(scanner.Source[token.Start : token.Start+token.Length])

	if tk != ";" {
		t.Errorf("Expected ';', got %v", tk)
	}
}
