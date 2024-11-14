package lexer

import (
	"fmt"
	"testing"

	"github.com/pattyshack/gt/lexutil"
	"github.com/pattyshack/gt/testing/expect"
	"github.com/pattyshack/gt/testing/suite"

	"github.com/pattyshack/pl/parser/lr"
)

type RawLexerSuite struct {
}

func TestRawLexer(t *testing.T) {
	suite.RunTests(t, &RawLexerSuite{})
}

func (s *RawLexerSuite) lex(
	t *testing.T,
	input string,
	expected ...lr.SymbolId,
) []lr.Token {
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
		s.lex(t, testInputPrefix, lr.AddToken, lr.NewlinesToken)
		s.lex(t, testInputPrefix+"-", lr.AddToken, lr.NewlinesToken, lr.SubToken)

		tokens := s.lex(
			t, testInputPrefix+"\r-",
			lr.AddToken, lr.NewlinesToken, lr.ParseErrorToken, lr.SubToken)

		newlines, ok := tokens[1].(*lr.TokenCount)
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
		s.lex(t, testInputPrefix, lr.AddToken, spacesToken)
		s.lex(t, testInputPrefix+"-", lr.AddToken, spacesToken, lr.SubToken)
	}
}

func (s *RawLexerSuite) TestLbraceTokens(t *testing.T) {
	s.lex(t, "{", lr.LbraceToken)
	s.lex(t, "+{-", lr.AddToken, lr.LbraceToken, lr.SubToken)
}

func (s *RawLexerSuite) TestRbraceTokens(t *testing.T) {
	s.lex(t, "}", lr.RbraceToken)
	s.lex(t, "+}-", lr.AddToken, lr.RbraceToken, lr.SubToken)
}

func (s *RawLexerSuite) TestLparenTokens(t *testing.T) {
	s.lex(t, "(", lr.LparenToken)
	s.lex(t, "+(-", lr.AddToken, lr.LparenToken, lr.SubToken)
}

func (s *RawLexerSuite) TestRparenTokens(t *testing.T) {
	s.lex(t, ")", lr.RparenToken)
	s.lex(t, "+)-", lr.AddToken, lr.RparenToken, lr.SubToken)
}

func (s *RawLexerSuite) TestLbracketTokens(t *testing.T) {
	s.lex(t, "[", lr.LbracketToken)
	s.lex(t, "+[-", lr.AddToken, lr.LbracketToken, lr.SubToken)
}

func (s *RawLexerSuite) TestRbracketTokens(t *testing.T) {
	s.lex(t, "]", lr.RbracketToken)
	s.lex(t, "+]-", lr.AddToken, lr.RbracketToken, lr.SubToken)
}

func (s *RawLexerSuite) TestDotTokens(t *testing.T) {
	s.lex(t, ".", lr.DotToken)
	s.lex(t, "+.-", lr.AddToken, lr.DotToken, lr.SubToken)

	s.lex(t, "...", lr.EllipsisToken)
	s.lex(t, "+...-", lr.AddToken, lr.EllipsisToken, lr.SubToken)
}

func (s *RawLexerSuite) TestCommaTokens(t *testing.T) {
	s.lex(t, ",", lr.CommaToken)
	s.lex(t, "+,-", lr.AddToken, lr.CommaToken, lr.SubToken)
}

func (s *RawLexerSuite) TestQuestionTokens(t *testing.T) {
	s.lex(t, "?", lr.QuestionToken)
	s.lex(t, "+?-", lr.AddToken, lr.QuestionToken, lr.SubToken)
}

func (s *RawLexerSuite) TestSemicolonTokens(t *testing.T) {
	s.lex(t, ";", lr.SemicolonToken)
	s.lex(t, "+;-", lr.AddToken, lr.SemicolonToken, lr.SubToken)
}

func (s *RawLexerSuite) TestColonTokens(t *testing.T) {
	s.lex(t, ":", lr.ColonToken)
	s.lex(t, "+:-", lr.AddToken, lr.ColonToken, lr.SubToken)
}

func (s *RawLexerSuite) TestDollarTokens(t *testing.T) {
	s.lex(t, "$[", lr.DollarToken, lr.LbracketToken)
	s.lex(t, "+$[-", lr.AddToken, lr.DollarToken, lr.LbracketToken, lr.SubToken)
	s.lex(t, "$", lr.DollarToken)
	s.lex(t, "+$-", lr.AddToken, lr.DollarToken, lr.SubToken)
}

func (s *RawLexerSuite) TestAddTokens(t *testing.T) {
	s.lex(t, "+", lr.AddToken)
	s.lex(t, "*+/", lr.MulToken, lr.AddToken, lr.DivToken)

	s.lex(t, "++", lr.AddOneAssignToken)
	s.lex(t, "*++/", lr.MulToken, lr.AddOneAssignToken, lr.DivToken)

	s.lex(t, "+=", lr.AddAssignToken)
	s.lex(t, "*+=/", lr.MulToken, lr.AddAssignToken, lr.DivToken)
}

func (s *RawLexerSuite) TestSubTokens(t *testing.T) {
	s.lex(t, "-", lr.SubToken)
	s.lex(t, "*-/", lr.MulToken, lr.SubToken, lr.DivToken)

	s.lex(t, "--", lr.SubOneAssignToken)
	s.lex(t, "*--/", lr.MulToken, lr.SubOneAssignToken, lr.DivToken)

	s.lex(t, "-=", lr.SubAssignToken)
	s.lex(t, "*-=/", lr.MulToken, lr.SubAssignToken, lr.DivToken)
}

func (s *RawLexerSuite) TestMulTokens(t *testing.T) {
	s.lex(t, "*", lr.MulToken)
	s.lex(t, "+*-", lr.AddToken, lr.MulToken, lr.SubToken)

	s.lex(t, "*=", lr.MulAssignToken)
	s.lex(t, "+*=-", lr.AddToken, lr.MulAssignToken, lr.SubToken)
}

func (s *RawLexerSuite) TestDivTokens(t *testing.T) {
	s.lex(t, "/", lr.DivToken)
	s.lex(t, "+/-", lr.AddToken, lr.DivToken, lr.SubToken)

	s.lex(t, "/=", lr.DivAssignToken)
	s.lex(t, "+/=-", lr.AddToken, lr.DivAssignToken, lr.SubToken)
}

