package token

type TokenType string

type Token struct {
  Type TokenType
  Literal string
}

const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

  LEFT_BRACE = "{"
  RIGHT_BRACE = "}"
  COMMA = ","
  COLON = ":"
  LEFT_BRACKET = "["
  RIGHT_BRACKET = "]"

  IDENT = "IDENT"
  INT = "INT"

  TRUE = "TRUE"
  FALSE = "FALSE"
)

