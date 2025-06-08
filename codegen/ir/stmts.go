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
		err := ir.GenerateFun(a)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ir *Ir) GenerateFun(fun ast.FunStmt) error {

	insts := gytes.NewByteBlock()

	f := MakeFunction(fun.Name, gytes.Access(gytes.ACC_PUBLIC|gytes.ACC_STATIC), make([]Argument, 0), 0, 0, ir.Type(fun.Returns), &insts)
	for paramName, paramType := range fun.Params {
		f.Arguments = append(f.Arguments, Argument{Name: paramName, Type: ir.Type(paramType)})
	}

	// for _, stmt := range fun.Body.Body {
	// 	// switch stmt.(type) {
	// 	// // case ast.ExprStmt:
	// 	// // 	// look inside of the stmt (Expr Field)
	// 	// // 	expr := stmt.(ast.ExprStmt).Expr
	// 	// }
	// }

	return nil
}
