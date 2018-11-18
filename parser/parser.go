package parser

import (
	"fmt"

	"github.com/rob2244/pascal-interpreter/lexer"
	"github.com/rob2244/pascal-interpreter/parser/ast"
)

type Parser struct {
	lexer        *lexer.Lexer
	currentToken *lexer.Token
}

func NewParser(lexer *lexer.Lexer) *Parser {
	token, _ := lexer.GetNextToken()
	return &Parser{lexer, token}
}

func (p *Parser) Parse() (*ast.Compound, error) {
	node := p.program()

	if p.currentToken.Type != lexer.EOF {
		return nil, fmt.Errorf("Expected EOF but got %s", p.currentToken.Type)
	}

	return node, nil
}

func (p *Parser) program() *ast.Compound {
	node := p.compoundStatement()

	p.eat(lexer.DOT)

	return node
}

func (p *Parser) compoundStatement() *ast.Compound {
	p.eat(lexer.BEGIN)
	nodes := p.statementList()
	p.eat(lexer.END)

	root := &ast.Compound{Children: make([]interface{}, 1)}

	for _, node := range nodes {
		root.Children = append(root.Children, node)
	}

	return root
}

func (p *Parser) statementList() []interface{} {
	node := p.statement()
	results := []interface{}{node}

	for p.currentToken.Type == lexer.SEMI {
		p.eat(lexer.SEMI)
		results = append(results, p.statement())
	}

	if p.currentToken.Type == lexer.ID {
		// TODO fix error handling
		panic("Invalid token type")
	}

	return results
}

func (p *Parser) statement() interface{} {
	if p.currentToken.Type == lexer.BEGIN {
		return p.compoundStatement()
	}

	if p.currentToken.Type == lexer.ID {
		return p.assignmentStatement()
	}

	return p.empty()
}

func (p *Parser) assignmentStatement() *ast.Assign {

	left := p.variable()
	token := p.currentToken
	p.eat(lexer.ASSIGN)
	right := p.expr()

	return &ast.Assign{Left: left, Token: token, Op: token.Value, Right: right}
}

func (p *Parser) variable() *ast.Var {
	node := &ast.Var{Token: p.currentToken}
	p.eat(lexer.ID)
	return node
}

func (p *Parser) empty() *ast.NoOp {
	return &ast.NoOp{}
}

func (p *Parser) expr() interface{} {
	node := p.term()

	for p.currentToken.Type == lexer.PLUS || p.currentToken.Type == lexer.MINUS {
		token := p.currentToken

		if token.Type == lexer.PLUS {
			p.eat(lexer.PLUS)
		}

		if token.Type == lexer.MINUS {
			p.eat(lexer.MINUS)
		}

		node = ast.NewBinOp(token, node, p.term())
	}

	return node
}

func (p *Parser) term() interface{} {
	node := p.factor()

	for p.currentToken.Type == lexer.MULTIPLY || p.currentToken.Type == lexer.DIVIDE {
		token := p.currentToken

		if token.Type == lexer.MULTIPLY {
			p.eat(lexer.MULTIPLY)
		}

		if token.Type == lexer.DIVIDE {
			p.eat(lexer.DIVIDE)
		}

		node = ast.NewBinOp(token, node, p.factor())
	}

	return node
}

func (p *Parser) factor() interface{} {
	token := p.currentToken

	if token.Type == lexer.PLUS {
		p.eat(lexer.PLUS)
		return &ast.UnaryOp{Token: token, Expr: p.factor()}
	}

	if token.Type == lexer.MINUS {
		p.eat(lexer.MINUS)
		return &ast.UnaryOp{Token: token, Expr: p.factor()}
	}

	if p.currentToken.Type == lexer.INTEGER {
		p.eat(lexer.INTEGER)
		return ast.NewNum(token)
	}

	if p.currentToken.Type == lexer.LPAREN {
		p.eat(lexer.LPAREN)
		node := p.expr()
		p.eat(lexer.RPAREN)

		return node
	}

	return p.variable()
}

func (p *Parser) eat(tokenType lexer.TokenType) error {
	if tokenType != p.currentToken.Type {
		return fmt.Errorf("Unexpected token found %s", tokenType)
	}

	p.currentToken, _ = p.lexer.GetNextToken()

	return nil
}
