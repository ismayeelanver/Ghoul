package ir

import (
	"ghoul/parser/ast"

	"github.com/chermehdi/gytes"
)

func (ir *Ir) Type(t ast.Type) gytes.JType {
	switch t := t.(type) {
	case ast.SymbolType:
		switch t.SymbolName {
		case "int":
			return gytes.JInt
		case "float":
			return gytes.JFloat
		case "double":
			return gytes.JDouble
		case "char":
			return gytes.JChar
		case "str":
			return gytes.JType{Name: "str", VMRep: "Ljava/lang/String;"}
		case "bool":
			return gytes.JBool
		case "byte":
			return gytes.JByte
		case "void":
			return gytes.JVoid
		}
	}
	return gytes.JType{Name: "Unknown", VMRep: "U"}
}

func (ir *Ir) InferType(expr ast.Expr) gytes.JType {
	switch e := expr.(type) {
	case ast.FloatExpr:
		return gytes.JFloat
	case ast.NumberExpr:
		return gytes.JInt
	case ast.CharExpr:
		return gytes.JChar
	case ast.BoolExpr:
		return gytes.JBool
	case ast.StringExpr:
		return gytes.JType{Name: "str", VMRep: "Ljava/lang/String;"}
	case ast.SymbolExpr:
		expr, ok := ir.Program.Variables[e.Symbol]
		if !ok {
			return gytes.JType{Name: "Unknown", VMRep: "U"}
		}
		return expr.Type
	case ast.BinaryExpr:
		panic("not implemented")
	default:
		return gytes.JType{Name: "Unknown", VMRep: "U"}
	}
}
