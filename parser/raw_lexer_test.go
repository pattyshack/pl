package parser

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/pattyshack/gt/testing/expect"
	"github.com/pattyshack/gt/testing/suite"
)

type RawLexerSuite struct {
}

func TestRawLexer(t *testing.T) {
	suite.RunTests(t, &RawLexerSuite{})
}

func (s *RawLexerSuite) lex(
	t *testing.T,
	input string,
	expected ...SymbolId,
) []Token {
	buffer := bytes.NewBufferString(input)

	lexer := NewRawLexer(
		"source.txt",
		buffer,
		RawLexerOptions{
			PreserveCommentContent: true,
			initialPeekWindowSize:  1,
		})

	tokens := []Token{}
	tokenIds := []SymbolId{}
	for {
		token, err := lexer.Next()
		if err == io.EOF {
			break
		}

		expect.Nil(t, err)

		tokens = append(tokens, token)
		tokenIds = append(tokenIds, token.Id())
	}

	expect.Equal(t, expected, tokenIds)

	return tokens
}

func (RawLexerSuite) idChar(idx int) string {
	charSet := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789_"
	return string(charSet[idx%len(charSet)])
}

func (s *RawLexerSuite) TestFloatLiteralToken(t *testing.T) {
	// TODO (reminder: need to check for dot prefix variant .123)
}

func (s *RawLexerSuite) TestRuneLiteralToken(t *testing.T) {
	// TODO
}

func (s *RawLexerSuite) TestStringLiteralToken(t *testing.T) {
	// TODO 4 variants
}

func (s *RawLexerSuite) TestNewlinesTokens(t *testing.T) {
	testInputPrefix := "+"

	// Test in loop to check for peek window resizing
	for i := 0; i < 151; i++ {
		if i%2 == 0 {
			testInputPrefix += "\n"
		} else {
			testInputPrefix += "\r\n"
		}
		s.lex(t, testInputPrefix, AddToken, NewlinesToken)
		s.lex(t, testInputPrefix+"-", AddToken, NewlinesToken, SubToken)

		tokens := s.lex(
			t, testInputPrefix+"\r-",
			AddToken, NewlinesToken, ParseErrorToken, SubToken)

		newlines, ok := tokens[1].(CountSymbol)
		expect.True(t, ok)
		expect.Equal(t, i+1, newlines.Count)

		parseError, ok := tokens[2].(ParseErrorSymbol)
		expect.True(t, ok)
		expect.Error(t, parseError.Error, "unexpected utf8 rune")
	}
}

func (s *RawLexerSuite) TestSpacesTokens(t *testing.T) {
	testInputPrefix := "+"

	// Test in loop to check for peek window resizing
	for i := 0; i < 151; i++ {
		testInputPrefix += " "
		s.lex(t, testInputPrefix, AddToken, spacesToken)
		s.lex(t, testInputPrefix+"-", AddToken, spacesToken, SubToken)
	}
}

func (s *RawLexerSuite) TestLbraceTokens(t *testing.T) {
	s.lex(t, "{", LbraceToken)
	s.lex(t, "+{-", AddToken, LbraceToken, SubToken)
}

func (s *RawLexerSuite) TestRbraceTokens(t *testing.T) {
	s.lex(t, "}", RbraceToken)
	s.lex(t, "+}-", AddToken, RbraceToken, SubToken)
}

func (s *RawLexerSuite) TestLparenTokens(t *testing.T) {
	s.lex(t, "(", LparenToken)
	s.lex(t, "+(-", AddToken, LparenToken, SubToken)
}

func (s *RawLexerSuite) TestRparenTokens(t *testing.T) {
	s.lex(t, ")", RparenToken)
	s.lex(t, "+)-", AddToken, RparenToken, SubToken)
}

func (s *RawLexerSuite) TestLbracketTokens(t *testing.T) {
	s.lex(t, "[", LbracketToken)
	s.lex(t, "+[-", AddToken, LbracketToken, SubToken)
}

func (s *RawLexerSuite) TestRbracketTokens(t *testing.T) {
	s.lex(t, "]", RbracketToken)
	s.lex(t, "+]-", AddToken, RbracketToken, SubToken)
}

