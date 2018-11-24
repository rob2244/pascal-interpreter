package ast

import "github.com/rob2244/pascal-interpreter/lexer"

type Type struct {
	Token *lexer.Token
	Value string
}

func NewType(token *lexer.Token) *Type {
	return &Type{token, token.Value}
}
