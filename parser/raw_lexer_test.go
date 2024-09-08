package parser

import (
	"fmt"
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
	return lex(t, NewRawLexer, input, expected...)
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

		newlines, ok := tokens[1].(TokenCount)
		expect.True(t, ok)
		expect.Equal(t, i+1, newlines.Count)

		expectError(t, tokens[2], "unexpected utf8 rune")
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

	s.lex(t, "...", EllipsisToken)
	s.lex(t, "+...-", AddToken, EllipsisToken, SubToken)
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
	expectError(t, tokens[0], "unexpected utf8 rune")

	tokens = s.lex(t, "+$-", AddToken, ParseErrorToken, SubToken)
	expectError(t, tokens[1], "unexpected utf8 rune")
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

func (s *RawLexerSuite) TestIdentifier(t *testing.T) {
	tokens := s.lex(t, "abc", IdentifierToken)
	expectValue(t, "abc", tokens[0])

	testId := ""
	// Test in loop to check for peek window resizing
	for i := 0; i < 151; i++ {
		testId += testIdxChar(i)
		tokens := s.lex(t, "+"+testId, AddToken, IdentifierToken)
		expectValue(t, testId, tokens[1])

		tokens = s.lex(t, "+"+testId+"-", AddToken, IdentifierToken, SubToken)
		expectValue(t, testId, tokens[1])
	}

	tokens = s.lex(t, "+世界", AddToken, IdentifierToken)
	expectValue(t, "世界", tokens[1])

	tokens = s.lex(t, "+世界-", AddToken, IdentifierToken, SubToken)
	expectValue(t, "世界", tokens[1])

	tokens = s.lex(
		// '+' + '世' + (first two bytes of '界') + '-'
		t, string([]byte{'+', 228, 184, 150, 231, 149, '-'}),
		AddToken, IdentifierToken, ParseErrorToken, ParseErrorToken, SubToken)
	expectValue(t, "世", tokens[1])

	expectError(t, tokens[2], "unexpected utf8 rune")
	expectError(t, tokens[3], "unexpected utf8 rune")
}

func (s *RawLexerSuite) TestLabelDecl(t *testing.T) {
	tokens := s.lex(t, "abc@", LabelDeclToken)
	expectValue(t, "abc@", tokens[0])

	testId := ""
	// Test in loop to check for peek window resizing
	for i := 0; i < 151; i++ {
		testId += testIdxChar(i)
		tokens := s.lex(t, "+"+testId+"@", AddToken, LabelDeclToken)
		expectValue(t, testId+"@", tokens[1])

		tokens = s.lex(t, "+"+testId+"@-", AddToken, LabelDeclToken, SubToken)
		expectValue(t, testId+"@", tokens[1])
	}

	tokens = s.lex(t, "+世界@", AddToken, LabelDeclToken)
	expectValue(t, "世界@", tokens[1])

	tokens = s.lex(t, "+世界@-", AddToken, LabelDeclToken, SubToken)
	expectValue(t, "世界@", tokens[1])
}

func (s *RawLexerSuite) TestJumpLabel(t *testing.T) {
	tokens := s.lex(t, "@abc", JumpLabelToken)
	expectValue(t, "@abc", tokens[0])

	tokens = s.lex(t, "+@-", AddToken, ParseErrorToken, SubToken)
	expectError(t, tokens[1], "no label name associated with @")

	tokens = s.lex(
		t, "@0abc",
		ParseErrorToken, IntegerLiteralToken, IdentifierToken)
	expectError(t, tokens[0], "no label name associated with @")

	tokens = s.lex(t, "+@label", AddToken, JumpLabelToken)
	expectValue(t, "@label", tokens[1])

	tokens = s.lex(t, "+@label-", AddToken, JumpLabelToken, SubToken)
	expectValue(t, "@label", tokens[1])
}

func (s *RawLexerSuite) TestKeywordTokens(t *testing.T) {
	for kw, symbolId := range keywords {
		s.lex(t, "+"+kw+"-", AddToken, symbolId, SubToken)

		tokens := s.lex(t, "+"+kw+"blah-", AddToken, IdentifierToken, SubToken)
		expectValue(t, kw+"blah", tokens[1])
	}
}

func (s *RawLexerSuite) TestLineComment(t *testing.T) {
	tokens := s.lex(t, "//", lineCommentToken)
	expectValue(t, "//", tokens[0])

	comment := "//"
	for i := 0; i < 101; i++ {
		tokens := s.lex(t, "+"+comment, AddToken, lineCommentToken)
		expectValue(t, comment, tokens[1])

		tokens = s.lex(
			t, "+"+comment+"\n",
			AddToken, lineCommentToken, NewlinesToken)
		expectValue(t, comment, tokens[1])

		tokens = s.lex(
			t, "+"+comment+"\r\n",
			AddToken, lineCommentToken, NewlinesToken)
		expectValue(t, comment, tokens[1])

		comment += testIdxChar(i)
	}
}

func (s *RawLexerSuite) TestBlockComment(t *testing.T) {
	tokens := s.lex(t, "/**/", blockCommentToken)
	expectValue(t, "/**/", tokens[0])

	commentBody := ""
	for i := 0; i < 151; i++ {
		tokens := s.lex(t, "+/*"+commentBody, AddToken, ParseErrorToken)
		expectError(t, tokens[1], "comment not terminated")

		tokens = s.lex(t, "+/*"+commentBody+"*", AddToken, ParseErrorToken)
		expectError(t, tokens[1], "comment not terminated")

		tokens = s.lex(t, "+/*"+commentBody+"*/", AddToken, blockCommentToken)
		expectValue(t, "/*"+commentBody+"*/", tokens[1])

		tokens = s.lex(
			t, "+/*"+commentBody+"*/-", AddToken,
			blockCommentToken, SubToken)
		expectValue(t, "/*"+commentBody+"*/", tokens[1])

		commentBody += testIdxChar(i)
	}

	comment := "/* scope 0 */"
	for i := 0; i < 10; i++ {
		tokens := s.lex(t, "+"+comment+"-", AddToken, blockCommentToken, SubToken)
		expectValue(t, comment, tokens[1])

		comment = fmt.Sprintf("/* scope %d %s %s */", i+1, comment, comment)
	}

	commentPrefix := ""
	for i := 0; i < 30; i++ {
		commentPrefix += fmt.Sprintf("/* scope %d ", i)

		comment := commentPrefix
		for j := i; j >= 1; j-- {
			comment += " */"

			tokens := s.lex(t, "+"+comment, AddToken, ParseErrorToken)
			expectError(t, tokens[1], "comment not terminated")
		}

		comment += " */"
		tokens := s.lex(t, "+"+comment+"-", AddToken, blockCommentToken, SubToken)
		expectValue(t, comment, tokens[1])
	}
}

func (s *RawLexerSuite) TestBinaryIntegerLiteral(t *testing.T) {
	tokens := s.lex(t, "0b101", IntegerLiteralToken)
	literal := expectValue(t, "0b101", tokens[0])
	expect.Equal(t, BinaryInteger, literal.SubType)

	tokens = s.lex(t, "+0b", AddToken, ParseErrorToken)
	expectError(t, tokens[1], "binary integer literal has no digits")

	tokens = s.lex(t, "+0b-", AddToken, ParseErrorToken, SubToken)
	expectError(t, tokens[1], "binary integer literal has no digits")

	tokens = s.lex(
		t, "+0B_-",
		AddToken, ParseErrorToken, UnderscoreToken, SubToken)
	expectError(t, tokens[1], "binary integer literal has no digits")

	tokens = s.lex(
		t, "+0B2-",
		AddToken, ParseErrorToken, IntegerLiteralToken, SubToken)
	expectError(t, tokens[1], "binary integer literal has no digits")

	expectValue(t, "2", tokens[2])

	tokens = s.lex(t, "+0b1", AddToken, IntegerLiteralToken)
	expectValue(t, "0b1", tokens[1])

	tokens = s.lex(t, "+0B0-", AddToken, IntegerLiteralToken, SubToken)
	literal = expectValue(t, "0B0", tokens[1])
	expect.Equal(t, BinaryInteger, literal.SubType)

	tokens = s.lex(
		t, "+0B02-",
		AddToken, IntegerLiteralToken, IntegerLiteralToken, SubToken)
	expectValue(t, "0B0", tokens[1])
	expectValue(t, "2", tokens[2])

	tokens = s.lex(t, "+0b_1-", AddToken, IntegerLiteralToken, SubToken)
	expectValue(t, "0b_1", tokens[1])

	tokens = s.lex(t, "+0B_100_010_111-", AddToken, IntegerLiteralToken, SubToken)
	expectValue(t, "0B_100_010_111", tokens[1])

	tokens = s.lex(t, "+0b011101000-", AddToken, IntegerLiteralToken, SubToken)
	expectValue(t, "0b011101000", tokens[1])
}

func (s *RawLexerSuite) TestZeroOPrefixedOctalIntegerLiteral(t *testing.T) {
	tokens := s.lex(t, "0o135", IntegerLiteralToken)
	literal := expectValue(t, "0o135", tokens[0])
	expect.Equal(t, ZeroOPrefixedOctalInteger, literal.SubType)

	tokens = s.lex(t, "+0o", AddToken, ParseErrorToken)
	expectError(t, tokens[1], "0o-prefixed octal integer literal has no digits")

	tokens = s.lex(t, "+0o-", AddToken, ParseErrorToken, SubToken)
	expectError(t, tokens[1], "0o-prefixed octal integer literal has no digits")

	tokens = s.lex(
		t, "+0O_-",
		AddToken, ParseErrorToken, UnderscoreToken, SubToken)
	expectError(t, tokens[1], "0o-prefixed octal integer literal has no digits")

	tokens = s.lex(
		t, "+0O8-",
		AddToken, ParseErrorToken, IntegerLiteralToken, SubToken)
	expectError(t, tokens[1], "0o-prefixed octal integer literal has no digits")

	expectValue(t, "8", tokens[2])

	tokens = s.lex(t, "+0o644", AddToken, IntegerLiteralToken)
	literal = expectValue(t, "0o644", tokens[1])
	expect.Equal(t, ZeroOPrefixedOctalInteger, literal.SubType)

	tokens = s.lex(t, "+0O7-", AddToken, IntegerLiteralToken, SubToken)
	expectValue(t, "0O7", tokens[1])
	expect.Equal(t, ZeroOPrefixedOctalInteger, literal.SubType)

	tokens = s.lex(
		t, "+0O78-",
		AddToken, IntegerLiteralToken, IntegerLiteralToken, SubToken)
	expectValue(t, "0O7", tokens[1])
	expect.Equal(t, ZeroOPrefixedOctalInteger, literal.SubType)
	expectValue(t, "8", tokens[2])

	tokens = s.lex(t, "+0o_1-", AddToken, IntegerLiteralToken, SubToken)
	expectValue(t, "0o_1", tokens[1])

	tokens = s.lex(t, "+0O_123_456_701-", AddToken, IntegerLiteralToken, SubToken)
	expectValue(t, "0O_123_456_701", tokens[1])

	tokens = s.lex(t, "+0o012345670-", AddToken, IntegerLiteralToken, SubToken)
	expectValue(t, "0o012345670", tokens[1])
}

func (s *RawLexerSuite) TestZeroPrefixedOctalIntegerLiteral(t *testing.T) {
	tokens := s.lex(t, "0246", IntegerLiteralToken)
	literal := expectValue(t, "0246", tokens[0])
	expect.Equal(t, ZeroPrefixedOctalInteger, literal.SubType)

	tokens = s.lex(
		t, "+08-",
		AddToken, IntegerLiteralToken, IntegerLiteralToken, SubToken)
	literal = expectValue(t, "0", tokens[1])
	expect.Equal(t, DecimalInteger, literal.SubType)
	expectValue(t, "8", tokens[2])

	tokens = s.lex(t, "+0644", AddToken, IntegerLiteralToken)
	expectValue(t, "0644", tokens[1])

	tokens = s.lex(t, "+07-", AddToken, IntegerLiteralToken, SubToken)
	expectValue(t, "07", tokens[1])

	tokens = s.lex(
		t, "+078-",
		AddToken, IntegerLiteralToken, IntegerLiteralToken, SubToken)
	expectValue(t, "07", tokens[1])
	expectValue(t, "8", tokens[2])

	tokens = s.lex(t, "+0_1-", AddToken, IntegerLiteralToken, SubToken)
	literal = expectValue(t, "0_1", tokens[1])
	expect.Equal(t, ZeroPrefixedOctalInteger, literal.SubType)

	tokens = s.lex(t, "+0_123_456_701-", AddToken, IntegerLiteralToken, SubToken)
	expectValue(t, "0_123_456_701", tokens[1])

	tokens = s.lex(t, "+0012345670-", AddToken, IntegerLiteralToken, SubToken)
	expectValue(t, "0012345670", tokens[1])
}

func (s *RawLexerSuite) TestDecimalIntegerLiteral(t *testing.T) {
	tokens := s.lex(t, "19", IntegerLiteralToken)
	literal := expectValue(t, "19", tokens[0])
	expect.Equal(t, DecimalInteger, literal.SubType)

	for i := 0; i < 20; i++ {
		value := fmt.Sprintf("%d", i)

		tokens := s.lex(
			t, "+"+value,
			AddToken, IntegerLiteralToken)
		literal = expectValue(t, value, tokens[1])
		expect.Equal(t, DecimalInteger, literal.SubType)
	}

	tokens = s.lex(
		t, "+1234567890-",
		AddToken, IntegerLiteralToken, SubToken)
	expectValue(t, "1234567890", tokens[1])

	tokens = s.lex(t, "+644", AddToken, IntegerLiteralToken)
	expectValue(t, "644", tokens[1])

	tokens = s.lex(t, "+1_2_3_4_5-", AddToken, IntegerLiteralToken, SubToken)
	expectValue(t, "1_2_3_4_5", tokens[1])

	tokens = s.lex(t, "+123_456_789_0-", AddToken, IntegerLiteralToken, SubToken)
	expectValue(t, "123_456_789_0", tokens[1])

	tokens = s.lex(
		t, "+10abc-",
		AddToken, IntegerLiteralToken, IdentifierToken, SubToken)
	expectValue(t, "10", tokens[1])
	expectValue(t, "abc", tokens[2])

	tokens = s.lex(
		t, "+5BCD-",
		AddToken, IntegerLiteralToken, IdentifierToken, SubToken)
	expectValue(t, "5", tokens[1])
	expectValue(t, "BCD", tokens[2])
}

func (s *RawLexerSuite) TestHexadecimalIntegerLiteral(t *testing.T) {
	tokens := s.lex(t, "0x19afAF", IntegerLiteralToken)
	literal := expectValue(t, "0x19afAF", tokens[0])
	expect.Equal(t, HexadecimalInteger, literal.SubType)

	tokens = s.lex(t, "+0x", AddToken, ParseErrorToken)
	expectError(t, tokens[1], "hexadecimal integer literal has no digits")

	tokens = s.lex(t, "+0x-", AddToken, ParseErrorToken, SubToken)
	expectError(t, tokens[1], "hexadecimal integer literal has no digits")

	tokens = s.lex(
		t, "+0X_-",
		AddToken, ParseErrorToken, UnderscoreToken, SubToken)
	expectError(t, tokens[1], "hexadecimal integer literal has no digits")
	expectValue(t, "_", tokens[2])

	tokens = s.lex(
		t, "+0Xg-",
		AddToken, ParseErrorToken, IdentifierToken, SubToken)
	expectError(t, tokens[1], "hexadecimal integer literal has no digits")
	expectValue(t, "g", tokens[2])

	tokens = s.lex(t, "+0x0123456789", AddToken, IntegerLiteralToken)
	expectValue(t, "0x0123456789", tokens[1])

	tokens = s.lex(t, "+0Xabcdef-", AddToken, IntegerLiteralToken, SubToken)
	literal = expectValue(t, "0Xabcdef", tokens[1])
	expect.Equal(t, HexadecimalInteger, literal.SubType)

	tokens = s.lex(
		t, "+0X_ABC_DEFG-",
		AddToken, IntegerLiteralToken, IdentifierToken, SubToken)
	expectValue(t, "0X_ABC_DEF", tokens[1])
	expectValue(t, "G", tokens[2])

	tokens = s.lex(t, "+0X123_ABC_abc-", AddToken, IntegerLiteralToken, SubToken)
	expectValue(t, "0X123_ABC_abc", tokens[1])
}

func (s *RawLexerSuite) TestRuneLiteral(t *testing.T) {
	expectInvalidEsc := func(token Token) {
		expectError(t, token, "invalid escaped")
	}

	expectNotTerminated := func(token Token) {
		expectError(t, token, "rune literal not terminated")
	}

	// Various errors
	tokens := s.lex(t, "''", ParseErrorToken)
	expectError(t, tokens[0], "empty rune literal or unescaped '")

	tokens = s.lex(t, "'ab'", ParseErrorToken)
	expectError(t, tokens[0], "more than one character in rune literal")

	tokens = s.lex(t, "'", ParseErrorToken)
	expectNotTerminated(tokens[0])

	tokens = s.lex(t, "'a", ParseErrorToken)
	expectNotTerminated(tokens[0])

	tokens = s.lex(t, "'ab", ParseErrorToken)
	expectNotTerminated(tokens[0])

	tokens = s.lex(t, "'\n", ParseErrorToken, NewlinesToken)
	expectNotTerminated(tokens[0])

	tokens = s.lex(t, "'\n'", ParseErrorToken, NewlinesToken, ParseErrorToken)
	expectNotTerminated(tokens[0])
	expectNotTerminated(tokens[2])

	tokens = s.lex(t, "'\\\n", ParseErrorToken, NewlinesToken)
	expectInvalidEsc(tokens[0])

	tokens = s.lex(t, "'\\\n'", ParseErrorToken, NewlinesToken, ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[2])

	// invalid escaped character \c
	tokens = s.lex(t, "'\\c'", ParseErrorToken, ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped octal \0 (too short)
	tokens = s.lex(t, "'\\0'", ParseErrorToken, ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped octal \07 (too short)
	tokens = s.lex(t, "'\\07'", ParseErrorToken, ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped octal \078 (wrong digit)
	tokens = s.lex(
		t, "'\\078'",
		ParseErrorToken, IntegerLiteralToken, ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectValue(t, "8", tokens[1])

	expectNotTerminated(tokens[2])

	// invalid escaped octal \400 (octal value = 256, i.e., out of bound)
	tokens = s.lex(t, "'\\400'", ParseErrorToken, ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \x (too short)
	tokens = s.lex(t, "'\\x'", ParseErrorToken, ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \xf (too short)
	tokens = s.lex(t, "'\\xf'", ParseErrorToken, ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \xfg (wrong digit)
	tokens = s.lex(
		t, "'\\xfg'",
		ParseErrorToken, IdentifierToken, ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectValue(t, "g", tokens[1])
	expectNotTerminated(tokens[2])

	// invalid escaped hex \u (too short)
	tokens = s.lex(t, "'\\u'", ParseErrorToken, ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \ua (too short)
	tokens = s.lex(t, "'\\ua'", ParseErrorToken, ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \uab (too short)
	tokens = s.lex(t, "'\\uab'", ParseErrorToken, ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \uabc (too short)
	tokens = s.lex(t, "'\\uabc'", ParseErrorToken, ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \uabcg (wrong digit)
	tokens = s.lex(
		t, "'\\uabcg'",
		ParseErrorToken, IdentifierToken, ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectValue(t, "g", tokens[1])
	expectNotTerminated(tokens[2])

	// invalid escaped hex \U (too short)
	tokens = s.lex(t, "'\\U'", ParseErrorToken, ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \Ua (too short)
	tokens = s.lex(t, "'\\Ua'", ParseErrorToken, ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \Uab (too short)
	tokens = s.lex(t, "'\\Uab'", ParseErrorToken, ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \Uabc (too short)
	tokens = s.lex(t, "'\\Uabc'", ParseErrorToken, ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \uabcd (too short)
	tokens = s.lex(t, "'\\Uabcd'", ParseErrorToken, ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \uabcde (too short)
	tokens = s.lex(t, "'\\Uabcde'", ParseErrorToken, ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \uabcdef (too short)
	tokens = s.lex(t, "'\\Uabcdef'", ParseErrorToken, ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \uabcdef0 (too short)
	tokens = s.lex(t, "'\\Uabcdef0'", ParseErrorToken, ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \uabcdef0g (wrong digit)
	tokens = s.lex(
		t, "'\\Uabcdef0g'",
		ParseErrorToken, IdentifierToken, ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectValue(t, "g", tokens[1])
	expectNotTerminated(tokens[2])

	// Basic unescaped byte / unicode

	tokens = s.lex(t, "'a'", RuneLiteralToken)
	expectValue(t, "'a'", tokens[0])

	tokens = s.lex(t, "'世'", RuneLiteralToken)
	expectValue(t, "'世'", tokens[0])

	// Escaped character

	for _, char := range "abfnrtv\\'\"`" {
		escapedChar := fmt.Sprintf("'\\%s'", string(char))
		tokens = s.lex(t, escapedChar, RuneLiteralToken)
		expectValue(t, escapedChar, tokens[0])
	}

	// Escaped octal

	// 0377 = 255 (largest valid octal)
	tokens = s.lex(t, "'\\377'", RuneLiteralToken)
	expectValue(t, "'\\377'", tokens[0])

	tokens = s.lex(t, "'\\000'", RuneLiteralToken)
	expectValue(t, "'\\000'", tokens[0])

	// Escaped \x hex

	tokens = s.lex(t, "'\\xf0'", RuneLiteralToken)
	expectValue(t, "'\\xf0'", tokens[0])

	// Escaped \u hex

	tokens = s.lex(t, "'\\u09af'", RuneLiteralToken)
	expectValue(t, "'\\u09af'", tokens[0])

	// Escaped \U hex

	tokens = s.lex(t, "'\\U09afAF12'", RuneLiteralToken)
	expectValue(t, "'\\U09afAF12'", tokens[0])
}

func (s *RawLexerSuite) TestSingleLineString(t *testing.T) {
	tokens := s.lex(t, `"abc`, ParseErrorToken)
	expectError(t, tokens[0], "string literal not terminated")

	tokens = s.lex(
		t, `"foo\400bar"`,
		ParseErrorToken, IdentifierToken, ParseErrorToken)
	expectError(t, tokens[0], "invalid escaped")
	expectValue(t, "bar", tokens[1])
	expectError(t, tokens[2], "string literal not terminated")

	tokens = s.lex(t, "`abc\n`", ParseErrorToken, NewlinesToken, ParseErrorToken)
	expectError(t, tokens[0], "string literal not terminated")
	expectError(t, tokens[2], "string literal not terminated")

	tokens = s.lex(t, "`mismatch\"", ParseErrorToken)
	expectError(t, tokens[0], "string literal not terminated")

	content := "a世\\`\\\"\\n\\000\\xaf\\u09AF\\U01234567"

	value := "`" + content + "\"`"
	tokens = s.lex(t, "+"+value+"-", AddToken, StringLiteralToken, SubToken)
	str := expectValue(t, value, tokens[1])
	expect.Equal(t, SingleLineString, str.SubType)

	value = "\"" + content + "`\""
	tokens = s.lex(t, "+"+value+"-", AddToken, StringLiteralToken, SubToken)
	str = expectValue(t, value, tokens[1])
	expect.Equal(t, SingleLineString, str.SubType)
}

func (s *RawLexerSuite) TestRawSingleLineString(t *testing.T) {
	tokens := s.lex(t, `r"abc`, ParseErrorToken)
	expectError(t, tokens[0], "string literal not terminated")

	tokens = s.lex(t, "r`abc\n`", ParseErrorToken, NewlinesToken, ParseErrorToken)
	expectError(t, tokens[0], "string literal not terminated")
	expectError(t, tokens[2], "string literal not terminated")

	tokens = s.lex(t, `r"abc"`, StringLiteralToken)
	str := expectValue(t, `r"abc"`, tokens[0])
	expect.Equal(t, RawSingleLineString, str.SubType)

	tokens = s.lex(t, "r`abc \\c \\400 ok`", StringLiteralToken)
	str = expectValue(t, "r`abc \\c \\400 ok`", tokens[0])
	expect.Equal(t, RawSingleLineString, str.SubType)

	tokens = s.lex(t, `r"foo\Uzzz"`, StringLiteralToken)
	expectValue(t, `r"foo\Uzzz"`, tokens[0])
}

func (s *RawLexerSuite) TestMultiLineString(t *testing.T) {
	tokens := s.lex(t, `"""abc" ""`, ParseErrorToken)
	expectError(t, tokens[0], "string literal not terminated")

	tokens = s.lex(
		t, `"""foo\400bar"""`,
		ParseErrorToken, IdentifierToken, ParseErrorToken)
	expectError(t, tokens[0], "invalid escaped")
	expectValue(t, "bar", tokens[1])
	expectError(t, tokens[2], "string literal not terminated")

	tokens = s.lex(t, "```mismatch\"\"\"", ParseErrorToken)
	expectError(t, tokens[0], "string literal not terminated")

	tokens = s.lex(t, "```\"\"\" \"\" \" ` ``\\\n abc\ndef```", StringLiteralToken)
	str := expectValue(t, "```\"\"\" \"\" \" ` ``\\\n abc\ndef```", tokens[0])
	expect.Equal(t, MultiLineString, str.SubType)

	tokens = s.lex(t, "\"\"\"``` \"\" \" ` `` abc\ndef\"\"\"", StringLiteralToken)
	expectValue(t, "\"\"\"``` \"\" \" ` `` abc\ndef\"\"\"", tokens[0])
}

func (s *RawLexerSuite) TestRawMultiLineString(t *testing.T) {
	tokens := s.lex(t, `r"""abc" ""`, ParseErrorToken)
	expectError(t, tokens[0], "string literal not terminated")

	tokens = s.lex(t, "r```mismatch\"\"\"", ParseErrorToken)
	expectError(t, tokens[0], "string literal not terminated")

	tokens = s.lex(t, `r"""foo\400bar"""`, StringLiteralToken)
	str := expectValue(t, `r"""foo\400bar"""`, tokens[0])
	expect.Equal(t, RawMultiLineString, str.SubType)

	tokens = s.lex(t, "r```\"\"\" \"\" \" ` `` abc\ndef```", StringLiteralToken)
	expectValue(t, "r```\"\"\" \"\" \" ` `` abc\ndef```", tokens[0])

	tokens = s.lex(
		t, "r\"\"\"``` \"\" \" ` `` abc\ndef\"\"\"",
		StringLiteralToken)
	expectValue(t, "r\"\"\"``` \"\" \" ` `` abc\ndef\"\"\"", tokens[0])
}

func (s *RawLexerSuite) TestDotDecimalFloat(t *testing.T) {
	s.lex(t, ".", DotToken)
	s.lex(t, ".a", DotToken, IdentifierToken)

	for i := 0; i < 10; i++ {
		value := fmt.Sprintf(".%d", i)
		tokens := s.lex(t, value, FloatLiteralToken)
		float := expectValue(t, value, tokens[0])
		expect.Equal(t, DecimalFloat, float.SubType)

		value = fmt.Sprintf(".01e%d", i)
		tokens = s.lex(t, value, FloatLiteralToken)
		float = expectValue(t, value, tokens[0])
		expect.Equal(t, DecimalFloat, float.SubType)

		value = fmt.Sprintf(".0_2E+0%d", i)
		tokens = s.lex(t, value, FloatLiteralToken)
		float = expectValue(t, value, tokens[0])
		expect.Equal(t, DecimalFloat, float.SubType)

		value = fmt.Sprintf(".0003e-0_0%d", i)
		tokens = s.lex(t, value, FloatLiteralToken)
		float = expectValue(t, value, tokens[0])
		expect.Equal(t, DecimalFloat, float.SubType)
	}

	tokens := s.lex(t, "._1", DotToken, IdentifierToken)
	expectValue(t, "_1", tokens[1])

	tokens = s.lex(t, ".1e", FloatLiteralToken, IdentifierToken)
	expectValue(t, ".1", tokens[0])
	expectValue(t, "e", tokens[1])

	tokens = s.lex(t, ".1e_1", FloatLiteralToken, IdentifierToken)
	expectValue(t, ".1", tokens[0])
	expectValue(t, "e_1", tokens[1])

	tokens = s.lex(
		t, ".1e+_1",
		FloatLiteralToken, IdentifierToken, AddToken, IdentifierToken)
	expectValue(t, ".1", tokens[0])
	expectValue(t, "e", tokens[1])
	expectValue(t, "_1", tokens[3])

	tokens = s.lex(t, ".1ea", FloatLiteralToken, IdentifierToken)
	expectValue(t, ".1", tokens[0])
	expectValue(t, "ea", tokens[1])

	tokens = s.lex(t, ".1p", FloatLiteralToken, IdentifierToken)
	expectValue(t, ".1", tokens[0])
	expectValue(t, "p", tokens[1])

	tokens = s.lex(t, ".1p5", FloatLiteralToken, IdentifierToken)
	expectValue(t, ".1", tokens[0])
	expectValue(t, "p5", tokens[1])

	tokens = s.lex(t, ".e10", DotToken, IdentifierToken)
	expectValue(t, "e10", tokens[1])

	s.lex(t, "+.123-", AddToken, FloatLiteralToken, SubToken)
}

func (s *RawLexerSuite) TestIntPrefixedDecimalFloat(t *testing.T) {
	tokens := s.lex(t, "123e+", IntegerLiteralToken, IdentifierToken, AddToken)
	expectValue(t, "123", tokens[0])
	expectValue(t, "e", tokens[1])

	tokens = s.lex(t, "123ea", IntegerLiteralToken, IdentifierToken)
	expectValue(t, "123", tokens[0])
	expectValue(t, "ea", tokens[1])

	tokens = s.lex(t, "123p0", IntegerLiteralToken, IdentifierToken)
	expectValue(t, "123", tokens[0])
	expectValue(t, "p0", tokens[1])

	for i := 0; i < 10; i++ {
		value := fmt.Sprintf("%d.", i)
		tokens = s.lex(t, value, FloatLiteralToken)
		float := expectValue(t, value, tokens[0])
		expect.Equal(t, DecimalFloat, float.SubType)

		value = fmt.Sprintf("1.%d", i)
		tokens = s.lex(t, value, FloatLiteralToken)
		float = expectValue(t, value, tokens[0])
		expect.Equal(t, DecimalFloat, float.SubType)

		value = fmt.Sprintf("2e%d", i)
		tokens = s.lex(t, value, FloatLiteralToken)
		float = expectValue(t, value, tokens[0])
		expect.Equal(t, DecimalFloat, float.SubType)

		value = fmt.Sprintf("3.E-%d", i)
		tokens = s.lex(t, value, FloatLiteralToken)
		float = expectValue(t, value, tokens[0])
		expect.Equal(t, DecimalFloat, float.SubType)

		value = fmt.Sprintf("4.0_%de+4_4", i)
		tokens = s.lex(t, value, FloatLiteralToken)
		float = expectValue(t, value, tokens[0])
		expect.Equal(t, DecimalFloat, float.SubType)
	}

	tokens = s.lex(t, "123.", FloatLiteralToken)
	float := expectValue(t, "123.", tokens[0])
	expect.Equal(t, DecimalFloat, float.SubType)

	tokens = s.lex(t, "123.456", FloatLiteralToken)
	float = expectValue(t, "123.456", tokens[0])
	expect.Equal(t, DecimalFloat, float.SubType)

	tokens = s.lex(t, "123e0", FloatLiteralToken)
	float = expectValue(t, "123e0", tokens[0])
	expect.Equal(t, DecimalFloat, float.SubType)

	tokens = s.lex(t, "123.456E789", FloatLiteralToken)
	float = expectValue(t, "123.456E789", tokens[0])
	expect.Equal(t, DecimalFloat, float.SubType)

	tokens = s.lex(t, "123.a", FloatLiteralToken, IdentifierToken)
	expectValue(t, "123.", tokens[0])
	expectValue(t, "a", tokens[1])

	tokens = s.lex(t, "123._0", FloatLiteralToken, IdentifierToken)
	expectValue(t, "123.", tokens[0])
	expectValue(t, "_0", tokens[1])

	tokens = s.lex(t, "123.0a", FloatLiteralToken, IdentifierToken)
	expectValue(t, "123.0", tokens[0])
	expectValue(t, "a", tokens[1])

	tokens = s.lex(t, "123.p0", FloatLiteralToken, IdentifierToken)
	expectValue(t, "123.", tokens[0])
	expectValue(t, "p0", tokens[1])
}

func (s *RawLexerSuite) TestHexaecimalFloat(t *testing.T) {
	tokens := s.lex(t, "0xabcp+", IntegerLiteralToken, IdentifierToken, AddToken)
	expectValue(t, "0xabc", tokens[0])
	expectValue(t, "p", tokens[1])

	tokens = s.lex(t, "0xdefpg", IntegerLiteralToken, IdentifierToken)
	expectValue(t, "0xdef", tokens[0])
	expectValue(t, "pg", tokens[1])

	tokens = s.lex(
		t, "0x123e+0",
		IntegerLiteralToken, AddToken, IntegerLiteralToken)
	expectValue(t, "0x123e", tokens[0])
	expectValue(t, "0", tokens[2])

	tokens = s.lex(t, "0x.gp0", ParseErrorToken, DotToken, IdentifierToken)
	expectValue(t, "gp0", tokens[2])

	// decimal place leading digit can't be '_'
	tokens = s.lex(t, "0x._0p0", ParseErrorToken, DotToken, IdentifierToken)
	expectValue(t, "_0p0", tokens[2])

	// missing exponent
	tokens = s.lex(t, "0x123.", IntegerLiteralToken, DotToken)
	expectValue(t, "0x123", tokens[0])

	/// missing exponent
	tokens = s.lex(t, "0x.123", ParseErrorToken, FloatLiteralToken)
	expectValue(t, ".123", tokens[1])

	/// missing exponent
	tokens = s.lex(
		t, "0x123.456",
		IntegerLiteralToken, FloatLiteralToken)
	expectValue(t, "0x123", tokens[0])
	expectValue(t, ".456", tokens[1])

	for i := 0; i < 10; i++ {
		value := fmt.Sprintf("0x%dpa", i)
		tokens = s.lex(t, value, FloatLiteralToken)
		float := expectValue(t, value, tokens[0])
		expect.Equal(t, HexadecimalFloat, float.SubType)

		value = fmt.Sprintf("0x_%dP+b", i)
		tokens = s.lex(t, value, FloatLiteralToken)
		float = expectValue(t, value, tokens[0])
		expect.Equal(t, HexadecimalFloat, float.SubType)

		value = fmt.Sprintf("0X%d.p-C", i)
		tokens = s.lex(t, value, FloatLiteralToken)
		float = expectValue(t, value, tokens[0])
		expect.Equal(t, HexadecimalFloat, float.SubType)

		value = fmt.Sprintf("0x%d.0_3p0_D", i)
		tokens = s.lex(t, value, FloatLiteralToken)
		float = expectValue(t, value, tokens[0])
		expect.Equal(t, HexadecimalFloat, float.SubType)

		value = fmt.Sprintf("0x.%dp0E", i)
		tokens = s.lex(t, value, FloatLiteralToken)
		float = expectValue(t, value, tokens[0])
		expect.Equal(t, HexadecimalFloat, float.SubType)
	}

	tokens = s.lex(t, "0x_0_123.ABC_789p456_DEF", FloatLiteralToken)
	expectValue(t, "0x_0_123.ABC_789p456_DEF", tokens[0])
}
