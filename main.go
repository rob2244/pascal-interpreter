package main

import (
	"fmt"

	"github.com/rob2244/pascal-interpreter/interpreter"
	"github.com/rob2244/pascal-interpreter/parser"

	"github.com/rob2244/pascal-interpreter/lexer"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)

	// for scanner.Scan() {
	// 	exp := scanner.Text()

	exp := ` 
	 BEGIN
	
	     BEGIN
	         number := 2;
	         a := number;
	         b := 10 * a + 10 * number / 4;
	         c := a - - b
	     END;
	
		x := 11;
	END.
	`

	l := lexer.NewLexer(exp)
	p := parser.NewParser(l)
	i := interpreter.NewInterpreter(p)

	fmt.Println(i.Interpret())
	//}
}
