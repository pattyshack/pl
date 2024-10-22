package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pattyshack/gt/argparse"

	"github.com/pattyshack/pl/build"
	"github.com/pattyshack/pl/build/cfg"
	"github.com/pattyshack/pl/util"
)

type Command struct {
	Tags string

	*argparse.Command
}

func (cmd *Command) Setup() {
	argparse.StringVar(&cmd.Tags, "tags", "", "Comma separated list of tags")

	cmd.SetCommandFunc(
		cmd.Build,
		argparse.PositionalArgument{
			Name:        "targets",
			Description: "list of build targets",
			NumExpected: 1,
			VarArgs:     true,
		})
}

func (cmd *Command) Build(
	args []string,
) error {
	config := cfg.NewConfig(strings.Split(cmd.Tags, ","))

	workspace, err := build.FindWorkspaceRoot(config)
	if err != nil {
		return err
	}

	pkgIds, err := workspace.MatchTargetPatterns(args...)
	if err != nil {
		return err
	}

	builder := build.NewBuilder(workspace, config)
	// TODO: don't hardcode build mode
	errs := builder.Build(build.BuildTest, pkgIds...)

	// TODO: rm
	for _, pkg := range builder.Packages() {
		fmt.Println("=================")
		fmt.Println("Package:", pkg.PackageID)
		fmt.Println("-----------------")
		if pkg.LibraryDefinitions != nil {
			fmt.Println("Library definitions:")
			fmt.Println(util.TreeString(pkg.LibraryDefinitions, "  "))
		}
		if pkg.BinaryDefinitions != nil {
			fmt.Println("Binary definitions:")
			fmt.Println(util.TreeString(pkg.BinaryDefinitions, "  "))
		}
		if pkg.TestDefinitions != nil {
			fmt.Println("Test definitions:")
			fmt.Println(util.TreeString(pkg.TestDefinitions, "  "))
		}
	}

	if len(errs) > 0 {
		fmt.Println("Errors:")
		for _, err := range errs {
			fmt.Println(err)
		}

		os.Exit(1)
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
		fmt.Println(err)
		os.Exit(1)
	}
}
