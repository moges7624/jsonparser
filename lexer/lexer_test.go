package lexer

import (
	"testing"

	"github.com/moges7624/jsonparser/token"
)

func TestNextToken(t *testing.T) {
  input := `{name:"moges","age":"twenty"}
  `
  tests := []struct {
    expectedType token.TokenType
    expectedLiteral string
  }{
    {token.LEFT_BRACE, "{"},
    {token.IDENT, "name"},
    {token.COLON, ":"},
    {token.IDENT, "\"moges\""},
    {token.COMMA, ","},
    {token.IDENT, "\"age\""},
    {token.COLON, ":"},
    {token.IDENT, "\"twenty\""},
  }

  l := New(input)

  for i, tt := range tests {
    tok := l.nextToken()

    if tok.Type != tt.expectedType {
      t.Fatalf("tests[%d] - token type wrong. Expected=%q, got=%q", i, tt.expectedType, tok.Type)
    }

    if tok.Literal != tt.expectedLiteral {
      t.Fatalf("tests[%d] - token literal wrong. Expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
    }
  }
}
