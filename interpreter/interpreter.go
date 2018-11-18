package interpreter

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/rob2244/pascal-interpreter/lexer"
	"github.com/rob2244/pascal-interpreter/parser"
	"github.com/rob2244/pascal-interpreter/parser/ast"
)

// Interpreter represents a simple Pascal Interpreter
type Interpreter struct {
	parser *parser.Parser
}

// GLOBAL_SCOPE acts as symbol table and user memory for program
var GLOBAL_SCOPE = make(map[string]interface{}, 1)

// NewInterpreter is the constructor function for the Pascal Interpreter
func NewInterpreter(parser *parser.Parser) *Interpreter {
	return &Interpreter{parser: parser}
}

// Interpret interprets your pascal code
func (i *Interpreter) Interpret() {
	tree, _ := i.parser.Parse()
	i.visit(tree)
}

func (i *Interpreter) visit(node interface{}) interface{} {
	nodeStringType := reflect.TypeOf(node).String()
	p := strings.Index(nodeStringType, ".")
	nodeStringType = nodeStringType[p+1:]
	nodeStringType = "Visit" + nodeStringType

	met := reflect.ValueOf(i).MethodByName(nodeStringType)
	value := met.Call([]reflect.Value{reflect.ValueOf(node)})

	if len(value) == 0 {
		return nil
	}

	return value[0].Interface()
}

func (i *Interpreter) VisitCompound(node *ast.Compound) {
	for _, n := range node.Children {
		i.visit(n)
	}
}

func (i *Interpreter) VisitNoOp(node *ast.NoOp) {
	return
}

func (i *Interpreter) VisitAssign(node *ast.Assign) {
	varName := node.Left.Value()
	GLOBAL_SCOPE[varName] = i.visit(node.Right)
}

func (i *Interpreter) VisitVar(node *ast.Var) (interface{}, error) {
	val, ok := GLOBAL_SCOPE[node.Value()]

	if !ok {
		return nil, fmt.Errorf("Value for %s is not in symbol table", node.Value())
	}

	return val, nil
}

func (i *Interpreter) VisitBinOp(node *ast.BinOp) (int, error) {
	switch node.Token.Type {
	case lexer.PLUS:
		return i.visit(node.Left).(int) + i.visit(node.Right).(int), nil
	case lexer.MINUS:
		return i.visit(node.Left).(int) - i.visit(node.Right).(int), nil
	case lexer.MULTIPLY:
		return i.visit(node.Left).(int) * i.visit(node.Right).(int), nil
	case lexer.DIVIDE:
		return i.visit(node.Left).(int) / i.visit(node.Right).(int), nil
	default:
		return 0, fmt.Errorf("Invalid Token")
	}
}

func (i *Interpreter) VisitNum(node *ast.Num) int {
	return node.Value
}

func (i *Interpreter) VisitUnaryOp(node *ast.UnaryOp) int {
	op := node.Token.Type

	if op == lexer.PLUS {
		return +i.visit(node.Expr).(int)
	}

	if op == lexer.MINUS {
		return -i.visit(node.Expr).(int)
	}

	return 0
}
