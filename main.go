package main

import (
	"fmt"

	"github.com/rob2244/pascal-interpreter/lexer"
)

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)

// 	for scanner.Scan() {
// 		exp := scanner.Text()

// 		l := lexer.NewLexer(exp)
// 		p := ast.NewParser(l)
// 		i := interpreter.NewTreeInterpreter(p)

// 		fmt.Println(i.Interpret())
// 	}
// }

func main() {
	lexer := lexer.NewLexer("BEGIN a := 2; END.")

	fmt.Println(lexer.GetNextToken())
	fmt.Println(lexer.GetNextToken())
	fmt.Println(lexer.GetNextToken())
	fmt.Println(lexer.GetNextToken())
	fmt.Println(lexer.GetNextToken())
	fmt.Println(lexer.GetNextToken())
	fmt.Println(lexer.GetNextToken())
	fmt.Println(lexer.GetNextToken())

}