func (s *RawLexerSuite) TestDotTokens(t *testing.T) {
	s.lex(t, ".", DotToken)
	s.lex(t, "+.-", AddToken, DotToken, SubToken)

	s.lex(t, "...", DotDotDotToken)
	s.lex(t, "+...-", AddToken, DotDotDotToken, SubToken)
}

func (s *RawLexerSuite) TestCommaTokens(t *testing.T) {
	s.lex(t, ",", CommaToken)
	s.lex(t, "+,-", AddToken, CommaToken, SubToken)
}

func (s *RawLexerSuite) TestQuestionTokens(t *testing.T) {
	s.lex(t, "?", QuestionToken)
	s.lex(t, "+?-", AddToken, QuestionToken, SubToken)
}

func (s *RawLexerSuite) TestSemicolonTokens(t *testing.T) {
	s.lex(t, ";", SemicolonToken)
	s.lex(t, "+;-", AddToken, SemicolonToken, SubToken)
}

func (s *RawLexerSuite) TestColonTokens(t *testing.T) {
	s.lex(t, ":", ColonToken)
	s.lex(t, "+:-", AddToken, ColonToken, SubToken)
}

func (s *RawLexerSuite) TestDollarTokens(t *testing.T) {
	s.lex(t, "$[", DollarLbracketToken)
	s.lex(t, "+$[-", AddToken, DollarLbracketToken, SubToken)

	tokens := s.lex(t, "$", ParseErrorToken)
	parseError, ok := tokens[0].(ParseErrorSymbol)
	expect.True(t, ok)
	expect.Error(t, parseError.Error, "unexpected utf8 rune")

	tokens = s.lex(t, "+$-", AddToken, ParseErrorToken, SubToken)
	parseError, ok = tokens[1].(ParseErrorSymbol)
	expect.True(t, ok)
	expect.Error(t, parseError.Error, "unexpected utf8 rune")
}

func (s *RawLexerSuite) TestAddTokens(t *testing.T) {
	s.lex(t, "+", AddToken)
	s.lex(t, "*+/", MulToken, AddToken, DivToken)

	s.lex(t, "++", AddOneAssignToken)
	s.lex(t, "*++/", MulToken, AddOneAssignToken, DivToken)

	s.lex(t, "+=", AddAssignToken)
	s.lex(t, "*+=/", MulToken, AddAssignToken, DivToken)
}

func (s *RawLexerSuite) TestSubTokens(t *testing.T) {
	s.lex(t, "-", SubToken)
	s.lex(t, "*-/", MulToken, SubToken, DivToken)

	s.lex(t, "--", SubOneAssignToken)
	s.lex(t, "*--/", MulToken, SubOneAssignToken, DivToken)

	s.lex(t, "-=", SubAssignToken)
	s.lex(t, "*-=/", MulToken, SubAssignToken, DivToken)
}

func (s *RawLexerSuite) TestMulTokens(t *testing.T) {
	s.lex(t, "*", MulToken)
	s.lex(t, "+*-", AddToken, MulToken, SubToken)

	s.lex(t, "*=", MulAssignToken)
	s.lex(t, "+*=-", AddToken, MulAssignToken, SubToken)
}

func (s *RawLexerSuite) TestDivTokens(t *testing.T) {
	s.lex(t, "/", DivToken)
	s.lex(t, "+/-", AddToken, DivToken, SubToken)

	s.lex(t, "/=", DivAssignToken)
	s.lex(t, "+/=-", AddToken, DivAssignToken, SubToken)
}

func (s *RawLexerSuite) TestModTokens(t *testing.T) {
	s.lex(t, "%", ModToken)
	s.lex(t, "+%-", AddToken, ModToken, SubToken)

	s.lex(t, "%=", ModAssignToken)
	s.lex(t, "+%=-", AddToken, ModAssignToken, SubToken)
}

func (s *RawLexerSuite) TestBitNegTokens(t *testing.T) {
	s.lex(t, "~", BitNegToken)
	s.lex(t, "+~-", AddToken, BitNegToken, SubToken)

	s.lex(t, "~=", BitNegAssignToken)
	s.lex(t, "+~=-", AddToken, BitNegAssignToken, SubToken)

	s.lex(t, "~~", TildeTildeToken)
	s.lex(t, "+~~-", AddToken, TildeTildeToken, SubToken)
}

