package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")

	var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
		fmt.Println("STDERR:", stderr.String())
		return
	}

	branch := out.String()
	fmt.Println("Current branch:", branch)
}