func (s *RawLexerSuite) TestModTokens(t *testing.T) {
	s.lex(t, "%", lr.ModToken)
	s.lex(t, "+%-", lr.AddToken, lr.ModToken, lr.SubToken)

	s.lex(t, "%=", lr.ModAssignToken)
	s.lex(t, "+%=-", lr.AddToken, lr.ModAssignToken, lr.SubToken)
}

func (s *RawLexerSuite) TestTildeTokens(t *testing.T) {
	s.lex(t, "~", lr.TildeToken)
	s.lex(t, "+~-", lr.AddToken, lr.TildeToken, lr.SubToken)

	s.lex(t, "~~", lr.TildeTildeToken)
	s.lex(t, "+~~-", lr.AddToken, lr.TildeTildeToken, lr.SubToken)
}

func (s *RawLexerSuite) TestBitAndTokens(t *testing.T) {
	s.lex(t, "&", lr.BitAndToken)
	s.lex(t, "+&-", lr.AddToken, lr.BitAndToken, lr.SubToken)

	s.lex(t, "&=", lr.BitAndAssignToken)
	s.lex(t, "+&=-", lr.AddToken, lr.BitAndAssignToken, lr.SubToken)
}

func (s *RawLexerSuite) TestBitXorTokens(t *testing.T) {
	s.lex(t, "^", lr.BitXorToken)
	s.lex(t, "+^-", lr.AddToken, lr.BitXorToken, lr.SubToken)

	s.lex(t, "^=", lr.BitXorAssignToken)
	s.lex(t, "+^=-", lr.AddToken, lr.BitXorAssignToken, lr.SubToken)
}

func (s *RawLexerSuite) TestBitOrTokens(t *testing.T) {
	s.lex(t, "|", lr.BitOrToken)
	s.lex(t, "+|-", lr.AddToken, lr.BitOrToken, lr.SubToken)

	s.lex(t, "|=", lr.BitOrAssignToken)
	s.lex(t, "+|=-", lr.AddToken, lr.BitOrAssignToken, lr.SubToken)
}

func (s *RawLexerSuite) TestEqualTokens(t *testing.T) {
	s.lex(t, "=", lr.AssignToken)
	s.lex(t, "(=)", lr.LparenToken, lr.AssignToken, lr.RparenToken)

	s.lex(t, "==", lr.EqualToken)
	s.lex(t, "(==)", lr.LparenToken, lr.EqualToken, lr.RparenToken)
}

func (s *RawLexerSuite) TestExclaimTokens(t *testing.T) {
	s.lex(t, "!", lr.ExclaimToken)
	s.lex(t, "+!-", lr.AddToken, lr.ExclaimToken, lr.SubToken)

	s.lex(t, "!=", lr.NotEqualToken)
	s.lex(t, "+!=-", lr.AddToken, lr.NotEqualToken, lr.SubToken)
}

func (s *RawLexerSuite) TestGreaterTokens(t *testing.T) {
	s.lex(t, ">", lr.GreaterToken)
	s.lex(t, "+>-", lr.AddToken, lr.GreaterToken, lr.SubToken)

	s.lex(t, ">=", lr.GreaterOrEqualToken)
	s.lex(t, "+>=-", lr.AddToken, lr.GreaterOrEqualToken, lr.SubToken)

	s.lex(t, ">>", lr.BitRshiftToken)
	s.lex(t, "+>>-", lr.AddToken, lr.BitRshiftToken, lr.SubToken)

	s.lex(t, ">>=", lr.BitRshiftAssignToken)
	s.lex(t, "+>>=-", lr.AddToken, lr.BitRshiftAssignToken, lr.SubToken)
}

func (s *RawLexerSuite) TestLessTokens(t *testing.T) {
	s.lex(t, "<", lr.LessToken)
	s.lex(t, "+</", lr.AddToken, lr.LessToken, lr.DivToken)

	s.lex(t, "<=", lr.LessOrEqualToken)
	s.lex(t, "+<=-", lr.AddToken, lr.LessOrEqualToken, lr.SubToken)

	s.lex(t, "<<", lr.BitLshiftToken)
	s.lex(t, "+<<-", lr.AddToken, lr.BitLshiftToken, lr.SubToken)

	s.lex(t, "<<=", lr.BitLshiftAssignToken)
	s.lex(t, "+<<=-", lr.AddToken, lr.BitLshiftAssignToken, lr.SubToken)

	s.lex(t, "<-", lr.ArrowToken)
	s.lex(t, "+<--", lr.AddToken, lr.ArrowToken, lr.SubToken)
}

func (s *RawLexerSuite) TestIdentifier(t *testing.T) {
	tokens := s.lex(t, "abc", lr.IdentifierToken)
	expectValue(t, "abc", tokens[0])

	testId := ""
	// Test in loop to check for peek window resizing
	for i := 0; i < 151; i++ {
		testId += testIdxChar(i)
		tokens := s.lex(t, "+"+testId, lr.AddToken, lr.IdentifierToken)
		expectValue(t, testId, tokens[1])

		tokens = s.lex(
			t, "+"+testId+"-",
			lr.AddToken, lr.IdentifierToken, lr.SubToken)
		expectValue(t, testId, tokens[1])
	}

	tokens = s.lex(t, "+世界", lr.AddToken, lr.IdentifierToken)
	expectValue(t, "世界", tokens[1])

	tokens = s.lex(t, "+世界-", lr.AddToken, lr.IdentifierToken, lr.SubToken)
	expectValue(t, "世界", tokens[1])

	tokens = s.lex(
		// '+' + '世' + (first two bytes of '界') + '-'
		t, string([]byte{'+', 228, 184, 150, 231, 149, '-'}),
		lr.AddToken, lr.IdentifierToken, lr.ParseErrorToken,
		lr.ParseErrorToken, lr.SubToken)
	expectValue(t, "世", tokens[1])

	expectError(t, tokens[2], "unexpected utf8 rune")
	expectError(t, tokens[3], "unexpected utf8 rune")
}

func (s *RawLexerSuite) TestPound(t *testing.T) {
	s.lex(t, "#[", lr.PoundToken, lr.LbracketToken)
}

