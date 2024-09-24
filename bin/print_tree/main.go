package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/pattyshack/gt/argparse"

	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser"
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
	fmt.Println(ast.ListTreeString(list.Elements, "", "", "Definition"))

	return nil
}

func (cmd *Command) printFunc(
	parse func(string) (*reducer.Reducer, interface{}, error),
	args []string,
) error {
	for _, fileName := range args {
		reducer, result, err := parse(fileName)
		if err != nil {
			fmt.Println("Parse error:", err)
			continue
		}

		fmt.Println("Parsed:")
		fmt.Println("-----")
		switch value := result.(type) {
		case ast.Node:
			fmt.Println(value.TreeString("", ""))
		case *ast.DefinitionList:
			ast.ListTreeString(value.Elements, "", "", "Definition")
		default:
			panic(fmt.Sprintf("unexpected value: %v", result))
		}

		if len(reducer.ParseErrors) == 0 {
			continue
		}

		fmt.Println("---------------------")
		fmt.Println("Embedded parse error:")
		fmt.Println("---------------------")
		for idx, err := range reducer.ParseErrors {
			fmt.Printf("error %d: %s\n", idx, err)
		}
	}

	return nil
}

func (cmd *Command) printExpr(args []string) error {
	return cmd.printFunc(
		func(
			fileName string,
		) (
			*reducer.Reducer,
			interface{},
			error,
		) {
			return parser.ParseExpr(fileName, cmd.ParserOptions)
		},
		args)
}

func (cmd *Command) printTypeExpr(args []string) error {
	return cmd.printFunc(
		func(
			fileName string,
		) (
			*reducer.Reducer,
			interface{},
			error,
		) {
			return parser.ParseTypeExpr(fileName, cmd.ParserOptions)
		},
		args)
}

func (cmd *Command) printStatement(args []string) error {
	return cmd.printFunc(
		func(
			fileName string,
		) (
			*reducer.Reducer,
			interface{},
			error,
		) {
			return parser.ParseStatement(fileName, cmd.ParserOptions)
		},
		args)
}

func (cmd *Command) printDefinition(args []string) error {
	return cmd.printFunc(
		func(
			fileName string,
		) (
			*reducer.Reducer,
			interface{},
			error,
		) {
			return parser.ParseDefinition(fileName, cmd.ParserOptions)
		},
		args)
}

func (cmd *Command) printSource(args []string) error {
	return cmd.printFunc(
		func(
			fileName string,
		) (
			*reducer.Reducer,
			interface{},
			error,
		) {
			return parser.ParseSource(fileName, cmd.ParserOptions)
		},
		args)
}

func (cmd *Command) printTokens(args []string) error {
	cmd.UseBasicLexer = true
	for _, fileName := range args {
		lexer, err := cmd.NewLexer(fileName, nil)
		if err != nil {
			continue
		}

		fmt.Println("Tokens:")
		fmt.Println("-----")
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

	cmd.NewReaderFunc = func(fileName string) (io.Reader, error) {
		fmt.Println("==========================")
		fmt.Println("File name:", fileName)
		fmt.Println("==========================")

		content, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return nil, err
		}

		if len(content) > 0 && content[len(content)-1] == '\n' {
			content = content[:len(content)-1]
		}

		fmt.Println("Content:")
		fmt.Println("--------")
		fmt.Println(string(content))
		fmt.Println("++++++++++++++++++++++++++++++++")

		return bytes.NewBuffer(content), nil
	}

	cmd.Setup()

	err := argparse.Execute()
	if err != nil {
		panic(err)
	}
}
