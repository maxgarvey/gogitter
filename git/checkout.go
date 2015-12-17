package git

import (
	// "fmt"
	"os/exec"
)

func Checkout(localWorkDir, branchName string, newBranch bool) (string, error) {
	// fmt.Printf("in Checkout method.\n") // debug

	// initialize checkout command
	var checkoutCommand *exec.Cmd
	if newBranch == true {
		checkoutCommand = exec.Command(
			"/usr/local/bin/git",
			"-C",
			localWorkDir,
			"checkout",
			"-b",
			branchName,
		)
	} else {
		checkoutCommand = exec.Command(
			"/usr/local/bin/git",
			"-C",
			localWorkDir,
			"checkout",
			branchName,
		)
	}
	// perform checkout, return combined out and error
	standardOut, err := checkoutCommand.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(standardOut), nil
}
