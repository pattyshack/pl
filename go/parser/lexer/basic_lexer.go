package lexer

import (
	"io"

	"github.com/pattyshack/gt/parseutil"

	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

func NewTrimSpacesLexer(
	sourceFileName string,
	sourceContent io.Reader,
	emitter *parseutil.Emitter,
	options LexerOptions,
) lr.Lexer {
	return parseutil.NewMergeTokenCountLexer(
		parseutil.NewTrimTokenLexer(
			NewRawLexer(sourceFileName, sourceContent, emitter, options),
			spacesToken),
		lr.NewlinesToken)
}

// Convert "adjacent" comments into comment groups.  blockComment is never
// adjacent to any other comment.  lineComment is adjacent to other
// lineComments if they are separated by a single newline.
type CommentGroupLexer struct {
	buffered *parseutil.BufferedReader[lr.Token]
	base     lr.Lexer
}

func NewCommentGroupLexer(
	sourceFileName string,
	sourceContent io.Reader,
	emitter *parseutil.Emitter,
	options LexerOptions,
) lr.Lexer {
	base := NewTrimSpacesLexer(sourceFileName, sourceContent, emitter, options)
	return &CommentGroupLexer{
		buffered: parseutil.NewBufferedReader(
			parseutil.NewLexerReader[lr.Token](base),
			10),
		base: base,
	}
}

func (lexer *CommentGroupLexer) CurrentLocation() parseutil.Location {
	peeked, err := lexer.buffered.Peek(1)
	if err != nil || len(peeked) == 0 {
		return lexer.base.CurrentLocation()
	}

	return peeked[0].Loc()
}

func (lexer *CommentGroupLexer) Next() (lr.Token, error) {
	token, err := lexer.buffered.Next()
	if err != nil {
		return nil, err
	}

	if token.Id() == blockCommentToken {
		return lr.CommentGroupToken{
			*ast.NewCommentGroup(token.(*lr.TokenValue)),
		}, nil
	} else if token.Id() != lineCommentToken {
		return token, nil
	}

	group := ast.NewCommentGroup(token.(*lr.TokenValue))
	for {
		peeked, err := lexer.buffered.Peek(2)
		if err != nil {
			break
		}

		if len(peeked) == 2 &&
			peeked[0].Id() == lr.NewlinesToken &&
			peeked[1].Id() == lineCommentToken {

			newlines := peeked[0].(*lr.TokenCount)
			if newlines.Count > 1 {
				break
			}

			group.Add(peeked[1].(*lr.TokenValue))

			_, err := lexer.buffered.Discard(2)
			if err != nil {
				panic("should never happen")
			}
		} else {
			break
		}
	}

	return lr.CommentGroupToken{*group}, nil
}

// Associate comment groups to real tokens whenever possible.
type AssociateCommentGroupsLexer struct {
	buffered *parseutil.BufferedReader[lr.Token]
	base     lr.Lexer
}

func NewAssociateCommentGroupsLexer(
	sourceFileName string,
	sourceContent io.Reader,
	emitter *parseutil.Emitter,
	options LexerOptions,
) lr.Lexer {
	base := NewCommentGroupLexer(sourceFileName, sourceContent, emitter, options)
	return &AssociateCommentGroupsLexer{
		buffered: parseutil.NewBufferedReader(
			parseutil.NewLexerReader[lr.Token](base),
			10),
		base: base,
	}
}

func (lexer *AssociateCommentGroupsLexer) CurrentLocation() parseutil.Location {
	peeked, err := lexer.buffered.Peek(1)
	if err != nil || len(peeked) == 0 {
		return lexer.base.CurrentLocation()
	}

	return peeked[0].Loc()
}

func (lexer *AssociateCommentGroupsLexer) Next() (lr.Token, error) {
	token, err := lexer.buffered.Next()
	if err != nil {
		return nil, err
	}

	var value *lr.TokenValue
	// Handle Leading comment groups
	if token.Id() == lr.CommentGroupTokenId {
		groups := ast.CommentGroups{
			Groups: []ast.CommentGroup{token.(lr.CommentGroupToken).CommentGroup},
		}

		found := false
		for !found {
			peeked, err := lexer.buffered.Peek(1)
			if len(peeked) == 0 || err != nil {
				// groups not associated with any real token
				return lr.CommentGroupsTok{groups}, nil
			}

			token = peeked[0]

			_, err = lexer.buffered.Discard(1)
			if err != nil {
				panic("should never happen")
			}

			switch token.Id() {
			case lr.NewlinesToken:
				// drop the newlines token
			case lr.CommentGroupTokenId:
				groups.Groups = append(
					groups.Groups,
					token.(lr.CommentGroupToken).CommentGroup)
			default:
				// Found a real token
				found = true
				value = token.(*lr.TokenValue)
				value.LeadingComment = groups
			}
		}
	} else if token.Id() == lr.NewlinesToken {
		return token, nil
	} else {
		value = token.(*lr.TokenValue)
	}

	// extract trailing comment groups
	for {
		peeked, err := lexer.buffered.Peek(1)
		if len(peeked) == 0 || err != nil {
			break
		}

		if peeked[0].Id() != lr.CommentGroupTokenId {
			break
		}

		value.TrailingComment.Groups = append(
			value.TrailingComment.Groups,
			peeked[0].(lr.CommentGroupToken).CommentGroup)

		_, err = lexer.buffered.Discard(1)
		if err != nil {
			panic("should never happen")
		}
	}

	return value, nil
}

// The basic Lexer conditionally emit newline from the underlying lexer to
// the parser to denote statement termination.  Newline is emitted if the line's
// final token token is
//  1. an identifier
//  2. a literal
//  3. a jump label
//  4. one of the keywords: `break`, `continue`, `fallthrough`, `return`,
//     `true`, or `false`
//  5. one of: `++`, `--`, `)`, `}`, or `]`
func NewBasicLexer(
	sourceFileName string,
	sourceContent io.Reader,
	emitter *parseutil.Emitter,
	options LexerOptions,
) lr.Lexer {
	return parseutil.NewImplicitTerminalLexer(
		NewAssociateCommentGroupsLexer(
			sourceFileName,
			sourceContent,
			emitter,
			options),
		lr.NewlinesToken,
		[]lr.SymbolId{
			lr.IdentifierToken, lr.UnderscoreToken,
			lr.IntegerLiteralToken, lr.StringLiteralToken,
			lr.RuneLiteralToken, lr.FloatLiteralToken,
			lr.ReturnToken, lr.BreakToken, lr.ContinueToken, lr.FallthroughToken,
			lr.TrueToken, lr.FalseToken,
			lr.AddOneAssignToken, lr.SubOneAssignToken,
			lr.RparenToken, lr.RbraceToken, lr.RbracketToken,
		})
}
