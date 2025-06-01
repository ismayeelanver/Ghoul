package ir

import (
	"ghoul/parser/ast"

	"github.com/chermehdi/gytes"
)

type Ir struct {
	GeneratedFrom string
	ProgramClass  string
	Program       Program
	Ast           *ast.BlockStmt
}

type Program struct {
	Variables map[string]Variable
	Functions map[string]Function
}

type Variable struct {
	Modifiers gytes.Access
	ValueExpr ast.Expr
	Value     string
	Type      gytes.JType
	Const     bool
}

type Function struct {
	Name         string
	Modifiers    gytes.Access
	Arguments    []Argument
	MaxLocals    int
	MaxStack     int
	ReturnType   gytes.JType
	Instructions []gytes.BytesBlock
}

type Argument struct {
	Name string
	Type gytes.JType
}

func MakeVariable(modifiers gytes.Access, value string, type_ gytes.JType, const_ bool, valueExpr ast.Expr) Variable {
	return Variable{
		Modifiers: modifiers,
		ValueExpr: valueExpr,
		Value:     value,
		Type:      type_,
		Const:     const_,
	}
}

func MakeFunction(name string, modifiers gytes.Access, arguments []Argument, maxLocals int, maxStack int, returnType gytes.JType, instructions []gytes.BytesBlock) Function {
	return Function{
		Name:         name,
		Modifiers:    modifiers,
		Arguments:    arguments,
		MaxLocals:    maxLocals,
		MaxStack:     maxStack,
		ReturnType:   returnType,
		Instructions: instructions,
	}
}
