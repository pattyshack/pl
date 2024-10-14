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
	builder.Build(pkgIds...)

	errs := builder.WaitForCompletion()

	// TODO: rm
	for id, pkg := range builder.Packages() {
		fmt.Println("=================")
		fmt.Println("Package:", id)
		fmt.Println("-----------------")
		if pkg.Definitions != nil {
			fmt.Println(util.TreeString(pkg.Definitions, ""))
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
