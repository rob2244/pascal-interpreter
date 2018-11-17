package ast

import "github.com/rob2244/pascal-interpreter/lexer"

// Var represents a variable declaration
type Var struct {
	Token *lexer.Token
}

// Value returns the variable name
func (v *Var) Value() string {
	return v.Token.Value
}
