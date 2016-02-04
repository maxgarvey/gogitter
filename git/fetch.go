package git

import (
	// "fmt"
	"os/exec"
)

func Fetch(gitBinary, localWorkDir, branches string) (string, error) {
	// fmt.Printf("in Fetch method.\n") // debug

	// initialize fetch command
	fetchCommand := exec.Command(
		gitBinary,
		"--git-dir",
		localWorkDir+"/.git",
		"fetch",
		branches,
	)
	// perform fetch, return combined out and error
	standardOut, err := fetchCommand.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(standardOut), nil
}
