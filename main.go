package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/maxgarvey/gogitter/git"
)

func main() {
	fmt.Printf("in main method\n")       // debug
	fmt.Printf("running diagnostic\n\n") // debug

	fmt.Printf("doing git clone\n") // debug

	// perform a git clone
	cloneRepository := "http://github.com/maxgarvey/junkZone"
	cloneRepository = "ssh://git@stash.nikedev.com/zzar/junkzone.git" // this one for nike

	gitBinary := "/usr/local/bin/git"

	localWorkDir := "/Users/mgarve/go/src/github.com/maxgarvey/gogitter/junkZone"
	cleanup := false
	cloneOutput, err := git.Clone(gitBinary, cloneRepository, localWorkDir, cleanup)
	if err != nil {
		fmt.Printf(
			"error doing clone: \n%s\n",
			err.Error(),
		)
	}
	fmt.Printf(
		"git.Clone output:\n%s\n",
		cloneOutput,
	)

	// perform a git checkout
	checkoutBranch := "master"
	newBranch := false
	checkoutOutput, err := git.Checkout(gitBinary, localWorkDir, checkoutBranch, newBranch)
	if err != nil {
		fmt.Printf(
			"error during checkout:\n%s\n",
			err.Error(),
		)
	}
	fmt.Printf(
		"git.Checkout output:\n%s\n",
		checkoutOutput,
	)

	// perform a git fetch
	branchesOfInterest := "--all"
	fetchOutput, err := git.Fetch(gitBinary, localWorkDir, branchesOfInterest)
	if err != nil {
		fmt.Printf(
			"error during fetch:\n%s\n",
			err.Error(),
		)
	}
	fmt.Printf(
		"git.Fetch output:\n%s\n",
		fetchOutput,
	)

	// perform a git branch
	branchName := "branch"
	deleteBranch := false
	newBranch = true
	branchOutput, err := git.Branch(gitBinary, localWorkDir, branchName, deleteBranch, newBranch)
	if err != nil {
		fmt.Printf(
			"error during branch:\n%s\n",
			err.Error(),
		)
	}
	fmt.Printf(
		"git.Branch output:\n%s\n",
		branchOutput,
	)
	// checkout the git branch
	checkoutOutput, err = git.Checkout(gitBinary, localWorkDir, branchName, false)
	if err != nil {
		fmt.Printf(
			"error during checkout:\n%s\n",
			err.Error(),
		)
	}
	fmt.Printf(
		"git.Checkout output:\n%s\n",
		checkoutOutput,
	)

	err = ioutil.WriteFile(localWorkDir+"/a.txt", []byte("blah"), 0644)
	if err != nil {
		panic(err)
	}

	// perform a git add
	addOutput, err := git.Add(
		gitBinary,
		localWorkDir,
		localWorkDir+"/a.txt",
	)
	if err != nil {
		fmt.Printf(
			"error during add:\n%s\n",
			err.Error(),
		)
	}
	fmt.Printf(
		"git.Add output:\n%s\n",
		addOutput,
	)

	// perform a git commit
	commitMessage := "my commit message"
	commmitOutput, err := git.Commit(gitBinary, localWorkDir, commitMessage)
	if err != nil {
		fmt.Printf(
			"error during commit:\n%s\n",
			err.Error(),
		)
	}
	fmt.Printf(
		"git.Commit output:\n%s\n",
		commmitOutput,
	)

	// perform the push; still requires github credentials
	// need to fix that somehow
	remote := "origin"
	pushBranch := branchName
	dryRun := false
	deleteBoo := true
	pushOutput, err := git.Push(gitBinary, localWorkDir, remote, pushBranch, dryRun, deleteBoo)
	if err != nil {
		fmt.Printf(
			"error during push:\n%s\n",
			err.Error(),
		)
	}
	fmt.Printf(
		"git.Push output:\n%s\n",
		pushOutput,
	)

	deleteBoo = false
	pushOutput, err = git.Push(gitBinary, localWorkDir, remote, pushBranch, dryRun, deleteBoo)
	if err != nil {
		fmt.Printf(
			"error during push:\n%s\n",
			err.Error(),
		)
	}
	fmt.Printf(
		"git.Push output:\n%s\n",
		pushOutput,
	)

	// cleanup
	cleanupRepoCommand := exec.Command(
		"rm",
		"-rf",
		localWorkDir,
	)
	_, err = cleanupRepoCommand.CombinedOutput()
	if err != nil {
		fmt.Printf(
			"error performing cleanup after git stuff:\n%s\n",
			err.Error(),
		)
	}
}
