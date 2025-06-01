package ir

import (
	"ghoul/parser/ast"

	"github.com/chermehdi/gytes"
	"github.com/davecgh/go-spew/spew"
)

func (ir *Ir) GenerateStmt(stmt ast.Stmt) error {
	switch a := stmt.(type) {
	case ast.VarStmt:
		if a.DataType == nil {
			if a.Expr == nil {
				err := spew.Errorf(`Unknown type for variable %s`, a.Name)
				return err
			}
			ir.Program.Variables[a.Name] = MakeVariable(gytes.Access(gytes.ACC_PUBLIC|gytes.ACC_STATIC), a.Name, ir.InferType(*a.Expr), a.Const, *a.Expr)
			break
		}
		ir.Program.Variables[a.Name] = MakeVariable(gytes.Access(gytes.ACC_PUBLIC|gytes.ACC_STATIC), a.Name, ir.Type(*a.DataType), a.Const, *a.Expr)
	case ast.FunStmt:
	}
	return nil
}