func (s *RawLexerSuite) TestLabelDecl(t *testing.T) {
	tokens := s.lex(t, "abc@", lr.IdentifierToken, lr.AtToken)
	expectValue(t, "abc", tokens[0])

	testId := ""
	// Test in loop to check for peek window resizing
	for i := 0; i < 151; i++ {
		testId += testIdxChar(i)
		tokens := s.lex(
			t, "+"+testId+"@",
			lr.AddToken, lr.IdentifierToken, lr.AtToken)
		expectValue(t, testId, tokens[1])

		tokens = s.lex(
			t, "+"+testId+"@-",
			lr.AddToken, lr.IdentifierToken, lr.AtToken, lr.SubToken)
		expectValue(t, testId, tokens[1])
	}

	tokens = s.lex(t, "+世界@", lr.AddToken, lr.IdentifierToken, lr.AtToken)
	expectValue(t, "世界", tokens[1])

	tokens = s.lex(t, "+世界@-",
		lr.AddToken, lr.IdentifierToken, lr.AtToken, lr.SubToken)
	expectValue(t, "世界", tokens[1])
}

func (s *RawLexerSuite) TestJumpLabel(t *testing.T) {
	tokens := s.lex(t, "@abc", lr.AtToken, lr.IdentifierToken)
	expectValue(t, "abc", tokens[1])

	s.lex(t, "+@-", lr.AddToken, lr.AtToken, lr.SubToken)

	s.lex(
		t, "@0abc",
		lr.AtToken, lr.IntegerLiteralToken, lr.IdentifierToken)

	tokens = s.lex(t, "+@label", lr.AddToken, lr.AtToken, lr.IdentifierToken)
	expectValue(t, "label", tokens[2])

	tokens = s.lex(
		t, "+@label-",
		lr.AddToken, lr.AtToken, lr.IdentifierToken, lr.SubToken)
	expectValue(t, "label", tokens[2])
}

func (s *RawLexerSuite) TestKeywordTokens(t *testing.T) {
	for kw, symbolId := range keywords {
		s.lex(t, "+"+kw+"-", lr.AddToken, symbolId, lr.SubToken)

		tokens := s.lex(
			t, "+"+kw+"blah-",
			lr.AddToken, lr.IdentifierToken, lr.SubToken)
		expectValue(t, kw+"blah", tokens[1])
	}
}

func (s *RawLexerSuite) TestLineComment(t *testing.T) {
	tokens := s.lex(t, "//", lineCommentToken)
	expectValue(t, "//", tokens[0])

	comment := "//"
	for i := 0; i < 101; i++ {
		tokens := s.lex(t, "+"+comment, lr.AddToken, lineCommentToken)
		expectValue(t, comment, tokens[1])

		tokens = s.lex(
			t, "+"+comment+"\n",
			lr.AddToken, lineCommentToken, lr.NewlinesToken)
		expectValue(t, comment, tokens[1])

		tokens = s.lex(
			t, "+"+comment+"\r\n",
			lr.AddToken, lineCommentToken, lr.NewlinesToken)
		expectValue(t, comment, tokens[1])

		comment += testIdxChar(i)
	}
}

func (s *RawLexerSuite) TestBlockComment(t *testing.T) {
	tokens := s.lex(t, "/**/", blockCommentToken)
	expectValue(t, "/**/", tokens[0])

	commentBody := ""
	for i := 0; i < 151; i++ {
		tokens := s.lex(t, "+/*"+commentBody, lr.AddToken, lr.ParseErrorToken)
		expectError(t, tokens[1], "comment not terminated")

		tokens = s.lex(t, "+/*"+commentBody+"*", lr.AddToken, lr.ParseErrorToken)
		expectError(t, tokens[1], "comment not terminated")

		tokens = s.lex(t, "+/*"+commentBody+"*/", lr.AddToken, blockCommentToken)
		expectValue(t, "/*"+commentBody+"*/", tokens[1])

		tokens = s.lex(
			t, "+/*"+commentBody+"*/-", lr.AddToken,
			blockCommentToken, lr.SubToken)
		expectValue(t, "/*"+commentBody+"*/", tokens[1])

		commentBody += testIdxChar(i)
	}

	comment := "/* scope 0 */"
	for i := 0; i < 10; i++ {
		tokens := s.lex(
			t, "+"+comment+"-",
			lr.AddToken, blockCommentToken, lr.SubToken)
		expectValue(t, comment, tokens[1])

		comment = fmt.Sprintf("/* scope %d %s %s */", i+1, comment, comment)
	}

	commentPrefix := ""
	for i := 0; i < 30; i++ {
		commentPrefix += fmt.Sprintf("/* scope %d ", i)

		comment := commentPrefix
		for j := i; j >= 1; j-- {
			comment += " */"

			tokens := s.lex(t, "+"+comment, lr.AddToken, lr.ParseErrorToken)
			expectError(t, tokens[1], "comment not terminated")
		}

		comment += " */"
		tokens := s.lex(
			t, "+"+comment+"-",
			lr.AddToken, blockCommentToken, lr.SubToken)
		expectValue(t, comment, tokens[1])
	}
}

func (s *RawLexerSuite) TestBinaryIntegerLiteral(t *testing.T) {
	tokens := s.lex(t, "0b101", lr.IntegerLiteralToken)
	literal := expectValue(t, "0b101", tokens[0])
	expect.Equal(t, lexutil.BinaryInteger, literal.SubType)

	tokens = s.lex(t, "+0b", lr.AddToken, lr.ParseErrorToken)
	expectError(t, tokens[1], "binary integer has no digits")

	tokens = s.lex(t, "+0b-", lr.AddToken, lr.ParseErrorToken, lr.SubToken)
	expectError(t, tokens[1], "binary integer has no digits")

	tokens = s.lex(
		t, "+0B_-",
		lr.AddToken, lr.ParseErrorToken, lr.UnderscoreToken, lr.SubToken)
	expectError(t, tokens[1], "binary integer has no digits")

	tokens = s.lex(
		t, "+0B2-",
		lr.AddToken, lr.ParseErrorToken, lr.IntegerLiteralToken, lr.SubToken)
	expectError(t, tokens[1], "binary integer has no digits")

	expectValue(t, "2", tokens[2])

	tokens = s.lex(t, "+0b1", lr.AddToken, lr.IntegerLiteralToken)
	expectValue(t, "0b1", tokens[1])

	tokens = s.lex(t, "+0B0-", lr.AddToken, lr.IntegerLiteralToken, lr.SubToken)
	literal = expectValue(t, "0B0", tokens[1])
	expect.Equal(t, lexutil.BinaryInteger, literal.SubType)

	tokens = s.lex(
		t, "+0B02-",
		lr.AddToken, lr.IntegerLiteralToken, lr.IntegerLiteralToken, lr.SubToken)
	expectValue(t, "0B0", tokens[1])
	expectValue(t, "2", tokens[2])

	tokens = s.lex(t, "+0b_1-", lr.AddToken, lr.IntegerLiteralToken, lr.SubToken)
	expectValue(t, "0b_1", tokens[1])

	tokens = s.lex(
		t, "+0B_100_010_111-",
		lr.AddToken, lr.IntegerLiteralToken, lr.SubToken)
	expectValue(t, "0B_100_010_111", tokens[1])

	tokens = s.lex(
		t, "+0b011101000-",
		lr.AddToken, lr.IntegerLiteralToken, lr.SubToken)
	expectValue(t, "0b011101000", tokens[1])
}

