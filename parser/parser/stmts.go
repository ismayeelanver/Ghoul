package parser

import (
	"fmt"
	"ghoul/parser/ast"
	"ghoul/parser/lexer"
)

func parse_stmt(p *parser) ast.Stmt {
	stmt_fn, exists := stmt_lu[p.currentTokenKind()]
	if exists {
		return stmt_fn(p)
	}

	expression := parse_expr(p, defalt_bp)
	if p.currentTokenKind() == lexer.EOF {
		return ast.ExprStmt{
			Expr: expression,
		}
	}

	p.expect(lexer.NEWLINE)

	return ast.ExprStmt{
		Expr: expression,
	}
}

func parse_inner_block(p *parser) ast.BlockStmt {
	body := make([]ast.Stmt, 0)

	p.skipNewlines()
	for p.hasTokens() && p.currentTokenKind() != lexer.RCURLY {
		body = append(body, parse_stmt(p))
		p.skipNewlines()
	}

	p.expect(lexer.RCURLY) // this will trigger an error if the token is \n or EOF
	return ast.BlockStmt{
		Body: body,
	}
}

func parse_var_stmt(p *parser) ast.Stmt {
	p.advance()
	var_name := p.expect(lexer.IDENTIFIER).Value

	if p.currentTokenKind() == lexer.NEWLINE || p.currentTokenKind() == lexer.EOF {
		return ast.VarStmt{
			Name:     var_name,
			Expr:     nil, // shorthand syntax, deal with it please
			DataType: nil, // deal with this, this means the type is not defined
			Const:    false,
		}
	}

	if p.currentTokenKind() == lexer.EQUAL {
		p.advance()
		expression := parse_expr(p, defalt_bp)

		return ast.VarStmt{
			Name:     var_name,
			Expr:     &expression,
			DataType: nil, // deal with this, this means the type is not defined
			Const:    false,
		}
	}

	data_type := parse_type(p, defalt_bp)
	if p.currentTokenKind() == lexer.NEWLINE || p.currentTokenKind() == lexer.EOF {
		return ast.VarStmt{
			Name:     var_name,
			Expr:     nil, // shorthand syntax, deal with it please
			DataType: &data_type,
			Const:    true,
		}
	}

	p.expect(lexer.EQUAL)
	expression := parse_expr(p, defalt_bp)
	return ast.VarStmt{
		Name:     var_name,
		Expr:     &expression,
		DataType: &data_type,
		Const:    false,
	}
}

func parse_const_stmt(p *parser) ast.Stmt {
	p.advance()
	var_name := p.expect(lexer.IDENTIFIER).Value

	if p.currentTokenKind() == lexer.NEWLINE || p.currentTokenKind() == lexer.EOF {
		return ast.VarStmt{
			Name:     var_name,
			Expr:     nil, // shorthand syntax, deal with it please
			DataType: nil, // deal with this, this means the type is not defined
			Const:    true,
		}
	}

	if p.currentTokenKind() == lexer.EQUAL {
		p.advance()
		expression := parse_expr(p, defalt_bp)

		return ast.VarStmt{
			Name:     var_name,
			Expr:     &expression,
			DataType: nil, // deal with this, this means the type is not defined
			Const:    true,
		}
	}

	data_type := parse_type(p, defalt_bp)
	if p.currentTokenKind() == lexer.NEWLINE || p.currentTokenKind() == lexer.EOF {
		return ast.VarStmt{
			Name:     var_name,
			Expr:     nil, // shorthand syntax, deal with it please
			DataType: &data_type,
			Const:    true,
		}
	}

	p.expect(lexer.EQUAL)
	expression := parse_expr(p, defalt_bp)

	return ast.VarStmt{
		Name:     var_name,
		Expr:     &expression,
		DataType: &data_type,
		Const:    true,
	}
}

