package frontend

/* LeftParen */
type LeftParen struct{}

func leftParenPrefixFunc() {
	expression()
	consume(TOKEN_RIGHT_PAREN, "Expect ')' after expression.")
}

func (g LeftParen) Prefix() ParseFn {
	return leftParenPrefixFunc
}

func (g LeftParen) Infix() ParseFn {
	return nil
}

func (g LeftParen) Precedence() int {
	return PREC_NONE
}

/* RightParen */
type RightParen struct{}

func (g RightParen) Prefix() ParseFn {
	return nil
}

func (g RightParen) Infix() ParseFn {
	return nil
}

func (g RightParen) Precedence() int {
	return PREC_NONE
}

/* LeftBrace */
type LeftBrace struct{}

func (g LeftBrace) Prefix() ParseFn {
	return nil
}

func (g LeftBrace) Infix() ParseFn {
	return nil
}

func (g LeftBrace) Precedence() int {
	return PREC_NONE
}

/* RightBrace */
type RightBrace struct{}

func (g RightBrace) Prefix() ParseFn {
	return nil
}

func (g RightBrace) Infix() ParseFn {
	return nil
}

func (g RightBrace) Precedence() int {
	return PREC_NONE
}

/* Comma */
type Comma struct{}

func (g Comma) Prefix() ParseFn {
	return nil
}

func (g Comma) Infix() ParseFn {
	return nil
}

func (g Comma) Precedence() int {
	return PREC_NONE
}

/* Dot */
type Dot struct{}

func (g Dot) Prefix() ParseFn {
	return nil
}

func (g Dot) Infix() ParseFn {
	return nil
}

func (g Dot) Precedence() int {
	return PREC_NONE
}

/* Minus */
type Minus struct{}

func (g Minus) Prefix() ParseFn {
	return Unary
}

func (g Minus) Infix() ParseFn {
	return Binary
}

func (g Minus) Precedence() int {
	return PREC_TERM
}

/* Plus */
type Plus struct{}

func (g Plus) Prefix() ParseFn {
	return nil
}

func (g Plus) Infix() ParseFn {
	return Binary
}

func (g Plus) Precedence() int {
	return PREC_TERM
}

/* Semicolon */
type Semicolon struct{}

func (g Semicolon) Prefix() ParseFn {
	return nil
}

func (g Semicolon) Infix() ParseFn {
	return nil
}

func (g Semicolon) Precedence() int {
	return PREC_NONE
}

/* Slash */
type Slash struct{}

func (g Slash) Prefix() ParseFn {
	return nil
}

func (g Slash) Infix() ParseFn {
	return Binary
}

func (g Slash) Precedence() int {
	return PREC_FACTOR
}

/* Star */
type Star struct{}

func (g Star) Prefix() ParseFn {
	return nil
}

func (g Star) Infix() ParseFn {
	return Binary
}

func (g Star) Precedence() int {
	return PREC_FACTOR
}

/* Bang */
type Bang struct{}

func (g Bang) Prefix() ParseFn {
	return nil
}

func (g Bang) Infix() ParseFn {
	return nil
}

func (g Bang) Precedence() int {
	return PREC_NONE
}

/* BangEqual */
type BangEqual struct{}

func (g BangEqual) Prefix() ParseFn {
	return nil
}

func (g BangEqual) Infix() ParseFn {
	return nil
}

func (g BangEqual) Precedence() int {
	return PREC_NONE
}

/* Greater */
type Greater struct{}

func (g Greater) Prefix() ParseFn {
	return nil
}

func (g Greater) Infix() ParseFn {
	return nil
}

func (g Greater) Precedence() int {
	return PREC_NONE
}

/* GreaterEqual */
type GreaterEqual struct{}

func (g GreaterEqual) Prefix() ParseFn {
	return nil
}

func (g GreaterEqual) Infix() ParseFn {
	return nil
}

func (g GreaterEqual) Precedence() int {
	return PREC_NONE
}

/* Less */
type Less struct{}

func (g Less) Prefix() ParseFn {
	return nil
}

func (g Less) Infix() ParseFn {
	return nil
}

func (g Less) Precedence() int {
	return PREC_NONE
}

/* LessEqual */
type LessEqual struct{}

func (g LessEqual) Prefix() ParseFn {
	return nil
}

func (g LessEqual) Infix() ParseFn {
	return nil
}

func (g LessEqual) Precedence() int {
	return PREC_NONE
}

/* Equal */
type Equal struct{}

func (g Equal) Prefix() ParseFn {
	return nil
}

func (g Equal) Infix() ParseFn {
	return nil
}

func (g Equal) Precedence() int {
	return PREC_NONE
}

/* EqualEqual */
type EqualEqual struct{}

func (g EqualEqual) Prefix() ParseFn {
	return nil
}