func (s *RawLexerSuite) TestBitAndTokens(t *testing.T) {
	s.lex(t, "&", BitAndToken)
	s.lex(t, "+&-", AddToken, BitAndToken, SubToken)

	s.lex(t, "&=", BitAndAssignToken)
	s.lex(t, "+&=-", AddToken, BitAndAssignToken, SubToken)
}

func (s *RawLexerSuite) TestBitXorTokens(t *testing.T) {
	s.lex(t, "^", BitXorToken)
	s.lex(t, "+^-", AddToken, BitXorToken, SubToken)

	s.lex(t, "^=", BitXorAssignToken)
	s.lex(t, "+^=-", AddToken, BitXorAssignToken, SubToken)
}

func (s *RawLexerSuite) TestBitOrTokens(t *testing.T) {
	s.lex(t, "|", BitOrToken)
	s.lex(t, "+|-", AddToken, BitOrToken, SubToken)

	s.lex(t, "|=", BitOrAssignToken)
	s.lex(t, "+|=-", AddToken, BitOrAssignToken, SubToken)
}

func (s *RawLexerSuite) TestEqualTokens(t *testing.T) {
	s.lex(t, "=", AssignToken)
	s.lex(t, "(=)", LparenToken, AssignToken, RparenToken)

	s.lex(t, "==", EqualToken)
	s.lex(t, "(==)", LparenToken, EqualToken, RparenToken)
}

func (s *RawLexerSuite) TestExclaimTokens(t *testing.T) {
	s.lex(t, "!", ExclaimToken)
	s.lex(t, "+!-", AddToken, ExclaimToken, SubToken)

	s.lex(t, "!=", NotEqualToken)
	s.lex(t, "+!=-", AddToken, NotEqualToken, SubToken)
}

func (s *RawLexerSuite) TestGreaterTokens(t *testing.T) {
	s.lex(t, ">", GreaterToken)
	s.lex(t, "+>-", AddToken, GreaterToken, SubToken)

	s.lex(t, ">=", GreaterOrEqualToken)
	s.lex(t, "+>=-", AddToken, GreaterOrEqualToken, SubToken)

	s.lex(t, ">>", BitRshiftToken)
	s.lex(t, "+>>-", AddToken, BitRshiftToken, SubToken)

	s.lex(t, ">>=", BitRshiftAssignToken)
	s.lex(t, "+>>=-", AddToken, BitRshiftAssignToken, SubToken)
}

func (s *RawLexerSuite) TestLessTokens(t *testing.T) {
	s.lex(t, "<", LessToken)
	s.lex(t, "+<-", AddToken, LessToken, SubToken)

	s.lex(t, "<=", LessOrEqualToken)
	s.lex(t, "+<=-", AddToken, LessOrEqualToken, SubToken)

	s.lex(t, "<<", BitLshiftToken)
	s.lex(t, "+<<-", AddToken, BitLshiftToken, SubToken)

	s.lex(t, "<<=", BitLshiftAssignToken)
	s.lex(t, "+<<=-", AddToken, BitLshiftAssignToken, SubToken)
}

func (s *RawLexerSuite) TestIdentifierToken(t *testing.T) {
	tokens := s.lex(t, "abc", IdentifierToken)
	value, ok := tokens[0].(ValueSymbol)
	expect.True(t, ok)
	expect.Equal(t, "abc", value.Value)

	testId := ""
	// Test in loop to check for peek window resizing
	for i := 0; i < 151; i++ {
		testId += s.idChar(i)
		tokens := s.lex(t, "+"+testId, AddToken, IdentifierToken)
		value, ok := tokens[1].(ValueSymbol)
		expect.True(t, ok)
		expect.Equal(t, testId, value.Value)

		tokens = s.lex(t, "+"+testId+"-", AddToken, IdentifierToken, SubToken)
		value, ok = tokens[1].(ValueSymbol)
		expect.True(t, ok)
		expect.Equal(t, testId, value.Value)
	}

	tokens = s.lex(t, "+世界", AddToken, IdentifierToken)
	value, ok = tokens[1].(ValueSymbol)
	expect.True(t, ok)
	expect.Equal(t, "世界", value.Value)

	tokens = s.lex(t, "+世界-", AddToken, IdentifierToken, SubToken)
	value, ok = tokens[1].(ValueSymbol)
	expect.True(t, ok)
	expect.Equal(t, "世界", value.Value)

	tokens = s.lex(
		// '+' + '世' + (first two bytes of '界') + '-'
		t, string([]byte{'+', 228, 184, 150, 231, 149, '-'}),
		AddToken, IdentifierToken, ParseErrorToken, ParseErrorToken, SubToken)
	value, ok = tokens[1].(ValueSymbol)
	expect.True(t, ok)
	expect.Equal(t, "世", value.Value)

	parseError, ok := tokens[2].(ParseErrorSymbol)
	expect.True(t, ok)
	expect.Error(t, parseError.Error, "unexpected utf8 rune")

	parseError, ok = tokens[3].(ParseErrorSymbol)
	expect.True(t, ok)
	expect.Error(t, parseError.Error, "unexpected utf8 rune")
}

