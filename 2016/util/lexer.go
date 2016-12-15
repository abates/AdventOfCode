/**
 * Adapted from
 *       Author: Rob Pike
 *         Date: 2011-08-30
 * Availability: https://talks.golang.org/2011/lex.slide#1
 */
package util

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type ItemType int

const (
	ItemError ItemType = 1
	EOF                = -1
)

type Item struct {
	Typ   ItemType
	Value string
}

type StateFn func(*Lexer) StateFn

type Lexer struct {
	input string
	start int
	pos   int
	width int
	items chan Item
	state StateFn
}

func Lex(input string, startState StateFn) *Lexer {
	l := &Lexer{
		input: input,
		state: startState,
		items: make(chan Item, 2),
	}
	return l
}

func (l *Lexer) Accept(valid string) bool {
	if strings.IndexRune(valid, l.Next()) >= 0 {
		return true
	}
	l.Backup()
	return false
}

func (l *Lexer) AcceptRun(valid string) {
	for strings.IndexRune(valid, l.Next()) >= 0 {
	}
	l.Backup()
}

func (l *Lexer) AcceptUntil(stop rune) {
	r := l.Next()
	for ; r != EOF && r != stop; r = l.Next() {
	}
	if r != EOF {
		l.Backup()
	}
}

func (l *Lexer) AcceptLength(length int) {
	l.pos += length
}

func (l *Lexer) AcceptInteger() {
	l.AcceptRun("0123456789")
}

func (l *Lexer) Backup() {
	l.pos -= l.width
}

func (l *Lexer) Emit(t ItemType) {
	if l.pos > l.start {
		l.items <- Item{t, l.input[l.start:l.pos]}
		l.start = l.pos
	}
}

func (l *Lexer) Errorf(format string, args ...interface{}) StateFn {
	return func(*Lexer) StateFn {
		l.items <- Item{
			ItemError,
			fmt.Sprintf(format, args...),
		}
		return nil
	}
}

func (l *Lexer) Ignore() {
	l.start = l.pos
}

func (l *Lexer) Next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return EOF
	}

	r, l.width =
		utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

func (l *Lexer) NextItem() *Item {
	for {
		select {
		case item := <-l.items:
			return &item
		default:
			if l.state == nil {
				return nil
			}
			l.state = l.state(l)
		}
	}
}

func (l *Lexer) Peek() rune {
	r := l.Next()
	l.Backup()
	return r
}
