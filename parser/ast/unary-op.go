package ast

import (
	"github.com/rob2244/pascal-interpreter/lexer"
)

type UnaryOp struct {
	Token *lexer.Token
	Expr  interface{}
}
