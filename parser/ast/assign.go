package ast

import "github.com/rob2244/pascal-interpreter/lexer"

// Assign represents an assignment statement
type Assign struct {
	Left  *Var
	Token *lexer.Token
	Op    string
	Right interface{}
}
