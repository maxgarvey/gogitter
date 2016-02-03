package git

import (
	// "fmt"
	"os/exec"
)

func Add(gitBinary, localWorkDir, fileToAdd string) (string, error) {
	// fmt.Printf("in Add method.\n") // debug

	// initialize add command
	addCommand := exec.Command(
		gitBinary,
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
