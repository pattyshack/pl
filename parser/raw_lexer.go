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
	commentGroupToken = SymbolId(-4)

	// s - single-line  vs m - multiple-line
	// i - interpreted  vs r - raw string
	// d - double quote vs b - back quote

	sidStringToken = SymbolId(-10)
	sibStringToken = SymbolId(-11)
	srdStringToken = SymbolId(-12)
	srbStringToken = SymbolId(-13)
	midStringToken = SymbolId(-14)
	mibStringToken = SymbolId(-15)
	mrdStringToken = SymbolId(-16)
	mrbStringToken = SymbolId(-17)

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

type LexerOptions struct {
	PreserveCommentContent bool

	InitialLookAheadBufferSize int

	// Initial peek window size used for lexing a variable length token
	// (Only used for testing)
	initialPeekWindowSize int
}

type RawLexer struct {
	LexerOptions

	lexutil.BufferedByteLocationReader
	*stringutil.InternPool
}

func NewRawLexer(
	sourceFileName string,
	sourceContent io.Reader,
	options LexerOptions,
) Lexer {
	if options.initialPeekWindowSize <= 0 {
		options.initialPeekWindowSize = defaultInitialPeekWindowSize
	}

	internPool := stringutil.NewInternPool()
	for kw, _ := range keywords {
		internPool.Intern(kw)
	}

	return &RawLexer{
		LexerOptions: options,
		BufferedByteLocationReader: lexutil.NewBufferedByteLocationReader(
			sourceFileName,
			sourceContent,
			options.InitialLookAheadBufferSize),
		InternPool: internPool,
	}
}

func (lexer *RawLexer) CurrentLocation() Location {
	return Location(lexer.Location)
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
	case '`':
		if len(peeked) >= 3 && peeked[1] == '`' && peeked[2] == '`' {
			return mibStringToken, 0, nil
		}
		return sibStringToken, 0, nil
	case '"':
		if len(peeked) >= 3 && peeked[1] == '"' && peeked[2] == '"' {
			return midStringToken, 0, nil
		}
		return sidStringToken, 0, nil
	case 'r':
		if len(peeked) > 1 {
			char := peeked[1]
			if char == '`' {
				if len(peeked) >= 4 && peeked[2] == '`' && peeked[3] == '`' {
					return mrbStringToken, 0, nil
				}

				return srbStringToken, 0, nil
			} else if char == '"' {
				if len(peeked) >= 4 && peeked[2] == '"' && peeked[3] == '"' {
					return mrdStringToken, 0, nil
				}

				return srdStringToken, 0, nil
			}
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
			return EllipsisToken, 3, nil
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

	return TokenCount{
		SymbolId:    spacesToken,
		StartEndPos: newStartEndPos(loc, Location(lexer.Location)),
		Count:       numSpaceBytes,
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

		return TokenCount{
			SymbolId:    NewlinesToken,
			StartEndPos: newStartEndPos(loc, Location(lexer.Location)),
			Count:       numNewlines,
		}, nil
	} else if foundInvalidNewline {
		_, err := lexer.Discard(1)
		if err != nil {
			panic("This should never happen")
		}

		return ParseErrorSymbol{
			StartEndPos: newStartEndPos(loc, Location(lexer.Location)),
			Error:       fmt.Errorf("unexpected utf8 rune"),
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

	return TokenValue{
		SymbolId:    lineCommentToken,
		StartEndPos: newStartEndPos(loc, Location(lexer.Location)),
		Value:       value,
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
			StartEndPos: newStartEndPos(loc, Location(lexer.Location)),
			Error:       fmt.Errorf("block comment not terminated"),
		}, nil
	}

	return TokenValue{
		SymbolId:    blockCommentToken,
		StartEndPos: newStartEndPos(loc, Location(lexer.Location)),
		Value:       value,
	}, nil
}

func (lexer *RawLexer) peekDigits(
	offset int,
	isDigit func(byte) bool,
	requireLeadingDigit bool,
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
				if requireLeadingDigit && numBytes == 0 {
					return numBytes, nil
				}

				if len(remaining) < 2 {
					if hasMore {
						// read more bytes
						break
					} else { // the int is followed by a '_'
						return numBytes, nil
					}
				}

				if isDigit(remaining[1]) {
					numBytes += 2
					remaining = remaining[2:]
				} else { // the int is followed by a '_'
					return numBytes, nil
				}
			} else {
				return numBytes, nil
			}
		}

		peekSize *= 2
	}

	return numBytes, nil
}

