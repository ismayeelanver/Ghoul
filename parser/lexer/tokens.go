package lexer

type TokenType uint

const (
	PLUS TokenType = iota
	DASH
	STAR
	SLASH
	MOD
	LPAREN
	RPAREN
	LSQUARE
	RSQUARE
	LCURLY
	RCURLY
	DOT
	DOTDOT
	COMMA
	COLON
	EQUAL
	EQUALSEQUALS
	LESS_EQUALS
	GREATER_EQUALS
	LESS
	GREATER
	NOT_EQUALS
	NOT
	AND
	OR
	PLUS_EQUALS
	MINUS_EQUALS

	NUMBER
	FLOAT
	STRING
	CHAR
	IDENTIFIER
	NEWLINE
	FUN
	IF
	ELSEIF
	ELSE
	VAR
	CONST
	IMPORT
	NEW
	PUB
	TRUE
	FALSE
	RETURN

	ERRINVALID
	ERRINVALIDCHAR
	ERRINVALIDSTRING
	ERRINVALIDNUMBER
	ERRINVALIDFLOAT
	ERRINVALIDIDENTIFIER

	EOF
)

func (t TokenType) String() string {
	switch t {
	case EOF:
		return "EOF"
	case PLUS:
		return "+"
	case DASH:
		return "-"
	case STAR:
		return "*"
	case SLASH:
		return "/"
	case MOD:
		return "%"
	case LPAREN:
		return "("
	case RPAREN:
		return ")"
	case LSQUARE:
		return "["
	case RSQUARE:
		return "]"
	case LCURLY:
		return "{"
	case RCURLY:
		return "}"
	case DOT:
		return "."
	case DOTDOT:
		return ".."
	case COMMA:
		return ","
	case COLON:
		return ":"
	case EQUAL:
		return "="
	case EQUALSEQUALS:
		return "=="
	case LESS_EQUALS:
		return "<="
	case GREATER_EQUALS:
		return ">="
	case LESS:
		return "<"
	case GREATER:
		return ">"
	case NOT_EQUALS:
		return "!="
	case NOT:
		return "!"
	case AND:
		return "&&"
	case OR:
		return "||"
	case PLUS_EQUALS:
		return "+="
	case MINUS_EQUALS:
		return "-="
	case NUMBER:
		return "number"
	case FLOAT:
		return "float"
	case STRING:
		return "string"
	case CHAR:
		return "char"
	case IDENTIFIER:
		return "identifier"
	case NEWLINE:
		return "newline"
	case FUN:
		return "fun"
	case IF:
		return "if"
	case ELSEIF:
		return "elseif"
	case ELSE:
		return "else"
	case VAR:
		return "var"
	case CONST:
		return "const"
	case IMPORT:
		return "import"
	case NEW:
		return "new"
	case PUB:
		return "pub"
	case TRUE:
		return "true"
	case FALSE:
		return "false"
	case ERRINVALID:
		return "ERRINVALID"
	case ERRINVALIDCHAR:
		return "ERRINVALIDCHAR"
	case ERRINVALIDSTRING:
		return "ERRINVALIDSTRING"
	case ERRINVALIDNUMBER:
		return "ERRINVALIDNUMBER"
	case ERRINVALIDFLOAT:
		return "ERRINVALIDFLOAT"
	case ERRINVALIDIDENTIFIER:
		return "ERRINVALIDIDENTIFIER"
	default:
		return t.String()
	}
}

type Position struct {
	Line   uint
	Column uint
}

type Token struct {
	Type  TokenType
	Value string
	Pos   Position
}

func makeToken(t TokenType, v string, pos Position) Token {
	return Token{
		Type:  t,
		Value: v,
		Pos:   pos,
	}
}
