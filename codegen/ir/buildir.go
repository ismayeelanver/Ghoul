package ir

import "ghoul/parser/ast"

func MakeIr(filename string, program *ast.BlockStmt) *Ir {
	return &Ir{
		GeneratedFrom: filename,
		ProgramClass:  "Program",
		Program: Program{
			Variables: make(map[string]Variable, 0),
			Functions: make(map[string]Function, 0),
		},
		Ast: program,
	}
}

func (ir *Ir) Generate() {
	for _, stmt := range ir.Ast.Body {
		err := ir.GenerateStmt(stmt)
		if err != nil {
			panic(err)
		}
	}
}
