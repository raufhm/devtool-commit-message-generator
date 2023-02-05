package main

import (
	"fmt"
	"github.com/raufhm/devtool-gitflow/repo"
)

func main() {
	commitGen := &repo.CommitGen{}
	commit := commitGen.Generate(commitGen.GetCurrentBranch(), commitGen.GetCommitMessage())
	if err := commitGen.Commit(commit.Branch, commit.Message); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Commit successful!")
}
