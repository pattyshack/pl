package lexer

import (
	"fmt"
	"io"
	"unicode/utf8"

	"github.com/pattyshack/gt/lexutil"
	"github.com/pattyshack/gt/stringutil"

	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
	"github.com/pattyshack/pl/util"
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
		"repeat":      lr.RepeatToken,
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
		"struct":      lr.StructToken,
		"enum":        lr.EnumToken,
		"trait":       lr.TraitToken,
		"func":        lr.FuncToken,
		"async":       lr.AsyncToken,
		"defer":       lr.DeferToken,
		"var":         lr.VarToken,
		"let":         lr.LetToken,
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
	*util.ErrorEmitter
	LexerOptions

	lexutil.BufferedByteLocationReader
	*stringutil.InternPool
}

func NewRawLexer(
	sourceFileName string,
	sourceContent io.Reader,
	emitter *util.ErrorEmitter,
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
		ErrorEmitter: emitter,
		LexerOptions: options,
		BufferedByteLocationReader: lexutil.NewBufferedByteLocationReader(
			sourceFileName,
			sourceContent,
			options.InitialLookAheadBufferSize),
		InternPool: internPool,
	}
}

func (lexer *RawLexer) CurrentLocation() lexutil.Location {
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
		if len(peeked) > 1 && peeked[1] == '[' {
			return lr.DollarLbracketToken, "$[", nil
		}
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
	numSpaceBytes, err := lexutil.PeekSpaces(
		lexer.BufferedByteLocationReader,
		lexer.initialPeekWindowSize)
	if err != nil {
		return nil, err
	}

	if numSpaceBytes == 0 {
		panic("This should never happen")
	}

	loc := lexer.Location

	_, err = lexer.Discard(numSpaceBytes)
	if err != nil {
		panic("This should never happen")
	}

	return lr.TokenCount{
		SymbolId:    spacesToken,
		StartEndPos: ast.NewStartEndPos(loc, lexer.Location),
		Count:       numSpaceBytes,
	}, nil
}

func (lexer *RawLexer) lexNewlinesToken() (lr.Token, error) {
	numNewlineBytes, numNewlines, foundInvalidNewline, err := lexutil.PeekNewlines(
		lexer.BufferedByteLocationReader,
		lexer.initialPeekWindowSize)
	if err != nil {
		return nil, err
	}

	loc := lexer.Location
	if numNewlineBytes > 0 {
		_, err := lexer.Discard(numNewlineBytes)
		if err != nil {
			panic("This should never happen")
		}

		return lr.TokenCount{
			SymbolId:    lr.NewlinesToken,
			StartEndPos: ast.NewStartEndPos(loc, lexer.Location),
			Count:       numNewlines,
		}, nil
	} else if foundInvalidNewline {
		_, err := lexer.Discard(1)
		if err != nil {
			panic("This should never happen")
		}

		return lr.NewParseErrorSymbol(
			loc, lexer.Location,
			"%s: unexpected utf8 rune", loc), nil
	}

	panic("This should never happen")
}

func (lexer *RawLexer) lexLineCommentToken() (lr.Token, error) {
	numCommentBytes, err := lexutil.PeekLineComment(
		lexer.BufferedByteLocationReader,
		lexer.initialPeekWindowSize)
	if err != nil {
		return nil, err
	}

	loc := lexer.Location

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

	return &lr.TokenValue{
		SymbolId:    lineCommentToken,
		StartEndPos: ast.NewStartEndPos(loc, lexer.Location),
		Value:       value,
	}, nil
}

func (lexer *RawLexer) lexBlockCommentToken() (lr.Token, error) {
	numCommentBytes, scopeLevel, err := lexutil.PeekBlockComment(
		lexer.BufferedByteLocationReader,
		true,
		lexer.initialPeekWindowSize)
	if err != nil {
		return nil, err
	}

	loc := lexer.Location

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
		return lr.NewParseErrorSymbol(
			loc, lexer.Location,
			"%s: block comment not terminated", loc), nil
	}

	return &lr.TokenValue{
		SymbolId:    blockCommentToken,
		StartEndPos: ast.NewStartEndPos(loc, lexer.Location),
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

func (lexer *RawLexer) lexIntegerOrFloatLiteralToken() (lr.Token, error) {
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

	loc := lexer.Location
	subType := lr.DecimalInteger
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

			subType = lr.BinaryInteger
			hasDigits = numBytes != 0
			totalBytes = numBytes + 2
		case 'o', 'O':
			numBytes, err := lexer.peekDigits(2, isOctalDigit, false)
			if err != nil {
				return nil, err
			}

			subType = lr.ZeroOPrefixedOctalInteger
			hasDigits = numBytes != 0
			totalBytes = numBytes + 2
		case 'x', 'X':
			numBytes, err := lexer.peekDigits(2, isHexadecimalDigit, false)
			if err != nil {
				return nil, err
			}

			subType = lr.HexadecimalInteger
			hasDigits = numBytes != 0
			totalBytes = numBytes + 2
		default:
			numBytes, err := lexer.peekDigits(1, isOctalDigit, false)
			if err != nil {
				return nil, err
			}

			if numBytes > 0 { // otherwise is a decimal "0"
				subType = lr.ZeroPrefixedOctalInteger
				hasDigits = true
				totalBytes = numBytes + 1
			}
		}
	}

	if subType == lr.DecimalInteger || subType == lr.HexadecimalInteger {
		token, err := lexer.maybeLexIntPrefixedFloat(
			totalBytes,
			subType == lr.HexadecimalInteger)
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
		return lr.NewParseErrorSymbol(
			loc, lexer.Location,
			"%s: %s has no digits", loc, subType), nil
	}

	return &lr.TokenValue{
		SymbolId:    lr.IntegerLiteralToken,
		StartEndPos: ast.NewStartEndPos(loc, lexer.Location),
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
	lr.Token,
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

	loc := lexer.Location
	value := string(peeked)

	_, err = lexer.Discard(numBytes)
	if err != nil {
		panic("should never happen")
	}

	subType := lr.DecimalFloat
	if isHex {
		subType = lr.HexadecimalFloat
	}

	return &lr.TokenValue{
		SymbolId:    lr.FloatLiteralToken,
		StartEndPos: ast.NewStartEndPos(loc, lexer.Location),
		Value:       value,
		SubType:     subType,
	}, nil
}

