package parser

import (
	"bytes"
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
		RawLexerOptions{})

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

func (s *RawLexerSuite) TestIdentifierToken(t *testing.T) {
	// TODO
}

func (s *RawLexerSuite) TestKeywordTokens(t *testing.T) {
	// TODO
}

func (s *RawLexerSuite) TestIntegerLiteralToken(t *testing.T) {
	// TODO (reminder: need to check for dot prefix variant .123)
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
	// TODO (also need to check lex error for '\r' without '\n')
}

func (s *RawLexerSuite) TestSpacesTokens(t *testing.T) {
	// TODO
}

func (s *RawLexerSuite) TestAtTokens(t *testing.T) {
	// TODO lex error for @ not followed by identifier
	// TODO jump label
}

func (s *RawLexerSuite) TestLbraceTokens(t *testing.T) {
	s.lex(t, "+{-", AddToken, LbraceToken, SubToken)
}

func (s *RawLexerSuite) TestRbraceTokens(t *testing.T) {
	s.lex(t, "+}-", AddToken, RbraceToken, SubToken)
}

func (s *RawLexerSuite) TestLparenTokens(t *testing.T) {
	s.lex(t, "+(-", AddToken, LparenToken, SubToken)
}

func (s *RawLexerSuite) TestRparenTokens(t *testing.T) {
	s.lex(t, "+)-", AddToken, RparenToken, SubToken)
}

func (s *RawLexerSuite) TestLbracketTokens(t *testing.T) {
	s.lex(t, "+[-", AddToken, LbracketToken, SubToken)
}

func (s *RawLexerSuite) TestRbracketTokens(t *testing.T) {
	s.lex(t, "+]-", AddToken, RbracketToken, SubToken)
}

func (s *RawLexerSuite) TestDotTokens(t *testing.T) {
	s.lex(t, "+.-", AddToken, DotToken, SubToken)
	s.lex(t, "+...-", AddToken, DotDotDotToken, SubToken)
}

func (s *RawLexerSuite) TestCommaTokens(t *testing.T) {
	s.lex(t, "+,-", AddToken, CommaToken, SubToken)
}

func (s *RawLexerSuite) TestQuestionTokens(t *testing.T) {
	s.lex(t, "+?-", AddToken, QuestionToken, SubToken)
}

func (s *RawLexerSuite) TestSemicolonTokens(t *testing.T) {
	s.lex(t, "+;-", AddToken, SemicolonToken, SubToken)
}

func (s *RawLexerSuite) TestColonTokens(t *testing.T) {
	s.lex(t, "+:-", AddToken, ColonToken, SubToken)
}

func (s *RawLexerSuite) TestDollarTokens(t *testing.T) {
	s.lex(t, "+$[-", AddToken, DollarLbracketToken, SubToken)

	tokens := s.lex(t, "+$-", AddToken, ParseErrorToken, SubToken)
	parseError, ok := tokens[1].(ParseError)

	expect.True(t, ok)

	expect.Error(t, parseError.Error, "unexpected utf8 rune")
}

func (s *RawLexerSuite) TestAddTokens(t *testing.T) {
	s.lex(t, "*+/", MulToken, AddToken, DivToken)
	s.lex(t, "*++/", MulToken, AddOneAssignToken, DivToken)
	s.lex(t, "*+=/", MulToken, AddAssignToken, DivToken)
}

func (s *RawLexerSuite) TestSubTokens(t *testing.T) {
	s.lex(t, "*-/", MulToken, SubToken, DivToken)
	s.lex(t, "*--/", MulToken, SubOneAssignToken, DivToken)
	s.lex(t, "*-=/", MulToken, SubAssignToken, DivToken)
}

func (s *RawLexerSuite) TestMulTokens(t *testing.T) {
	s.lex(t, "+*-", AddToken, MulToken, SubToken)
	s.lex(t, "+*=-", AddToken, MulAssignToken, SubToken)
}

func (s *RawLexerSuite) TestDivTokens(t *testing.T) {
	// TODO line and block comments
	s.lex(t, "+/-", AddToken, DivToken, SubToken)
	s.lex(t, "+/=-", AddToken, DivAssignToken, SubToken)
}

func (s *RawLexerSuite) TestModTokens(t *testing.T) {
	s.lex(t, "+%-", AddToken, ModToken, SubToken)
	s.lex(t, "+%=-", AddToken, ModAssignToken, SubToken)
}

func (s *RawLexerSuite) TestBitNegTokens(t *testing.T) {
	s.lex(t, "+~-", AddToken, BitNegToken, SubToken)
	s.lex(t, "+~=-", AddToken, BitNegAssignToken, SubToken)
	s.lex(t, "+~~-", AddToken, TildeTildeToken, SubToken)
}

func (s *RawLexerSuite) TestBitAndTokens(t *testing.T) {
	s.lex(t, "+&-", AddToken, BitAndToken, SubToken)
	s.lex(t, "+&=-", AddToken, BitAndAssignToken, SubToken)
}

func (s *RawLexerSuite) TestBitXorTokens(t *testing.T) {
	s.lex(t, "+^-", AddToken, BitXorToken, SubToken)
	s.lex(t, "+^=-", AddToken, BitXorAssignToken, SubToken)
}

func (s *RawLexerSuite) TestBitOrTokens(t *testing.T) {
	s.lex(t, "+|-", AddToken, BitOrToken, SubToken)
	s.lex(t, "+|=-", AddToken, BitOrAssignToken, SubToken)
}

func (s *RawLexerSuite) TestEqualTokens(t *testing.T) {
	s.lex(t, "(=)", LparenToken, AssignToken, RparenToken)
	s.lex(t, "(==)", LparenToken, EqualToken, RparenToken)
}

func (s *RawLexerSuite) TestExclaimTokens(t *testing.T) {
	s.lex(t, "+!-", AddToken, ExclaimToken, SubToken)
	s.lex(t, "+!=-", AddToken, NotEqualToken, SubToken)
}

func (s *RawLexerSuite) TestGreaterTokens(t *testing.T) {
	s.lex(t, "+>-", AddToken, GreaterToken, SubToken)
	s.lex(t, "+>=-", AddToken, GreaterOrEqualToken, SubToken)
	s.lex(t, "+>>-", AddToken, BitRshiftToken, SubToken)
	s.lex(t, "+>>=-", AddToken, BitRshiftAssignToken, SubToken)
}

func (s *RawLexerSuite) TestLessTokens(t *testing.T) {
	s.lex(t, "+<-", AddToken, LessToken, SubToken)
	s.lex(t, "+<=-", AddToken, LessOrEqualToken, SubToken)
	s.lex(t, "+<<-", AddToken, BitLshiftToken, SubToken)
	s.lex(t, "+<<=-", AddToken, BitLshiftAssignToken, SubToken)
}
