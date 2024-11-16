package lexer

import (
	"fmt"
	"io"
	"unicode/utf8"

	"github.com/pattyshack/gt/parseutil"
	"github.com/pattyshack/gt/stringutil"

	"github.com/pattyshack/pl/parser/lr"
)

const (
	spacesToken       = lr.SymbolId(-1) // [ \t]+
	lineCommentToken  = lr.SymbolId(-2)
	blockCommentToken = lr.SymbolId(-3)

	// s - single-line  vs m - multiple-line
	// i - interpreted  vs r - raw string
	// d - double quote vs b - back quote

	sidStringToken = lr.SymbolId(-10)
	sibStringToken = lr.SymbolId(-11)
	srdStringToken = lr.SymbolId(-12)
	srbStringToken = lr.SymbolId(-13)
	midStringToken = lr.SymbolId(-14)
	mibStringToken = lr.SymbolId(-15)
	mrdStringToken = lr.SymbolId(-16)
	mrbStringToken = lr.SymbolId(-17)

	defaultInitialPeekWindowSize = 64
)

var (
	keywords = map[string]lr.SymbolId{
		"true":        lr.TrueToken,
		"false":       lr.FalseToken,
		"if":          lr.IfToken,
		"else":        lr.ElseToken,
		"switch":      lr.SwitchToken,
		"select":      lr.SelectToken,
		"case":        lr.CaseToken,
		"default":     lr.DefaultToken,
		"for":         lr.ForToken,
		"do":          lr.DoToken,
		"in":          lr.InToken,
		"return":      lr.ReturnToken,
		"break":       lr.BreakToken,
		"continue":    lr.ContinueToken,
		"fallthrough": lr.FallthroughToken,
		"import":      lr.ImportToken,
		"unsafe":      lr.UnsafeToken,
		"type":        lr.TypeToken,
		"implements":  lr.ImplementsToken,
		"alias":       lr.AliasToken,
		"struct":      lr.StructToken,
		"enum":        lr.EnumToken,
		"trait":       lr.TraitToken,
		"func":        lr.FuncToken,
		"async":       lr.AsyncToken,
		"defer":       lr.DeferToken,
		"var":         lr.VarToken,
		"let":         lr.LetToken,
		"make":        lr.MakeToken,
		"as":          lr.AsToken,
		"and":         lr.AndToken,
		"or":          lr.OrToken,
		"not":         lr.NotToken,
		"_":           lr.UnderscoreToken,
	}
)

type LexerOptions struct {
	PreserveCommentContent bool

	InitialLookAheadBufferSize int

	// Initial peek window size used for lexing a variable length token
	// (Only used for testing)
	initialPeekWindowSize int
}

type RawLexer struct {
	*parseutil.Emitter
	LexerOptions

	parseutil.BufferedByteLocationReader
	*stringutil.InternPool
}

func NewRawLexer(
	sourceFileName string,
	sourceContent io.Reader,
	emitter *parseutil.Emitter,
	options LexerOptions,
) lr.Lexer {
	if options.initialPeekWindowSize <= 0 {
		options.initialPeekWindowSize = defaultInitialPeekWindowSize
	}

	internPool := stringutil.NewInternPool()
	for kw, _ := range keywords {
		internPool.Intern(kw)
	}

	return &RawLexer{
		Emitter:      emitter,
		LexerOptions: options,
		BufferedByteLocationReader: parseutil.NewBufferedByteLocationReader(
			sourceFileName,
			sourceContent,
			options.InitialLookAheadBufferSize),
		InternPool: internPool,
	}
}

func (lexer *RawLexer) CurrentLocation() parseutil.Location {
	return lexer.Location
}