func (s *RawLexerSuite) TestZeroOPrefixedOctalIntegerLiteral(t *testing.T) {
	tokens := s.lex(t, "0o135", lr.IntegerLiteralToken)
	literal := expectValue(t, "0o135", tokens[0])
	expect.Equal(t, lexutil.ZeroOPrefixedOctalInteger, literal.SubType)

	tokens = s.lex(t, "+0o", lr.AddToken, lr.ParseErrorToken)
	expectError(t, tokens[1], "0o-prefixed octal integer has no digits")

	tokens = s.lex(t, "+0o-", lr.AddToken, lr.ParseErrorToken, lr.SubToken)
	expectError(t, tokens[1], "0o-prefixed octal integer has no digits")

	tokens = s.lex(
		t, "+0O_-",
		lr.AddToken, lr.ParseErrorToken, lr.UnderscoreToken, lr.SubToken)
	expectError(t, tokens[1], "0o-prefixed octal integer has no digits")

	tokens = s.lex(
		t, "+0O8-",
		lr.AddToken, lr.ParseErrorToken, lr.IntegerLiteralToken, lr.SubToken)
	expectError(t, tokens[1], "0o-prefixed octal integer has no digits")

	expectValue(t, "8", tokens[2])

	tokens = s.lex(t, "+0o644", lr.AddToken, lr.IntegerLiteralToken)
	literal = expectValue(t, "0o644", tokens[1])
	expect.Equal(t, lexutil.ZeroOPrefixedOctalInteger, literal.SubType)

	tokens = s.lex(t, "+0O7-", lr.AddToken, lr.IntegerLiteralToken, lr.SubToken)
	expectValue(t, "0O7", tokens[1])
	expect.Equal(t, lexutil.ZeroOPrefixedOctalInteger, literal.SubType)

	tokens = s.lex(
		t, "+0O78-",
		lr.AddToken, lr.IntegerLiteralToken, lr.IntegerLiteralToken, lr.SubToken)
	expectValue(t, "0O7", tokens[1])
	expect.Equal(t, lexutil.ZeroOPrefixedOctalInteger, literal.SubType)
	expectValue(t, "8", tokens[2])

	tokens = s.lex(t, "+0o_1-", lr.AddToken, lr.IntegerLiteralToken, lr.SubToken)
	expectValue(t, "0o_1", tokens[1])

	tokens = s.lex(
		t, "+0O_123_456_701-",
		lr.AddToken, lr.IntegerLiteralToken, lr.SubToken)
	expectValue(t, "0O_123_456_701", tokens[1])

	tokens = s.lex(
		t, "+0o012345670-",
		lr.AddToken, lr.IntegerLiteralToken, lr.SubToken)
	expectValue(t, "0o012345670", tokens[1])
}

func (s *RawLexerSuite) TestZeroPrefixedOctalIntegerLiteral(t *testing.T) {
	tokens := s.lex(t, "0246", lr.IntegerLiteralToken)
	literal := expectValue(t, "0246", tokens[0])
	expect.Equal(t, lexutil.ZeroPrefixedOctalInteger, literal.SubType)

	tokens = s.lex(
		t, "+08-",
		lr.AddToken, lr.IntegerLiteralToken, lr.IntegerLiteralToken, lr.SubToken)
	literal = expectValue(t, "0", tokens[1])
	expect.Equal(t, lexutil.DecimalInteger, literal.SubType)
	expectValue(t, "8", tokens[2])

	tokens = s.lex(t, "+0644", lr.AddToken, lr.IntegerLiteralToken)
	expectValue(t, "0644", tokens[1])

	tokens = s.lex(t, "+07-", lr.AddToken, lr.IntegerLiteralToken, lr.SubToken)
	expectValue(t, "07", tokens[1])

	tokens = s.lex(
		t, "+078-",
		lr.AddToken, lr.IntegerLiteralToken, lr.IntegerLiteralToken, lr.SubToken)
	expectValue(t, "07", tokens[1])
	expectValue(t, "8", tokens[2])

	tokens = s.lex(t, "+0_1-", lr.AddToken, lr.IntegerLiteralToken, lr.SubToken)
	literal = expectValue(t, "0_1", tokens[1])
	expect.Equal(t, lexutil.ZeroPrefixedOctalInteger, literal.SubType)

	tokens = s.lex(
		t, "+0_123_456_701-",
		lr.AddToken, lr.IntegerLiteralToken, lr.SubToken)
	expectValue(t, "0_123_456_701", tokens[1])

	tokens = s.lex(
		t, "+0012345670-",
		lr.AddToken, lr.IntegerLiteralToken, lr.SubToken)
	expectValue(t, "0012345670", tokens[1])
}

