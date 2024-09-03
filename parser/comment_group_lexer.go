package parser

import (
	"io"

	"github.com/pattyshack/gt/lexutil"
)

type CommentGroupLexer struct {
	*lexutil.BufferedReader[Token]

	base Lexer
}

func NewCommentGroupLexer(
	sourceFileName string,
	sourceContent io.Reader,
	options LexerOptions,
) Lexer {
	base := NewTrimSpacesLexer(sourceFileName, sourceContent, options)
	return &CommentGroupLexer{
		BufferedReader: lexutil.NewBufferedReader(
			lexutil.NewLexerReader[Token](base),
			10),
		base: base,
	}
}

func (lexer *CommentGroupLexer) CurrentLocation() Location {
	peeked, err := lexer.Peek(1)
	if err != nil || len(peeked) == 0 {
		return lexer.base.CurrentLocation()
	}

	return peeked[0].Loc()
}

func (lexer *CommentGroupLexer) Next() (Token, error) {
	token, err := readToken(lexer)
	if err != nil {
		return nil, err
	}

	if token.Id() == blockCommentToken {
		return CommentGroup{token.(ValueSymbol)}, nil
	} else if token.Id() != lineCommentToken {
		return token, nil
	}

	group := CommentGroup{token.(ValueSymbol)}
	for {
		peeked, err := lexer.Peek(2)
		if err != nil {
			break
		}

		if len(peeked) == 2 &&
			peeked[0].Id() == NewlinesToken &&
			peeked[1].Id() == lineCommentToken {

			newlines := peeked[0].(CountSymbol)
			if newlines.Count > 1 {
				break
			}

			group = append(group, peeked[1].(ValueSymbol))

			_, err := lexer.Discard(2)
			if err != nil {
				panic("should never happen")
			}
		} else {
			break
		}
	}

	return group, nil
}
