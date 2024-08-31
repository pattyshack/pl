package parser

import (
	"fmt"
	"io"
	"unicode"
	"unicode/utf8"

	"github.com/pattyshack/gt/lexutil"
	"github.com/pattyshack/gt/stringutil"
)

const (
	spacesToken       = SymbolId(-1) // [ \t]+
	lineCommentToken  = SymbolId(-2)
	blockCommentToken = SymbolId(-3)

	defaultInitialPeekWindowSize = 64
)

var (
	keywords = map[string]SymbolId{
		"true":        TrueToken,
		"false":       FalseToken,
		"if":          IfToken,
		"else":        ElseToken,
		"switch":      SwitchToken,
		"case":        CaseToken,
		"default":     DefaultToken,
		"for":         ForToken,
		"do":          DoToken,
		"in":          InToken,
		"return":      ReturnToken,
		"break":       BreakToken,
		"continue":    ContinueToken,
		"fallthrough": FallthroughToken,
		"package":     PackageToken,
		"import":      ImportToken,
		"as":          AsToken,
		"unsafe":      UnsafeToken,
		"type":        TypeToken,
		"implements":  ImplementsToken,
		"struct":      StructToken,
		"enum":        EnumToken,
		"trait":       TraitToken,
		"func":        FuncToken,
		"async":       AsyncToken,
		"defer":       DeferToken,
		"var":         VarToken,
		"let":         LetToken,
		"and":         AndToken,
		"or":          OrToken,
		"not":         NotToken,
	}
)

