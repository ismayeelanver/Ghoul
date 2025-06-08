package lexer

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Lexer struct {
	input    string
	Filename string
	idx      uint
	pos      Position
}

func MakeLexer(filename string) Lexer {
	return Lexer{
		idx:      0,
		pos:      Position{Line: 1, Column: 1},
		input:    "",
		Filename: filename,
	}
}

func (l *Lexer) Lex() []Token {
	tokens := []Token{}

	file, err := os.Open(l.Filename)
	if err != nil {
		panic(fmt.Sprintf("error: unable to open file %s: No such file or directory", l.Filename))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l.input += scanner.Text()
		l.input += "\n"
	}

	l.input = l.input[:len(l.input)-1]

	for l.idx < uint(len(l.input)) {
		ch := rune(l.input[l.idx])

		if unicode.IsDigit(ch) {
			start := l.idx
			dotSeen := false

			for l.idx < uint(len(l.input)) {
				curr := rune(l.input[l.idx])
				if unicode.IsDigit(curr) {
					l.idx++
					l.pos.Column++
				} else if curr == '_' {
					if l.idx+1 >= uint(len(l.input)) || (!unicode.IsDigit(rune(l.input[l.idx+1])) && rune(l.input[l.idx+1]) != '.') {
						tokens = append(tokens, makeToken(ERRINVALIDNUMBER, l.input[start:l.idx], l.pos))
						continue
					}
					l.idx++
				} else if curr == '.' && !dotSeen {
					if l.idx+1 >= uint(len(l.input)) || !unicode.IsDigit(rune(l.input[l.idx+1])) {
						tokens = append(tokens, makeToken(ERRINVALIDFLOAT, l.input[start:l.idx], l.pos))
						continue
					}
					dotSeen = true
					l.idx++
					l.pos.Column++
				} else {
					break
				}
			}

			raw := l.input[start:l.idx]
			if raw[len(raw)-1] == '_' {
				panic(fmt.Sprintf("error: numeric literal cannot end with underscore\n\tat line %d, column %d", l.pos.Line, l.pos.Column))
			}

			cleaned := ""
			for _, r := range raw {
				if r != '_' {
					cleaned += string(r)
				}
			}

			tokens = append(tokens, makeToken(NUMBER, cleaned, l.pos))
			continue
		}

		if unicode.IsLetter(ch) || ch == '_' {
			start := l.idx
			for l.idx < uint(len(l.input)) {
				next := rune(l.input[l.idx])
				if unicode.IsLetter(next) || unicode.IsDigit(next) || next == '_' {
					l.idx++
					l.pos.Column++
				} else {
					break
				}
			}
			value := l.input[start:l.idx]
			switch value {
			case "fun":
				tokens = append(tokens, makeToken(FUN, "fun", l.pos))
			case "if":
				tokens = append(tokens, makeToken(IF, "if", l.pos))
			case "elif":
				tokens = append(tokens, makeToken(ELSEIF, "elif", l.pos))
			case "else":
				tokens = append(tokens, makeToken(ELSE, "else", l.pos))
			case "var":
				tokens = append(tokens, makeToken(VAR, "var", l.pos))
			case "pub":
				tokens = append(tokens, makeToken(PUB, "pub", l.pos))
			case "const":
				tokens = append(tokens, makeToken(CONST, "const", l.pos))
			case "import":
				tokens = append(tokens, makeToken(IMPORT, "import", l.pos))
			case "new":
				tokens = append(tokens, makeToken(NEW, "new", l.pos))
			case "and":
				tokens = append(tokens, makeToken(AND, "and", l.pos))
			case "or":
				tokens = append(tokens, makeToken(OR, "or", l.pos))
			case "true":
				tokens = append(tokens, makeToken(TRUE, "true", l.pos))
			case "false":
				tokens = append(tokens, makeToken(FALSE, "false", l.pos))
			case "return":
				tokens = append(tokens, makeToken(RETURN, "return", l.pos))
			case "while":
				tokens = append(tokens, makeToken(WHILE, "while", l.pos))
			case "do":
				tokens = append(tokens, makeToken(DO, "do", l.pos))
			case "for":
				tokens = append(tokens, makeToken(FOR, "for", l.pos))
			default:
				tokens = append(tokens, makeToken(IDENTIFIER, value, l.pos))
			}
			continue
		}

		if unicode.IsSpace(ch) {
			if ch == '\n' {
				tokens = append(tokens, makeToken(NEWLINE, "\n", l.pos))
				l.pos.Line++
				l.pos.Column = 1
			} else {
				l.pos.Column++
			}
			l.idx++
			continue
		}

		switch ch {
		case '=':
			if l.Peek() == '=' {
				l.idx += 2
				l.pos.Column += 2
				tokens = append(tokens, makeToken(EQUALSEQUALS, "==", l.pos))
			} else {
				l.idx++
				l.pos.Column++
				tokens = append(tokens, makeToken(EQUAL, string(ch), l.pos))
			}
		case '!':
			if l.Peek() == '=' {
				l.idx += 2
				l.pos.Column += 2
				tokens = append(tokens, makeToken(NOT_EQUALS, "!=", l.pos))
			} else {
				l.idx++
				l.pos.Column++
				tokens = append(tokens, makeToken(NOT, string(ch), l.pos))
			}
		case '+':
			if l.Peek() == '=' {
				l.idx += 2
				l.pos.Column += 2
				tokens = append(tokens, makeToken(PLUS_EQUALS, "+=", l.pos))
			} else {
				l.idx++
				l.pos.Column++
				tokens = append(tokens, makeToken(PLUS, string(ch), l.pos))
			}
		case '-':
			if l.Peek() == '=' {
				l.idx += 2
				l.pos.Column += 2
				tokens = append(tokens, makeToken(MINUS_EQUALS, "-=", l.pos))
			} else {
				l.idx++
				l.pos.Column++
				tokens = append(tokens, makeToken(DASH, string(ch), l.pos))
			}
		case '*':
			l.idx++
			l.pos.Column++
			tokens = append(tokens, makeToken(STAR, string(ch), l.pos))
		case '/':
			l.idx++
			l.pos.Column++
			tokens = append(tokens, makeToken(SLASH, string(ch), l.pos))
		case '%':
			l.idx++
			l.pos.Column++
			tokens = append(tokens, makeToken(MOD, string(ch), l.pos))
		case '<':
			if l.Peek() == '=' {
				l.idx += 2
				l.pos.Column += 2
				tokens = append(tokens, makeToken(LESS_EQUALS, "<=", l.pos))
			} else {
				l.idx++
				l.pos.Column++
				tokens = append(tokens, makeToken(LESS, string(ch), l.pos))
			}
		case '>':
			if l.Peek() == '=' {
				l.idx += 2
				l.pos.Column += 2
				tokens = append(tokens, makeToken(GREATER_EQUALS, ">=", l.pos))
			} else {
				l.idx++
				l.pos.Column++
				tokens = append(tokens, makeToken(GREATER, string(ch), l.pos))
			}
		case '(':
			l.idx++
			l.pos.Column++
			tokens = append(tokens, makeToken(LPAREN, string(ch), l.pos))
		case ')':
			l.idx++
			l.pos.Column++
			tokens = append(tokens, makeToken(RPAREN, string(ch), l.pos))
		case '[':
			l.idx++
			l.pos.Column++
			tokens = append(tokens, makeToken(LSQUARE, string(ch), l.pos))
		case ']':
			l.idx++
			l.pos.Column++
			tokens = append(tokens, makeToken(RSQUARE, string(ch), l.pos))
		case '{':
			l.idx++
			l.pos.Column++
			tokens = append(tokens, makeToken(LCURLY, string(ch), l.pos))
		case '}':
			l.idx++
			l.pos.Column++
			tokens = append(tokens, makeToken(RCURLY, string(ch), l.pos))
		case ';':
			l.idx++
			l.pos.Column++
			tokens = append(tokens, makeToken(SEMICOLON, string(ch), l.pos))
		case ',':
			l.idx++
			l.pos.Column++
			tokens = append(tokens, makeToken(COMMA, string(ch), l.pos))
		case ':':
			l.idx++
			l.pos.Column++
			tokens = append(tokens, makeToken(COLON, string(ch), l.pos))
		case '.':
			if l.Peek() == '.' {
				l.idx += 2
				l.pos.Column += 2
				tokens = append(tokens, makeToken(DOTDOT, "..", l.pos))
			} else {
				l.idx++
				l.pos.Column++
				tokens = append(tokens, makeToken(DOT, ".", l.pos))
			}

		case '\'':
			{
				start := l.idx
				l.idx++ // Skip opening quote
				l.pos.Column++

				for l.idx < uint(len(l.input)) {
					ch := rune(l.input[l.idx])
					if ch == '\'' {
						break
					}
					l.idx++
					l.pos.Column++
				}

				if l.idx >= uint(len(l.input)) || rune(l.input[l.idx]) != '"' {
					tokens = append(tokens, makeToken(ERRINVALIDCHAR, l.input[start:l.idx], l.pos))
					continue
				}

				l.idx++ // Skip closing quote
				l.pos.Column++

				raw := l.input[start:l.idx]
				value, err := strconv.Unquote(raw)
				if err != nil {
					tokens = append(tokens, makeToken(ERRINVALIDCHAR, raw, l.pos))
					continue
				}
				tokens = append(tokens, makeToken(STRING, value, l.pos))
			}

		case '#':
			{
				for l.idx < uint(len(l.input)) {
					ch := rune(l.input[l.idx])
					if ch == '\n' {
						break
					}
					l.idx++
					l.pos.Column++
				}

				// If end of input reached after comment, break out cleanly
				if l.idx >= uint(len(l.input)) {
					break
				}

				// Handle newline after comment
				if rune(l.input[l.idx]) == '\n' {
					tokens = append(tokens, makeToken(NEWLINE, "\n", l.pos))
					l.idx++
					l.pos.Line++
					l.pos.Column = 1
				}

				continue

			}

		case '"':
			{
				start := l.idx
				l.idx++ // Skip opening quote
				l.pos.Column++

				for l.idx < uint(len(l.input)) {
					ch := rune(l.input[l.idx])
					if ch == '"' {
						break
					}
					l.idx++
					l.pos.Column++
				}

				if l.idx >= uint(len(l.input)) || rune(l.input[l.idx]) != '"' {
					tokens = append(tokens, makeToken(ERRINVALIDCHAR, l.input[start:l.idx], l.pos))
					continue
				}

				l.idx++ // Skip closing quote
				l.pos.Column++

				raw := l.input[start:l.idx]
				value, err := strconv.Unquote(raw)
				if err != nil {
					tokens = append(tokens, makeToken(ERRINVALIDSTRING, raw, l.pos))
					continue
				}
				tokens = append(tokens, makeToken(STRING, value, l.pos))
			}

		default:
			ch := rune(l.input[l.idx])
			l.idx++
			l.pos.Column++
			tokens = append(tokens, makeToken(ERRINVALID, string(ch), l.pos))

		}
	}

	tokens = append(tokens, makeToken(EOF, "", Position{Line: 0, Column: 0}))
	return tokens
}

func (l *Lexer) Peek() rune {
	if int(l.idx+1) >= len(l.input) {
		return 0
	}
	return rune(l.input[l.idx+1])
}
