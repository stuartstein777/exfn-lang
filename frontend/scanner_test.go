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

func TestScanTokenWithNewLinesAndComments(t *testing.T) {
	InitScanner(
		`1234.567 * 789 
	= 987654;
	// this is a comment
	987-654=333;`)
	_, token := ScanToken()
	tk := string(scanner.Source[token.Start : token.Start+token.Length])
	if tk != "1234.567" {
		t.Errorf("Expected '1234.567', got _%v_", tk)
	}

	_, token = ScanToken()
	tk = string(scanner.Source[token.Start : token.Start+token.Length])
	if tk != "*" {
		t.Errorf("Expected '*', got _%v_", tk)
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
		t.Errorf("Expected ';', got _%v_", tk)
	}

	//_, token = ScanToken()
	_, token = ScanToken()
	tk = string(scanner.Source[token.Start : token.Start+token.Length])
	if tk != "987" {
		t.Errorf("Expected '987', got _%v_", tk)
	}

	_, token = ScanToken()
	tk = string(scanner.Source[token.Start : token.Start+token.Length])
	if tk != "-" {
		t.Errorf("Expected '-', got _%v_", tk)
	}

	_, token = ScanToken()
	tk = string(scanner.Source[token.Start : token.Start+token.Length])
	if tk != "654" {
		t.Errorf("Expected '654', got _%v_", tk)
	}

	_, token = ScanToken()
	tk = string(scanner.Source[token.Start : token.Start+token.Length])
	if tk != "=" {
		t.Errorf("Expected '=', got _%v_", tk)
	}

	_, token = ScanToken()
	tk = string(scanner.Source[token.Start : token.Start+token.Length])
	if tk != "333" {
		t.Errorf("Expected '333', got _%v_", tk)
	}

	_, token = ScanToken()
	tk = string(scanner.Source[token.Start : token.Start+token.Length])
	if tk != ";" {
		t.Errorf("Expected ';', got _%v_", tk)
	}
}

func TestFindingAndKeyword(t *testing.T) {
	InitScanner("and")
	_, token := ScanToken()

	if token.Type != TOKEN_AND {
		t.Errorf("Expected And, got %v", token.Type)
	}

	tk := string(scanner.Source[token.Start : token.Start+token.Length])
	if tk != "and" {
		t.Errorf("Expected 'and', got _%v_", tk)
	}
}

func TestFindingFalseKeyword(t *testing.T) {
	InitScanner("false")
	_, token := ScanToken()

	if token.Type != TOKEN_FALSE {
		t.Errorf("Expected False, got %v", token.Type)
	}

	tk := string(scanner.Source[token.Start : token.Start+token.Length])
	if tk != "false" {
		t.Errorf("Expected 'false', got _%v_", tk)
	}
}

func TestFindingForKeyword(t *testing.T) {
	InitScanner("for")
	_, token := ScanToken()

	if token.Type != TOKEN_FOR {
		t.Errorf("Expected for, got %v", token.Type)
	}

	tk := string(scanner.Source[token.Start : token.Start+token.Length])
	if tk != "for" {
		t.Errorf("Expected 'for', got _%v_", tk)
	}
}

func TestFortReturnsIdentifier(t *testing.T) {
	InitScanner("fort")
	_, token := ScanToken()

	if token.Type != TOKEN_IDENTIFIER {
		t.Errorf("Expected fort, got %v", token.Type)
	}

	tk := string(scanner.Source[token.Start : token.Start+token.Length])
	if tk != "fort" {
		t.Errorf("Expected 'fort', got _%v_", tk)
	}
}

func TestFunReturnsFunKeyword(t *testing.T) {
	InitScanner("fun")
	_, token := ScanToken()

	if token.Type != TOKEN_FUN {
		t.Errorf("Expected fun, got %v", token.Type)
	}

	tk := string(scanner.Source[token.Start : token.Start+token.Length])
	if tk != "fun" {
		t.Errorf("Expected 'fun', got _%v_", tk)
	}
}

func TestOrReturnsOrKeyword(t *testing.T) {
	InitScanner("or")
	_, token := ScanToken()

	if token.Type != TOKEN_OR {
		t.Errorf("Expected or, got %v", token.Type)
	}

	tk := string(scanner.Source[token.Start : token.Start+token.Length])
	if tk != "or" {
		t.Errorf("Expected 'or', got _%v_", tk)
	}
}

func TestClassReturnsClassKeyword(t *testing.T) {
	InitScanner("class")
	_, token := ScanToken()

	if token.Type != TOKEN_CLASS {
		t.Errorf("Expected class, got %v", token.Type)
	}

	tk := string(scanner.Source[token.Start : token.Start+token.Length])
	if tk != "class" {
		t.Errorf("Expected 'class', got _%v_", tk)
	}
}

func TestClassesReturnsClassesIdentifier(t *testing.T) {
	InitScanner("classes")
	_, token := ScanToken()

	if token.Type != TOKEN_IDENTIFIER {
		t.Errorf("Expected classes, got %v", token.Type)
	}

	tk := string(scanner.Source[token.Start : token.Start+token.Length])
	if tk != "classes" {
		t.Errorf("Expected 'classes', got _%v_", tk)
	}
}

func TestClasReturnsClasIdentifier(t *testing.T) {
	InitScanner("clas")
	_, token := ScanToken()

	if token.Type != TOKEN_IDENTIFIER {
		t.Errorf("Expected clas, got %v", token.Type)
	}

	tk := string(scanner.Source[token.Start : token.Start+token.Length])
	if tk != "clas" {
		t.Errorf("Expected 'clas', got _%v_", tk)
	}
}