func (s *RawLexerSuite) TestLabelDeclToken(t *testing.T) {
	tokens := s.lex(t, "abc@", LabelDeclToken)
	value, ok := tokens[0].(ValueSymbol)
	expect.True(t, ok)
	expect.Equal(t, "abc@", value.Value)

	testId := ""
	// Test in loop to check for peek window resizing
	for i := 0; i < 151; i++ {
		testId += s.idChar(i)
		tokens := s.lex(t, "+"+testId+"@", AddToken, LabelDeclToken)
		value, ok := tokens[1].(ValueSymbol)
		expect.True(t, ok)
		expect.Equal(t, testId+"@", value.Value)

		tokens = s.lex(t, "+"+testId+"@-", AddToken, LabelDeclToken, SubToken)
		value, ok = tokens[1].(ValueSymbol)
		expect.True(t, ok)
		expect.Equal(t, testId+"@", value.Value)
	}

	tokens = s.lex(t, "+世界@", AddToken, LabelDeclToken)
	value, ok = tokens[1].(ValueSymbol)
	expect.True(t, ok)
	expect.Equal(t, "世界@", value.Value)

	tokens = s.lex(t, "+世界@-", AddToken, LabelDeclToken, SubToken)
	value, ok = tokens[1].(ValueSymbol)
	expect.True(t, ok)
	expect.Equal(t, "世界@", value.Value)
}

func (s *RawLexerSuite) TestJumpLabelToken(t *testing.T) {
	tokens := s.lex(t, "@abc", JumpLabelToken)
	value, ok := tokens[0].(ValueSymbol)
	expect.True(t, ok)
	expect.Equal(t, "@abc", value.Value)

	tokens = s.lex(t, "+@-", AddToken, ParseErrorToken, SubToken)
	parseError, ok := tokens[1].(ParseErrorSymbol)
	expect.True(t, ok)
	expect.Error(t, parseError.Error, "no label name associated with @")

	//
	// TODO check for invalid identifier (first char is a number) @0abc
	//

	tokens = s.lex(t, "+@label", AddToken, JumpLabelToken)
	symbol, ok := tokens[1].(ValueSymbol)
	expect.True(t, ok)
	expect.Equal(t, "@label", symbol.Value)

	tokens = s.lex(t, "+@label-", AddToken, JumpLabelToken, SubToken)
	symbol, ok = tokens[1].(ValueSymbol)
	expect.True(t, ok)
	expect.Equal(t, "@label", symbol.Value)
}

func (s *RawLexerSuite) TestKeywordTokens(t *testing.T) {
	for kw, symbolId := range keywords {
		s.lex(t, "+"+kw+"-", AddToken, symbolId, SubToken)

		tokens := s.lex(t, "+"+kw+"blah-", AddToken, IdentifierToken, SubToken)
		symbol, ok := tokens[1].(ValueSymbol)
		expect.True(t, ok)
		expect.Equal(t, kw+"blah", symbol.Value)
	}
}