// variable length token will return 0 as length
func (lexer *RawLexer) peekNextToken() (lr.SymbolId, string, error) {
	peeked, err := lexer.Peek(7)
	if len(peeked) > 0 && err == io.EOF {
		err = nil
	}
	if err != nil {
		return 0, "", err
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

		return lr.IdentifierToken, "", nil
	}

	if '0' <= char && char <= '9' {
		return lr.IntegerLiteralToken, "", nil // int or float
	}

	switch char {
	case '`':
		if len(peeked) >= 3 && peeked[1] == '`' && peeked[2] == '`' {
			return mibStringToken, "", nil
		}
		return sibStringToken, "", nil
	case '"':
		if len(peeked) >= 3 && peeked[1] == '"' && peeked[2] == '"' {
			return midStringToken, "", nil
		}
		return sidStringToken, "", nil
	case 'r':
		if len(peeked) > 1 {
			char := peeked[1]
			if char == '`' {
				if len(peeked) >= 4 && peeked[2] == '`' && peeked[3] == '`' {
					return mrbStringToken, "", nil
				}

				return srbStringToken, "", nil
			} else if char == '"' {
				if len(peeked) >= 4 && peeked[2] == '"' && peeked[3] == '"' {
					return mrdStringToken, "", nil
				}

				return srdStringToken, "", nil
			}
		}

		return lr.IdentifierToken, "", nil
	case '\'':
		return lr.RuneLiteralToken, "", nil
	case '\r', '\n':
		return lr.NewlinesToken, "", nil
	case ' ', '\t':
		return spacesToken, "", nil
	case '#':
		return lr.PoundToken, "#", nil
	case '@':
		return lr.AtToken, "@", nil

	case '{':
		return lr.LbraceToken, "{", nil
	case '}':
		return lr.RbraceToken, "}", nil
	case '(':
		return lr.LparenToken, "(", nil
	case ')':
		return lr.RparenToken, ")", nil
	case '[':
		return lr.LbracketToken, "[", nil
	case ']':
		return lr.RbracketToken, "]", nil
	case '.':
		if len(peeked) > 2 && peeked[1] == '.' && peeked[2] == '.' {
			return lr.EllipsisToken, "...", nil
		}

		if len(peeked) > 1 && '0' <= peeked[1] && peeked[1] <= '9' {
			return lr.FloatLiteralToken, "", nil
		}

		return lr.DotToken, ".", nil
	case ',':
		return lr.CommaToken, ",", nil
	case '?':
		return lr.QuestionToken, "?", nil
	case ';':
		return lr.SemicolonToken, ";", nil
	case ':':
		return lr.ColonToken, ":", nil
	case '$':
		return lr.DollarToken, "$", nil
	case '+':
		if len(peeked) > 1 {
			if peeked[1] == '+' {
				return lr.AddOneAssignToken, "++", nil

			} else if peeked[1] == '=' {
				return lr.AddAssignToken, "+=", nil
			}
		}

		return lr.AddToken, "+", nil
	case '-':
		if len(peeked) > 1 {
			if peeked[1] == '-' {
				return lr.SubOneAssignToken, "--", nil

			} else if peeked[1] == '=' {
				return lr.SubAssignToken, "-=", nil
			}
		}

		return lr.SubToken, "-", nil
	case '*':
		if len(peeked) > 1 && peeked[1] == '=' {
			return lr.MulAssignToken, "*=", nil
		}

		return lr.MulToken, "*", nil
	case '/':
		if len(peeked) > 1 {
			if peeked[1] == '/' {
				return lineCommentToken, "", nil
			} else if peeked[1] == '*' {
				return blockCommentToken, "", nil
			} else if peeked[1] == '=' {
				return lr.DivAssignToken, "/=", nil
			}
		}

		return lr.DivToken, "/", nil
	case '%':
		if len(peeked) > 1 && peeked[1] == '=' {
			return lr.ModAssignToken, "%=", nil
		}

		return lr.ModToken, "%", nil
	case '~':
		if len(peeked) > 1 {
			if peeked[1] == '~' {
				return lr.TildeTildeToken, "~~", nil
			}
		}

		return lr.TildeToken, "~", nil
	case '&':
		if len(peeked) > 1 && peeked[1] == '=' {
			return lr.BitAndAssignToken, "&=", nil
		}

		return lr.BitAndToken, "&", nil
	case '^':
		if len(peeked) > 1 && peeked[1] == '=' {
			return lr.BitXorAssignToken, "^=", nil
		}

		return lr.BitXorToken, "^", nil
	case '|':
		if len(peeked) > 1 && peeked[1] == '=' {
			return lr.BitOrAssignToken, "|=", nil
		}

		return lr.BitOrToken, "|", nil
	case '=':
		if len(peeked) > 1 && peeked[1] == '=' {
			return lr.EqualToken, "==", nil
		}

		return lr.AssignToken, "=", nil
	case '!':
		if len(peeked) > 1 && peeked[1] == '=' {
			return lr.NotEqualToken, "!=", nil
		}

		return lr.ExclaimToken, "!", nil
	case '<':
		if len(peeked) > 1 {
			if peeked[1] == '=' {
				return lr.LessOrEqualToken, "<=", nil

			} else if peeked[1] == '<' {
				if len(peeked) > 2 && peeked[2] == '=' {
					return lr.BitLshiftAssignToken, "<<=", nil
				}

				return lr.BitLshiftToken, "<<", nil
			} else if peeked[1] == '-' {
				return lr.ArrowToken, "<-", nil
			}
		}

		return lr.LessToken, "<", nil
	case '>':
		if len(peeked) > 1 {
			if peeked[1] == '=' {
				return lr.GreaterOrEqualToken, ">=", nil

			} else if peeked[1] == '>' {
				if len(peeked) > 2 && peeked[2] == '=' {
					return lr.BitRshiftAssignToken, ">>=", nil
				}

				return lr.BitRshiftToken, ">>", nil
			}
		}

		return lr.GreaterToken, ">", nil
	}

	utf8Char, size := utf8.DecodeRune(peeked)
	if size == 1 || utf8Char == utf8.RuneError { // unexpected token
		return lr.ParseErrorToken, string(peeked[:size]), nil
	}

	return lr.IdentifierToken, "", nil
}

