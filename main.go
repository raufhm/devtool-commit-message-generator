package main

import (
	"fmt"
	"github.com/raufhm/devtool-gitflow/repo"
)

// main is the main function that performs the git commit using the CommitGen struct.
// It gets the current branch and the commit message, and passes them to the CommitGen struct to perform the git commit.
func main() {
	cg := &repo.CommitGen{}
	commit := cg.New(cg.GetCurrentBranch(), cg.GetCommitMessage())
	if err := cg.CommitGen(commit.Branch, commit.Message); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("CommitGen successful!")
}
