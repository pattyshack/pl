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
		comments = append(comments, value.Val())
	}

	expect.Equal(t, expectedComments, comments)
}

func expectCommentGroups(
	t *testing.T,
	groups CommentGroups,
	expected ...[]string,
) {
	commentGroups := [][]string{}
	for _, group := range groups.Groups {
		comments := []string{}
		for _, value := range group {
			comments = append(comments, value.Val())
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
	expect.Nil(t, a.LeadingComment.Groups)
	expectCommentGroups(
		t,
		a.TrailingComment,
		[]string{"/*a1*/"},
		[]string{"//a2l1", "//a2l2"})

	b := expectValue(t, "b", tokens[2])
	expectCommentGroups(
		t,
		b.LeadingComment,
		[]string{"/*b1*/"},
		[]string{"//b2"},
		[]string{"//b3l1", "//b3l2"},
		[]string{"/*b4*/"})
	expectCommentGroups(
		t,
		b.TrailingComment,
		[]string{"/*b5*/"},
		[]string{"/*b6*/"},
		[]string{"/*b7*/"})

	c := expectValue(t, "c", tokens[4])
	expectCommentGroups(
		t,
		c.LeadingComment,
		[]string{"//c1"})
	expect.Nil(t, c.TrailingComment.Groups)

	floating, ok := tokens[6].(CommentGroups)
	expect.True(t, ok)
	expectCommentGroups(
		t,
		floating,
		[]string{"//?1l1", "//?1l2"},
		[]string{"//?2l1", "//?2l2", "//?2l3"})
}

type TerminalNewlinesLexerSuite struct {
}

func TestTerminalNewlinesLexer(t *testing.T) {
	suite.RunTests(t, &TerminalNewlinesLexerSuite{})
}

func (s *TerminalNewlinesLexerSuite) lex(
	t *testing.T,
	input string,
	expected ...SymbolId,
) []Token {
	return lex(t, NewLexer, input, expected...)
}

func (s *TerminalNewlinesLexerSuite) TestEOF(t *testing.T) {
	s.lex(t, "+", AddToken)
	s.lex(t, "+\n", AddToken)
	s.lex(t, ",", CommaToken)
	s.lex(t, ",\n", CommaToken)
	s.lex(t, "// comment", CommentGroupsToken)
	s.lex(t, "// comment\n", CommentGroupsToken)
	s.lex(t, "a", IdentifierToken)
	s.lex(t, "a\n", IdentifierToken, NewlinesToken)
	s.lex(t, "a // comment \n", IdentifierToken, NewlinesToken)
}

func (s *TerminalNewlinesLexerSuite) TestIdentifier(t *testing.T) {
	s.lex(
		t, "a // foo \n+ b",
		IdentifierToken, NewlinesToken, AddToken, IdentifierToken)
	s.lex(
		t, "a + b\n c",
		IdentifierToken, AddToken, IdentifierToken, NewlinesToken, IdentifierToken)
	s.lex(
		t, "a + _\n c",
		IdentifierToken, AddToken, UnderscoreToken, NewlinesToken, IdentifierToken)
}

func (s *TerminalNewlinesLexerSuite) TestLiteral(t *testing.T) {
	s.lex(
		t, "1 \n + b",
		IntegerLiteralToken, NewlinesToken, AddToken, IdentifierToken)
	s.lex(
		t, "a + 0x123\n c",
		IdentifierToken, AddToken, IntegerLiteralToken, NewlinesToken,
		IdentifierToken)
	s.lex(
		t, "1. \n + b",
		FloatLiteralToken, NewlinesToken, AddToken, IdentifierToken)
	s.lex(
		t, "'a' \n + b",
		RuneLiteralToken, NewlinesToken, AddToken, IdentifierToken)
	s.lex(
		t, "\"string\" \n + b",
		StringLiteralToken, NewlinesToken, AddToken, IdentifierToken)
	s.lex(
		t, "true \n + b",
		TrueToken, NewlinesToken, AddToken, IdentifierToken)
	s.lex(
		t, "false \n + b",
		FalseToken, NewlinesToken, AddToken, IdentifierToken)
}

func (s *TerminalNewlinesLexerSuite) TestJumps(t *testing.T) {
	s.lex(t, "return @label\n}",
		ReturnToken, JumpLabelToken, NewlinesToken, RbraceToken)
	s.lex(t, "return @label // comment \n}",
		ReturnToken, JumpLabelToken, NewlinesToken, RbraceToken)
	s.lex(t, "return\n}", ReturnToken, NewlinesToken, RbraceToken)
	s.lex(t, "break\n}", BreakToken, NewlinesToken, RbraceToken)
	s.lex(t, "continue\n}", ContinueToken, NewlinesToken, RbraceToken)
	s.lex(t, "fallthrough\n}", FallthroughToken, NewlinesToken, RbraceToken)
}

func (s *TerminalNewlinesLexerSuite) TestSuffixUnary(t *testing.T) {
	s.lex(t, "a++\n}",
		IdentifierToken, AddOneAssignToken, NewlinesToken, RbraceToken)
	s.lex(t, "b--\n}",
		IdentifierToken, SubOneAssignToken, NewlinesToken, RbraceToken)
}

func (s *TerminalNewlinesLexerSuite) TestEndScope(t *testing.T) {
	s.lex(t, "}\na", RbraceToken, NewlinesToken, IdentifierToken)
	s.lex(t, ")\na", RparenToken, NewlinesToken, IdentifierToken)
	s.lex(t, "]\na", RbracketToken, NewlinesToken, IdentifierToken)
}

func (s *TerminalNewlinesLexerSuite) TestNonTerminalNewlines(t *testing.T) {
	s.lex(t, ";\na", SemicolonToken, IdentifierToken)
	s.lex(t, ".\na", DotToken, IdentifierToken)
	s.lex(t, ",\na", CommaToken, IdentifierToken)
	s.lex(t, "=\na", AssignToken, IdentifierToken)
	s.lex(t, "+=\na", AddAssignToken, IdentifierToken)
	s.lex(t, "<\na", LessToken, IdentifierToken)
	s.lex(t, "+\na", AddToken, IdentifierToken)
	s.lex(t, "-\na", SubToken, IdentifierToken)
	s.lex(t, "*\na", MulToken, IdentifierToken)
	s.lex(t, "/\na", DivToken, IdentifierToken)
	s.lex(t, "%\na", ModToken, IdentifierToken)
	s.lex(t, "and\na", AndToken, IdentifierToken)
	s.lex(t, "or\na", OrToken, IdentifierToken)
	s.lex(t, "not\na", NotToken, IdentifierToken)
	s.lex(t, "async\na", AsyncToken, IdentifierToken)
	s.lex(t, "defer\na", DeferToken, IdentifierToken)
	s.lex(t, "func\na", FuncToken, IdentifierToken)
	s.lex(t, "type\na", TypeToken, IdentifierToken)
	s.lex(t, "var\na", VarToken, IdentifierToken)
	s.lex(t, "let\na", LetToken, IdentifierToken)
}