func (s *RawLexerSuite) TestLineCommentToken(t *testing.T) {
	tokens := s.lex(t, "//", lineCommentToken)
	value, ok := tokens[0].(ValueSymbol)
	expect.True(t, ok)
	expect.Equal(t, "//", value.Value)

	comment := "//"
	for i := 0; i < 101; i++ {
		tokens := s.lex(t, "+"+comment, AddToken, lineCommentToken)
		value, ok := tokens[1].(ValueSymbol)
		expect.True(t, ok)
		expect.Equal(t, comment, value.Value)

		tokens = s.lex(
			t, "+"+comment+"\n",
			AddToken, lineCommentToken, NewlinesToken)
		value, ok = tokens[1].(ValueSymbol)
		expect.True(t, ok)
		expect.Equal(t, comment, value.Value)

		tokens = s.lex(
			t, "+"+comment+"\r\n",
			AddToken, lineCommentToken, NewlinesToken)
		value, ok = tokens[1].(ValueSymbol)
		expect.True(t, ok)
		expect.Equal(t, comment, value.Value)

		comment += s.idChar(i)
	}
}

func (s *RawLexerSuite) TestBlockCommentToken(t *testing.T) {
	tokens := s.lex(t, "/**/", blockCommentToken)
	value, ok := tokens[0].(ValueSymbol)
	expect.True(t, ok)
	expect.Equal(t, "/**/", value.Value)

	commentBody := ""
	for i := 0; i < 151; i++ {
		tokens := s.lex(t, "+/*"+commentBody, AddToken, ParseErrorToken)
		parseError, ok := tokens[1].(ParseErrorSymbol)
		expect.True(t, ok)
		expect.Error(t, parseError.Error, "comment not terminated")

		tokens = s.lex(t, "+/*"+commentBody+"*", AddToken, ParseErrorToken)
		parseError, ok = tokens[1].(ParseErrorSymbol)
		expect.True(t, ok)
		expect.Error(t, parseError.Error, "comment not terminated")

		tokens = s.lex(t, "+/*"+commentBody+"*/", AddToken, blockCommentToken)
		value, ok := tokens[1].(ValueSymbol)
		expect.True(t, ok)
		expect.Equal(t, value.Value, "/*"+commentBody+"*/")

		tokens = s.lex(
			t, "+/*"+commentBody+"*/-", AddToken,
			blockCommentToken, SubToken)
		value, ok = tokens[1].(ValueSymbol)
		expect.True(t, ok)
		expect.Equal(t, value.Value, "/*"+commentBody+"*/")

		commentBody += s.idChar(i)
	}

	comment := "/* scope 0 */"
	for i := 0; i < 10; i++ {
		tokens := s.lex(t, "+"+comment+"-", AddToken, blockCommentToken, SubToken)
		value, ok := tokens[1].(ValueSymbol)
		expect.True(t, ok)
		expect.Equal(t, value.Value, comment)

		comment = fmt.Sprintf("/* scope %d %s %s */", i+1, comment, comment)
	}

	commentPrefix := ""
	for i := 0; i < 30; i++ {
		commentPrefix += fmt.Sprintf("/* scope %d ", i)

		comment := commentPrefix
		for j := i; j >= 1; j-- {
			comment += " */"

			tokens := s.lex(t, "+"+comment, AddToken, ParseErrorToken)
			parseError, ok := tokens[1].(ParseErrorSymbol)
			expect.True(t, ok)
			expect.Error(t, parseError.Error, "comment not terminated")
		}

		comment += " */"
		tokens := s.lex(t, "+"+comment+"-", AddToken, blockCommentToken, SubToken)
		value, ok := tokens[1].(ValueSymbol)
		expect.True(t, ok)
		expect.Equal(t, value.Value, comment)
	}
}

