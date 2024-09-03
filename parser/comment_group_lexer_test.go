package parser

import (
	"testing"

	"github.com/pattyshack/gt/testing/expect"
	"github.com/pattyshack/gt/testing/suite"
)

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
		IdentifierToken, CommentGroupToken, CommentGroupToken,
		IdentifierToken, CommentGroupToken)
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
		IdentifierToken, CommentGroupToken, NewlinesToken, CommentGroupToken,
		NewlinesToken, CommentGroupToken, CommentGroupToken,
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
