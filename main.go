package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

func main() {
	msg, err := getCommitMsg()
	if err != nil {
		panic(err)
	}

	branch, err := cmdCurrentBranch()
	if err != nil {
		panic(err)
	}

	var stdout, stderr bytes.Buffer
	cmd := exec.Command("git", "commit", "-m", fmt.Sprintf("%s: %s", branch, msg))
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(stdout.String())

}

func getCommitMsg() (string, error) {
	// get commit message
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return input, nil
}

func cmdCurrentBranch() (string, error) {
	// get branch name until ticket number
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
		fmt.Println("STDERR:", stderr.String())
		return "", err
	}

	re := regexp.MustCompile(`(\w+/\w+-\d+).*`)
	matches := re.FindStringSubmatch(out.String())
	if len(matches) > 1 {
		return matches[1], nil
	}

	return "", errors.New("unable to get current branch")
}
