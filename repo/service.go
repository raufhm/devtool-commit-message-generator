package repo

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

// CommitGenerator is an interface for the commit generator
type CommitGenerator interface {
	Generate(branch, message string) *CommitMessageGen
	GetCommitMessage() string
	GetCurrentBranch() string
	Commit(branch, message string) error
}

type CommitMessageGen struct {
	Branch  string
	Message string
}

type CommitGen struct{}

// New is a constructor that creates a new instance of CommitMessageGen
func (c *CommitGen) New(branch, message string) *CommitMessageGen {
	return &CommitMessageGen{
		Branch:  branch,
		Message: message,
	}
}

// GetCommitMessage prompts the user for input to get the commit message
func (c *CommitGen) GetCommitMessage() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter commit message: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return input
}

// GetCurrentBranch gets the current branch of the Git repository
func (c *CommitGen) GetCurrentBranch() string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
		fmt.Println("STDERR:", stderr.String())
		panic(err)
	}

	re := regexp.MustCompile(`(\w+/\w+-\d+).*`)
	matches := re.FindStringSubmatch(out.String())
	if len(matches) > 1 {
		return matches[1]
	} else {
		panic("not able to get current branch")
	}
}

// CommitGen generates a new Git commit with the provided branch and message
func (c *CommitGen) CommitGen(branch, message string) error {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("git", "commit", "-am", fmt.Sprintf("%s: %s", branch, message))
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	fmt.Println(stdout.String())
	return nil
}
