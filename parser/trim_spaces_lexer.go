package parser

import (
	"io"

	"github.com/pattyshack/gt/lexutil"
)

type TokenPeekDiscarder interface {
	Peek(int) ([]Token, error)
	Discard(int) (int, error)
}

func readToken(reader TokenPeekDiscarder) (Token, error) {
	peeked, err := reader.Peek(1)
	if err != nil {
		return nil, err
	}

	if len(peeked) != 1 {
		panic("should never happen")
	}

	token := peeked[0]

	_, err = reader.Discard(1)
	if err != nil {
		panic("should never happen")
	}

	return token, nil
}

type TrimSpacesLexer struct {
	*lexutil.BufferedReader[Token]
	base Lexer
}

func NewTrimSpacesLexer(
	sourceFileName string,
	sourceContent io.Reader,
	options LexerOptions,
) Lexer {
	base := NewRawLexer(sourceFileName, sourceContent, options)
	return &TrimSpacesLexer{
		BufferedReader: lexutil.NewBufferedReader(
			lexutil.NewLexerReader[Token](base),
			10),
		base: base,
	}
}

func (lexer *TrimSpacesLexer) CurrentLocation() Location {
	peeked, err := lexer.Peek(1)
	if err != nil || len(peeked) == 0 {
		return (lexer.base.CurrentLocation())
	}

	return peeked[0].Loc()
}

func (lexer *TrimSpacesLexer) Next() (Token, error) {
	token, err := readToken(lexer)
	if err != nil {
		return nil, err
	}

	if token.Id() == spacesToken {
		token, err = readToken(lexer)
		if err != nil {
			return nil, err
		}
	}

	if token.Id() != NewlinesToken {
		return token, nil
	}

	newlines := token.(CountSymbol)
	for {
		peeked, err := lexer.Peek(2)
		if err != nil {
			break
		}

		if len(peeked) == 2 &&
			peeked[0].Id() == spacesToken &&
			peeked[1].Id() == NewlinesToken {

			other := peeked[1].(CountSymbol)
			newlines.Count += other.Count
			newlines.EndPos = other.EndPos

			_, err = lexer.Discard(2)
			if err != nil {
				panic("should never happen")
			}
		} else {
			break
		}
	}

	return newlines, nil
}