func (lexer *RawLexer) lexDotDecimalFloatLiteralToken() (lr.Token, error) {
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

	loc := lexer.Location
	value := string(peeked)

	_, err = lexer.Discard(numBytes)
	if err != nil {
		panic("should never happen")
	}

	return &lr.TokenValue{
		SymbolId:    lr.FloatLiteralToken,
		StartEndPos: ast.NewStartEndPos(loc, lexer.Location),
		Value:       value,
		SubType:     lr.DecimalFloat,
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

func (lexer *RawLexer) lexRuneLiteralToken() (lr.Token, error) {
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

	loc := lexer.Location

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
		return lr.NewParseErrorSymbol(
			loc, lexer.Location,
			"%s: %s", loc, result.errorMsg), nil
	}

	return &lr.TokenValue{
		SymbolId:    lr.RuneLiteralToken,
		StartEndPos: ast.NewStartEndPos(loc, lexer.Location),
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
	lr.Token,
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

	loc := lexer.Location

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
		return lr.NewParseErrorSymbol(
			loc, lexer.Location,
			"%s: %s", loc, result.errorMsg), nil
	}

	return &lr.TokenValue{
		SymbolId:    lr.StringLiteralToken,
		StartEndPos: ast.NewStartEndPos(loc, lexer.Location),
		Value:       value,
		SubType:     subType,
	}, nil
}

func (lexer *RawLexer) lexIdentifierOrKeywords() (
	*lr.TokenValue,
	error,
) {
	size, err := lexutil.PeekIdentifier(
		lexer.BufferedByteLocationReader,
		0,
		lexer.initialPeekWindowSize)
	if err != nil {
		return nil, err
	}

	if size == 0 {
		panic("Should never hapapen")
	}

	symbolId := lr.IdentifierToken

	peeked, _ := lexer.Peek(size)

	loc := lexer.Location
	value := lexer.InternBytes(peeked)

	_, err = lexer.Discard(size)
	if err != nil {
		panic("Should never happen")
	}

	kwSymbolId, ok := keywords[value]
	if ok {
		symbolId = kwSymbolId
	}

	return &lr.TokenValue{
		SymbolId:    symbolId,
		StartEndPos: ast.NewStartEndPos(loc, lexer.Location),
		Value:       value,
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
			loc, lexer.Location,
			"%s: unexpected utf8 rune", loc), nil
	}

	if size > 0 {
		loc := lexer.Location

		_, err := lexer.Discard(size)
		if err != nil {
			panic("Should never happen")
		}

		return &lr.TokenValue{
			SymbolId:    symbolId,
			StartEndPos: ast.NewStartEndPos(loc, lexer.Location),
			Value:       value,
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
		return lexer.lexDotDecimalFloatLiteralToken()
	case lr.RuneLiteralToken:
		return lexer.lexRuneLiteralToken()
	case sibStringToken:
		return lexer.lexStringLiteralToken(
			lr.SingleLineString,
			'`',
			1,     // start marker length
			1,     // end marker length
			false, // allow multiline
			true)  // allow escaped
	case sidStringToken:
		return lexer.lexStringLiteralToken(
			lr.SingleLineString,
			'"',
			1,     // start marker length
			1,     // end marker length
			false, // allow multiline
			true)  // allow escaped
	case srbStringToken:
		return lexer.lexStringLiteralToken(
			lr.RawSingleLineString,
			'`',
			2,     // start marker length
			1,     // end marker length
			false, // allow multiline
			false) // allow escaped
	case srdStringToken:
		return lexer.lexStringLiteralToken(
			lr.RawSingleLineString,
			'"',
			2,     // start marker length
			1,     // end marker length
			false, // allow multiline
			false) // allow escaped
	case mibStringToken:
		return lexer.lexStringLiteralToken(
			lr.MultiLineString,
			'`',
			3,    // start marker length
			3,    // end marker length
			true, // allow multiline
			true) // allow escaped
	case midStringToken:
		return lexer.lexStringLiteralToken(
			lr.MultiLineString,
			'"',
			3,    // start marker length
			3,    // end marker length
			true, // allow multiline
			true) // allow escaped
	case mrbStringToken:
		return lexer.lexStringLiteralToken(
			lr.RawMultiLineString,
			'`',
			4,     // start marker length
			3,     // end marker length
			true,  // allow multiline
			false) // allow escaped
	case mrdStringToken:
		return lexer.lexStringLiteralToken(
			lr.RawMultiLineString,
			'"',
			4,     // start marker length
			3,     // end marker length
			true,  // allow multiline
			false) // allow escaped
	case lr.IdentifierToken:
		return lexer.lexIdentifierOrKeywords()
	}

	panic(fmt.Sprintf("RawLexer: unhandled variable length token: %v", symbolId))
}