func (g EqualEqual) Infix() ParseFn {
	return nil
}

func (g EqualEqual) Precedence() int {
	return PREC_NONE
}

/* Identifier */
type Identifier struct{}

func (g Identifier) Prefix() ParseFn {
	return nil
}

func (g Identifier) Infix() ParseFn {
	return nil
}

func (g Identifier) Precedence() int {
	return PREC_NONE
}

/* String */
type String struct{}

func (g String) Prefix() ParseFn {
	return nil
}

func (g String) Infix() ParseFn {
	return nil
}

func (g String) Precedence() int {
	return PREC_NONE
}

/* Number */
type NumberPrec struct{}

func (g NumberPrec) Prefix() ParseFn {
	return Number
}

func (g NumberPrec) Infix() ParseFn {
	return nil
}

func (g NumberPrec) Precedence() int {
	return PREC_NONE
}

/* And */
type And struct{}

func (g And) Prefix() ParseFn {
	return nil
}

func (g And) Infix() ParseFn {
	return nil
}

func (g And) Precedence() int {
	return PREC_NONE
}

/* Class */
type Class struct{}

func (g Class) Prefix() ParseFn {
	return nil
}

func (g Class) Infix() ParseFn {
	return nil
}

func (g Class) Precedence() int {
	return PREC_NONE
}

/* Else */
type Else struct{}

func (g Else) Prefix() ParseFn {
	return nil
}

func (g Else) Infix() ParseFn {
	return nil
}

func (g Else) Precedence() int {
	return PREC_NONE
}

/* False */
type False struct{}

func (g False) Prefix() ParseFn {
	return nil
}

func (g False) Infix() ParseFn {
	return nil
}

func (g False) Precedence() int {
	return PREC_NONE
}

/* For */
type For struct{}

func (g For) Prefix() ParseFn {
	return nil
}

func (g For) Infix() ParseFn {
	return nil
}

func (g For) Precedence() int {
	return PREC_NONE
}

/* Fun */
type Fun struct{}

func (g Fun) Prefix() ParseFn {
	return nil
}

func (g Fun) Infix() ParseFn {
	return nil
}

func (g Fun) Precedence() int {
	return PREC_NONE
}

/* If */
type If struct{}

func (g If) Prefix() ParseFn {
	return nil
}

func (g If) Infix() ParseFn {
	return nil
}

func (g If) Precedence() int {
	return PREC_NONE
}

/* Nil */
type Nil struct{}

func (g Nil) Prefix() ParseFn {
	return nil
}

func (g Nil) Infix() ParseFn {
	return nil
}

func (g Nil) Precedence() int {
	return PREC_NONE
}

/* Or */
type Or struct{}

func (g Or) Prefix() ParseFn {
	return nil
}

func (g Or) Infix() ParseFn {
	return nil
}

func (g Or) Precedence() int {
	return PREC_NONE
}

/* Print */
type Print struct{}

func (g Print) Prefix() ParseFn {
	return nil
}

func (g Print) Infix() ParseFn {
	return nil
}

func (g Print) Precedence() int {
	return PREC_NONE
}

/* Return */
type Return struct{}

func (g Return) Prefix() ParseFn {
	return nil
}

func (g Return) Infix() ParseFn {
	return nil
}

func (g Return) Precedence() int {
	return PREC_NONE
}

/* Super */
type Super struct{}

func (g Super) Prefix() ParseFn {
	return nil
}

func (g Super) Infix() ParseFn {
	return nil
}

func (g Super) Precedence() int {
	return PREC_NONE
}

/* True */
type True struct{}

func (g True) Prefix() ParseFn {
	return nil
}

func (g True) Infix() ParseFn {
	return nil
}

func (g True) Precedence() int {
	return PREC_NONE
}

/* this */
type This struct{}

func (g This) Prefix() ParseFn {
	return nil
}

func (g This) Infix() ParseFn {
	return nil
}

func (g This) Precedence() int {
	return PREC_NONE
}

/* Var */
type Var struct{}

func (g Var) Prefix() ParseFn {
	return nil
}

func (g Var) Infix() ParseFn {
	return nil
}

func (g Var) Precedence() int {
	return PREC_NONE
}

/* While */
type While struct{}

func (g While) Prefix() ParseFn {
	return nil
}

func (g While) Infix() ParseFn {
	return nil
}

func (g While) Precedence() int {
	return PREC_NONE
}

/* Error */
type Error struct{}

func (g Error) Prefix() ParseFn {
	return nil
}

func (g Error) Infix() ParseFn {
	return nil
}

func (g Error) Precedence() int {
	return PREC_NONE
}

/* EOF */
type EOF struct{}

func (g EOF) Prefix() ParseFn {
	return nil
}

func (g EOF) Infix() ParseFn {
	return nil
}

func (g EOF) Precedence() int {
	return PREC_NONE
}