func (lexer *RawLexer) lexSpacesToken() (lr.Token, error) {
	token, err := parseutil.MaybeTokenizeSpaces(
		lexer.BufferedByteLocationReader,
		lexer.initialPeekWindowSize,
		spacesToken)
	if err != nil {
		return nil, err
	}

	if token == nil {
		panic("This should never happen")
	}

	return token, nil
}

func (lexer *RawLexer) lexNewlinesToken() (lr.Token, error) {
	token, foundInvalidNewline, err := parseutil.MaybeTokenizeNewlines(
		lexer.BufferedByteLocationReader,
		lexer.initialPeekWindowSize,
		lr.NewlinesToken)
	if err != nil {
		return nil, err
	}

	if token == nil {
		panic("This should never happen")
	}

	if foundInvalidNewline {
		return lr.NewParseErrorSymbol(
			token.StartEndPos,
			"unexpected utf8 rune"), nil
	}

	return token, nil
}

func (lexer *RawLexer) lexLineCommentToken() (lr.Token, error) {
	token, err := parseutil.MaybeTokenizeLineComment(
		lexer.BufferedByteLocationReader,
		lexer.initialPeekWindowSize,
		lineCommentToken,
		lexer.PreserveCommentContent)
	if err != nil {
		return nil, err
	}

	if token == nil {
		panic("should never happen")
	}

	return &lr.TokenValue{
		TokenValue: *token,
	}, nil
}

func (lexer *RawLexer) lexBlockCommentToken() (lr.Token, error) {
	token, notTerminated, err := parseutil.MaybeTokenizeBlockComment(
		lexer.BufferedByteLocationReader,
		true,
		lexer.initialPeekWindowSize,
		blockCommentToken,
		lexer.PreserveCommentContent)
	if err != nil {
		return nil, err
	}

	if token == nil {
		panic("should never happend")
	}

	if notTerminated {
		return lr.NewParseErrorSymbol(
			token.StartEndPos,
			"block comment not terminated"), nil
	}

	return &lr.TokenValue{
		TokenValue: *token,
	}, nil
}

func (lexer *RawLexer) lexIntegerOrFloatLiteralToken() (lr.Token, error) {
	token, hasNoDigits, err := parseutil.MaybeTokenizeIntegerOrFloatLiteral(
		lexer.BufferedByteLocationReader,
		lexer.initialPeekWindowSize,
		lexer.InternPool,
		lr.IntegerLiteralToken,
		lr.FloatLiteralToken)
	if err != nil {
		return nil, err
	}

	if token == nil {
		panic("should never happen")
	}

	if hasNoDigits {
		return lr.NewParseErrorSymbol(
			token.StartEndPos,
			"%s has no digits",
			token.SubType), nil
	}

	return &lr.TokenValue{
		TokenValue: *token,
	}, nil
}

func (lexer *RawLexer) lexRuneLiteralToken() (lr.Token, error) {
	token, errMsg, err := parseutil.MaybeTokenizeRuneLiteral(
		lexer.BufferedByteLocationReader,
		lexer.initialPeekWindowSize,
		lexer.InternPool,
		lr.RuneLiteralToken)
	if err != nil {
		return nil, err
	}

	if token == nil {
		panic("should never happen")
	}

	if errMsg != "" {
		return lr.NewParseErrorSymbol(token.StartEndPos, errMsg), nil
	}

	return &lr.TokenValue{
		TokenValue: *token,
	}, nil
}

