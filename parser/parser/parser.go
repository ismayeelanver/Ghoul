package parser

import (
	"fmt"
	"ghoul/parser/ast"
	"ghoul/parser/lexer"
)

type parser struct {
	tokens     []lexer.Token
	pos        int
	Filename   string
	inFunction bool
}

func (p *parser) hasTokens() bool {
	return p.pos < len(p.tokens) && p.currentTokenKind() != lexer.EOF
}

func (p *parser) currentTokenKind() lexer.TokenType {
	return p.tokens[p.pos].Type
}

func (p *parser) currentToken() lexer.Token {
	return p.tokens[p.pos]
}

func (p *parser) expect(kind lexer.TokenType) lexer.Token {
	k := p.currentTokenKind()
	if k != kind {
		panic(fmt.Sprintf("error: expected %d token, got %d token\n\tat %s:%d:%d\n", kind, k, p.Filename, p.currentToken().Pos.Line, p.currentToken().Pos.Column))
	}

	return p.advance()
}

func (p *parser) advance() lexer.Token {
	tk := p.currentToken()
	p.pos++
	return tk
}

func (p *parser) skipNewlines() {
	for p.currentTokenKind() == lexer.NEWLINE {
		p.advance()
	}
}

func createParser(tokens []lexer.Token, filename string) *parser {
	createTokenLookups()
	createTokenTypeLookups()
	return &parser{
		tokens:   tokens,
		pos:      0,
		Filename: filename,
		inFunction: false,
	}
}

func Parse(source_tokens []lexer.Token, filename string) *ast.BlockStmt {
	p := createParser(source_tokens, filename)
	body := make([]ast.Stmt, 0)

	p.skipNewlines()
	for p.hasTokens() {
		body = append(body, parse_stmt(p))
		p.skipNewlines()
	}

	return &ast.BlockStmt{
		Body: body,
	}
}
