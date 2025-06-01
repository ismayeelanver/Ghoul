package main

import (
	"ghoul/codegen/ir"
	"ghoul/parser/lexer"
	"ghoul/parser/parser"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	lex := lexer.MakeLexer("example.gh")

	tokens := lex.Lex()

	a := parser.Parse(tokens, lex.Filename)

	ir := ir.MakeIr(lex.Filename, a)

	ir.Generate()

	spew.Dump(ir)

}
