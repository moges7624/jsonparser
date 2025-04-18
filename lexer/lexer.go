package lexer

import (
	// "fmt"

	"github.com/moges7624/jsonparser/token"
)

type Lexer struct {
  input string
  position int
  readPosition int
  ch byte
}

func New(input string) *Lexer {
  l := &Lexer{input: input}
  l.readChar()
  return l
}

func (l *Lexer) readChar() {
  if l.readPosition >= len(l.input) {
    l.ch = 0
  } else {
    l.ch = l.input[l.readPosition]
  }

  l.position = l.readPosition
  l.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
  return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) nextToken() token.Token {
  var tok token.Token

  switch l.ch {
  case '{':
    tok = newToken(token.LEFT_BRACE, l.ch)
  case '}':
    tok = newToken(token.RIGHT_BRACE, l.ch)
  case ',':
    tok = newToken(token.COMMA, l.ch)
  case ':':
    tok = newToken(token.COLON, l.ch)
  case '[':
    tok = newToken(token.LEFT_BRACKET, l.ch)
  case ']':
    tok = newToken(token.RIGHT_BRACKET, l.ch)
  default:
    // check if ident and read
    if isLetter(l.ch) {
      tok.Literal = l.readIdentifier()
      tok.Type = token.IDENT
      return tok
    } else {
      tok = newToken(token.ILLEGAL, l.ch)
      return tok
    }
    // check if number and read
  }

  l.readChar()
  return tok
}

func isLetter(ch byte) bool {
  return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 
  'Z' || ch == '"' 
}

func (l *Lexer) readIdentifier() string {
  position := l.position

  for isLetter(l.ch) {
    l.readChar()
  }
  
  return l.input[position:l.position]
}
