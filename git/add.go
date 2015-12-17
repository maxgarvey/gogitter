package git

import (
	// "fmt"
	"os/exec"
)

func Add(localWorkDir, fileToAdd string) (string, error) {
	// fmt.Printf("in Add method.\n") // debug

	// initialize add command
	addCommand := exec.Command(
		"/usr/local/bin/git",
		"-C",
		localWorkDir,
		"add",
		fileToAdd,
	)

	// perform add, return combined out and error
	standardOut, err := addCommand.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(standardOut), err
}
