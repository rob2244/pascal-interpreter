package ast

import (
	"strconv"

	"github.com/rob2244/pascal-interpreter/lexer"
)

type Num struct {
	Token *lexer.Token
	Value int
}

func NewNum(token *lexer.Token) *Num {
	num, _ := strconv.Atoi(token.Value)
	return &Num{token, num}
}