type RawLexerOptions struct {
	PreserveCommentContent bool

	InitialLookAheadBufferSize int

	// Initial peek window size used for lexing a variable length token
	// (Only used for testing)
	initialPeekWindowSize int
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
	if options.initialPeekWindowSize <= 0 {
		options.initialPeekWindowSize = defaultInitialPeekWindowSize
	}

	internPool := stringutil.NewInternPool()
	for kw, _ := range keywords {
		internPool.Intern(kw)
	}

	return RawLexer{
		RawLexerOptions: options,
		BufferedByteLocationReader: lexutil.NewBufferedByteLocationReader(
			sourceFileName,
			sourceContent,
			options.InitialLookAheadBufferSize),
		InternPool: internPool,
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

func (lexer *RawLexer) lexSpacesToken() (Token, error) {
	hasMore := true
	peekSize := lexer.initialPeekWindowSize
	numSpaceBytes := 0

	for hasMore {
		peeked, err := lexer.Peek(peekSize)
		if len(peeked) > 0 && err == io.EOF {
			hasMore = false
			err = nil
		}
		if err != nil {
			return nil, err
		}

		for numSpaceBytes < len(peeked) {
			char := peeked[numSpaceBytes]
			if char == ' ' || char == '\t' {
				numSpaceBytes++
			} else {
				hasMore = false
				break
			}
		}

		peekSize *= 2
	}

	if numSpaceBytes == 0 {
		panic("This should never happen")
	}

	loc := Location(lexer.Location)

	_, err := lexer.Discard(numSpaceBytes)
	if err != nil {
		panic("This should never happen")
	}

	return GenericSymbol{
		SymbolId: spacesToken,
		Location: loc,
	}, nil
}

func (lexer *RawLexer) lexNewlinesToken() (Token, error) {
	hasMore := true
	peekSize := lexer.initialPeekWindowSize
	numNewlineBytes := 0
	numNewlines := 0
	foundInvalidNewline := false

	for hasMore {
		peeked, err := lexer.Peek(peekSize)
		if len(peeked) > 0 && err == io.EOF {
			hasMore = false
			err = nil
		}
		if err != nil {
			return nil, err
		}

		for numNewlineBytes < len(peeked) {
			char := peeked[numNewlineBytes]
			if char == '\n' {
				numNewlineBytes++
				numNewlines++

			} else if char == '\r' {
				if numNewlineBytes+1 >= len(peeked) {
					// If hasMore is true, then we may have read half of the \r\n pair.
					// The next outer loop iteration will take care of this. If hasMore
					// is false, then this is a lone \r and it should not be included
					// as part of this token.
					break
				}

				if peeked[numNewlineBytes+1] == '\n' {
					numNewlineBytes += 2
					numNewlines++
				} else { // '\r' not paired with '\n'
					foundInvalidNewline = true
					hasMore = false
					break
				}
			} else {
				hasMore = false
				break
			}
		}

		peekSize *= 2
	}

	loc := Location(lexer.Location)
	if numNewlineBytes > 0 {
		_, err := lexer.Discard(numNewlineBytes)
		if err != nil {
			panic("This should never happen")
		}

		return CountSymbol{
			SymbolId: NewlinesToken,
			Location: loc,
			Count:    numNewlines,
		}, nil
	} else if foundInvalidNewline {
		_, err := lexer.Discard(1)
		if err != nil {
			panic("This should never happen")
		}

		return ParseErrorSymbol{
			Error:    fmt.Errorf("unexpected utf8 rune"),
			Location: loc,
		}, nil
	}

	panic("This should never happen")
}

func (lexer *RawLexer) lexLineCommentToken() (Token, error) {
	hasMore := true
	peekSize := lexer.initialPeekWindowSize
	numCommentBytes := 2 // "//" already peeked by peekNextToken

	for hasMore {
		peeked, err := lexer.Peek(peekSize)
		if len(peeked) > 0 && err == io.EOF {
			hasMore = false
			err = nil
		}
		if err != nil {
			return nil, err
		}

		for numCommentBytes < len(peeked) {
			char := peeked[numCommentBytes]
			if char == '\n' || char == '\r' {
				hasMore = false
				break
			}

			numCommentBytes++
		}

		peekSize *= 2
	}

	loc := Location(lexer.Location)

	value := ""
	if lexer.PreserveCommentContent {
		peeked, _ := lexer.Peek(numCommentBytes)
		// Don't intern comment string since duplicates are unlikely
		value = string(peeked)
	}

	_, err := lexer.Discard(numCommentBytes)
	if err != nil {
		panic("This should never happen")
	}

	return ValueSymbol{
		SymbolId: lineCommentToken,
		Location: loc,
		Value:    value,
	}, nil
}

func (lexer *RawLexer) lexBlockCommentToken() (Token, error) {
	panic("TODO")
	// err msg "comment not terminated"
}

func (lexer *RawLexer) lexIntegerOrFloatLiteralToken() (Token, error) {
	panic("TODO")
}

func (lexer *RawLexer) lexDotDecimalFloatLiteralToken() (Token, error) {
	panic("TODO")
}

func (lexer *RawLexer) lexRuneLiteralToken() (Token, error) {
	panic("TODO")
}

func (lexer *RawLexer) lexStringLiteralToken() (Token, error) {
	panic("TODO")
}

// offset = 0 for Identifier, 1 for jump label
func (lexer *RawLexer) peekIdentifier(offset int) (int, error) {
	hasMore := true
	peekSize := lexer.initialPeekWindowSize
	numIdentifierBytes := 0

	for hasMore {
		peeked, err := lexer.Peek(offset + peekSize)
		if len(peeked) > 0 && err == io.EOF {
			hasMore = false
			err = nil
		}
		if err != nil {
			return 0, err
		}

		if len(peeked) < offset+numIdentifierBytes {
			panic("should never happen")
		}

		remaining := peeked[offset+numIdentifierBytes:]

		for len(remaining) > 0 {
			utf8Char, size := utf8.DecodeRune(remaining)
			if utf8Char == utf8.RuneError {
				if len(remaining) < 4 && hasMore {
					// The rune may have been chopped off in the middle by Peek.  Read
					// more bytes and try again.
					break
				} else {
					// Encountered a real invalid utf8 byte
					hasMore = false
					break
				}
			}

			if size == 0 {
				panic("should never happen")
			}

			if unicode.IsLetter(utf8Char) ||
				utf8Char == '_' ||
				(numIdentifierBytes > 0 && unicode.IsNumber(utf8Char)) {

				numIdentifierBytes += size
				remaining = remaining[size:]
			} else {
				hasMore = false
				break
			}
		}

		peekSize *= 2
	}

	return numIdentifierBytes, nil
}

func (lexer *RawLexer) lexJumpLabelToken() (Token, error) {
	// Note: we can skip checking the first byte, which we alreay know is '@'

	size, err := lexer.peekIdentifier(1)
	if err != nil {
		return nil, err
	}

	loc := Location(lexer.Location)

	if size == 0 {
		_, err := lexer.Discard(1)
		if err != nil {
			panic("Should never happen")
		}

		return ParseErrorSymbol{
			Error:    fmt.Errorf("invalid label. no label name."),
			Location: loc,
		}, nil
	}

	size++

	jumpLabelBytes, err := lexer.Peek(size)
	if err != nil || len(jumpLabelBytes) != size {
		panic("Should never happen")
	}

	jumpLabel := lexer.InternBytes(jumpLabelBytes)

	_, err = lexer.Discard(size)
	if err != nil {
		panic("Should never happen")
	}

	return ValueSymbol{
		SymbolId: JumpLabelToken,
		Location: loc,
		Value:    jumpLabel,
	}, nil
}

func (lexer *RawLexer) lexIdentifierKeywordsOrLabelDeclToken() (
	ValueSymbol,
	error,
) {
	size, err := lexer.peekIdentifier(0)
	if err != nil {
		return ValueSymbol{}, err
	}

	if size == 0 {
		panic("Should never hapapen")
	}

	symbolId := IdentifierToken

	peeked, _ := lexer.Peek(size + 1)
	if size+1 == len(peeked) && peeked[size] == '@' {
		size++
		symbolId = LabelDeclToken
	}

	loc := Location(lexer.Location)
	value := lexer.InternBytes(peeked[:size])

	_, err = lexer.Discard(size)
	if err != nil {
		panic("Should never happen")
	}

	kwSymbolId, ok := keywords[value]
	if ok {
		symbolId = kwSymbolId
	}

	return ValueSymbol{
		SymbolId: symbolId,
		Location: loc,
		Value:    value,
	}, nil
}

func (lexer *RawLexer) Next() (Token, error) {
	symbolId, size, err := lexer.peekNextToken()
	if err != nil {
		return nil, err
	}

	if symbolId == ParseErrorToken {
		loc := Location(lexer.Location)

		_, err := lexer.Discard(size)
		if err != nil {
			panic("Should never happen")
		}

		return ParseErrorSymbol{
			Error:    fmt.Errorf("unexpected utf8 rune"),
			Location: loc,
		}, nil
	}

	if size > 0 {
		loc := Location(lexer.Location)

		_, err := lexer.Discard(size)
		if err != nil {
			panic("Should never happen")
		}

		return GenericSymbol{
			SymbolId: symbolId,
			Location: loc,
		}, nil
	}

	// variable length tokens
	switch symbolId {
	case spacesToken:
		return lexer.lexSpacesToken()
	case NewlinesToken:
		return lexer.lexNewlinesToken()
	case lineCommentToken:
		return lexer.lexLineCommentToken()
	case blockCommentToken:
		return lexer.lexBlockCommentToken()
	case IntegerLiteralToken:
		return lexer.lexIntegerOrFloatLiteralToken()
	case FloatLiteralToken:
		return lexer.lexDotDecimalFloatLiteralToken()
	case RuneLiteralToken:
		return lexer.lexRuneLiteralToken()
	case StringLiteralToken:
		return lexer.lexStringLiteralToken()
	case JumpLabelToken:
		return lexer.lexJumpLabelToken()
	case IdentifierToken:
		return lexer.lexIdentifierKeywordsOrLabelDeclToken()
	}

	panic(fmt.Sprintf("RawLexer: unhandled variable length token: %v", symbolId))
}
