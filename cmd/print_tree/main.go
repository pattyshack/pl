package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/pattyshack/gt/argparse"
	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/parser"
	"github.com/pattyshack/pl/util"
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

	exprCmd := cmd.AddSubcommand("source", "print source")
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

	emitter := &lexutil.ErrorEmitter{}
	list, _, _ := parser.ParsePackage(args, emitter, cmd.ParserOptions)

	parseErrors := emitter.Errors()

	fmt.Println("Tree:")
	fmt.Println("-----")
	buffer := &bytes.Buffer{}
	util.PrintTree(buffer, list, "")
	fmt.Println(buffer.String())

	if len(parseErrors) > 0 {
		fmt.Println("---------------------")
		fmt.Println("Parse errors:")
		fmt.Println("---------------------")
		for idx, err := range parseErrors {
			fmt.Printf("error %d: %s\n", idx, err)
		}
	}

	return nil
}

func (cmd *Command) printSource(
	args []string,
) error {
	for _, fileName := range args {
		emitter := &lexutil.ErrorEmitter{}
		result := parser.ParseSource(fileName, emitter, cmd.ParserOptions)
		parseErrors := emitter.Errors()

		if result != nil {
			fmt.Println("Parsed:")
			fmt.Println("-----")
			buffer := &bytes.Buffer{}
			util.PrintTree(buffer, result, "")
			fmt.Println(buffer.String())
		}

		if len(parseErrors) > 0 {
			fmt.Println("---------------------")
			fmt.Println("Parse errors:")
			fmt.Println("---------------------")
			for idx, err := range parseErrors {
				fmt.Printf("error %d: %s\n", idx, err)
			}
		}
	}

	return nil
}

func (cmd *Command) printTokens(args []string) error {
	cmd.UseBasicLexer = true
	for _, fileName := range args {
		lexer, err := cmd.NewLexer(fileName, &lexutil.ErrorEmitter{}, nil)
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

	cmd.OpenFunc = func(fileName string) (io.Reader, error) {
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