func (s *RawLexerSuite) TestBinaryIntegerLiteral(t *testing.T) {
	tokens := s.lex(t, "0b101", IntegerLiteralToken)
	literal, ok := tokens[0].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "0b101", literal.Value)
	expect.Equal(t, BinaryInteger, literal.IntegerSubType)

	tokens = s.lex(t, "+0b", AddToken, ParseErrorToken)
	parseError, ok := tokens[1].(ParseErrorSymbol)
	expect.True(t, ok)
	expect.Error(t, parseError.Error, "binary integer has no digits")

	tokens = s.lex(t, "+0b-", AddToken, ParseErrorToken, SubToken)
	parseError, ok = tokens[1].(ParseErrorSymbol)
	expect.True(t, ok)
	expect.Error(t, parseError.Error, "binary integer has no digits")

	tokens = s.lex(
		t, "+0B_-",
		AddToken, ParseErrorToken, IdentifierToken, SubToken)
	parseError, ok = tokens[1].(ParseErrorSymbol)
	expect.True(t, ok)
	expect.Error(t, parseError.Error, "binary integer has no digits")

	tokens = s.lex(
		t, "+0B2-",
		AddToken, ParseErrorToken, IntegerLiteralToken, SubToken)
	parseError, ok = tokens[1].(ParseErrorSymbol)
	expect.True(t, ok)
	expect.Error(t, parseError.Error, "binary integer has no digits")

	literal, ok = tokens[2].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "2", literal.Value)
	expect.Equal(t, DecimalInteger, literal.IntegerSubType)

	tokens = s.lex(t, "+0b1", AddToken, IntegerLiteralToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "0b1", literal.Value)
	expect.Equal(t, BinaryInteger, literal.IntegerSubType)

	tokens = s.lex(t, "+0B0-", AddToken, IntegerLiteralToken, SubToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "0B0", literal.Value)
	expect.Equal(t, BinaryInteger, literal.IntegerSubType)

	tokens = s.lex(
		t, "+0B02-",
		AddToken, IntegerLiteralToken, IntegerLiteralToken, SubToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "0B0", literal.Value)
	expect.Equal(t, BinaryInteger, literal.IntegerSubType)

	literal, ok = tokens[2].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "2", literal.Value)
	expect.Equal(t, DecimalInteger, literal.IntegerSubType)

	tokens = s.lex(t, "+0b_1-", AddToken, IntegerLiteralToken, SubToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "0b_1", literal.Value)
	expect.Equal(t, BinaryInteger, literal.IntegerSubType)

	tokens = s.lex(t, "+0B_100_010_111-", AddToken, IntegerLiteralToken, SubToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "0B_100_010_111", literal.Value)
	expect.Equal(t, BinaryInteger, literal.IntegerSubType)

	tokens = s.lex(t, "+0b011101000-", AddToken, IntegerLiteralToken, SubToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "0b011101000", literal.Value)
	expect.Equal(t, BinaryInteger, literal.IntegerSubType)
}

func (s *RawLexerSuite) TestZeroOPrefixedOctalIntegerLiteral(t *testing.T) {
	tokens := s.lex(t, "0o135", IntegerLiteralToken)
	literal, ok := tokens[0].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "0o135", literal.Value)
	expect.Equal(t, ZeroOPrefixedOctalInteger, literal.IntegerSubType)

	tokens = s.lex(t, "+0o", AddToken, ParseErrorToken)
	parseError, ok := tokens[1].(ParseErrorSymbol)
	expect.True(t, ok)
	expect.Error(t, parseError.Error, "0o-prefixed octal integer has no digits")

	tokens = s.lex(t, "+0o-", AddToken, ParseErrorToken, SubToken)
	parseError, ok = tokens[1].(ParseErrorSymbol)
	expect.True(t, ok)
	expect.Error(t, parseError.Error, "0o-prefixed octal integer has no digits")

	tokens = s.lex(
		t, "+0O_-",
		AddToken, ParseErrorToken, IdentifierToken, SubToken)
	parseError, ok = tokens[1].(ParseErrorSymbol)
	expect.True(t, ok)
	expect.Error(t, parseError.Error, "0o-prefixed octal integer has no digits")

	tokens = s.lex(
		t, "+0O8-",
		AddToken, ParseErrorToken, IntegerLiteralToken, SubToken)
	parseError, ok = tokens[1].(ParseErrorSymbol)
	expect.True(t, ok)
	expect.Error(t, parseError.Error, "0o-prefixed octal integer has no digits")

	literal, ok = tokens[2].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "8", literal.Value)
	expect.Equal(t, DecimalInteger, literal.IntegerSubType)

	tokens = s.lex(t, "+0o644", AddToken, IntegerLiteralToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "0o644", literal.Value)
	expect.Equal(t, ZeroOPrefixedOctalInteger, literal.IntegerSubType)

	tokens = s.lex(t, "+0O7-", AddToken, IntegerLiteralToken, SubToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "0O7", literal.Value)
	expect.Equal(t, ZeroOPrefixedOctalInteger, literal.IntegerSubType)

	tokens = s.lex(
		t, "+0O78-",
		AddToken, IntegerLiteralToken, IntegerLiteralToken, SubToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "0O7", literal.Value)
	expect.Equal(t, ZeroOPrefixedOctalInteger, literal.IntegerSubType)

	literal, ok = tokens[2].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "8", literal.Value)
	expect.Equal(t, DecimalInteger, literal.IntegerSubType)

	tokens = s.lex(t, "+0o_1-", AddToken, IntegerLiteralToken, SubToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "0o_1", literal.Value)
	expect.Equal(t, ZeroOPrefixedOctalInteger, literal.IntegerSubType)

	tokens = s.lex(t, "+0O_123_456_701-", AddToken, IntegerLiteralToken, SubToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "0O_123_456_701", literal.Value)
	expect.Equal(t, ZeroOPrefixedOctalInteger, literal.IntegerSubType)

	tokens = s.lex(t, "+0o012345670-", AddToken, IntegerLiteralToken, SubToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "0o012345670", literal.Value)
	expect.Equal(t, ZeroOPrefixedOctalInteger, literal.IntegerSubType)
}

