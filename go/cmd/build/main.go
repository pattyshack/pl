package main

import (
	"fmt"
	"os"

	"github.com/pattyshack/pl/build"
	"github.com/pattyshack/pl/util"
)

func main() {
	workspace, err := build.FindWorkspaceRoot()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pkgIds, err := workspace.MatchTargetPatterns(os.Args[1:]...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	builder := build.NewBuilder(workspace)
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
}
