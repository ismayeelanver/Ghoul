package ast

import "ghoul/parser/lexer"

type SymbolExpr struct {
	Symbol 	string
}

func (s SymbolExpr) expr() {}


type NumberExpr struct {
	Int uint
}

func (n NumberExpr) expr() {}

type FloatExpr struct {
	Float float64
}

func (n FloatExpr) expr() {}

type StringExpr struct {
	Value string
}

func (n StringExpr) expr() {}

type CharExpr struct {
	Value string
}

func (n CharExpr) expr() {}

type BoolExpr struct {
	Value bool
}

func (n BoolExpr) expr() {}

type BinaryExpr struct {
	Left  *Expr
	Right *Expr
	Op    lexer.TokenType
}

func (n BinaryExpr) expr() {}

type NewExpr struct {
	DataType Type
}

func (n NewExpr) expr() {}

type CallExpr struct {
	Args   []Expr
	Callee Expr
}

func (n CallExpr) expr() {}