func (s *RawLexerSuite) TestZeroPrefixedOctalIntegerLiteral(t *testing.T) {
	tokens := s.lex(t, "0246", IntegerLiteralToken)
	literal, ok := tokens[0].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "0246", literal.Value)
	expect.Equal(t, ZeroPrefixedOctalInteger, literal.IntegerSubType)

	tokens = s.lex(
		t, "+08-",
		AddToken, IntegerLiteralToken, IntegerLiteralToken, SubToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "0", literal.Value)
	expect.Equal(t, DecimalInteger, literal.IntegerSubType)

	literal, ok = tokens[2].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "8", literal.Value)
	expect.Equal(t, DecimalInteger, literal.IntegerSubType)

	tokens = s.lex(t, "+0644", AddToken, IntegerLiteralToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "0644", literal.Value)
	expect.Equal(t, ZeroPrefixedOctalInteger, literal.IntegerSubType)

	tokens = s.lex(t, "+07-", AddToken, IntegerLiteralToken, SubToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "07", literal.Value)
	expect.Equal(t, ZeroPrefixedOctalInteger, literal.IntegerSubType)

	tokens = s.lex(
		t, "+078-",
		AddToken, IntegerLiteralToken, IntegerLiteralToken, SubToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "07", literal.Value)
	expect.Equal(t, ZeroPrefixedOctalInteger, literal.IntegerSubType)

	literal, ok = tokens[2].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "8", literal.Value)
	expect.Equal(t, DecimalInteger, literal.IntegerSubType)

	tokens = s.lex(t, "+0_1-", AddToken, IntegerLiteralToken, SubToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "0_1", literal.Value)
	expect.Equal(t, ZeroPrefixedOctalInteger, literal.IntegerSubType)

	tokens = s.lex(t, "+0_123_456_701-", AddToken, IntegerLiteralToken, SubToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "0_123_456_701", literal.Value)
	expect.Equal(t, ZeroPrefixedOctalInteger, literal.IntegerSubType)

	tokens = s.lex(t, "+0012345670-", AddToken, IntegerLiteralToken, SubToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "0012345670", literal.Value)
	expect.Equal(t, ZeroPrefixedOctalInteger, literal.IntegerSubType)
}