func (s *RawLexerSuite) TestDecimalIntegerLiteral(t *testing.T) {
	tokens := s.lex(t, "19", lr.IntegerLiteralToken)
	literal := expectValue(t, "19", tokens[0])
	expect.Equal(t, lexutil.DecimalInteger, literal.SubType)

	for i := 0; i < 20; i++ {
		value := fmt.Sprintf("%d", i)

		tokens := s.lex(
			t, "+"+value,
			lr.AddToken, lr.IntegerLiteralToken)
		literal = expectValue(t, value, tokens[1])
		expect.Equal(t, lexutil.DecimalInteger, literal.SubType)
	}

	tokens = s.lex(
		t, "+1234567890-",
		lr.AddToken, lr.IntegerLiteralToken, lr.SubToken)
	expectValue(t, "1234567890", tokens[1])

	tokens = s.lex(t, "+644", lr.AddToken, lr.IntegerLiteralToken)
	expectValue(t, "644", tokens[1])

	tokens = s.lex(
		t, "+1_2_3_4_5-",
		lr.AddToken, lr.IntegerLiteralToken, lr.SubToken)
	expectValue(t, "1_2_3_4_5", tokens[1])

	tokens = s.lex(
		t, "+123_456_789_0-",
		lr.AddToken, lr.IntegerLiteralToken, lr.SubToken)
	expectValue(t, "123_456_789_0", tokens[1])

	tokens = s.lex(
		t, "+10abc-",
		lr.AddToken, lr.IntegerLiteralToken, lr.IdentifierToken, lr.SubToken)
	expectValue(t, "10", tokens[1])
	expectValue(t, "abc", tokens[2])

	tokens = s.lex(
		t, "+5BCD-",
		lr.AddToken, lr.IntegerLiteralToken, lr.IdentifierToken, lr.SubToken)
	expectValue(t, "5", tokens[1])
	expectValue(t, "BCD", tokens[2])
}

func (s *RawLexerSuite) TestHexadecimalIntegerLiteral(t *testing.T) {
	tokens := s.lex(t, "0x19afAF", lr.IntegerLiteralToken)
	literal := expectValue(t, "0x19afAF", tokens[0])
	expect.Equal(t, lexutil.HexadecimalInteger, literal.SubType)

	tokens = s.lex(t, "+0x", lr.AddToken, lr.ParseErrorToken)
	expectError(t, tokens[1], "hexadecimal integer has no digits")

	tokens = s.lex(t, "+0x-", lr.AddToken, lr.ParseErrorToken, lr.SubToken)
	expectError(t, tokens[1], "hexadecimal integer has no digits")

	tokens = s.lex(
		t, "+0X_-",
		lr.AddToken, lr.ParseErrorToken, lr.UnderscoreToken, lr.SubToken)
	expectError(t, tokens[1], "hexadecimal integer has no digits")
	expectValue(t, "_", tokens[2])

	tokens = s.lex(
		t, "+0Xg-",
		lr.AddToken, lr.ParseErrorToken, lr.IdentifierToken, lr.SubToken)
	expectError(t, tokens[1], "hexadecimal integer has no digits")
	expectValue(t, "g", tokens[2])

	tokens = s.lex(t, "+0x0123456789", lr.AddToken, lr.IntegerLiteralToken)
	expectValue(t, "0x0123456789", tokens[1])

	tokens = s.lex(
		t, "+0Xabcdef-",
		lr.AddToken, lr.IntegerLiteralToken, lr.SubToken)
	literal = expectValue(t, "0Xabcdef", tokens[1])
	expect.Equal(t, lexutil.HexadecimalInteger, literal.SubType)

	tokens = s.lex(
		t, "+0X_ABC_DEFG-",
		lr.AddToken, lr.IntegerLiteralToken, lr.IdentifierToken, lr.SubToken)
	expectValue(t, "0X_ABC_DEF", tokens[1])
	expectValue(t, "G", tokens[2])

	tokens = s.lex(
		t, "+0X123_ABC_abc-",
		lr.AddToken, lr.IntegerLiteralToken, lr.SubToken)
	expectValue(t, "0X123_ABC_abc", tokens[1])
}

