package parser

import (
	"fmt"
	"io"
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
	numSpaceBytes, err := lexutil.PeekSpaces(
		lexer.BufferedByteLocationReader,
		lexer.initialPeekWindowSize)
	if err != nil {
		return nil, err
	}

	if numSpaceBytes == 0 {
		panic("This should never happen")
	}

	loc := Location(lexer.Location)

	_, err = lexer.Discard(numSpaceBytes)
	if err != nil {
		panic("This should never happen")
	}

	return GenericSymbol{
		SymbolId: spacesToken,
		Location: loc,
	}, nil
}

func (lexer *RawLexer) lexNewlinesToken() (Token, error) {
	numNewlineBytes, numNewlines, foundInvalidNewline, err := lexutil.PeekNewlines(
		lexer.BufferedByteLocationReader,
		lexer.initialPeekWindowSize)
	if err != nil {
		return nil, err
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
	numCommentBytes, err := lexutil.PeekLineComment(
		lexer.BufferedByteLocationReader,
		lexer.initialPeekWindowSize)
	if err != nil {
		return nil, err
	}

	loc := Location(lexer.Location)

	value := ""
	if lexer.PreserveCommentContent {
		peeked, _ := lexer.Peek(numCommentBytes)
		// Don't intern comment string since duplicates are unlikely
		value = string(peeked)
	}

	_, err = lexer.Discard(numCommentBytes)
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
	numCommentBytes, scopeLevel, err := lexutil.PeekBlockComment(
		lexer.BufferedByteLocationReader,
		true,
		lexer.initialPeekWindowSize)
	if err != nil {
		return nil, err
	}

	loc := Location(lexer.Location)

	value := ""
	if lexer.PreserveCommentContent && scopeLevel == 0 {
		peeked, _ := lexer.Peek(numCommentBytes)
		// Don't intern comment string since duplicates are unlikely
		value = string(peeked)
	}

	_, err = lexer.Discard(numCommentBytes)
	if err != nil {
		panic("This should never happen")
	}

	if scopeLevel > 0 {
		return ParseErrorSymbol{
			Error:    fmt.Errorf("block comment not terminated"),
			Location: loc,
		}, nil
	}

	return ValueSymbol{
		SymbolId: blockCommentToken,
		Location: loc,
		Value:    value,
	}, nil
}

func (lexer *RawLexer) peekDigits(
	offset int,
	isDigit func(byte) bool,
) (
	int,
	error,
) {
	hasMore := true
	peekSize := lexer.initialPeekWindowSize
	numBytes := 0

	for hasMore {
		peeked, err := lexer.Peek(offset + peekSize)
		if len(peeked) >= offset && err == io.EOF {
			hasMore = false
			err = nil
		}
		if err != nil {
			return 0, err
		}

		remaining := peeked[offset+numBytes:]
		for len(remaining) > 0 {
			char := remaining[0]
			if isDigit(char) {
				numBytes++
				remaining = remaining[1:]
			} else if char == '_' {
				if len(remaining) < 2 {
					if hasMore {
						// read more bytes
						break
					} else { // the int is followed by a '_'
						hasMore = false
						break
					}
				}

				if isDigit(remaining[1]) {
					numBytes += 2
					remaining = remaining[2:]
				} else { // the int is followed by a '_'
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

	return numBytes, nil
}

func (lexer *RawLexer) isBinaryDigit(char byte) bool {
	return char == '0' || char == '1'
}

func (lexer *RawLexer) isOctalDigit(char byte) bool {
	return '0' <= char && char <= '7'
}

func (lexer *RawLexer) isDecimalDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (lexer *RawLexer) isHexadecimalDigit(char byte) bool {
	return '0' <= char && char <= '9' ||
		'A' <= char && char <= 'F' ||
		'a' <= char && char <= 'f'
}

func (lexer *RawLexer) lexIntegerOrFloatLiteralToken() (Token, error) {
	peeked, err := lexer.Peek(2)
	if len(peeked) > 0 && err == io.EOF {
		err = nil
	}
	if err != nil {
		return nil, err
	}

	char := peeked[0]
	if !lexer.isDecimalDigit(char) {
		panic("should never happen")
	}

	loc := Location(lexer.Location)
	subType := DecimalInteger
	hasDigits := true
	totalBytes := 1

	if char != '0' {
		numBytes, err := lexer.peekDigits(1, lexer.isDecimalDigit)
		if err != nil {
			return nil, err
		}

		totalBytes = numBytes + 1
	} else if len(peeked) > 1 {
		switch peeked[1] {
		case 'b', 'B':
			numBytes, err := lexer.peekDigits(2, lexer.isBinaryDigit)
			if err != nil {
				return nil, err
			}

			subType = BinaryInteger
			hasDigits = numBytes != 0
			totalBytes = numBytes + 2
		case 'o', 'O':
			numBytes, err := lexer.peekDigits(2, lexer.isOctalDigit)
			if err != nil {
				return nil, err
			}

			subType = ZeroOPrefixedOctalInteger
			hasDigits = numBytes != 0
			totalBytes = numBytes + 2
		case 'x', 'X':
			numBytes, err := lexer.peekDigits(2, lexer.isHexadecimalDigit)
			if err != nil {
				return nil, err
			}

			subType = HexadecimalInteger
			hasDigits = numBytes != 0
			totalBytes = numBytes + 2
		default:
			numBytes, err := lexer.peekDigits(1, lexer.isOctalDigit)
			if err != nil {
				return nil, err
			}

			if numBytes > 0 { // otherwise is a decimal "0"
				subType = ZeroPrefixedOctalInteger
				hasDigits = true
				totalBytes = numBytes + 1
			}
		}
	}

	// TODO handle float

	value := ""
	if hasDigits {
		peeked, err = lexer.Peek(totalBytes)
		if err != nil {
			panic("should never happen")
		}

		value = string(peeked)
	}

	_, err = lexer.Discard(totalBytes)
	if err != nil {
		panic("should never happen")
	}

	if !hasDigits {
		return ParseErrorSymbol{
			Location: loc,
			Error:    fmt.Errorf("%s integer has no digits", subType),
		}, nil
	}

	return IntegerLiteralSymbol{
		Location:       loc,
		Value:          value,
		IntegerSubType: subType,
	}, nil
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

func (lexer *RawLexer) lexJumpLabelToken() (Token, error) {
	// Note: we can skip checking the first byte, which we alreay know is '@'

	size, err := lexutil.PeekIdentifier(
		lexer.BufferedByteLocationReader,
		1,
		lexer.initialPeekWindowSize)
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
			Error:    fmt.Errorf("no label name associated with @"),
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
	size, err := lexutil.PeekIdentifier(
		lexer.BufferedByteLocationReader,
		0,
		lexer.initialPeekWindowSize)
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