func (s *RawLexerSuite) TestDecimalIntegerLiteral(t *testing.T) {
	tokens := s.lex(t, "19", IntegerLiteralToken)
	literal, ok := tokens[0].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "19", literal.Value)
	expect.Equal(t, DecimalInteger, literal.IntegerSubType)

	for i := 0; i < 20; i++ {
		value := fmt.Sprintf("%d", i)

		tokens := s.lex(
			t, "+"+value,
			AddToken, IntegerLiteralToken)
		literal, ok := tokens[1].(IntegerLiteralSymbol)
		expect.True(t, ok)
		expect.Equal(t, value, literal.Value)
		expect.Equal(t, DecimalInteger, literal.IntegerSubType)
	}

	tokens = s.lex(
		t, "+1234567890-",
		AddToken, IntegerLiteralToken, SubToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "1234567890", literal.Value)
	expect.Equal(t, DecimalInteger, literal.IntegerSubType)

	tokens = s.lex(t, "+644", AddToken, IntegerLiteralToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "644", literal.Value)
	expect.Equal(t, DecimalInteger, literal.IntegerSubType)

	tokens = s.lex(t, "+1_2_3_4_5-", AddToken, IntegerLiteralToken, SubToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "1_2_3_4_5", literal.Value)
	expect.Equal(t, DecimalInteger, literal.IntegerSubType)

	tokens = s.lex(t, "+123_456_789_0-", AddToken, IntegerLiteralToken, SubToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "123_456_789_0", literal.Value)
	expect.Equal(t, DecimalInteger, literal.IntegerSubType)

	tokens = s.lex(
		t, "+10abc-",
		AddToken, IntegerLiteralToken, IdentifierToken, SubToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "10", literal.Value)
	expect.Equal(t, DecimalInteger, literal.IntegerSubType)

	value, ok := tokens[2].(ValueSymbol)
	expect.True(t, ok)
	expect.Equal(t, "abc", value.Value)

	tokens = s.lex(
		t, "+5BCD-",
		AddToken, IntegerLiteralToken, IdentifierToken, SubToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "5", literal.Value)
	expect.Equal(t, DecimalInteger, literal.IntegerSubType)

	value, ok = tokens[2].(ValueSymbol)
	expect.True(t, ok)
	expect.Equal(t, "BCD", value.Value)
}

func (s *RawLexerSuite) TestHexadecimalIntegerLiteral(t *testing.T) {
	tokens := s.lex(t, "0x19afAF", IntegerLiteralToken)
	literal, ok := tokens[0].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "0x19afAF", literal.Value)
	expect.Equal(t, HexadecimalInteger, literal.IntegerSubType)

	tokens = s.lex(t, "+0x", AddToken, ParseErrorToken)
	parseError, ok := tokens[1].(ParseErrorSymbol)
	expect.True(t, ok)
	expect.Error(t, parseError.Error, "hexadecimal integer has no digits")

	tokens = s.lex(t, "+0x-", AddToken, ParseErrorToken, SubToken)
	parseError, ok = tokens[1].(ParseErrorSymbol)
	expect.True(t, ok)
	expect.Error(t, parseError.Error, "hexadecimal integer has no digits")

	tokens = s.lex(
		t, "+0X_-",
		AddToken, ParseErrorToken, IdentifierToken, SubToken)
	parseError, ok = tokens[1].(ParseErrorSymbol)
	expect.True(t, ok)
	expect.Error(t, parseError.Error, "hexadecimal integer has no digits")

	value, ok := tokens[2].(ValueSymbol)
	expect.True(t, ok)
	expect.Equal(t, "_", value.Value)

	tokens = s.lex(
		t, "+0Xg-",
		AddToken, ParseErrorToken, IdentifierToken, SubToken)
	parseError, ok = tokens[1].(ParseErrorSymbol)
	expect.True(t, ok)
	expect.Error(t, parseError.Error, "hexadecimal integer has no digits")

	value, ok = tokens[2].(ValueSymbol)
	expect.True(t, ok)
	expect.Equal(t, "g", value.Value)

	tokens = s.lex(t, "+0x0123456789", AddToken, IntegerLiteralToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "0x0123456789", literal.Value)
	expect.Equal(t, HexadecimalInteger, literal.IntegerSubType)

	tokens = s.lex(t, "+0Xabcdef-", AddToken, IntegerLiteralToken, SubToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "0Xabcdef", literal.Value)
	expect.Equal(t, HexadecimalInteger, literal.IntegerSubType)

	tokens = s.lex(
		t, "+0X_ABC_DEFG-",
		AddToken, IntegerLiteralToken, IdentifierToken, SubToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "0X_ABC_DEF", literal.Value)
	expect.Equal(t, HexadecimalInteger, literal.IntegerSubType)

	value, ok = tokens[2].(ValueSymbol)
	expect.True(t, ok)
	expect.Equal(t, "G", value.Value)

	tokens = s.lex(t, "+0X123_ABC_abc-", AddToken, IntegerLiteralToken, SubToken)
	literal, ok = tokens[1].(IntegerLiteralSymbol)
	expect.True(t, ok)
	expect.Equal(t, "0X123_ABC_abc", literal.Value)
	expect.Equal(t, HexadecimalInteger, literal.IntegerSubType)
}
