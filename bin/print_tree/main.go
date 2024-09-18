package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/pattyshack/gt/argparse"
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lexer"
	"github.com/pattyshack/pl/parser/lr"
	"github.com/pattyshack/pl/parser/reducer"
)

type Command struct {
	*argparse.Command

	useBasicLexer bool
}

func (cmd *Command) Setup() {
	argparse.BoolVar(
		&cmd.useBasicLexer,
		"use-basic-lexer",
		false,
		"when true, use basic lexer for strict bottom-up lr parsing")

	exprCmd := cmd.AddSubcommand("expr", "print expression")
	exprCmd.SetCommandFunc(
		cmd.printExpr,
		argparse.PositionalArgument{
			Name:        "files",
			Description: "list of expression file name paths",
			NumExpected: 1,
			VarArgs:     true,
		})

	exprCmd = cmd.AddSubcommand("type_expr", "print type expression")
	exprCmd.SetCommandFunc(
		cmd.printTypeExpr,
		argparse.PositionalArgument{
			Name:        "files",
			Description: "list of expression file name paths",
			NumExpected: 1,
			VarArgs:     true,
		})

	exprCmd = cmd.AddSubcommand("statement", "print statement")
	exprCmd.SetCommandFunc(
		cmd.printStatement,
		argparse.PositionalArgument{
			Name:        "files",
			Description: "list of expression file name paths",
			NumExpected: 1,
			VarArgs:     true,
		})

	exprCmd = cmd.AddSubcommand("definition", "print definition")
	exprCmd.SetCommandFunc(
		cmd.printDefinition,
		argparse.PositionalArgument{
			Name:        "files",
			Description: "list of expression file name paths",
			NumExpected: 1,
			VarArgs:     true,
		})

	exprCmd = cmd.AddSubcommand("source", "print source")
	exprCmd.SetCommandFunc(
		cmd.printSource,
		argparse.PositionalArgument{
			Name:        "files",
			Description: "list of expression file name paths",
			NumExpected: 1,
			VarArgs:     true,
		})

	exprCmd = cmd.AddSubcommand("tokens", "print token")
	exprCmd.SetCommandFunc(
		cmd.printTokens,
		argparse.PositionalArgument{
			Name:        "files",
			Description: "list of file name paths",
			NumExpected: 1,
			VarArgs:     true,
		})
}

func (cmd *Command) newLexer(
	fileName string,
	reader io.Reader,
	options lexer.LexerOptions,
	reducer lr.Reducer,
) lr.Lexer {
	if cmd.useBasicLexer {
		return lexer.NewBasicLexer(fileName, reader, options)
	}

	return lexer.NewLexer(fileName, reader, options, reducer)
}

func (cmd *Command) printFunc(
	parse func(lr.Lexer, lr.Reducer) (ast.Node, error),
	args []string,
) error {
	for _, fileName := range args {
		fmt.Println("==========================")
		fmt.Println("File name:", fileName)
		fmt.Println("==========================")

		content, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error opening file:", err)
			continue
		}

		fmt.Println("Content:")
		fmt.Println("--------")
		fmt.Println(string(content[:len(content)-1]))
		fmt.Println("++++++++++++++++++++++++++++++++")

		buffer := bytes.NewBuffer(content[:len(content)-1])

		reducer := reducer.NewReducer()
		lexer := cmd.newLexer(fileName, buffer, lexer.LexerOptions{}, reducer)

		expr, err := parse(lexer, reducer)
		if err != nil {
			fmt.Println("Parse error:", err)
			continue
		}

		fmt.Println("Tree:")
		fmt.Println("-----")
		fmt.Println(expr.TreeString("", ""))
	}

	return nil
}

func (cmd *Command) printExpr(args []string) error {
	return cmd.printFunc(
		func(lexer lr.Lexer, reducer lr.Reducer) (ast.Node, error) {
			return lr.ParseExpr(lexer, reducer)
		},
		args)
}

func (cmd *Command) printTypeExpr(args []string) error {
	return cmd.printFunc(
		func(lexer lr.Lexer, reducer lr.Reducer) (ast.Node, error) {
			return lr.ParseTypeExpr(lexer, reducer)
		},
		args)
}

func (cmd *Command) printStatement(args []string) error {
	return cmd.printFunc(
		func(lexer lr.Lexer, reducer lr.Reducer) (ast.Node, error) {
			return lr.ParseStatement(lexer, reducer)
		},
		args)
}

func (cmd *Command) printDefinition(args []string) error {
	return cmd.printFunc(
		func(lexer lr.Lexer, reducer lr.Reducer) (ast.Node, error) {
			return lr.ParseDefinition(lexer, reducer)
		},
		args)
}

func (cmd *Command) printSource(args []string) error {
	return cmd.printFunc(
		func(lexer lr.Lexer, reducer lr.Reducer) (ast.Node, error) {
			return lr.ParseSource(lexer, reducer)
		},
		args)
}

func (cmd *Command) printTokens(args []string) error {
	for _, fileName := range args {
		fmt.Println("==========================")
		fmt.Println("File name:", fileName)
		fmt.Println("==========================")

		content, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error opening file:", err)
			continue
		}

		fmt.Println("Content:")
		fmt.Println("--------")
		fmt.Println(string(content[:len(content)-1]))
		fmt.Println("++++++++++++++++++++++++++++++++")
		fmt.Println("Tokens:")
		fmt.Println("-----")

		buffer := bytes.NewBuffer(content[:len(content)-1])

		lexer := lexer.NewBasicLexer(fileName, buffer, lexer.LexerOptions{})
		for {
			token, err := lexer.Next()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Lex error:", err)
				break
			}

			fmt.Println(token)
		}
	}

	return nil
}

func main() {
	cmd := &Command{
		Command: argparse.CommandLine,
	}

	cmd.Setup()

	err := argparse.Execute()
	if err != nil {
		panic(err)
	}
}
