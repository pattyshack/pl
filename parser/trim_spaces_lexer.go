package parser

import (
	"io"

	"github.com/pattyshack/gt/lexutil"
)

type TrimSpacesLexer struct {
	*lexutil.BufferedReader[Token]
}

func NewTrimSpacesLexer(
	sourceFileName string,
	sourceContent io.Reader,
	options LexerOptions,
) TrimSpacesLexer {
	return TrimSpacesLexer{
		BufferedReader: lexutil.NewBufferedReader(
			lexutil.NewLexerReader[Token](
				NewRawLexer(sourceFileName, sourceContent, options)),
			100),
	}
}

func (lexer *TrimSpacesLexer) read() (Token, error) {
	peeked, err := lexer.Peek(1)
	if err != nil {
		return nil, err
	}

	if len(peeked) != 1 {
		panic("should never happen")
	}

	token := peeked[0]

	_, err = lexer.Discard(1)
	if err != nil {
		panic("should never happen")
	}

	return token, nil
}

func (lexer *TrimSpacesLexer) Next() (Token, error) {
	token, err := lexer.read()
	if err != nil {
		return nil, err
	}

	if token.Id() == spacesToken {
		token, err = lexer.read()
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
