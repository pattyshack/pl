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

func expectCommentGroups(
	t *testing.T,
	groups CommentGroups,
	expected ...[]string,
) {
	commentGroups := [][]string{}
	for _, group := range groups {
		comments := []string{}
		for _, value := range group {
			comments = append(comments, value.Value)
		}

		commentGroups = append(commentGroups, comments)
	}

	expect.Equal(t, expected, commentGroups)
}

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

type AssociateCommentGroupsLexerSuite struct {
}

func TestAssociateCommentGroupsLexer(t *testing.T) {
	suite.RunTests(t, &AssociateCommentGroupsLexerSuite{})
}

func (s *AssociateCommentGroupsLexerSuite) lex(
	t *testing.T,
	input string,
	expected ...SymbolId,
) []Token {
	return lex(t, NewAssociateCommentGroupsLexer, input, expected...)
}

func (s *AssociateCommentGroupsLexerSuite) TestAssociation(t *testing.T) {
	tokens := lex(
		t,
		NewAssociateCommentGroupsLexer,
		`a /*a1*/ //a2l1
    //a2l2

    /*b1*/
    //b2

    //b3l1
    //b3l2

    /*b4*/ b /*b5*/ /*b6*/ /*b7*/

    //c1
    c

    //?1l1
    //?1l2

    //?2l1
    //?2l2
    //?2l3
    `,
		IdentifierToken, NewlinesToken,
		IdentifierToken, NewlinesToken,
		IdentifierToken, NewlinesToken,
		CommentGroupsToken)

	a := expectValue(t, "a", tokens[0])
	expect.Nil(t, a.LeadingComments)
	expectCommentGroups(
		t,
		a.TrailingComments,
		[]string{"/*a1*/"},
		[]string{"//a2l1", "//a2l2"})

	b := expectValue(t, "b", tokens[2])
	expectCommentGroups(
		t,
		b.LeadingComments,
		[]string{"/*b1*/"},
		[]string{"//b2"},
		[]string{"//b3l1", "//b3l2"},
		[]string{"/*b4*/"})
	expectCommentGroups(
		t,
		b.TrailingComments,
		[]string{"/*b5*/"},
		[]string{"/*b6*/"},
		[]string{"/*b7*/"})

	c := expectValue(t, "c", tokens[4])
	expectCommentGroups(
		t,
		c.LeadingComments,
		[]string{"//c1"})
	expect.Nil(t, c.TrailingComments)

	floating, ok := tokens[6].(CommentGroups)
	expect.True(t, ok)
	expectCommentGroups(
		t,
		floating,
		[]string{"//?1l1", "//?1l2"},
		[]string{"//?2l1", "//?2l2", "//?2l3"})
}
