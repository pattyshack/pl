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

// Discard spacesToken and merge adjacent NewlinesTokens
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

// Convert "adjacent" comments into comment groups.  blockComment is never
// adjacent to any other comment.  lineComment is adjacent to other
// lineComments if they are separated by a single newline.
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

// Associate comment groups to real tokens whenever possible.
type AssociateCommentGroupsLexer struct {
	*lexutil.BufferedReader[Token]
	base Lexer
}

func NewAssociateCommentGroupsLexer(
	sourceFileName string,
	sourceContent io.Reader,
	options LexerOptions,
) Lexer {
	base := NewCommentGroupLexer(sourceFileName, sourceContent, options)
	return &AssociateCommentGroupsLexer{
		BufferedReader: lexutil.NewBufferedReader(
			lexutil.NewLexerReader[Token](base),
			10),
		base: base,
	}
}

func (lexer *AssociateCommentGroupsLexer) CurrentLocation() Location {
	peeked, err := lexer.Peek(1)
	if err != nil || len(peeked) == 0 {
		return lexer.base.CurrentLocation()
	}

	return peeked[0].Loc()
}

func (lexer *AssociateCommentGroupsLexer) Next() (Token, error) {
	token, err := readToken(lexer)
	if err != nil {
		return nil, err
	}

	var value ValueSymbol
	// Handle Leading comment groups
	if token.Id() == commentGroupToken {
		groups := CommentGroups{token.(CommentGroup)}

		found := false
		for !found {
			peeked, err := lexer.Peek(1)
			if len(peeked) == 0 || err != nil {
				// groups not associated with any real token
				return groups, nil
			}

			token = peeked[0]

			_, err = lexer.Discard(1)
			if err != nil {
				panic("should never happen")
			}

			switch token.Id() {
			case NewlinesToken:
				// drop the newlines token
			case commentGroupToken:
				groups = append(groups, token.(CommentGroup))
			default:
				// Found a real token
				found = true
				value = token.(ValueSymbol)
				value.LeadingComments = groups
			}
		}
	} else if token.Id() == NewlinesToken {
		return token, nil
	} else {
		value = token.(ValueSymbol)
	}

	// extract trailing comment groups
	for {
		peeked, err := lexer.Peek(1)
		if len(peeked) == 0 || err != nil {
			break
		}

		if peeked[0].Id() != commentGroupToken {
			break
		}

		value.TrailingComments = append(
			value.TrailingComments,
			peeked[0].(CommentGroup))

		_, err = lexer.Discard(1)
		if err != nil {
			panic("should never happen")
		}
	}

	return value, nil
}
