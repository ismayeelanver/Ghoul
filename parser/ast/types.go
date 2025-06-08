package ast

type SymbolType struct {
	SymbolName string
}

func (s SymbolType) type_kind() {}

type ArrayType struct {
	Underlying Type
	Size       *Expr
}

func (a ArrayType) type_kind() {}

func MakeVoidType() SymbolType {
	return SymbolType{
		SymbolName: "void",
	}
}