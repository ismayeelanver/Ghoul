package parser

import (
	"fmt"
	"ghoul/parser/ast"
	"ghoul/parser/lexer"
	"strconv"
)

func parse_expr(p *parser, bp binding_power) ast.Expr {
	tokenKind := p.currentTokenKind()

	nud_fn, exists := nud_lu[tokenKind]

	if !exists {
		panic(fmt.Sprintf("nud Cannot create expression from %d\n\tat %s:%d:%d\n", tokenKind, p.Filename, p.currentToken().Pos.Line, p.currentToken().Pos.Column))
	}

	left := nud_fn(p)

	for bp_lu[p.currentTokenKind()] > bp {
		tokenKind = p.currentTokenKind()
		led_fn, exists := led_lu[tokenKind]

		if !exists {
			panic(fmt.Sprintf("led Cannot create expression from %d\n", tokenKind))
		}

		left = led_fn(p, left, bp)
	}

	return left
}

func parse_primary_expr(p *parser) ast.Expr {
	switch p.currentTokenKind() {
	case lexer.LPAREN:
		p.advance()
		value := parse_expr(p, defalt_bp)
		p.expect(lexer.RPAREN)

		return value
	case lexer.NUMBER:
		number, _ := strconv.ParseInt(p.advance().Value, 10, 64)
		return ast.NumberExpr{
			Int: uint(number),
		}
	case lexer.FLOAT:
		number, _ := strconv.ParseFloat(p.advance().Value, 64)
		return ast.FloatExpr{
			Float: number,
		}
	case lexer.STRING:
		return ast.StringExpr{
			Value: p.advance().Value,
		}
	case lexer.CHAR:
		return ast.CharExpr{
			Value: p.advance().Value,
		}
	case lexer.IDENTIFIER:
		return ast.SymbolExpr{
			Symbol: p.advance().Value,
		}
	case lexer.TRUE:
		p.advance()
		return ast.BoolExpr{
			Value: true,
		}
	case lexer.FALSE:
		p.advance()
		return ast.BoolExpr{
			Value: false,
		}
	case lexer.NEW:
		p.advance()
		data_type := parse_type(p, defalt_bp)
		return ast.NewExpr{
			DataType: data_type,
		}
	default:
		panic(fmt.Sprintf("Cannot create primary expression from %d\n", p.currentTokenKind()))

	}
}

func parse_binary_expr(p *parser, left ast.Expr, bp binding_power) ast.Expr {
	operator := p.advance()
	right := parse_expr(p, bp)

	return ast.BinaryExpr{
		Left:  &left,
		Right: &right,
		Op:    operator.Type,
	}
}

func parse_call_expr(p *parser, left ast.Expr, bp binding_power) ast.Expr {
	p.advance() // consume '('
	args := []ast.Expr{}

	for p.hasTokens() && p.currentTokenKind() != lexer.RPAREN {
		args = append(args, parse_expr(p, defalt_bp))

		if p.currentTokenKind() == lexer.COMMA {
			p.advance()
		}
	}
	p.expect(lexer.RPAREN)

	return &ast.CallExpr{
		Callee: left,
		Args:   args,
	}
}