func parse_if_stmt(p *parser) ast.Stmt {
	p.advance()
	condition := parse_expr(p, defalt_bp)

	p.expect(lexer.LCURLY)
	body := parse_inner_block(p)

	switch p.currentTokenKind() {
	case lexer.ELSEIF:
		next := parse_if_stmt(p).(ast.IfStmt)
		return ast.IfStmt{
			Condition:  condition,
			Body:       body,
			Consequent: &next,
		}
	case lexer.ELSE:
		p.advance()
		p.expect(lexer.LCURLY)
		else_body := parse_inner_block(p)
		return ast.IfStmt{
			Condition: condition,
			Body:      body,
			Consequent: &ast.IfStmt{
				Condition: ast.BoolExpr{
					Value: true,
				},
				Body:       else_body,
				Consequent: nil,
			},
		}
	default:
		return ast.IfStmt{
			Condition:  condition,
			Body:       body,
			Consequent: nil,
		}
	}
}

func parse_fun_stmt(p *parser) ast.Stmt {
	p.advance()
	name := p.expect(lexer.IDENTIFIER).Value

	p.expect(lexer.LPAREN)
	params := make(map[string]ast.Type, 0)
	for p.hasTokens() && p.currentTokenKind() != lexer.RPAREN {
		argument_name := p.expect(lexer.IDENTIFIER).Value
		argument_type := parse_type(p, defalt_bp)
		params[argument_name] = argument_type
		if !p.hasTokens() || p.currentTokenKind() == lexer.RPAREN {
			break
		}

		p.expect(lexer.COMMA)
	}

	p.expect(lexer.RPAREN)
	var returns ast.Type = ast.MakeVoidType()
	if p.currentTokenKind() != lexer.LCURLY {
		returns = parse_type(p, defalt_bp)
	}

	p.expect(lexer.LCURLY)
	prev := p.inFunction
	p.inFunction = true
	body := parse_inner_block(p)
	p.inFunction = prev

	return ast.FunStmt{
		Name:    name,
		Params:  params,
		Body:    body,
		Returns: returns,
	}
}

func parse_while_stmt(p *parser) ast.Stmt {
	p.advance()
	condition := parse_expr(p, defalt_bp)
	for p.currentTokenKind() == lexer.NEWLINE {
		p.advance()
	}

	p.expect(lexer.LCURLY)
	body := parse_inner_block(p)

	return ast.WhileStmt{
		Condition: condition,
		Body:      body,
	}
}

func parse_do_stmt(p *parser) ast.Stmt {
	p.advance()
	for p.currentTokenKind() == lexer.NEWLINE {
		p.advance()
	}

	p.expect(lexer.LCURLY)
	body := parse_inner_block(p)

	p.expect(lexer.WHILE)
	condition := parse_expr(p, defalt_bp)

	return ast.DoStmt{
		Condition: condition,
		Body:      body,
	}
}

func parse_for_stmt(p *parser) ast.Stmt {
	p.advance()
	p.expect(lexer.LPAREN)

	stmt1 := parse_stmt(p)
	p.expect(lexer.SEMICOLON)

	condition := parse_expr(p, defalt_bp)
	p.expect(lexer.SEMICOLON)

	stmt2 := parse_stmt(p)
	p.expect(lexer.RPAREN)

	for p.currentTokenKind() == lexer.NEWLINE {
		p.advance()
	}

	p.expect(lexer.LCURLY)
	body := parse_inner_block(p)

	return ast.ForStmt{
		Stmt1:     stmt1,
		Condition: condition,
		Stmt2:     stmt2,
		Body:      body,
	}
}

func parse_return_stmt(p *parser) ast.Stmt {
	if !p.inFunction {
		panic(fmt.Sprintf("Return statements are only allowed inside functions\n\tat %s:%d:%d\n",
			p.Filename, p.currentToken().Pos.Line, p.currentToken().Pos.Column))
	}
	p.advance()
	if p.currentToken().Type == lexer.NEWLINE {
		return ast.ReturnStmt{
			Expr: nil,
		}
	}
	expression := parse_expr(p, defalt_bp)
	return ast.ReturnStmt{
		Expr: expression,
	}
}
