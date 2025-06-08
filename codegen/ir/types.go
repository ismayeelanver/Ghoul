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
		case "long":
			return gytes.JLong
		case "short":
			return gytes.JShort
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
	case ast.ArrayType:
		switch t.Underlying.(type) {
		case ast.SymbolType:
			ty := ir.ArrayType(t.Underlying)
			return ty
		}
	}
	return gytes.JType{Name: "Unknown", VMRep: "U"}
}


// here
func (ir *Ir) ArrayType(t ast.Type) gytes.JType {
	switch t := t.(type) {
	case ast.SymbolType:
		switch t.SymbolName {
		case "int":
			return gytes.JAInt	
		case "float":
			return gytes.JAFloat
		case "double":
			return gytes.JADouble
		case "long":
			return gytes.JALong
		case "char":
			return gytes.JAChar
		case "bool":
			return gytes.JAByte
		case "byte":
			return gytes.JABool
		case "short":
			return gytes.JAShort
		case "str":
			return gytes.JType{Name: "[]str", VMRep: "[Ljava/lang/String;"}
		}
		default: {
			return gytes.JType{Name: "Unknown", VMRep: "U"}
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
