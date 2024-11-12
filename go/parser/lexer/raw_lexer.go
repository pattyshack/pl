package lexer

import (
	"fmt"
	"io"
	"unicode/utf8"

	"github.com/pattyshack/gt/lexutil"
	"github.com/pattyshack/gt/stringutil"

	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/errors"
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
	*errors.Emitter
	LexerOptions

	lexutil.BufferedByteLocationReader
	*stringutil.InternPool
}

func NewRawLexer(
	sourceFileName string,
	sourceContent io.Reader,
	emitter *errors.Emitter,
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
			"unexpected utf8 rune"), nil
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
			"block comment not terminated"), nil
	}

	return &lr.TokenValue{
		SymbolId:    blockCommentToken,
		StartEndPos: ast.NewStartEndPos(loc, lexer.Location),
		Value:       value,
	}, nil
}

func peekDigits(
	reader lexutil.BufferedByteLocationReader,
	initialPeekWindowSize int,
	offset int,
	isDigit func(rune) bool,
	requireLeadingDigit bool,
) (
	int,
	error,
) {
	hasMore := true
	peekSize := initialPeekWindowSize
	numBytes := 0

	for hasMore {
		peeked, err := reader.Peek(offset + peekSize)
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
			if isDigit(rune(char)) {
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

				if isDigit(rune(remaining[1])) {
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

type PeekIntegerOrFloatResult struct {
	NumBytes   int
	IsNegative bool
	IsFloat    bool
	SubType    lr.LiteralSubType
	HasDigits  bool
}

func PeekIntegerOrFloat(
	reader lexutil.BufferedByteLocationReader,
	initialPeekWindowSize int,
) (
	PeekIntegerOrFloatResult,
	error,
) {
	peeked, err := reader.Peek(3)
	if len(peeked) > 0 && err == io.EOF {
		err = nil
	}
	if err != nil {
		return PeekIntegerOrFloatResult{}, err
	}

	isNegative := false
	totalBytes := 0

	if peeked[0] == '-' {
		isNegative = true
		totalBytes = 1
		peeked = peeked[1:]
	}

	if len(peeked) == 0 {
		return PeekIntegerOrFloatResult{}, err
	}

	char := peeked[0]
	if char == '.' {
		totalFloatBytes, err := peekFloat(
			reader,
			initialPeekWindowSize,
			totalBytes,
			false)
		if err != nil {
			return PeekIntegerOrFloatResult{}, err
		}

		if totalFloatBytes == 0 {
			return PeekIntegerOrFloatResult{}, nil
		} else {
			return PeekIntegerOrFloatResult{
				NumBytes:   totalFloatBytes,
				IsNegative: isNegative,
				IsFloat:    true,
				SubType:    lr.DecimalFloat,
				HasDigits:  true,
			}, nil
		}
	}

	if !lexutil.IsDecimalDigit(rune(char)) {
		return PeekIntegerOrFloatResult{}, nil
	}

	subType := lr.DecimalInteger
	hasDigits := true
	totalBytes++

	if char != '0' {
		numDigits, err := peekDigits(
			reader,
			initialPeekWindowSize,
			totalBytes,
			lexutil.IsDecimalDigit,
			false)
		if err != nil {
			return PeekIntegerOrFloatResult{}, err
		}

		totalBytes += numDigits
	} else if len(peeked) > 1 {
		switch peeked[1] {
		case 'b', 'B':
			totalBytes++
			numDigits, err := peekDigits(
				reader,
				initialPeekWindowSize,
				totalBytes,
				lexutil.IsBinaryDigit,
				false)
			if err != nil {
				return PeekIntegerOrFloatResult{}, err
			}

			subType = lr.BinaryInteger
			hasDigits = numDigits != 0
			totalBytes += numDigits
		case 'o', 'O':
			totalBytes++
			numDigits, err := peekDigits(
				reader,
				initialPeekWindowSize,
				totalBytes,
				lexutil.IsOctalDigit,
				false)
			if err != nil {
				return PeekIntegerOrFloatResult{}, err
			}

			subType = lr.ZeroOPrefixedOctalInteger
			hasDigits = numDigits != 0
			totalBytes += numDigits
		case 'x', 'X':
			totalBytes++
			numDigits, err := peekDigits(
				reader,
				initialPeekWindowSize,
				totalBytes,
				lexutil.IsHexadecimalDigit,
				false)
			if err != nil {
				return PeekIntegerOrFloatResult{}, err
			}

			subType = lr.HexadecimalInteger
			hasDigits = numDigits != 0
			totalBytes += numDigits
		default:
			numDigits, err := peekDigits(
				reader,
				initialPeekWindowSize,
				totalBytes,
				lexutil.IsOctalDigit,
				false)
			if err != nil {
				return PeekIntegerOrFloatResult{}, err
			}

			if numDigits > 0 { // otherwise is a decimal "0"
				subType = lr.ZeroPrefixedOctalInteger
				totalBytes += numDigits
			}
		}
	}

	result := PeekIntegerOrFloatResult{
		NumBytes:   totalBytes,
		IsNegative: isNegative,
		SubType:    subType,
		HasDigits:  hasDigits,
	}

	if subType == lr.DecimalInteger || subType == lr.HexadecimalInteger {
		isHex := subType == lr.HexadecimalInteger
		totalFloatBytes, err := peekFloat(
			reader,
			initialPeekWindowSize,
			totalBytes,
			isHex)
		if err != nil {
			return result, nil // we still have a valid integer
		}

		if totalFloatBytes > 0 {
			result.NumBytes = totalFloatBytes
			result.IsFloat = true
			result.HasDigits = true
			result.SubType = lr.DecimalFloat
			if isHex {
				result.SubType = lr.HexadecimalFloat
			}
		}
	}

	return result, nil
}

func (lexer *RawLexer) lexIntegerOrFloatLiteralToken() (lr.Token, error) {
	result, err := PeekIntegerOrFloat(
		lexer.BufferedByteLocationReader,
		lexer.initialPeekWindowSize)
	if err != nil {
		return nil, err
	}

	if result.NumBytes == 0 {
		panic("should never happen")
	}

	loc := lexer.Location

	value := ""
	if result.HasDigits {
		peeked, err := lexer.Peek(result.NumBytes)
		if err != nil {
			panic("should never happen")
		}

		if len(peeked) < 3 { // intern short ints
			value = lexer.InternBytes(peeked)
		} else {
			value = string(peeked)
		}
	}

	_, err = lexer.Discard(result.NumBytes)
	if err != nil {
		panic("should never happen")
	}

	if !result.HasDigits {
		return lr.NewParseErrorSymbol(
			loc, lexer.Location,
			"%s has no digits", result.SubType), nil
	}

	symbolId := lr.IntegerLiteralToken
	if result.IsFloat {
		symbolId = lr.FloatLiteralToken
	}

	return &lr.TokenValue{
		SymbolId:    symbolId,
		StartEndPos: ast.NewStartEndPos(loc, lexer.Location),
		Value:       value,
		SubType:     result.SubType,
	}, nil
}

func peekFloat(
	reader lexutil.BufferedByteLocationReader,
	initialPeekWindowSize int,
	leadingIntBytes int,
	isHex bool, // false = decimal
) (
	int,
	error,
) {
	isDigit := lexutil.IsDecimalDigit
	intPrefixLen := 0 // no prefix
	lowerExp := byte('e')
	upperExp := byte('E')
	if isHex {
		isDigit = lexutil.IsHexadecimalDigit
		intPrefixLen = 2 // "0x" or "0X"
		lowerExp = 'p'
		upperExp = 'P'
	}

	// peek for (hexa)decimal points of the form DOT digits

	floatBytes := leadingIntBytes
	peeked, err := reader.Peek(floatBytes + 1)
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

		digitBytes, err := peekDigits(
			reader,
			initialPeekWindowSize,
			floatBytes,
			isDigit,
			true)
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

	peeked, err = reader.Peek(floatBytes + 2)
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

		exponentDigits, err := peekDigits(
			reader,
			initialPeekWindowSize,
			floatBytes+prefix,
			isDigit,
			true)
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

func (lexer *RawLexer) lexRuneLiteralToken() (lr.Token, error) {
	result, err := lexutil.PeekString(
		lexer.BufferedByteLocationReader,
		lexer.initialPeekWindowSize,
		"rune",
		0, // skip leading bytes
		'\'',
		1,     // marker length
		false, // allow multiline
		true)  // allow escaped
	if err != nil {
		return nil, err
	}

	loc := lexer.Location

	value := ""
	if result.ErrorMsg == "" {
		if result.ContentLength > 1 {
			result.ErrorMsg = "more than one character in rune literal"
		} else if result.ContentLength < 1 {
			result.ErrorMsg = "empty rune literal or unescaped '"
		} else {
			peeked, err := lexer.Peek(result.NumBytes)
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

	_, err = lexer.Discard(result.NumBytes)
	if err != nil {
		panic("should never happen")
	}

	if result.ErrorMsg != "" {
		return lr.NewParseErrorSymbol(loc, lexer.Location, result.ErrorMsg), nil
	}

	return &lr.TokenValue{
		SymbolId:    lr.RuneLiteralToken,
		StartEndPos: ast.NewStartEndPos(loc, lexer.Location),
		Value:       value,
	}, nil
}

func (lexer *RawLexer) lexStringLiteralToken(
	subType lr.LiteralSubType,
	skipLeadingBytes int,
	marker byte,
	markerLength int,
	allowMultiline bool,
	allowEscaped bool,
) (
	lr.Token,
	error,
) {
	result, err := lexutil.PeekString(
		lexer.BufferedByteLocationReader,
		lexer.initialPeekWindowSize,
		string(subType),
		skipLeadingBytes,
		marker,
		markerLength,
		allowMultiline,
		allowEscaped)
	if err != nil {
		return nil, err
	}

	loc := lexer.Location

	value := ""
	if result.ErrorMsg == "" {
		peeked, err := lexer.Peek(result.NumBytes)
		if err != nil {
			panic("should never happen")
		}

		if len(peeked) == skipLeadingBytes+2*markerLength { // intern empty string
			value = lexer.InternBytes(peeked)
		} else {
			value = string(peeked)
		}
	}

	_, err = lexer.Discard(result.NumBytes)
	if err != nil {
		panic("should never happen")
	}

	if result.ErrorMsg != "" {
		return lr.NewParseErrorSymbol(loc, lexer.Location, result.ErrorMsg), nil
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
			"unexpected utf8 rune"), nil
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
		return lexer.lexIntegerOrFloatLiteralToken()
	case lr.RuneLiteralToken:
		return lexer.lexRuneLiteralToken()
	case sibStringToken:
		return lexer.lexStringLiteralToken(
			lr.SingleLineString,
			0, // skip leading bytes
			'`',
			1,     // marker length
			false, // allow multiline
			true)  // allow escaped
	case sidStringToken:
		return lexer.lexStringLiteralToken(
			lr.SingleLineString,
			0, // skip leading bytes
			'"',
			1,     // marker length
			false, // allow multiline
			true)  // allow escaped
	case srbStringToken:
		return lexer.lexStringLiteralToken(
			lr.RawSingleLineString,
			1, // skip leading bytes
			'`',
			1,     // marker length
			false, // allow multiline
			false) // allow escaped
	case srdStringToken:
		return lexer.lexStringLiteralToken(
			lr.RawSingleLineString,
			1, // skip leading bytes
			'"',
			1,     // marker length
			false, // allow multiline
			false) // allow escaped
	case mibStringToken:
		return lexer.lexStringLiteralToken(
			lr.MultiLineString,
			0, // skip leading bytes
			'`',
			3,    // marker length
			true, // allow multiline
			true) // allow escaped
	case midStringToken:
		return lexer.lexStringLiteralToken(
			lr.MultiLineString,
			0, // skip leading bytes
			'"',
			3,    // marker length
			true, // allow multiline
			true) // allow escaped
	case mrbStringToken:
		return lexer.lexStringLiteralToken(
			lr.RawMultiLineString,
			1, // skip leading bytes
			'`',
			3,     // marker length
			true,  // allow multiline
			false) // allow escaped
	case mrdStringToken:
		return lexer.lexStringLiteralToken(
			lr.RawMultiLineString,
			1, // skip leading bytes
			'"',
			3,     // marker length
			true,  // allow multiline
			false) // allow escaped
	case lr.IdentifierToken:
		return lexer.lexIdentifierOrKeywords()
	}

	panic(fmt.Sprintf("RawLexer: unhandled variable length token: %v", symbolId))
}
