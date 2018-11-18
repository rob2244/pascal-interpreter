package main

import (
	"fmt"

	"github.com/rob2244/pascal-interpreter/interpreter"
	"github.com/rob2244/pascal-interpreter/lexer"
	"github.com/rob2244/pascal-interpreter/parser"
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
	         ____ := 10 * a + 10 * number DIV 4;
	         c := a - - ____
	     END;
	
		x := 11;
	END.
	`

	l := lexer.NewLexer(exp)
	p := parser.NewParser(l)

	i := interpreter.NewInterpreter(p)

	i.Interpret()
	fmt.Println(interpreter.GLOBAL_SCOPE)

	//}
}
