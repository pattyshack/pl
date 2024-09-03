package parser

import (
	"bytes"
	"io"
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
	buffer := bytes.NewBufferString(input)

	lexer := NewTrimSpacesLexer(
		"source.txt",
		buffer,
		LexerOptions{
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

func (s *TrimSpacesLexerSuite) expectValue(
	t *testing.T,
	expectedValue string,
	token Token,
) ValueSymbol {
	value, ok := token.(ValueSymbol)
	expect.True(t, ok)
	expect.Equal(t, expectedValue, value.Value)
	return value
}

func (s *TrimSpacesLexerSuite) expectCount(
	t *testing.T,
	expectedCount int,
	token Token,
) {
	count, ok := token.(CountSymbol)
	expect.True(t, ok)
	expect.Equal(t, expectedCount, count.Count)
}

func (s *TrimSpacesLexerSuite) TestTrimLeadingSpaces(t *testing.T) {
	tokens := s.lex(
		t, " a  b   \n\n c",
		IdentifierToken, IdentifierToken, NewlinesToken, IdentifierToken)
	s.expectValue(t, "a", tokens[0])
	s.expectValue(t, "b", tokens[1])
	s.expectCount(t, 2, tokens[2])
	s.expectValue(t, "c", tokens[3])
}

func (s *TrimSpacesLexerSuite) TestMergeNewlines(t *testing.T) {
	tokens := s.lex(
		t, " a// asdf \n  \n  b   \n\n   \n\n\n \t\n c // comment\n ",
		IdentifierToken, lineCommentToken, NewlinesToken, IdentifierToken,
		NewlinesToken, IdentifierToken, lineCommentToken, NewlinesToken)
	s.expectValue(t, "a", tokens[0])
	s.expectValue(t, "// asdf ", tokens[1])
	s.expectCount(t, 2, tokens[2])
	s.expectValue(t, "b", tokens[3])
	s.expectCount(t, 6, tokens[4])
	s.expectValue(t, "c", tokens[5])
	s.expectValue(t, "// comment", tokens[6])
	s.expectCount(t, 1, tokens[7])
}
