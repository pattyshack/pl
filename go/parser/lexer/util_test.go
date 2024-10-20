package lexer

import (
	"bytes"
	"io"
	"testing"

	"github.com/pattyshack/gt/testing/expect"

	"github.com/pattyshack/pl/errors"
	"github.com/pattyshack/pl/parser/lr"
)

func expectError(t *testing.T, token lr.Token, errMsg string) {
	pe, ok := token.(*lr.ParseErrorSymbol)
	expect.True(t, ok)
	expect.Error(t, pe.Error, errMsg)
}

func expectValue(
	t *testing.T,
	expectedValue string,
	token lr.Token,
) *lr.TokenValue {
	value, ok := token.(*lr.TokenValue)
	expect.True(t, ok)
	expect.Equal(t, expectedValue, value.Value)
	return value
}

func expectCount(
	t *testing.T,
	expectedCount int,
	token lr.Token,
) {
	count, ok := token.(lr.TokenCount)
	expect.True(t, ok)
	expect.Equal(t, expectedCount, count.Count)
}

func lex[T lr.Lexer](
	t *testing.T,
	newLexer func(string, io.Reader, *errors.Emitter, LexerOptions) T,
	input string,
	expected ...lr.SymbolId,
) []lr.Token {
	buffer := bytes.NewBufferString(input)

	lexer := newLexer(
		"source.txt",
		buffer,
		&errors.Emitter{},
		LexerOptions{
			PreserveCommentContent: true,
			initialPeekWindowSize:  1,
		})

	tokens := []lr.Token{}
	tokenIds := []lr.SymbolId{}
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
