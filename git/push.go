package git

import (
	// "fmt"
	"os/exec"
)

func Push(localWorkDir, remote, branchName string, dryRun, deleteBoo bool) (string, error) {
	// fmt.Printf("in Push method.\n") // debug

	if deleteBoo == true {
		branchName = ":" + branchName
	}

	// initialize push command
	var pushCommand *exec.Cmd
	if dryRun == true {
		pushCommand = exec.Command(
			"/usr/local/bin/git",
			"-C",
			localWorkDir,
			"push",
			"--dry-run",
			remote,
			branchName,
		)
	} else {
		pushCommand = exec.Command(
			"/usr/local/bin/git",
			"-C",
			localWorkDir,
			"push",
			remote,
			branchName,
		)
	}
	// perform push, return combined out and error
	standardOut, err := pushCommand.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(standardOut), nil
}
