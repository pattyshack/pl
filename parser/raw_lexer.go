package parser

import (
	//"fmt"
	"io"
	"unicode/utf8"

	"github.com/pattyshack/gt/lexutil"
	"github.com/pattyshack/gt/stringutil"
)

const (
	spacesToken       = SymbolId(-1) // [ \t]+
	lineCommentToken  = SymbolId(-2)
	blockCommentToken = SymbolId(-3)
)

type RawLexerOptions struct {
	PreserveCommentContent bool

	InitialLookAheadBufferSize int
}

type RawLexer struct {
	RawLexerOptions

	lexutil.BufferedByteLocationReader
	*stringutil.InternPool
}

func NewRawLexer(
	sourceFileName string,
	sourceContent io.Reader,
	options RawLexerOptions,
) RawLexer {
	return RawLexer{
		RawLexerOptions: options,
		BufferedByteLocationReader: lexutil.NewBufferedByteLocationReader(
			sourceFileName,
			sourceContent,
			options.InitialLookAheadBufferSize),
		InternPool: stringutil.NewInternPool(),
	}
}

// variable length token will return 0 as length
func (lexer *RawLexer) peekNextToken() (SymbolId, int, error) {
	peeked, err := lexer.Peek(7)
	if len(peeked) > 0 && err == io.EOF {
		err = nil
	}
	if err != nil {
		return 0, 0, err
	}

	if len(peeked) == 0 {
		panic("This should never happen")
	}

	char := peeked[0]

	if ('a' <= char && char < 'r') ||
		// note that 'r' could be a raw string literal
		('r' < char && char <= 'z') ||
		('A' <= char && char <= 'Z') ||
		char == '_' {

		return IdentifierToken, 0, nil
	}

	if '0' <= char && char <= '9' {
		return IntegerLiteralToken, 0, nil // int or float
	}

	switch char {
	case '`', '"':
		return StringLiteralToken, 0, nil
	case 'r':
		if len(peeked) > 1 && (peeked[1] == '`' || peeked[1] == '"') {
			return StringLiteralToken, 0, nil
		}

		return IdentifierToken, 0, nil
	case '\'':
		return RuneLiteralToken, 0, nil
	case '\r', '\n':
		return NewlinesToken, 0, nil
	case ' ', '\t':
		return spacesToken, 0, nil
	case '@':
		return JumpLabelToken, 0, nil

	case '{':
		return LbraceToken, 1, nil
	case '}':
		return RbraceToken, 1, nil
	case '(':
		return LparenToken, 1, nil
	case ')':
		return RparenToken, 1, nil
	case '[':
		return LbracketToken, 1, nil
	case ']':
		return RbracketToken, 1, nil
	case '.':
		if len(peeked) > 2 && peeked[1] == '.' && peeked[2] == '.' {
			return DotDotDotToken, 3, nil
		}

		if len(peeked) > 1 && '0' <= peeked[1] && peeked[1] <= '9' {
			return FloatLiteralToken, 0, nil
		}

		return DotToken, 1, nil
	case ',':
		return CommaToken, 1, nil
	case '?':
		return QuestionToken, 1, nil
	case ';':
		return SemicolonToken, 1, nil
	case ':':
		return ColonToken, 1, nil
	case '$':
		if len(peeked) > 1 && peeked[1] == '[' {
			return DollarLbracketToken, 2, nil
		}
	case '+':
		if len(peeked) > 1 {
			if peeked[1] == '+' {
				return AddOneAssignToken, 2, nil

			} else if peeked[1] == '=' {
				return AddAssignToken, 2, nil
			}
		}

		return AddToken, 1, nil
	case '-':
		if len(peeked) > 1 {
			if peeked[1] == '-' {
				return SubOneAssignToken, 2, nil

			} else if peeked[1] == '=' {
				return SubAssignToken, 2, nil
			}
		}

		return SubToken, 1, nil
	case '*':
		if len(peeked) > 1 && peeked[1] == '=' {
			return MulAssignToken, 2, nil
		}

		return MulToken, 1, nil
	case '/':
		if len(peeked) > 1 {
			if peeked[1] == '/' {
				return lineCommentToken, 0, nil
			} else if peeked[1] == '*' {
				return blockCommentToken, 0, nil
			} else if peeked[1] == '=' {
				return DivAssignToken, 2, nil
			}
		}

		return DivToken, 1, nil
	case '%':
		if len(peeked) > 1 && peeked[1] == '=' {
			return ModAssignToken, 2, nil
		}

		return ModToken, 1, nil
	case '~':
		if len(peeked) > 1 {
			if peeked[1] == '=' {
				return BitNegAssignToken, 2, nil

			} else if peeked[1] == '~' {
				return TildeTildeToken, 2, nil
			}
		}

		return BitNegToken, 1, nil
	case '&':
		if len(peeked) > 1 && peeked[1] == '=' {
			return BitAndAssignToken, 2, nil
		}

		return BitAndToken, 1, nil
	case '^':
		if len(peeked) > 1 && peeked[1] == '=' {
			return BitXorAssignToken, 2, nil
		}

		return BitXorToken, 1, nil
	case '|':
		if len(peeked) > 1 && peeked[1] == '=' {
			return BitOrAssignToken, 2, nil
		}

		return BitOrToken, 1, nil
	case '=':
		if len(peeked) > 1 && peeked[1] == '=' {
			return EqualToken, 2, nil
		}

		return AssignToken, 1, nil
	case '!':
		if len(peeked) > 1 && peeked[1] == '=' {
			return NotEqualToken, 2, nil
		}

		return ExclaimToken, 1, nil
	case '<':
		if len(peeked) > 1 {
			if peeked[1] == '=' {
				return LessOrEqualToken, 2, nil

			} else if peeked[1] == '<' {
				if len(peeked) > 2 && peeked[2] == '=' {
					return BitLshiftAssignToken, 3, nil
				}

				return BitLshiftToken, 2, nil
			}
		}

		return LessToken, 1, nil
	case '>':
		if len(peeked) > 1 {
			if peeked[1] == '=' {
				return GreaterOrEqualToken, 2, nil

			} else if peeked[1] == '>' {
				if len(peeked) > 2 && peeked[2] == '=' {
					return BitRshiftAssignToken, 3, nil
				}

				return BitRshiftToken, 2, nil
			}
		}

		return GreaterToken, 1, nil
	}

	utf8Char, size := utf8.DecodeRune(peeked)
	if size == 1 || utf8Char == utf8.RuneError { // unexpected token
		return ParseErrorToken, size, nil
	}

	return IdentifierToken, 0, nil
}

func (lexer *RawLexer) Next() (Token, error) {
	symbolId, size, err := lexer.peekNextToken()
	if err != nil {
		return nil, err
	}

	if symbolId == ParseErrorToken {
		// ParseErrorToken
		panic("TODO")
	}

	if size > 0 {
		lexer.Discard(size)
		return &GenericSymbol{
			SymbolId: symbolId,
			Location: Location(lexer.Location),
		}, nil
	}

	// variable length tokens
	panic("TODO")
}
