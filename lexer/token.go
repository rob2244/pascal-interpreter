package lexer

import (
	"fmt"
)

// TokenType defines the available tokens
type TokenType string

const (
	EOF          TokenType = "EOF"
	INTEGER      TokenType = "INTEGER"
	REAL         TokenType = "REAL"
	PLUS         TokenType = "PLUS"
	MINUS        TokenType = "MINUS"
	INTEGERDIV   TokenType = "INTEGERDIV"
	FLOATDIV     TokenType = "FLOATDIV"
	MULTIPLY     TokenType = "MULTIPLY"
	LPAREN       TokenType = "LPAREN"
	RPAREN       TokenType = "RPAREN"
	BEGIN        TokenType = "BEGIN"
	END          TokenType = "END"
	DOT          TokenType = "DOT"
	ASSIGN       TokenType = "ASSIGN"
	SEMI         TokenType = "SEMI"
	ID           TokenType = "ID"
	PROGRAM      TokenType = "PROGRAM"
	VAR          TokenType = "VAR"
	COLON        TokenType = "COLON"
	COMMA        TokenType = "COMMA"
	INTEGERCONST TokenType = "INTEGERCONST"
	REALCONST    TokenType = "REALCONST"
)

// Token represents a syntax token
type Token struct {
	Type  TokenType
	Value string
}

func (t *Token) String() string {
	return fmt.Sprintf("%s, %s", t.Type, t.Value)
}
