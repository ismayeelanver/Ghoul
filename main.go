package main

import (
	"ghoul/parser/lexer"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	lex := lexer.MakeLexer("example.gh")

	tokens := lex.Lex()
	spew.Dump(tokens)

	// a := parser.Parse(tokens, lex.Filename)
	// for _, stmt := range a.Body {
	// 	spew.Dump(stmt)
	// }

	// ir := ir.MakeIr(lex.Filename, a)

	// ir.Generate()

	// ir.Ast = nil // for looking at the ir clearly

	// spew.Dump(ir)

}
