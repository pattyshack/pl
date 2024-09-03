package parser

import (
	"testing"

	"github.com/pattyshack/gt/testing/suite"
)

type TrimSpacesLexerSuite struct {
}

func TestTrimSpacesLexer(t *testing.T) {
	suite.RunTests(t, &TrimSpacesLexerSuite{})
}

func (s *TrimSpacesLexerSuite) lex(
	t *testing.T,
	input string,
	expected ...SymbolId,
) []Token {
	return lex(t, NewTrimSpacesLexer, input, expected...)
}

func (s *TrimSpacesLexerSuite) TestTrimLeadingSpaces(t *testing.T) {
	tokens := s.lex(
		t, " a  b   \n\n c",
		IdentifierToken, IdentifierToken, NewlinesToken, IdentifierToken)
	expectValue(t, "a", tokens[0])
	expectValue(t, "b", tokens[1])
	expectCount(t, 2, tokens[2])
	expectValue(t, "c", tokens[3])
}

func (s *TrimSpacesLexerSuite) TestMergeNewlines(t *testing.T) {
	tokens := s.lex(
		t, " a// asdf \n  \n  b   \n\n   \n\n\n \t\n c // comment\n ",
		IdentifierToken, lineCommentToken, NewlinesToken, IdentifierToken,
		NewlinesToken, IdentifierToken, lineCommentToken, NewlinesToken)
	expectValue(t, "a", tokens[0])
	expectValue(t, "// asdf ", tokens[1])
	expectCount(t, 2, tokens[2])
	expectValue(t, "b", tokens[3])
	expectCount(t, 6, tokens[4])
	expectValue(t, "c", tokens[5])
	expectValue(t, "// comment", tokens[6])
	expectCount(t, 1, tokens[7])
}