func isBinaryDigit(char byte) bool {
	return char == '0' || char == '1'
}

func isOctalDigit(char byte) bool {
	return '0' <= char && char <= '7'
}

func isDecimalDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func isHexadecimalDigit(char byte) bool {
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
	if !isDecimalDigit(char) {
		panic("should never happen")
	}

	loc := Location(lexer.Location)
	subType := DecimalInteger
	hasDigits := true
	totalBytes := 1

	if char != '0' {
		numBytes, err := lexer.peekDigits(1, isDecimalDigit, false)
		if err != nil {
			return nil, err
		}

		totalBytes = numBytes + 1
	} else if len(peeked) > 1 {
		switch peeked[1] {
		case 'b', 'B':
			numBytes, err := lexer.peekDigits(2, isBinaryDigit, false)
			if err != nil {
				return nil, err
			}

			subType = BinaryInteger
			hasDigits = numBytes != 0
			totalBytes = numBytes + 2
		case 'o', 'O':
			numBytes, err := lexer.peekDigits(2, isOctalDigit, false)
			if err != nil {
				return nil, err
			}

			subType = ZeroOPrefixedOctalInteger
			hasDigits = numBytes != 0
			totalBytes = numBytes + 2
		case 'x', 'X':
			numBytes, err := lexer.peekDigits(2, isHexadecimalDigit, false)
			if err != nil {
				return nil, err
			}

			subType = HexadecimalInteger
			hasDigits = numBytes != 0
			totalBytes = numBytes + 2
		default:
			numBytes, err := lexer.peekDigits(1, isOctalDigit, false)
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

	if subType == DecimalInteger || subType == HexadecimalInteger {
		token, err := lexer.maybeLexIntPrefixedFloat(
			totalBytes,
			subType == HexadecimalInteger)
		if err != nil {
			return nil, err
		}

		if token != nil {
			return token, nil
		}
	}

	value := ""
	if hasDigits {
		peeked, err = lexer.Peek(totalBytes)
		if err != nil {
			panic("should never happen")
		}

		if len(peeked) < 3 { // intern short ints
			value = lexer.InternBytes(peeked)
		} else {
			value = string(peeked)
		}
	}

	_, err = lexer.Discard(totalBytes)
	if err != nil {
		panic("should never happen")
	}

	if !hasDigits {
		return ParseErrorSymbol{
			StartEndPos: newStartEndPos(loc, Location(lexer.Location)),
			Error:       fmt.Errorf("%s has no digits", subType),
		}, nil
	}

	return TokenValue{
		SymbolId:    IntegerLiteralToken,
		StartEndPos: newStartEndPos(loc, Location(lexer.Location)),
		Value:       value,
		SubType:     subType,
	}, nil
}

func (lexer *RawLexer) maybePeekFloat(
	leadingIntBytes int,
	isHex bool, // false = decimal
) (
	int,
	error,
) {
	isDigit := isDecimalDigit
	intPrefixLen := 0 // no prefix
	lowerExp := byte('e')
	upperExp := byte('E')
	if isHex {
		isDigit = isHexadecimalDigit
		intPrefixLen = 2 // "0x" or "0X"
		lowerExp = 'p'
		upperExp = 'P'
	}

	// peek for (hexa)decimal points of the form DOT digits

	floatBytes := leadingIntBytes
	peeked, err := lexer.Peek(floatBytes + 1)
	if len(peeked) > 0 && err == io.EOF {
		err = nil
	}
	if err != nil {
		return 0, err
	}

	if len(peeked) <= floatBytes {
		// can't be a float without decimal point or exponent
		return 0, nil
	}

	if peeked[floatBytes] == '.' {
		floatBytes++

		digitBytes, err := lexer.peekDigits(floatBytes, isDigit, true)
		if err != nil {
			return 0, err
		}

		if digitBytes == 0 && leadingIntBytes == intPrefixLen {
			// ".", without leading and trailing digits, is not a valid float
			return 0, err
		}

		floatBytes += digitBytes
	}

	// peek for exponent of the form [eEpP][+-]?digits

	peeked, err = lexer.Peek(floatBytes + 2)
	if len(peeked) > 0 && err == io.EOF {
		err = nil
	}
	if err != nil {
		return 0, err
	}

	if len(peeked) <= floatBytes {
		if isHex { // exponent is not optional for hexadecimal
			return 0, nil
		} else if floatBytes > leadingIntBytes {
			// a valid decimal float without exponent
			return floatBytes, nil
		}
	}

	exp := peeked[floatBytes]
	if exp == lowerExp || exp == upperExp {
		prefix := 1
		if len(peeked) == floatBytes+2 {
			char := peeked[floatBytes+1]
			if char == '+' || char == '-' {
				prefix = 2
			}
		}

		exponentDigits, err := lexer.peekDigits(floatBytes+prefix, isDigit, true)
		if err != nil {
			return 0, err
		}

		if exponentDigits > 0 {
			floatBytes += prefix + exponentDigits
		} else if isHex { // exponent is not optional for hexadecimal
			return 0, err
		}
	} else if isHex { // exponent is not optional for hexadecimal
		return 0, err
	}

	if floatBytes > leadingIntBytes {
		return floatBytes, nil
	}

	return 0, nil
}

func (lexer *RawLexer) maybeLexIntPrefixedFloat(
	leadingIntBytes int,
	isHex bool, // false is decimal
) (
	Token,
	error,
) {
	numBytes, err := lexer.maybePeekFloat(leadingIntBytes, isHex)
	if err != nil {
		return nil, err
	}

	if numBytes == 0 {
		return nil, nil
	}

	peeked, err := lexer.Peek(numBytes)
	if err != nil {
		panic("should never happen")
	}

	loc := Location(lexer.Location)
	value := string(peeked)

	_, err = lexer.Discard(numBytes)
	if err != nil {
		panic("should never happen")
	}

	subType := DecimalFloat
	if isHex {
		subType = HexadecimalFloat
	}

	return TokenValue{
		SymbolId:    FloatLiteralToken,
		StartEndPos: newStartEndPos(loc, Location(lexer.Location)),
		Value:       value,
		SubType:     subType,
	}, nil
}

func (lexer *RawLexer) lexDotDecimalFloatLiteralToken() (Token, error) {
	numBytes, err := lexer.maybePeekFloat(0, false)
	if err != nil {
		return nil, err
	}

	if numBytes == 0 {
		panic("should never happen")
	}

	peeked, err := lexer.Peek(numBytes)
	if err != nil {
		panic("should never happen")
	}

	loc := Location(lexer.Location)
	value := string(peeked)

	_, err = lexer.Discard(numBytes)
	if err != nil {
		panic("should never happen")
	}

	return TokenValue{
		SymbolId:    FloatLiteralToken,
		StartEndPos: newStartEndPos(loc, Location(lexer.Location)),
		Value:       value,
		SubType:     DecimalFloat,
	}, nil
}

type peekStringResult struct {
	numBytes      int
	contentLength int
	errorMsg      string
}

func (lexer *RawLexer) peekString(
	typeName string,
	marker byte,
	startMarkerLength int,
	endMarkerLength int,
	allowMultiline bool,
	allowEscaped bool,
) (
	peekStringResult,
	error,
) {
	peekSize := lexer.initialPeekWindowSize
	hasMore := true

	result := peekStringResult{
		// Skip over the start marker since we've already peek the marker
		numBytes: startMarkerLength,
	}

	for hasMore {
		peeked, err := lexer.Peek(startMarkerLength + peekSize)
		if len(peeked) > 0 && err == io.EOF {
			hasMore = false
			err = nil
		}
		if err != nil {
			return result, err
		}

		remaining := peeked[result.numBytes:]

		for len(remaining) > 0 {
			// Ensure we can process the longest rune content: \U[0-9a-fA-F]{8}
			if len(remaining) < 10 && hasMore {
				// read more bytes
				break
			}

			char := remaining[0]
			if char == marker {
				result.numBytes++
				remaining = remaining[1:]

				count := 1
				for ; count < endMarkerLength && len(remaining) > 0; count++ {
					if remaining[0] == marker {
						result.numBytes++
						remaining = remaining[1:]
					} else {
						break
					}
				}

				if count == endMarkerLength {
					return result, nil
				} else {
					result.contentLength += count
				}
			} else if char == '\n' {
				if allowMultiline {
					result.numBytes++
					remaining = remaining[1:]
				} else { // don't include the newline
					result.errorMsg = typeName + " not terminated"
					return result, nil
				}
			} else if allowEscaped && char == '\\' { // escape sequence
				result.numBytes++
				remaining = remaining[1:]

				if len(remaining) == 0 {
					result.errorMsg = "invalid escaped character in " + typeName
					return result, nil
				}

				char := remaining[0]
				switch char {
				case 'a', 'b', 'f', 'n', 'r', 't', 'v', '\\', '\'', '"', '`':
					// valid escape
					// Note: if we replace '\'', '"', and '`' with marker, then this
					// would behave like golang's escape.
					result.numBytes++
					result.contentLength++
					remaining = remaining[1:]
				case '\n':
					if allowMultiline { // valid escape (line continuation)
						result.numBytes++
						result.contentLength++
						remaining = remaining[1:]
					} else { // don't include the newline as part of this token.
						result.errorMsg = "invalid escaped character in " + typeName
						return result, nil
					}
				default:
					result.numBytes++
					remaining = remaining[1:]

					value := int(char - '0')
					verifyOctal := false
					isDigit := isOctalDigit
					length := 2
					if isOctalDigit(char) { // \[0-7]{3}
						// need to check the remaining 2 octal bytes, and
						verifyOctal = true
					} else if char == 'x' { // \x[0-9a-fA-F]{2}
						isDigit = isHexadecimalDigit
					} else if char == 'u' { // \u[0-9a-fA-F]{4}
						isDigit = isHexadecimalDigit
						length = 4
					} else if char == 'U' { // \U[0-9a-fA-F]{8}
						isDigit = isHexadecimalDigit
						length = 8
					} else { // invalid escape
						result.errorMsg = "invalid escaped character in " + typeName
						return result, nil
					}

					count := 0
					for ; count < length && len(remaining) > 0; count++ {
						if isDigit(remaining[0]) {
							if verifyOctal {
								value <<= 3
								value |= int(remaining[0] - '0')
							}

							result.numBytes++
							remaining = remaining[1:]
						} else {
							break
						}
					}

					if count == length {
						result.contentLength++
					} else {
						result.errorMsg = "invalid escaped unicode value in " + typeName
						return result, nil
					}

					if verifyOctal && value > 255 {
						result.errorMsg = fmt.Sprintf(
							"invalid escaped octal value (%d > 255)  in %s",
							value,
							typeName)
						return result, nil
					}
				}
			} else {
				utf8Char, size := utf8.DecodeRune(remaining)

				result.numBytes += size
				remaining = remaining[size:]

				if utf8Char == utf8.RuneError {
					result.errorMsg = "invalid unicode rune in " + typeName
					return result, nil
				} else {
					result.contentLength++
				}
			}
		}

		peekSize *= 2
	}

	result.errorMsg = typeName + " not terminated"
	return result, nil
}

func (lexer *RawLexer) lexRuneLiteralToken() (Token, error) {
	result, err := lexer.peekString(
		"rune literal",
		'\'',
		1,     // start marker length
		1,     // end marker length
		false, // allow multiline
		true)  // allow escaped
	if err != nil {
		return nil, err
	}

	loc := Location(lexer.Location)

	value := ""
	if result.errorMsg == "" {
		if result.contentLength > 1 {
			result.errorMsg = "more than one character in rune literal"
		} else if result.contentLength < 1 {
			result.errorMsg = "empty rune literal or unescaped '"
		} else {
			peeked, err := lexer.Peek(result.numBytes)
			if err != nil {
				panic("should never happen")
			}

			if len(peeked) == 3 ||
				(len(peeked) == 4 && peeked[1] == '/') { // intern ascii
				value = lexer.InternBytes(peeked)
			} else {
				value = string(peeked)
			}
		}
	}

	_, err = lexer.Discard(result.numBytes)
	if err != nil {
		panic("should never happen")
	}

	if result.errorMsg != "" {
		return ParseErrorSymbol{
			StartEndPos: newStartEndPos(loc, Location(lexer.Location)),
			Error:       fmt.Errorf(result.errorMsg),
		}, nil
	}

	return TokenValue{
		SymbolId:    RuneLiteralToken,
		StartEndPos: newStartEndPos(loc, Location(lexer.Location)),
		Value:       value,
	}, nil
}

func (lexer *RawLexer) lexStringLiteralToken(
	subType string,
	marker byte,
	startMarkerLength int,
	endMarkerLength int,
	allowMultiline bool,
	allowEscaped bool,
) (
	Token,
	error,
) {
	result, err := lexer.peekString(
		subType,
		marker,
		startMarkerLength,
		endMarkerLength,
		allowMultiline,
		allowEscaped)
	if err != nil {
		return nil, err
	}

	loc := Location(lexer.Location)

	value := ""
	if result.errorMsg == "" {
		peeked, err := lexer.Peek(result.numBytes)
		if err != nil {
			panic("should never happen")
		}

		if len(peeked) == startMarkerLength+endMarkerLength { // intern empty string
			value = lexer.InternBytes(peeked)
		} else {
			value = string(peeked)
		}
	}

	_, err = lexer.Discard(result.numBytes)
	if err != nil {
		panic("should never happen")
	}

	if result.errorMsg != "" {
		return ParseErrorSymbol{
			StartEndPos: newStartEndPos(loc, Location(lexer.Location)),
			Error:       fmt.Errorf(result.errorMsg),
		}, nil
	}

	return TokenValue{
		SymbolId:    StringLiteralToken,
		StartEndPos: newStartEndPos(loc, Location(lexer.Location)),
		Value:       value,
		SubType:     subType,
	}, nil
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
			StartEndPos: newStartEndPos(loc, Location(lexer.Location)),
			Error:       fmt.Errorf("no label name associated with @"),
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

	return TokenValue{
		SymbolId:    JumpLabelToken,
		StartEndPos: newStartEndPos(loc, Location(lexer.Location)),
		Value:       jumpLabel,
	}, nil
}

func (lexer *RawLexer) lexIdentifierKeywordsOrLabelDeclToken() (
	TokenValue,
	error,
) {
	size, err := lexutil.PeekIdentifier(
		lexer.BufferedByteLocationReader,
		0,
		lexer.initialPeekWindowSize)
	if err != nil {
		return TokenValue{}, err
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

	return TokenValue{
		SymbolId:    symbolId,
		StartEndPos: newStartEndPos(loc, Location(lexer.Location)),
		Value:       value,
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
			StartEndPos: newStartEndPos(loc, Location(lexer.Location)),
			Error:       fmt.Errorf("unexpected utf8 rune"),
		}, nil
	}

	if size > 0 {
		loc := Location(lexer.Location)

		_, err := lexer.Discard(size)
		if err != nil {
			panic("Should never happen")
		}

		return TokenValue{
			SymbolId:    symbolId,
			StartEndPos: newStartEndPos(loc, Location(lexer.Location)),
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
	case sibStringToken:
		return lexer.lexStringLiteralToken(
			SingleLineString,
			'`',
			1,     // start marker length
			1,     // end marker length
			false, // allow multiline
			true)  // allow escaped
	case sidStringToken:
		return lexer.lexStringLiteralToken(
			SingleLineString,
			'"',
			1,     // start marker length
			1,     // end marker length
			false, // allow multiline
			true)  // allow escaped
	case srbStringToken:
		return lexer.lexStringLiteralToken(
			RawSingleLineString,
			'`',
			2,     // start marker length
			1,     // end marker length
			false, // allow multiline
			false) // allow escaped
	case srdStringToken:
		return lexer.lexStringLiteralToken(
			RawSingleLineString,
			'"',
			2,     // start marker length
			1,     // end marker length
			false, // allow multiline
			false) // allow escaped
	case mibStringToken:
		return lexer.lexStringLiteralToken(
			MultiLineString,
			'`',
			3,    // start marker length
			3,    // end marker length
			true, // allow multiline
			true) // allow escaped
	case midStringToken:
		return lexer.lexStringLiteralToken(
			MultiLineString,
			'"',
			3,    // start marker length
			3,    // end marker length
			true, // allow multiline
			true) // allow escaped
	case mrbStringToken:
		return lexer.lexStringLiteralToken(
			RawMultiLineString,
			'`',
			4,     // start marker length
			3,     // end marker length
			true,  // allow multiline
			false) // allow escaped
	case mrdStringToken:
		return lexer.lexStringLiteralToken(
			RawMultiLineString,
			'"',
			4,     // start marker length
			3,     // end marker length
			true,  // allow multiline
			false) // allow escaped
	case JumpLabelToken:
		return lexer.lexJumpLabelToken()
	case IdentifierToken:
		return lexer.lexIdentifierKeywordsOrLabelDeclToken()
	}

	panic(fmt.Sprintf("RawLexer: unhandled variable length token: %v", symbolId))
}