func (s *RawLexerSuite) TestRuneLiteral(t *testing.T) {
	expectInvalidEsc := func(token lr.Token) {
		expectError(t, token, "invalid escaped")
	}

	expectNotTerminated := func(token lr.Token) {
		expectError(t, token, "rune not terminated")
	}

	// Various errors
	tokens := s.lex(t, "''", lr.ParseErrorToken)
	expectError(t, tokens[0], "empty rune literal or unescaped '")

	tokens = s.lex(t, "'ab'", lr.ParseErrorToken)
	expectError(t, tokens[0], "more than one character in rune literal")

	tokens = s.lex(t, "'", lr.ParseErrorToken)
	expectNotTerminated(tokens[0])

	tokens = s.lex(t, "'a", lr.ParseErrorToken)
	expectNotTerminated(tokens[0])

	tokens = s.lex(t, "'ab", lr.ParseErrorToken)
	expectNotTerminated(tokens[0])

	tokens = s.lex(t, "'\n", lr.ParseErrorToken, lr.NewlinesToken)
	expectNotTerminated(tokens[0])

	tokens = s.lex(
		t, "'\n'",
		lr.ParseErrorToken, lr.NewlinesToken, lr.ParseErrorToken)
	expectNotTerminated(tokens[0])
	expectNotTerminated(tokens[2])

	tokens = s.lex(t, "'\\\n", lr.ParseErrorToken, lr.NewlinesToken)
	expectInvalidEsc(tokens[0])

	tokens = s.lex(
		t, "'\\\n'",
		lr.ParseErrorToken, lr.NewlinesToken, lr.ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[2])

	// invalid escaped character \c
	tokens = s.lex(t, "'\\c'", lr.ParseErrorToken, lr.ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped octal \0 (too short)
	tokens = s.lex(t, "'\\0'", lr.ParseErrorToken, lr.ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped octal \07 (too short)
	tokens = s.lex(t, "'\\07'", lr.ParseErrorToken, lr.ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped octal \078 (wrong digit)
	tokens = s.lex(
		t, "'\\078'",
		lr.ParseErrorToken, lr.IntegerLiteralToken, lr.ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectValue(t, "8", tokens[1])

	expectNotTerminated(tokens[2])

	// invalid escaped octal \400 (octal value = 256, i.e., out of bound)
	tokens = s.lex(t, "'\\400'", lr.ParseErrorToken, lr.ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \x (too short)
	tokens = s.lex(t, "'\\x'", lr.ParseErrorToken, lr.ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \xf (too short)
	tokens = s.lex(t, "'\\xf'", lr.ParseErrorToken, lr.ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \xfg (wrong digit)
	tokens = s.lex(
		t, "'\\xfg'",
		lr.ParseErrorToken, lr.IdentifierToken, lr.ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectValue(t, "g", tokens[1])
	expectNotTerminated(tokens[2])

	// invalid escaped hex \u (too short)
	tokens = s.lex(t, "'\\u'", lr.ParseErrorToken, lr.ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \ua (too short)
	tokens = s.lex(t, "'\\ua'", lr.ParseErrorToken, lr.ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \uab (too short)
	tokens = s.lex(t, "'\\uab'", lr.ParseErrorToken, lr.ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \uabc (too short)
	tokens = s.lex(t, "'\\uabc'", lr.ParseErrorToken, lr.ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \uabcg (wrong digit)
	tokens = s.lex(
		t, "'\\uabcg'",
		lr.ParseErrorToken, lr.IdentifierToken, lr.ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectValue(t, "g", tokens[1])
	expectNotTerminated(tokens[2])

	// invalid escaped hex \U (too short)
	tokens = s.lex(t, "'\\U'", lr.ParseErrorToken, lr.ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \Ua (too short)
	tokens = s.lex(t, "'\\Ua'", lr.ParseErrorToken, lr.ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \Uab (too short)
	tokens = s.lex(t, "'\\Uab'", lr.ParseErrorToken, lr.ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \Uabc (too short)
	tokens = s.lex(t, "'\\Uabc'", lr.ParseErrorToken, lr.ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \uabcd (too short)
	tokens = s.lex(t, "'\\Uabcd'", lr.ParseErrorToken, lr.ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \uabcde (too short)
	tokens = s.lex(t, "'\\Uabcde'", lr.ParseErrorToken, lr.ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \uabcdef (too short)
	tokens = s.lex(t, "'\\Uabcdef'", lr.ParseErrorToken, lr.ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \uabcdef0 (too short)
	tokens = s.lex(t, "'\\Uabcdef0'", lr.ParseErrorToken, lr.ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectNotTerminated(tokens[1])

	// invalid escaped hex \uabcdef0g (wrong digit)
	tokens = s.lex(
		t, "'\\Uabcdef0g'",
		lr.ParseErrorToken, lr.IdentifierToken, lr.ParseErrorToken)
	expectInvalidEsc(tokens[0])
	expectValue(t, "g", tokens[1])
	expectNotTerminated(tokens[2])

	// Basic unescaped byte / unicode

	tokens = s.lex(t, "'a'", lr.RuneLiteralToken)
	expectValue(t, "'a'", tokens[0])

	tokens = s.lex(t, "'世'", lr.RuneLiteralToken)
	expectValue(t, "'世'", tokens[0])

	// Escaped character

	for _, char := range "abfnrtv\\'\"`" {
		escapedChar := fmt.Sprintf("'\\%s'", string(char))
		tokens = s.lex(t, escapedChar, lr.RuneLiteralToken)
		expectValue(t, escapedChar, tokens[0])
	}

	// Escaped octal

	// 0377 = 255 (largest valid octal)
	tokens = s.lex(t, "'\\377'", lr.RuneLiteralToken)
	expectValue(t, "'\\377'", tokens[0])

	tokens = s.lex(t, "'\\000'", lr.RuneLiteralToken)
	expectValue(t, "'\\000'", tokens[0])

	// Escaped \x hex

	tokens = s.lex(t, "'\\xf0'", lr.RuneLiteralToken)
	expectValue(t, "'\\xf0'", tokens[0])

	// Escaped \u hex

	tokens = s.lex(t, "'\\u09af'", lr.RuneLiteralToken)
	expectValue(t, "'\\u09af'", tokens[0])

	// Escaped \U hex

	tokens = s.lex(t, "'\\U09afAF12'", lr.RuneLiteralToken)
	expectValue(t, "'\\U09afAF12'", tokens[0])
}

func (s *RawLexerSuite) TestSingleLineString(t *testing.T) {
	tokens := s.lex(t, `"abc`, lr.ParseErrorToken)
	expectError(t, tokens[0], "string not terminated")

	tokens = s.lex(
		t, `"foo\400bar"`,
		lr.ParseErrorToken, lr.IdentifierToken, lr.ParseErrorToken)
	expectError(t, tokens[0], "invalid escaped")
	expectValue(t, "bar", tokens[1])
	expectError(t, tokens[2], "string not terminated")

	tokens = s.lex(
		t, "`abc\n`",
		lr.ParseErrorToken, lr.NewlinesToken, lr.ParseErrorToken)
	expectError(t, tokens[0], "string not terminated")
	expectError(t, tokens[2], "string not terminated")

	tokens = s.lex(t, "`mismatch\"", lr.ParseErrorToken)
	expectError(t, tokens[0], "string not terminated")

	content := "a世\\`\\\"\\n\\000\\xaf\\u09AF\\U01234567"

	value := "`" + content + "\"`"
	tokens = s.lex(
		t, "+"+value+"-",
		lr.AddToken, lr.StringLiteralToken, lr.SubToken)
	str := expectValue(t, value, tokens[1])
	expect.Equal(t, lexutil.SingleLineString, str.SubType)

	value = "\"" + content + "`\""
	tokens = s.lex(
		t, "+"+value+"-",
		lr.AddToken, lr.StringLiteralToken, lr.SubToken)
	str = expectValue(t, value, tokens[1])
	expect.Equal(t, lexutil.SingleLineString, str.SubType)
}

func (s *RawLexerSuite) TestRawSingleLineString(t *testing.T) {
	tokens := s.lex(t, `r"abc`, lr.ParseErrorToken)
	expectError(t, tokens[0], "string not terminated")

	tokens = s.lex(
		t, "r`abc\n`",
		lr.ParseErrorToken, lr.NewlinesToken, lr.ParseErrorToken)
	expectError(t, tokens[0], "string not terminated")
	expectError(t, tokens[2], "string not terminated")

	tokens = s.lex(t, `r"abc"`, lr.StringLiteralToken)
	str := expectValue(t, `r"abc"`, tokens[0])
	expect.Equal(t, lexutil.RawSingleLineString, str.SubType)

	tokens = s.lex(t, "r`abc \\c \\400 ok`", lr.StringLiteralToken)
	str = expectValue(t, "r`abc \\c \\400 ok`", tokens[0])
	expect.Equal(t, lexutil.RawSingleLineString, str.SubType)

	tokens = s.lex(t, `r"foo\Uzzz"`, lr.StringLiteralToken)
	expectValue(t, `r"foo\Uzzz"`, tokens[0])
}

func (s *RawLexerSuite) TestMultiLineString(t *testing.T) {
	tokens := s.lex(t, `"""abc" ""`, lr.ParseErrorToken)
	expectError(t, tokens[0], "string not terminated")

	tokens = s.lex(
		t, `"""foo\400bar"""`,
		lr.ParseErrorToken, lr.IdentifierToken, lr.ParseErrorToken)
	expectError(t, tokens[0], "invalid escaped")
	expectValue(t, "bar", tokens[1])
	expectError(t, tokens[2], "string not terminated")

	tokens = s.lex(t, "```mismatch\"\"\"", lr.ParseErrorToken)
	expectError(t, tokens[0], "string not terminated")

	tokens = s.lex(
		t, "```\"\"\" \"\" \" ` ``\\\n abc\ndef```",
		lr.StringLiteralToken)
	str := expectValue(t, "```\"\"\" \"\" \" ` ``\\\n abc\ndef```", tokens[0])
	expect.Equal(t, lexutil.MultiLineString, str.SubType)

	tokens = s.lex(
		t, "\"\"\"``` \"\" \" ` `` abc\ndef\"\"\"",
		lr.StringLiteralToken)
	expectValue(t, "\"\"\"``` \"\" \" ` `` abc\ndef\"\"\"", tokens[0])
}

func (s *RawLexerSuite) TestRawMultiLineString(t *testing.T) {
	tokens := s.lex(t, `r"""abc" ""`, lr.ParseErrorToken)
	expectError(t, tokens[0], "string not terminated")

	tokens = s.lex(t, "r```mismatch\"\"\"", lr.ParseErrorToken)
	expectError(t, tokens[0], "string not terminated")

	tokens = s.lex(t, `r"""foo\400bar"""`, lr.StringLiteralToken)
	str := expectValue(t, `r"""foo\400bar"""`, tokens[0])
	expect.Equal(t, lexutil.RawMultiLineString, str.SubType)

	tokens = s.lex(
		t, "r```\"\"\" \"\" \" ` `` abc\ndef```",
		lr.StringLiteralToken)
	expectValue(t, "r```\"\"\" \"\" \" ` `` abc\ndef```", tokens[0])

	tokens = s.lex(
		t, "r\"\"\"``` \"\" \" ` `` abc\ndef\"\"\"",
		lr.StringLiteralToken)
	expectValue(t, "r\"\"\"``` \"\" \" ` `` abc\ndef\"\"\"", tokens[0])
}

func (s *RawLexerSuite) TestDotDecimalFloat(t *testing.T) {
	s.lex(t, ".", lr.DotToken)
	s.lex(t, ".a", lr.DotToken, lr.IdentifierToken)

	for i := 0; i < 10; i++ {
		value := fmt.Sprintf(".%d", i)
		tokens := s.lex(t, value, lr.FloatLiteralToken)
		float := expectValue(t, value, tokens[0])
		expect.Equal(t, lexutil.DecimalFloat, float.SubType)

		value = fmt.Sprintf(".01e%d", i)
		tokens = s.lex(t, value, lr.FloatLiteralToken)
		float = expectValue(t, value, tokens[0])
		expect.Equal(t, lexutil.DecimalFloat, float.SubType)

		value = fmt.Sprintf(".0_2E+0%d", i)
		tokens = s.lex(t, value, lr.FloatLiteralToken)
		float = expectValue(t, value, tokens[0])
		expect.Equal(t, lexutil.DecimalFloat, float.SubType)

		value = fmt.Sprintf(".0003e-0_0%d", i)
		tokens = s.lex(t, value, lr.FloatLiteralToken)
		float = expectValue(t, value, tokens[0])
		expect.Equal(t, lexutil.DecimalFloat, float.SubType)
	}

	tokens := s.lex(t, "._1", lr.DotToken, lr.IdentifierToken)
	expectValue(t, "_1", tokens[1])

	tokens = s.lex(t, ".1e", lr.FloatLiteralToken, lr.IdentifierToken)
	expectValue(t, ".1", tokens[0])
	expectValue(t, "e", tokens[1])

	tokens = s.lex(t, ".1e_1", lr.FloatLiteralToken, lr.IdentifierToken)
	expectValue(t, ".1", tokens[0])
	expectValue(t, "e_1", tokens[1])

	tokens = s.lex(
		t, ".1e+_1",
		lr.FloatLiteralToken, lr.IdentifierToken, lr.AddToken, lr.IdentifierToken)
	expectValue(t, ".1", tokens[0])
	expectValue(t, "e", tokens[1])
	expectValue(t, "_1", tokens[3])

	tokens = s.lex(t, ".1ea", lr.FloatLiteralToken, lr.IdentifierToken)
	expectValue(t, ".1", tokens[0])
	expectValue(t, "ea", tokens[1])

	tokens = s.lex(t, ".1p", lr.FloatLiteralToken, lr.IdentifierToken)
	expectValue(t, ".1", tokens[0])
	expectValue(t, "p", tokens[1])

	tokens = s.lex(t, ".1p5", lr.FloatLiteralToken, lr.IdentifierToken)
	expectValue(t, ".1", tokens[0])
	expectValue(t, "p5", tokens[1])

	tokens = s.lex(t, ".e10", lr.DotToken, lr.IdentifierToken)
	expectValue(t, "e10", tokens[1])

	s.lex(t, "+.123-", lr.AddToken, lr.FloatLiteralToken, lr.SubToken)
}

func (s *RawLexerSuite) TestIntPrefixedDecimalFloat(t *testing.T) {
	tokens := s.lex(
		t, "123e+",
		lr.IntegerLiteralToken, lr.IdentifierToken, lr.AddToken)
	expectValue(t, "123", tokens[0])
	expectValue(t, "e", tokens[1])

	tokens = s.lex(t, "123ea", lr.IntegerLiteralToken, lr.IdentifierToken)
	expectValue(t, "123", tokens[0])
	expectValue(t, "ea", tokens[1])

	tokens = s.lex(t, "123p0", lr.IntegerLiteralToken, lr.IdentifierToken)
	expectValue(t, "123", tokens[0])
	expectValue(t, "p0", tokens[1])

	for i := 0; i < 10; i++ {
		value := fmt.Sprintf("%d.", i)
		tokens = s.lex(t, value, lr.FloatLiteralToken)
		float := expectValue(t, value, tokens[0])
		expect.Equal(t, lexutil.DecimalFloat, float.SubType)

		value = fmt.Sprintf("1.%d", i)
		tokens = s.lex(t, value, lr.FloatLiteralToken)
		float = expectValue(t, value, tokens[0])
		expect.Equal(t, lexutil.DecimalFloat, float.SubType)

		value = fmt.Sprintf("2e%d", i)
		tokens = s.lex(t, value, lr.FloatLiteralToken)
		float = expectValue(t, value, tokens[0])
		expect.Equal(t, lexutil.DecimalFloat, float.SubType)

		value = fmt.Sprintf("3.E-%d", i)
		tokens = s.lex(t, value, lr.FloatLiteralToken)
		float = expectValue(t, value, tokens[0])
		expect.Equal(t, lexutil.DecimalFloat, float.SubType)

		value = fmt.Sprintf("4.0_%de+4_4", i)
		tokens = s.lex(t, value, lr.FloatLiteralToken)
		float = expectValue(t, value, tokens[0])
		expect.Equal(t, lexutil.DecimalFloat, float.SubType)
	}

	tokens = s.lex(t, "123.", lr.FloatLiteralToken)
	float := expectValue(t, "123.", tokens[0])
	expect.Equal(t, lexutil.DecimalFloat, float.SubType)

	tokens = s.lex(t, "123.456", lr.FloatLiteralToken)
	float = expectValue(t, "123.456", tokens[0])
	expect.Equal(t, lexutil.DecimalFloat, float.SubType)

	tokens = s.lex(t, "123e0", lr.FloatLiteralToken)
	float = expectValue(t, "123e0", tokens[0])
	expect.Equal(t, lexutil.DecimalFloat, float.SubType)

	tokens = s.lex(t, "123.456E789", lr.FloatLiteralToken)
	float = expectValue(t, "123.456E789", tokens[0])
	expect.Equal(t, lexutil.DecimalFloat, float.SubType)

	tokens = s.lex(t, "123.a", lr.FloatLiteralToken, lr.IdentifierToken)
	expectValue(t, "123.", tokens[0])
	expectValue(t, "a", tokens[1])

	tokens = s.lex(t, "123._0", lr.FloatLiteralToken, lr.IdentifierToken)
	expectValue(t, "123.", tokens[0])
	expectValue(t, "_0", tokens[1])

	tokens = s.lex(t, "123.0a", lr.FloatLiteralToken, lr.IdentifierToken)
	expectValue(t, "123.0", tokens[0])
	expectValue(t, "a", tokens[1])

	tokens = s.lex(t, "123.p0", lr.FloatLiteralToken, lr.IdentifierToken)
	expectValue(t, "123.", tokens[0])
	expectValue(t, "p0", tokens[1])
}

func (s *RawLexerSuite) TestHexaecimalFloat(t *testing.T) {
	tokens := s.lex(
		t, "0xabcp+",
		lr.IntegerLiteralToken, lr.IdentifierToken, lr.AddToken)
	expectValue(t, "0xabc", tokens[0])
	expectValue(t, "p", tokens[1])

	tokens = s.lex(t, "0xdefpg", lr.IntegerLiteralToken, lr.IdentifierToken)
	expectValue(t, "0xdef", tokens[0])
	expectValue(t, "pg", tokens[1])

	tokens = s.lex(
		t, "0x123e+0",
		lr.IntegerLiteralToken, lr.AddToken, lr.IntegerLiteralToken)
	expectValue(t, "0x123e", tokens[0])
	expectValue(t, "0", tokens[2])

	tokens = s.lex(
		t, "0x.gp0",
		lr.ParseErrorToken, lr.DotToken, lr.IdentifierToken)
	expectValue(t, "gp0", tokens[2])

	// decimal place leading digit can't be '_'
	tokens = s.lex(
		t, "0x._0p0",
		lr.ParseErrorToken, lr.DotToken, lr.IdentifierToken)
	expectValue(t, "_0p0", tokens[2])

	// missing exponent
	tokens = s.lex(t, "0x123.", lr.IntegerLiteralToken, lr.DotToken)
	expectValue(t, "0x123", tokens[0])

	/// missing exponent
	tokens = s.lex(t, "0x.123", lr.ParseErrorToken, lr.FloatLiteralToken)
	expectValue(t, ".123", tokens[1])

	/// missing exponent
	tokens = s.lex(
		t, "0x123.456",
		lr.IntegerLiteralToken, lr.FloatLiteralToken)
	expectValue(t, "0x123", tokens[0])
	expectValue(t, ".456", tokens[1])

	for i := 0; i < 10; i++ {
		value := fmt.Sprintf("0x%dpa", i)
		tokens = s.lex(t, value, lr.FloatLiteralToken)
		float := expectValue(t, value, tokens[0])
		expect.Equal(t, lexutil.HexadecimalFloat, float.SubType)

		value = fmt.Sprintf("0x_%dP+b", i)
		tokens = s.lex(t, value, lr.FloatLiteralToken)
		float = expectValue(t, value, tokens[0])
		expect.Equal(t, lexutil.HexadecimalFloat, float.SubType)

		value = fmt.Sprintf("0X%d.p-C", i)
		tokens = s.lex(t, value, lr.FloatLiteralToken)
		float = expectValue(t, value, tokens[0])
		expect.Equal(t, lexutil.HexadecimalFloat, float.SubType)

		value = fmt.Sprintf("0x%d.0_3p0_D", i)
		tokens = s.lex(t, value, lr.FloatLiteralToken)
		float = expectValue(t, value, tokens[0])
		expect.Equal(t, lexutil.HexadecimalFloat, float.SubType)

		value = fmt.Sprintf("0x.%dp0E", i)
		tokens = s.lex(t, value, lr.FloatLiteralToken)
		float = expectValue(t, value, tokens[0])
		expect.Equal(t, lexutil.HexadecimalFloat, float.SubType)
	}

	tokens = s.lex(t, "0x_0_123.ABC_789p456_DEF", lr.FloatLiteralToken)
	expectValue(t, "0x_0_123.ABC_789p456_DEF", tokens[0])
}
