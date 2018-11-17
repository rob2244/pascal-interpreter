package interpreter

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/rob2244/interpreter/ast"
	"github.com/rob2244/pascal-interpreter/lexer"
)

// Interpreter represents a simple Pascal Interpreter
type Interpreter struct {
	parser *ast.Parser
}

// NewInterpreter is the constructor function for the Pascal Interpreter
func NewInterpreter(parser *ast.Parser) *Interpreter {
	return &Interpreter{parser: parser}
}

// Interpret interprets your pascal code
func (i *Interpreter) Interpret() int64 {
	tree := i.parser.Parse()
	return i.visit(tree)
}

func (i *Interpreter) visit(node interface{}) int64 {
	nodeStringType := reflect.TypeOf(node).String()
	p := strings.Index(nodeStringType, ".")
	nodeStringType = nodeStringType[p+1:]
	nodeStringType = "Visit" + nodeStringType

	met := reflect.ValueOf(i).MethodByName(nodeStringType)
	value := met.Call([]reflect.Value{reflect.ValueOf(node)})

	return value[0].Int()
}

func (i *Interpreter) VisitBinOp(node *ast.BinOp) (int64, error) {
	switch node.Token.Type {
	case lexer.PLUS:
		return i.visit(node.Left) + i.visit(node.Right), nil
	case lexer.MINUS:
		return i.visit(node.Left) - i.visit(node.Right), nil
	case lexer.MULTIPLY:
		return i.visit(node.Left) * i.visit(node.Right), nil
	case lexer.DIVIDE:
		return i.visit(node.Left) / i.visit(node.Right), nil
	default:
		return 0, fmt.Errorf("Invalid Token")
	}
}

func (i *Interpreter) VisitNum(node *ast.Num) int64 {
	return node.Value
}

func (i *Interpreter) VisitUnaryOp(node *ast.UnaryOp) int64 {
	op := node.Token.Type

	if op == lexer.PLUS {
		return +i.visit(node.Expr)
	}

	if op == lexer.MINUS {
		return -i.visit(node.Expr)
	}

	return 0
}
