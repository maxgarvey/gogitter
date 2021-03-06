package git

import (
	// "fmt"
	"os/exec"
)

func Commit(gitBinary, localWorkDir, commitMessage string) (string, error) {
	// fmt.Printf("in Commit method.\n") // debug

	// initialize commit command
	commitCommand := exec.Command(
		gitBinary,
		"-C",
		localWorkDir,
		"commit",
		"-m",
		commitMessage,
	)

	// perform commit, return combined out and error
	standardOut, err := commitCommand.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(standardOut), nil
}
