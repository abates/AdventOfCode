package main

import (
	"reflect"
	"testing"

	"github.com/abates/lang/lex"
)

func TestLexing(t *testing.T) {
	tests := []struct {
		input    string
		tokens   []lex.TokenType
		literals []string
	}{
		{"pbga (66)", []lex.TokenType{TokenPID, TokenSize, TokenEOL}, []string{"pbga", "66", ""}},
		{"fwft (72) -> ktlj, cntj, xhth", []lex.TokenType{TokenPID, TokenSize, TokenChildPID, TokenChildPID, TokenChildPID, TokenEOL}, []string{"fwft", "72", "ktlj", "cntj", "xhth", ""}},
		{"xhth (57) \nebii (61)", []lex.TokenType{TokenPID, TokenSize, TokenEOL, TokenPID, TokenSize, TokenEOL}, []string{"xhth", "57", "", "ebii", "61", ""}},
	}

	for i, test := range tests {
		tokens := make([]lex.TokenType, 0)
		literals := make([]string, 0)
		lexer := lex.New(test.input, lexPID)
		for token := lexer.NextToken(); token.Type != lex.EOF; token = lexer.NextToken() {
			tokens = append(tokens, token.Type)
			literals = append(literals, token.Literal)
		}

		if !reflect.DeepEqual(test.tokens, tokens) {
			t.Errorf("tests[%d] expected %+v got %+v", i, test.tokens, tokens)
		}

		if !reflect.DeepEqual(test.literals, literals) {
			t.Errorf("tests[%d] expected %+v got %+v", i, test.literals, literals)
		}
	}
}
