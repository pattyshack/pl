package parser

import (
	"bytes"
	"io"
	"os"

	"github.com/pattyshack/pl/errors"
	"github.com/pattyshack/pl/parser/lexer"
	"github.com/pattyshack/pl/parser/lr"
)

type ParserOptions struct {
	lexer.LexerOptions

	OpenFunc func(string) (io.Reader, error)

	UseBasicLexer bool

	UseLRParseSource bool

	DisableSyntaxAnalysis bool
}

func (options ParserOptions) NewReader(fileName string) (io.Reader, error) {
	if options.OpenFunc != nil {
		return options.OpenFunc(fileName)
	}

	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(content), nil
}

func (options ParserOptions) NewLexer(
	fileName string,
	emitter *errors.Emitter,
	reducer lr.Reducer,
) (
	lr.Lexer,
	error,
) {
	reader, err := options.NewReader(fileName)
	if err != nil {
		return nil, err
	}

	if options.UseBasicLexer {
		return lexer.NewBasicLexer(
			fileName,
			reader,
			emitter,
			options.LexerOptions), nil
	}
	return lexer.NewLexer(
		fileName,
		reader,
		emitter,
		reducer,
		options.LexerOptions), nil
}
