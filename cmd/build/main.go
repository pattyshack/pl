package main

import (
	"fmt"
	"os"

	"github.com/pattyshack/pl/build"
)

func main() {
	workspace, err := build.FindWorkspaceRoot()
	if err != nil {
		panic(err)
	}

	for _, arg := range os.Args[1:] {
		fmt.Println("Pattern:", arg)
		ids, err := workspace.MatchTargetPattern(arg)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("Matches:")
		for _, id := range ids {
			fmt.Println(" ", id)
		}
	}
}
