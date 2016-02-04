package git

import (
	// "fmt"
	"os/exec"
)

func Branch(gitBinary, localWorkDir, branchName string, deleteBranch, newBranch bool) (string, error) {
	// fmt.Printf("in Branch method.\n") // debug

	// initialize checkout command
	var branchCommand *exec.Cmd
	if deleteBranch == true {
		branchCommand = exec.Command(
			gitBinary,
			"--git-dir",
			localWorkDir+"/.git",
			"branch",
			"-D",
			branchName,
		)
	} else if newBranch == true {
		branchCommand = exec.Command(
			gitBinary,
			"--git-dir",
			localWorkDir+"/.git",
			"branch",
			branchName,
		)
	} else {
		branchCommand = exec.Command(
			gitBinary,
			"--git-dir",
			localWorkDir+"/.git",
			"branch",
			branchName,
		)
	}
	// perform branch, return combined out and error
	standardOut, err := branchCommand.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(standardOut), nil
}
