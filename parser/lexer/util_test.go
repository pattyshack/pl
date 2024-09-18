package lexer

import (
	"bytes"
	"io"
	"testing"

	"github.com/pattyshack/gt/testing/expect"

	. "github.com/pattyshack/pl/parser/lr"
)

func expectError(t *testing.T, token Token, errMsg string) {
	pe, ok := token.(*ParseErrorSymbol)
	expect.True(t, ok)
	expect.Error(t, pe.Errors[0], errMsg)
}

func expectValue(
	t *testing.T,
	expectedValue string,
	token Token,
) *TokenValue {
	value, ok := token.(*TokenValue)
	expect.True(t, ok)
	expect.Equal(t, expectedValue, value.Value)
	return value
}

func expectCount(
	t *testing.T,
	expectedCount int,
	token Token,
) {
	count, ok := token.(TokenCount)
	expect.True(t, ok)
	expect.Equal(t, expectedCount, count.Count)
}

func lex(
	t *testing.T,
	newLexer func(string, io.Reader, LexerOptions) Lexer,
	input string,
	expected ...SymbolId,
) []Token {
	buffer := bytes.NewBufferString(input)

	lexer := newLexer(
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

func testIdxChar(idx int) string {
	charSet := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789_"
	return string(charSet[idx%len(charSet)])
}
