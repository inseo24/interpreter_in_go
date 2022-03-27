package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// 특수 타입
	ILLEGAL = "ILLEGAL" // 렉서가 어떤 토큰이나 문자를 알 수 없을 때
	EOF     = "EOF"     // EOF : End Of File, 파일의 끝. 파서에게 이만 멈춰도 좋다~라고 얘기하는 용도.

	// 식별자 + 리터럴
	IDENT = "IDENT" // add, foobar, x, y ...(identifier)
	INT   = "INT"

	// 연산자
	ASSIGN = "="
	PLUS   = "+"

	// 구분자
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// 예약어
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
