package main

import (
	"fmt"
	"github.com/raufhm/devtool-gitflow/repo"
)

func main() {
	commitGen := &repo.CommitGen{}
	commit := commitGen.New(
		commitGen.GetCurrentBranch(),
		commitGen.GetCommitMessage())

	if err := commitGen.CommitGen(commit.Branch, commit.Message); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("CommitGen successful!")
}
