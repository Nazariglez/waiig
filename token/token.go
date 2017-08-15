// Created by nazarigonzalez on 15/8/17.

package token

type TokenType string

type Token struct {
  Type TokenType
  Literal string
}

const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

  //identifiers + literals
  IDENT = "IDENT"
  INT = "INT"

  //operators
  ASSIGN = "="
  PLUS = "+"

  //delimiters
  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  //keywords
  FUNCTION = "FUNCTION"
  LET = "LET"
)