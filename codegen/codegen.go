package codegen

import (
	"fmt"
	"ghoul/parser/ast"
	"os"
	"path/filepath"
	"strings"
)

type Generator struct {
	ast        *ast.BlockStmt
	filename   string
}

func NewGenerator(ast *ast.BlockStmt, filename string) *Generator {
	return &Generator{
		ast:        ast,
		filename:   filename,
	}
}

func (g *Generator) Generate(f string) {
	file, err := os.Create(f)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	className := strings.TrimSuffix(filepath.Base(f), filepath.Ext(f))

	// Class header
	file.WriteString(fmt.Sprintf(".class public %s\n", className))
	file.WriteString(".super java/lang/Object\n")

	// Default constructor
	file.WriteString("\n.method public <init>()V\n")
	file.WriteString("    aload_0\n")
	file.WriteString("    invokespecial java/lang/Object/<init>()V\n")
	file.WriteString("    return\n")
	file.WriteString(".end method\n")
}

