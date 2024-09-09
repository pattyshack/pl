package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/pattyshack/gt/argparse"
	"github.com/pattyshack/pl/parser"
)

type Command struct {
	*argparse.Command
}

func (cmd *Command) Setup() {
	exprCmd := cmd.AddSubcommand("expr", "print expression")
	exprCmd.SetCommandFunc(
		cmd.printExpr,
		argparse.PositionalArgument{
			Name:        "files",
			Description: "list of expression file name paths",
			NumExpected: 1,
			VarArgs:     true,
		})
}

func (cmd *Command) printExpr(args []string) error {
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

		reducer := parser.NewReducer()
		lexer := parser.NewLexer(fileName, buffer, parser.LexerOptions{})

		expr, err := parser.ParseExpr(lexer, reducer)
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
