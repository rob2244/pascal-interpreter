package lexer

import (
	"fmt"
)

// TokenType defines the available tokens
type TokenType string

const (
	EOF      TokenType = "EOF"
	INTEGER  TokenType = "INTEGER"
	PLUS     TokenType = "PLUS"
	MINUS    TokenType = "MINUS"
	DIVIDE   TokenType = "DIVIDE"
	MULTIPLY TokenType = "MULTIPLY"
	LPAREN   TokenType = "LPAREN"
	RPAREN   TokenType = "RPAREN"
	BEGIN    TokenType = "BEGIN"
	END      TokenType = "END"
	DOT      TokenType = "DOT"
	ASSIGN   TokenType = "ASSIGN"
	SEMI     TokenType = "SEMI"
	ID       TokenType = "ID"
)

// Token represents a syntax token
type Token struct {
	Type  TokenType
	Value string
}

func (t *Token) String() string {
	return fmt.Sprintf("%s, %s", t.Type, t.Value)
}