func TestIfReturnsIfIdentifier(t *testing.T) {
	InitScanner("if")
	_, token := ScanToken()

	if token.Type != TOKEN_IF {
		t.Errorf("Expected if, got %v", token.Type)
	}

	tk := string(scanner.Source[token.Start : token.Start+token.Length])
	if tk != "if" {
		t.Errorf("Expected 'if', got _%v_", tk)
	}
}

func TestPrintReturnsPrintIdentifier(t *testing.T) {
	InitScanner("print")
	_, token := ScanToken()

	if token.Type != TOKEN_PRINT {
		t.Errorf("Expected print, got %v", token.Type)
	}

	tk := string(scanner.Source[token.Start : token.Start+token.Length])
	if tk != "print" {
		t.Errorf("Expected 'print', got _%v_", tk)
	}
}

func TestReturnReturnsReturnIdentifier(t *testing.T) {
	InitScanner("return")
	_, token := ScanToken()

	if token.Type != TOKEN_RETURN {
		t.Errorf("Expected return, got %v", token.Type)
	}

	tk := string(scanner.Source[token.Start : token.Start+token.Length])
	if tk != "return" {
		t.Errorf("Expected 'return', got _%v_", tk)
	}
}

func TestSuperReturnsSuperIdentifier(t *testing.T) {
	InitScanner("super")
	_, token := ScanToken()

	if token.Type != TOKEN_SUPER {
		t.Errorf("Expected super, got %v", token.Type)
	}

	tk := string(scanner.Source[token.Start : token.Start+token.Length])
	if tk != "super" {
		t.Errorf("Expected 'super', got _%v_", tk)
	}
}

func TestThisReturnsThisIdentifier(t *testing.T) {
	InitScanner("this")
	_, token := ScanToken()

	if token.Type != TOKEN_THIS {
		t.Errorf("Expected this, got %v", token.Type)
	}

	tk := string(scanner.Source[token.Start : token.Start+token.Length])
	if tk != "this" {
		t.Errorf("Expected 'this', got _%v_", tk)
	}
}

func TestTrueReturnsTrueIdentifier(t *testing.T) {
	InitScanner("true")
	_, token := ScanToken()

	if token.Type != TOKEN_TRUE {
		t.Errorf("Expected true, got %v", token.Type)
	}

	tk := string(scanner.Source[token.Start : token.Start+token.Length])
	if tk != "true" {
		t.Errorf("Expected 'true', got _%v_", tk)
	}
}

func TestVarReturnsVarIdentifier(t *testing.T) {
	InitScanner("var")
	_, token := ScanToken()

	if token.Type != TOKEN_VAR {
		t.Errorf("Expected var, got %v", token.Type)
	}

	tk := string(scanner.Source[token.Start : token.Start+token.Length])
	if tk != "var" {
		t.Errorf("Expected 'var', got _%v_", tk)
	}
}

func TestWhileReturnsWhileIdentifier(t *testing.T) {
	InitScanner("while")
	_, token := ScanToken()

	if token.Type != TOKEN_WHILE {
		t.Errorf("Expected while, got %v", token.Type)
	}

	tk := string(scanner.Source[token.Start : token.Start+token.Length])
	if tk != "while" {
		t.Errorf("Expected 'while', got _%v_", tk)
	}
}

func TestScanningPortionOfSourceCode(t *testing.T) {
	InitScanner(`fun foo()
	{
		var x = 5;
		print(x);
	}`)

	_, token := ScanToken()
	if token.Type != TOKEN_FUN {
		t.Errorf("Expected fun, got %v", token.Type)
	}

	_, token = ScanToken()
	if token.Type != TOKEN_IDENTIFIER {
		t.Errorf("Expected foo, got %v", token.Type)
	}

	_, token = ScanToken()
	if token.Type != TOKEN_LEFT_PAREN {
		t.Errorf("Expected (, got %v", token.Type)
	}

	_, token = ScanToken()
	if token.Type != TOKEN_RIGHT_PAREN {
		t.Errorf("Expected ), got %v", token.Type)
	}

	_, token = ScanToken()
	if token.Type != TOKEN_LEFT_BRACE {
		t.Errorf("Expected {, got %v", token.Type)
	}

	_, token = ScanToken()
	if token.Type != TOKEN_VAR {
		t.Errorf("Expected var, got %v", token.Type)
	}

	_, token = ScanToken()
	if token.Type != TOKEN_IDENTIFIER {
		t.Errorf("Expected Identifer, got %v", token.Type)
	}

	_, token = ScanToken()
	if token.Type != TOKEN_EQUAL {
		t.Errorf("Expected =, got %v", token.Type)
	}

	_, token = ScanToken()
	if token.Type != TOKEN_NUMBER {
		t.Errorf("Expected number, got %v", token.Type)
	}

	_, token = ScanToken()
	if token.Type != TOKEN_SEMICOLON {
		t.Errorf("Expected ;, got %v", token.Type)
	}

	_, token = ScanToken()
	if token.Type != TOKEN_PRINT {
		t.Errorf("Expected print, got %v", token.Type)
	}

	_, token = ScanToken()
	if token.Type != TOKEN_LEFT_PAREN {
		t.Errorf("Expected (, got %v", token.Type)
	}

	_, token = ScanToken()
	if token.Type != TOKEN_IDENTIFIER {
		t.Errorf("Expected Identifier, got %v", token.Type)
	}

	_, token = ScanToken()
	if token.Type != TOKEN_RIGHT_PAREN {
		t.Errorf("Expected ), got %v", token.Type)
	}

	_, token = ScanToken()
	if token.Type != TOKEN_SEMICOLON {
		t.Errorf("Expected ;, got %v", token.Type)
	}

	_, token = ScanToken()
	if token.Type != TOKEN_RIGHT_BRACE {
		t.Errorf("Expected }, got %v", token.Type)
	}
}
