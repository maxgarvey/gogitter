package git

import (
	// "fmt"
	"os/exec"
)

func Clone(gitBinary, url, destinationDirectory string, cleanup bool) (string, error) {
	// fmt.Printf("in Clone method.\n") // debug

	// perform cleanup
	if cleanup {
		cleanupCommand := exec.Command(
			"/bin/rm",
			"-rf",
			destinationDirectory,
		)
		_, err := cleanupCommand.CombinedOutput()
		if err != nil {
			return "", err
		}
	}

	// initialize clone command
	cloneCommand := exec.Command(
		gitBinary,
		"clone",
		url,
		destinationDirectory,
	)
	// perform clone, return combined out and error
	standardOut, err := cloneCommand.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(standardOut), nil
}
