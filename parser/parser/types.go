package parser

import (
	"fmt"
	"ghoul/parser/ast"
	"ghoul/parser/lexer"
)

type type_nud_handler func(p *parser) ast.Type // h
type type_led_handler func(p *parser, left ast.Type, bp binding_power) ast.Type

type type_nud_lku map[lexer.TokenType]type_nud_handler
type type_led_lku map[lexer.TokenType]type_led_handler
type type_bp_lku map[lexer.TokenType]binding_power

var type_bp_lu = type_bp_lku{}
var type_led_lu = type_led_lku{}
var type_nud_lu = type_nud_lku{}

func type_led(kind lexer.TokenType, bp binding_power, led_fn type_led_handler) {
	type_bp_lu[kind] = bp
	type_led_lu[kind] = led_fn
}

func type_nud(kind lexer.TokenType, bp binding_power, nud_fn type_nud_handler) {
	type_bp_lu[kind] = bp
	type_nud_lu[kind] = nud_fn
}

func createTokenTypeLookups() {
	type_nud(lexer.IDENTIFIER, primary, parse_symbol_type)
	type_nud(lexer.LSQUARE, primary, parse_nud_array_type)

	type_led(lexer.LSQUARE, member, parse_led_array_type)
}

func parse_type(p *parser, bp binding_power) ast.Type {
	tokenKind := p.currentTokenKind()
	nud_fn, exists := type_nud_lu[tokenKind]
	if !exists {
		panic(fmt.Sprintf("nud Cannot create type for token %d at %s:%d:%d\n", tokenKind, p.Filename, p.currentToken().Pos.Line, p.advance().Pos.Column))
	}

	left := nud_fn(p)
	for type_bp_lu[p.currentTokenKind()] > bp {
		tokenKind = p.currentTokenKind()
		led_fn, exists := type_led_lu[tokenKind]
		if !exists {
			panic(fmt.Sprintf("led Cannot create type for token %d at %s:%d:%d\n", tokenKind, p.Filename, p.currentToken().Pos.Line, p.advance().Pos.Column))
		}

		left = led_fn(p, left, bp)
	}

	return left
}

func parse_symbol_type(p *parser) ast.Type {
	return ast.SymbolType{
		SymbolName: p.advance().Value,
	}
}

func parse_nud_array_type(p *parser) ast.Type {
	p.advance()
	if p.currentTokenKind() == lexer.RSQUARE {
		p.advance()
		return ast.ArrayType{
			Underlying: parse_type(p, defalt_bp),
			Size:       nil, // the array size are mutable, aka a slice in golang
		}
	}

	size := parse_expr(p, defalt_bp)
	p.expect(lexer.RSQUARE)
	return ast.ArrayType{
		Underlying: parse_type(p, defalt_bp),
		Size:       &size,
	}
}

func parse_led_array_type(p *parser, left ast.Type, bp binding_power) ast.Type {
	p.advance()
	if p.currentTokenKind() == lexer.RSQUARE {
		p.advance()
		return ast.ArrayType{
			Underlying: left,
			Size:       nil, // the array size are mutable, aka a slice in golang
		}
	}

	size := parse_expr(p, defalt_bp)
	p.expect(lexer.RSQUARE)
	return ast.ArrayType{
		Underlying: left,
		Size:       &size,
	}
}
