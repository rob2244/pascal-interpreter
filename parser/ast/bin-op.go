package ast

import "github.com/rob2244/pascal-interpreter/lexer"

type BinOp struct {
	Left  interface{}
	Token *lexer.Token
	Op    string
	Right interface{}
}

func NewBinOp(token *lexer.Token, left, right interface{}) *BinOp {
	return &BinOp{left, token, token.Value, right}
}
