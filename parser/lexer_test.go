package parser

import (
	"testing"

	"github.com/pattyshack/gt/testing/expect"
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

func expectCommentGroup(
	t *testing.T,
	token Token,
	expectedComments ...string,
) {
	group, ok := token.(CommentGroup)
	expect.True(t, ok)

	comments := []string{}
	for _, value := range group {
		comments = append(comments, value.Value)
	}

	expect.Equal(t, expectedComments, comments)
}

type CommentGroupLexerSuite struct {
}

func TestCommentGroupLexer(t *testing.T) {
	suite.RunTests(t, &CommentGroupLexerSuite{})
}

func (s *CommentGroupLexerSuite) lex(
	t *testing.T,
	input string,
	expected ...SymbolId,
) []Token {
	return lex(t, NewCommentGroupLexer, input, expected...)
}

func (s *CommentGroupLexerSuite) TestBlockCommentGroups(t *testing.T) {
	tokens := s.lex(
		t, " a  /* group1 */ /* group2 */ b /* group3 */",
		IdentifierToken, commentGroupToken, commentGroupToken,
		IdentifierToken, commentGroupToken)
	expectValue(t, "a", tokens[0])
	expectCommentGroup(t, tokens[1], "/* group1 */")
	expectCommentGroup(t, tokens[2], "/* group2 */")
	expectValue(t, "b", tokens[3])
	expectCommentGroup(t, tokens[4], "/* group3 */")
}

func (s *CommentGroupLexerSuite) TestLineGroups(t *testing.T) {
	tokens := s.lex(
		t,
		" a //g1l1\n  //g1l2\n //g1l3\n\n //g2l1\n"+
			"\t//g2l2\n /*g3*/ //g4l1\n//g4l2\nb",
		IdentifierToken, commentGroupToken, NewlinesToken, commentGroupToken,
		NewlinesToken, commentGroupToken, commentGroupToken,
		NewlinesToken, IdentifierToken)
	expectValue(t, "a", tokens[0])
	expectCommentGroup(t, tokens[1], "//g1l1", "//g1l2", "//g1l3")
	expectCount(t, 2, tokens[2])
	expectCommentGroup(t, tokens[3], "//g2l1", "//g2l2")
	expectCount(t, 1, tokens[4])
	expectCommentGroup(t, tokens[5], "/*g3*/")
	expectCommentGroup(t, tokens[6], "//g4l1", "//g4l2")
	expectCount(t, 1, tokens[7])
	expectValue(t, "b", tokens[8])
}
