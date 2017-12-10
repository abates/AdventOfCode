package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/abates/lang/lex"
)

const (
	TokenSize     = lex.TokenType("Size")
	TokenPID      = lex.TokenType("PID")
	TokenChildPID = lex.TokenType("ChildPID")
	TokenEOL      = lex.TokenType("EOL")
)

func lexStatement(lexer *lex.Lexer) lex.StateFn {
	lexer.IgnoreWhitespace()
	r := lexer.Peek()
	if lex.IsAlpha(r) {
		return lexPID
	}
	return nil
}

func lexChildPID(lexer *lex.Lexer) lex.StateFn {
	lexer.AcceptAlpha()
	lexer.Emit(TokenChildPID)
	lexer.IgnoreWhitespace()
	if lexer.Peek() == ',' {
		lexer.Next()
		lexer.IgnoreWhitespace()
		if lex.IsAlpha(lexer.Peek()) {
			return lexChildPID
		}
		return lexer.Errorf("Expected alpha got %q", lexer.Peek())
	}
	lexer.Emit(TokenEOL)
	return lexStatement
}

func lexArrow(lexer *lex.Lexer) lex.StateFn {
	if lexer.Next() == '>' {
		lexer.IgnoreWhitespace()
		if lex.IsAlpha(lexer.Peek()) {
			return lexChildPID
		}
		return lexer.Errorf("Expected alpha got %q", lexer.Peek())
	}
	lexer.Backup()
	return lexer.Errorf("Expected '>' got %q", lexer.Peek())
}

func lexPID(lexer *lex.Lexer) lex.StateFn {
	lexer.AcceptAlpha()
	lexer.Emit(TokenPID)
	lexer.IgnoreWhitespace()
	if lexer.Next() == '(' {
		lexer.Ignore()
		return lexSize
	}
	lexer.Backup()
	return lexer.Errorf("Expected ( got %q", lexer.Peek())
}

func lexSize(lexer *lex.Lexer) lex.StateFn {
	lexer.AcceptDigits()
	lexer.Emit(TokenSize)
	lexer.IgnoreWhitespace()
	if lexer.Next() == ')' {
		lexer.IgnoreWhitespace()
		if lexer.Peek() == '-' {
			lexer.Next()
			return lexArrow
		}
		lexer.Ignore()
		lexer.Emit(TokenEOL)
		return lexStatement
	}
	return lexer.Errorf("Expected ) got %s", string(lexer.Peek()))
}

type Program struct {
	ID       string
	Size     int
	Parent   *Program
	Children []*Program
}

func (p *Program) Weight() int {
	weight := p.Size
	for _, child := range p.Children {
		weight += child.Weight()
	}
	return weight
}

func (p *Program) FindUnbalanced() {
	if len(p.Children) > 0 {
		weights := make(map[int]int)
		for _, child := range p.Children {
			weights[child.Weight()] += 1
		}

		weight := 0
		for w, count := range weights {
			if count > 1 {
				weight = w
			}
		}

		for _, child := range p.Children {
			if child.Weight() != weight {
				fmt.Printf("Node %s is unbalanced (need %d).  Its weight is %d and it's total weight is %d\n", child.ID, weight, child.Size, child.Weight())
				child.FindUnbalanced()
			}
		}
	}
}

type Programs struct {
	programs map[string]*Program
}

func (ps *Programs) Find(pid string) (program *Program) {
	program, found := ps.programs[pid]
	if !found {
		program = &Program{
			ID:       pid,
			Children: make([]*Program, 0),
		}
		ps.programs[pid] = program
	}
	return program
}

func (ps *Programs) parseStatement(id string, lexer *lex.Lexer) {
	sizeToken := lexer.NextToken()
	size, _ := strconv.Atoi(sizeToken.Literal)
	program := ps.Find(id)
	program.Size = size
	children := make([]*Program, 0)
	for token := lexer.NextToken(); token.Type != TokenEOL; token = lexer.NextToken() {
		child := ps.Find(token.Literal)
		child.Parent = program
		children = append(children, child)
	}

	program.Children = children
}

func main() {
	ps := &Programs{
		programs: make(map[string]*Program),
	}

	b, _ := ioutil.ReadAll(os.Stdin)
	lexer := lex.New(string(b), lexPID)
	for token := lexer.NextToken(); token.Type != lex.EOF; token = lexer.NextToken() {
		if token.Type == TokenPID {
			ps.parseStatement(token.Literal, lexer)
		}
	}

	var root *Program
	for _, program := range ps.programs {
		if program.Parent == nil {
			root = program
			fmt.Printf("Root: %s\n", program.ID)
			break
		}
	}

	root.FindUnbalanced()
}