func (lexer *RawLexer) lexStringLiteralToken(
	subType parseutil.LiteralSubType,
	useBacktickMarker bool,
) (
	lr.Token,
	error,
) {
	token, errMsg, err := parseutil.MaybeTokenizeStringLiteral(
		lexer.BufferedByteLocationReader,
		lexer.initialPeekWindowSize,
		lexer.InternPool,
		lr.StringLiteralToken,
		subType,
		useBacktickMarker)
	if err != nil {
		return nil, err
	}

	if token == nil {
		panic("should never happen")
	}

	if errMsg != "" {
		return lr.NewParseErrorSymbol(token.StartEndPos, errMsg), nil
	}

	return &lr.TokenValue{
		TokenValue: *token,
	}, nil
}

func (lexer *RawLexer) lexIdentifierOrKeywords() (
	*lr.TokenValue,
	error,
) {
	token, err := parseutil.MaybeTokenizeIdentifier(
		lexer.BufferedByteLocationReader,
		lexer.initialPeekWindowSize,
		lexer.InternPool,
		lr.IdentifierToken)
	if err != nil {
		return nil, err
	}

	if token == nil {
		panic("Should never hapapen")
	}

	kwSymbolId, ok := keywords[token.Value]
	if ok {
		token.SymbolId = kwSymbolId
	}

	return &lr.TokenValue{
		TokenValue: *token,
	}, nil
}

func (lexer *RawLexer) Next() (lr.Token, error) {
	symbolId, value, err := lexer.peekNextToken()
	if err != nil {
		return nil, err
	}

	size := len(value)

	if symbolId == lr.ParseErrorToken {
		loc := lexer.Location

		_, err := lexer.Discard(size)
		if err != nil {
			panic("Should never happen")
		}

		return lr.NewParseErrorSymbol(
			parseutil.NewStartEndPos(loc, lexer.Location),
			"unexpected utf8 rune"), nil
	}

	if size > 0 {
		loc := lexer.Location

		_, err := lexer.Discard(size)
		if err != nil {
			panic("Should never happen")
		}

		return &lr.TokenValue{
			TokenValue: parseutil.TokenValue[lr.SymbolId]{
				SymbolId:    symbolId,
				StartEndPos: parseutil.NewStartEndPos(loc, lexer.Location),
				Value:       value,
			},
		}, nil
	}

	// variable length tokens
	switch symbolId {
	case spacesToken:
		return lexer.lexSpacesToken()
	case lr.NewlinesToken:
		return lexer.lexNewlinesToken()
	case lineCommentToken:
		return lexer.lexLineCommentToken()
	case blockCommentToken:
		return lexer.lexBlockCommentToken()
	case lr.IntegerLiteralToken:
		return lexer.lexIntegerOrFloatLiteralToken()
	case lr.FloatLiteralToken:
		return lexer.lexIntegerOrFloatLiteralToken()
	case lr.RuneLiteralToken:
		return lexer.lexRuneLiteralToken()
	case sibStringToken:
		return lexer.lexStringLiteralToken(
			parseutil.SingleLineString,
			true)
	case sidStringToken:
		return lexer.lexStringLiteralToken(
			parseutil.SingleLineString,
			false)
	case srbStringToken:
		return lexer.lexStringLiteralToken(
			parseutil.RawSingleLineString,
			true)
	case srdStringToken:
		return lexer.lexStringLiteralToken(
			parseutil.RawSingleLineString,
			false)
	case mibStringToken:
		return lexer.lexStringLiteralToken(
			parseutil.MultiLineString,
			true)
	case midStringToken:
		return lexer.lexStringLiteralToken(
			parseutil.MultiLineString,
			false)
	case mrbStringToken:
		return lexer.lexStringLiteralToken(
			parseutil.RawMultiLineString,
			true)
	case mrdStringToken:
		return lexer.lexStringLiteralToken(
			parseutil.RawMultiLineString,
			false)
	case lr.IdentifierToken:
		return lexer.lexIdentifierOrKeywords()
	}

	panic(fmt.Sprintf("RawLexer: unhandled variable length token: %v", symbolId))
}
