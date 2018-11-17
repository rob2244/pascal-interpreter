package lexer

import (
	"fmt"
	"strings"
	"unicode"
)

// Lexer represents a lexer
type Lexer struct {
	text string
	pos  int
}

var reservedKeywords = map[string]*Token{
	"BEGIN": &Token{BEGIN, "BEGIN"},
	"END":   &Token{END, "END"},
}

// NONE Represents end of character stream
const NONE = ""

// NewLexer returns a new lexer
func NewLexer(txt string) *Lexer {
	return &Lexer{pos: 0, text: txt}
}

// GetNextToken retrieves the next token
func (l *Lexer) GetNextToken() (*Token, error) {
	for l.currentChar() != NONE {
		if isWhitespace(l.currentChar()) {
			l.skipWhitespace()
		}

		if unicode.IsLetter([]rune(l.currentChar())[0]) {
			return l.id(), nil
		}

		if unicode.IsDigit([]rune(l.currentChar())[0]) {
			return &Token{INTEGER, l.integer()}, nil
		}

		if l.currentChar() == ":" && l.peek() == "=" {
			l.advance()
			l.advance()
			return &Token{ASSIGN, ":="}, nil
		}

		if l.currentChar() == ";" {
			l.advance()
			return &Token{SEMI, ";"}, nil
		}

		if l.currentChar() == "." {
			l.advance()
			return &Token{DOT, "."}, nil
		}

		if l.currentChar() == "*" {
			l.advance()
			return &Token{MULTIPLY, l.currentChar()}, nil
		}

		if l.currentChar() == "/" {
			l.advance()
			return &Token{DIVIDE, l.currentChar()}, nil
		}

		if l.currentChar() == "+" {
			l.advance()
			return &Token{PLUS, l.currentChar()}, nil
		}

		if l.currentChar() == "-" {
			l.advance()
			return &Token{MINUS, l.currentChar()}, nil
		}

		if l.currentChar() == "(" {
			l.advance()
			return &Token{LPAREN, l.currentChar()}, nil
		}

		if l.currentChar() == ")" {
			l.advance()
			return &Token{RPAREN, l.currentChar()}, nil
		}

		return nil, fmt.Errorf("Unrecognized character %v", l.currentChar())
	}

	return &Token{EOF, NONE}, nil
}

func (l *Lexer) currentChar() string {
	if l.pos > len(l.text)-1 {
		return NONE
	}

	return string(l.text[l.pos])
}

func (l *Lexer) integer() string {
	var numString string

	for ; l.pos < len(l.text) && unicode.IsDigit([]rune(l.currentChar())[0]); l.pos++ {
		numString += string(l.currentChar())
	}

	return numString
}

func (l *Lexer) id() *Token {
	result := strings.Builder{}

	for char := l.currentChar(); char != NONE && isAlphaNumeric(char); char = l.currentChar() {
		result.WriteString(char)
		l.pos++
	}

	val, ok := reservedKeywords[result.String()]

	if ok {
		return val
	}

	return &Token{ID, result.String()}
}

func (l *Lexer) skipWhitespace() {
	if l.pos > len(l.text)-1 {
		return
	}

	for ; l.currentChar() == " "; l.pos++ {

	}
}

func (l *Lexer) peek() string {
	pos := l.pos + 1
	if pos > len(l.text)-1 {
		return NONE
	}

	return string(l.text[pos])
}

func (l *Lexer) advance() {
	l.pos++
}

func isAlphaNumeric(s string) bool {
	r := []rune(s)[0]
	return unicode.IsNumber(r) || unicode.IsLetter(r)
}

func isWhitespace(s string) bool {
	return s == " "
}
