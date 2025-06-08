package parser

import (
	"ghoul/parser/ast"
	"ghoul/parser/lexer"
)

type binding_power int

const (
	defalt_bp binding_power = iota
	comma
	assignment
	logical
	relational
	additive
	multiplicative
	unary
	call
	member
	primary
)

type stmt_handler func(p *parser) ast.Stmt
type nud_handler func(p *parser) ast.Expr // h
type led_handler func(p *parser, left ast.Expr, bp binding_power) ast.Expr

type stmt_lku map[lexer.TokenType]stmt_handler
type nud_lku map[lexer.TokenType]nud_handler
type led_lku map[lexer.TokenType]led_handler
type bp_lku map[lexer.TokenType]binding_power

var bp_lu = bp_lku{}
var led_lu = led_lku{}
var nud_lu = nud_lku{}
var stmt_lu = stmt_lku{}

func led(kind lexer.TokenType, bp binding_power, led_fn led_handler) {
	bp_lu[kind] = bp
	led_lu[kind] = led_fn
}

func nud(kind lexer.TokenType, bp binding_power, nud_fn nud_handler) {
	bp_lu[kind] = bp
	nud_lu[kind] = nud_fn
}

func stmt(kind lexer.TokenType, stmt_fn stmt_handler) {
	bp_lu[kind] = defalt_bp
	stmt_lu[kind] = stmt_fn
}

func createTokenLookups() {
	// led(lexer.EQUAL, assignment, parse_assignment_expr)
	// led(lexer.PLUS_EQUALS, assignment, parse_assignment_expr)
	// led(lexer.MINUS_EQUALS, assignment, parse_assignment_expr)

	led(lexer.AND, logical, parse_binary_expr)
	led(lexer.OR, logical, parse_binary_expr)
	// led(lexer.DOTDOT, logical, parse_range_expr)

	led(lexer.LESS, relational, parse_binary_expr)
	led(lexer.LESS_EQUALS, relational, parse_binary_expr)
	led(lexer.GREATER, relational, parse_binary_expr)
	led(lexer.GREATER_EQUALS, relational, parse_binary_expr)
	led(lexer.EQUALSEQUALS, relational, parse_binary_expr)
	led(lexer.NOT_EQUALS, relational, parse_binary_expr)

	led(lexer.PLUS, additive, parse_binary_expr)
	led(lexer.DASH, additive, parse_binary_expr)
	led(lexer.SLASH, multiplicative, parse_binary_expr)
	led(lexer.STAR, multiplicative, parse_binary_expr)
	led(lexer.MOD, multiplicative, parse_binary_expr)

	led(lexer.LPAREN, call, parse_call_expr)

	nud(lexer.NUMBER, primary, parse_primary_expr)
	nud(lexer.STRING, primary, parse_primary_expr)
	nud(lexer.CHAR, primary, parse_primary_expr)
	nud(lexer.FLOAT, primary, parse_primary_expr)
	nud(lexer.TRUE, primary, parse_primary_expr)
	nud(lexer.FALSE, primary, parse_primary_expr)
	nud(lexer.IDENTIFIER, primary, parse_primary_expr)
	nud(lexer.LPAREN, primary, parse_primary_expr)
	nud(lexer.NEW, primary, parse_primary_expr)

	stmt(lexer.VAR, parse_var_stmt)
	stmt(lexer.CONST, parse_const_stmt)
	stmt(lexer.IF, parse_if_stmt)
	stmt(lexer.FUN, parse_fun_stmt)
	stmt(lexer.RETURN, parse_return_stmt)
	stmt(lexer.WHILE, parse_while_stmt)
	stmt(lexer.DO, parse_do_stmt)
	stmt(lexer.FOR, parse_for_stmt)
}
