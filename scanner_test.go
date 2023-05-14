package main

import "testing"

func TestBasicScanToken(t *testing.T) {
	source := "var x = 5;"
	token, _, _ := ScanToken(source, 6)

	if token.Type != EQUAL {
		t.Errorf("Expected EQUAL, got %v", token.Type)
	}
}

func TestBasicScanTokenLookaheadWithBang(t *testing.T) {
	source := "y = x != 5;"
	token, _, _ := ScanToken(source, 6)

	if token.Type != BANG_EQUAL {
		t.Errorf("Expected BANG_EQUAL, got %v", token.Type)
	}
}

func TestBasicScanTokenNoLookaheadWithBang(t *testing.T) {
	source := "y = !x;"
	token, _, _ := ScanToken(source, 4)

	if token.Type != BANG {
		t.Errorf("Expected BANG, got %v", token.Type)
	}
}
