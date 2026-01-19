package lexer

import "github.com/goholic/gvm/token"

// Deterministic Finite Automation
// How I understand?

// I see l. I move to the "Potential Keyword" state.
// I see e. I stay in the "Potential Keyword" state.
// I see t. I stay in the "Potential Keyword" state.
// I see (space). The game ends. I check my current location.
// I am at let.
// I yell: "I found a Keyword!"

// If the input was lettuce instead of let
// ... l, e, t ...

// I see t. I stay in the state.
// I see u.
// I realize this isn't the keyword let.
// I switch to the "Identifier" state (variable name).

// define a struct to hold
// input, position, readposition (curr + 1), ch (curr)
type Lexer struct {
	input        string
	position     int
	readposition int
	ch           byte
}

// Lexer Constructos
func New(input string) *Lexer {
	// pointer to specific Lexer
	// with input & others = default zero
	l := &Lexer{input: input}

	// TODO:
	// initialize the first char
	// postion = 0
	// readPosition = 0
	// "let"
	// l.ch = l.input[0] = l
	// l.position = l.readPosition = 0
	// l.readPosition = 1

	// ^ Done
	l.readChar()

	// return the ptr
	return l
}

func (l *Lexer) readChar() {
	// curr + 1 does not exist
	// we reached the end
	if l.readposition >= len(l.input) {
		l.ch = 0 // EOF
	} else {
		l.ch = l.input[l.readposition]
	}

	l.position = l.readposition
	l.readposition += 1
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// utility
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

// nextToken
// look at curr and emit token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// TODO:
	// Eat Spaces, tabs, \n
	l.skipWhiteSpace()

	// TODO:
	// newToken
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		tok = newToken(token.BANG, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:

		// TODO:
		// isLettr
		if isLetter(l.ch) {
			// TODO:
			// l.readIdentifier
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookUpIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok

}
