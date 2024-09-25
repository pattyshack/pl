package lexer

import (
	"testing"

	"github.com/pattyshack/gt/testing/expect"
	"github.com/pattyshack/gt/testing/suite"

	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

func expectCommentGroup(
	t *testing.T,
	token lr.Token,
	expectedComments ...string,
) {
	group, ok := token.(lr.CommentGroupToken)
	expect.True(t, ok)
	expect.Equal(t, expectedComments, group.Comments)
}

func expectCommentGroups(
	t *testing.T,
	groups ast.CommentGroups,
	expected ...[]string,
) {
	commentGroups := [][]string{}
	for _, group := range groups.Groups {
		commentGroups = append(commentGroups, group.Comments)
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
	expected ...lr.SymbolId,
) []lr.Token {
	return lex(t, NewTrimSpacesLexer, input, expected...)
}

func (s *TrimSpacesLexerSuite) TestTrimLeadingSpaces(t *testing.T) {
	tokens := s.lex(
		t, " a  b   \n\n c",
		lr.IdentifierToken, lr.IdentifierToken, lr.NewlinesToken,
		lr.IdentifierToken)
	expectValue(t, "a", tokens[0])
	expectValue(t, "b", tokens[1])
	expectCount(t, 2, tokens[2])
	expectValue(t, "c", tokens[3])
}

func (s *TrimSpacesLexerSuite) TestMergeNewlines(t *testing.T) {
	tokens := s.lex(
		t, " a// asdf \n  \n  b   \n\n   \n\n\n \t\n c // comment\n ",
		lr.IdentifierToken, lineCommentToken, lr.NewlinesToken, lr.IdentifierToken,
		lr.NewlinesToken, lr.IdentifierToken, lineCommentToken, lr.NewlinesToken)
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
	expected ...lr.SymbolId,
) []lr.Token {
	return lex(t, NewCommentGroupLexer, input, expected...)
}

func (s *CommentGroupLexerSuite) TestBlockCommentGroups(t *testing.T) {
	tokens := s.lex(
		t, " a  /* group1 */ /* group2 */ b /* group3 */",
		lr.IdentifierToken, lr.CommentGroupTokenId, lr.CommentGroupTokenId,
		lr.IdentifierToken, lr.CommentGroupTokenId)
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
		lr.IdentifierToken, lr.CommentGroupTokenId,
		lr.NewlinesToken, lr.CommentGroupTokenId,
		lr.NewlinesToken, lr.CommentGroupTokenId, lr.CommentGroupTokenId,
		lr.NewlinesToken, lr.IdentifierToken)
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
	expected ...lr.SymbolId,
) []lr.Token {
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
		lr.IdentifierToken, lr.NewlinesToken,
		lr.IdentifierToken, lr.NewlinesToken,
		lr.IdentifierToken, lr.NewlinesToken,
		lr.CommentGroupsToken)

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

	floating, ok := tokens[6].(lr.CommentGroupsTok)
	expect.True(t, ok)
	expectCommentGroups(
		t,
		floating.CommentGroups,
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
	expected ...lr.SymbolId,
) []lr.Token {
	return lex(t, NewBasicLexer, input, expected...)
}

func (s *TerminalNewlinesLexerSuite) TestEOF(t *testing.T) {
	s.lex(t, "+", lr.AddToken)
	s.lex(t, "+\n", lr.AddToken)
	s.lex(t, ",", lr.CommaToken)
	s.lex(t, ",\n", lr.CommaToken)
	s.lex(t, "// comment", lr.CommentGroupsToken)
	s.lex(t, "// comment\n", lr.CommentGroupsToken)
	s.lex(t, "a", lr.IdentifierToken)
	s.lex(t, "a\n", lr.IdentifierToken, lr.NewlinesToken)
	s.lex(t, "a // comment \n", lr.IdentifierToken, lr.NewlinesToken)
}

func (s *TerminalNewlinesLexerSuite) TestIdentifier(t *testing.T) {
	s.lex(
		t, "a // foo \n+ b",
		lr.IdentifierToken, lr.NewlinesToken, lr.AddToken, lr.IdentifierToken)
	s.lex(
		t, "a + b\n c",
		lr.IdentifierToken, lr.AddToken, lr.IdentifierToken, lr.NewlinesToken,
		lr.IdentifierToken)
	s.lex(
		t, "a + _\n c",
		lr.IdentifierToken, lr.AddToken, lr.UnderscoreToken, lr.NewlinesToken,
		lr.IdentifierToken)
}

func (s *TerminalNewlinesLexerSuite) TestLiteral(t *testing.T) {
	s.lex(
		t, "1 \n + b",
		lr.IntegerLiteralToken, lr.NewlinesToken, lr.AddToken, lr.IdentifierToken)
	s.lex(
		t, "a + 0x123\n c",
		lr.IdentifierToken, lr.AddToken, lr.IntegerLiteralToken, lr.NewlinesToken,
		lr.IdentifierToken)
	s.lex(
		t, "1. \n + b",
		lr.FloatLiteralToken, lr.NewlinesToken, lr.AddToken, lr.IdentifierToken)
	s.lex(
		t, "'a' \n + b",
		lr.RuneLiteralToken, lr.NewlinesToken, lr.AddToken, lr.IdentifierToken)
	s.lex(
		t, "\"string\" \n + b",
		lr.StringLiteralToken, lr.NewlinesToken, lr.AddToken, lr.IdentifierToken)
	s.lex(
		t, "true \n + b",
		lr.TrueToken, lr.NewlinesToken, lr.AddToken, lr.IdentifierToken)
	s.lex(
		t, "false \n + b",
		lr.FalseToken, lr.NewlinesToken, lr.AddToken, lr.IdentifierToken)
}

func (s *TerminalNewlinesLexerSuite) TestJumps(t *testing.T) {
	s.lex(t, "return @label\n}",
		lr.ReturnToken, lr.JumpLabelToken, lr.NewlinesToken, lr.RbraceToken)
	s.lex(t, "return @label // comment \n}",
		lr.ReturnToken, lr.JumpLabelToken, lr.NewlinesToken, lr.RbraceToken)
	s.lex(t, "return\n}", lr.ReturnToken, lr.NewlinesToken, lr.RbraceToken)
	s.lex(t, "break\n}", lr.BreakToken, lr.NewlinesToken, lr.RbraceToken)
	s.lex(t, "continue\n}", lr.ContinueToken, lr.NewlinesToken, lr.RbraceToken)
	s.lex(
		t, "fallthrough\n}",
		lr.FallthroughToken, lr.NewlinesToken, lr.RbraceToken)
}

func (s *TerminalNewlinesLexerSuite) TestSuffixUnary(t *testing.T) {
	s.lex(t, "a++\n}",
		lr.IdentifierToken, lr.AddOneAssignToken, lr.NewlinesToken, lr.RbraceToken)
	s.lex(t, "b--\n}",
		lr.IdentifierToken, lr.SubOneAssignToken, lr.NewlinesToken, lr.RbraceToken)
}

func (s *TerminalNewlinesLexerSuite) TestEndScope(t *testing.T) {
	s.lex(t, "}\na", lr.RbraceToken, lr.NewlinesToken, lr.IdentifierToken)
	s.lex(t, ")\na", lr.RparenToken, lr.NewlinesToken, lr.IdentifierToken)
	s.lex(t, "]\na", lr.RbracketToken, lr.NewlinesToken, lr.IdentifierToken)
}

func (s *TerminalNewlinesLexerSuite) TestNonTerminalNewlines(t *testing.T) {
	s.lex(t, ";\na", lr.SemicolonToken, lr.IdentifierToken)
	s.lex(t, ".\na", lr.DotToken, lr.IdentifierToken)
	s.lex(t, ",\na", lr.CommaToken, lr.IdentifierToken)
	s.lex(t, "=\na", lr.AssignToken, lr.IdentifierToken)
	s.lex(t, "+=\na", lr.AddAssignToken, lr.IdentifierToken)
	s.lex(t, "<\na", lr.LessToken, lr.IdentifierToken)
	s.lex(t, "+\na", lr.AddToken, lr.IdentifierToken)
	s.lex(t, "-\na", lr.SubToken, lr.IdentifierToken)
	s.lex(t, "*\na", lr.MulToken, lr.IdentifierToken)
	s.lex(t, "/\na", lr.DivToken, lr.IdentifierToken)
	s.lex(t, "%\na", lr.ModToken, lr.IdentifierToken)
	s.lex(t, "and\na", lr.AndToken, lr.IdentifierToken)
	s.lex(t, "or\na", lr.OrToken, lr.IdentifierToken)
	s.lex(t, "not\na", lr.NotToken, lr.IdentifierToken)
	s.lex(t, "async\na", lr.AsyncToken, lr.IdentifierToken)
	s.lex(t, "defer\na", lr.DeferToken, lr.IdentifierToken)
	s.lex(t, "func\na", lr.FuncToken, lr.IdentifierToken)
	s.lex(t, "type\na", lr.TypeToken, lr.IdentifierToken)
	s.lex(t, "var\na", lr.VarToken, lr.IdentifierToken)
	s.lex(t, "let\na", lr.LetToken, lr.IdentifierToken)
}
