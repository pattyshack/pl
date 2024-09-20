package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/pattyshack/gt/argparse"

	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser"
	"github.com/pattyshack/pl/parser/lexer"
	"github.com/pattyshack/pl/parser/lr"
	"github.com/pattyshack/pl/parser/reducer"
)

type Command struct {
	*argparse.Command

	parser.ParserOptions
}

func (cmd *Command) Setup() {
	argparse.BoolVar(
		&cmd.UseBasicLexer,
		"use-basic-lexer",
		false,
		"when true, use basic lexer for strict bottom-up lr parsing")

	argparse.BoolVar(
		&cmd.UseLRParseSource,
		"use-lr-parse-source",
		false,
		"when true, use lr.ParseSource instead of parser.ParseSource for source")

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

	exprCmd = cmd.AddSubcommand("package", "print package")
	exprCmd.SetCommandFunc(
		cmd.printPackage,
		argparse.PositionalArgument{
			Name:        "files",
			Description: "list of file name paths",
			NumExpected: 1,
			VarArgs:     true,
		})
}

func (cmd *Command) printPackage(
	args []string,
) error {
	fmt.Println("Files:")
	for _, filename := range args {
		fmt.Println(" ", filename)
	}
	fmt.Println("==========================")

	_, list, err := parser.ParsePackage(args, cmd.ParserOptions)

	if err != nil {
		fmt.Println("Parse error:", err)
		return err
	}

	fmt.Println("Tree:")
	fmt.Println("-----")
	fmt.Println(list.TreeString("", ""))

	return nil
}

func (cmd *Command) printFunc(
	parse func(string) (ast.Node, error),
	args []string,
) error {
	cmd.NewReaderFunc = func(fileName string) (io.Reader, error) {
		fmt.Println("==========================")
		fmt.Println("File name:", fileName)
		fmt.Println("==========================")

		content, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return nil, err
		}

		fmt.Println("Content:")
		fmt.Println("--------")
		fmt.Println(string(content[:len(content)-1]))
		fmt.Println("++++++++++++++++++++++++++++++++")

		return bytes.NewBuffer(content[:len(content)-1]), nil
	}

	for _, fileName := range args {
		expr, err := parse(fileName)
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
		func(
			fileName string,
		) (
			ast.Node,
			error,
		) {
			reducer := reducer.NewReducer()
			lexer, err := cmd.NewLexer(fileName, reducer)
			if err != nil {
				return nil, err
			}
			return lr.ParseExpr(lexer, reducer)
		},
		args)
}

func (cmd *Command) printTypeExpr(args []string) error {
	return cmd.printFunc(
		func(
			fileName string,
		) (
			ast.Node,
			error,
		) {
			reducer := reducer.NewReducer()
			lexer, err := cmd.NewLexer(fileName, reducer)
			if err != nil {
				return nil, err
			}
			return lr.ParseTypeExpr(lexer, reducer)
		},
		args)
}

func (cmd *Command) printStatement(args []string) error {
	return cmd.printFunc(
		func(
			fileName string,
		) (
			ast.Node,
			error,
		) {
			reducer := reducer.NewReducer()
			lexer, err := cmd.NewLexer(fileName, reducer)
			if err != nil {
				return nil, err
			}
			return lr.ParseStatement(lexer, reducer)
		},
		args)
}

func (cmd *Command) printDefinition(args []string) error {
	return cmd.printFunc(
		func(
			fileName string,
		) (
			ast.Node,
			error,
		) {
			reducer := reducer.NewReducer()
			lexer, err := cmd.NewLexer(fileName, reducer)
			if err != nil {
				return nil, err
			}
			return lr.ParseDefinition(lexer, reducer)
		},
		args)
}

func (cmd *Command) printSource(args []string) error {
	return cmd.printFunc(
		func(
			fileName string,
		) (
			ast.Node,
			error,
		) {
			_, src, err := parser.ParseSource(fileName, cmd.ParserOptions)
			return src, err
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
