package main

import (
	"log"

	"github.com/svtter/commit/pkg"
)

// change branch and pull from master
func changePull(branchName string) {
	r := pkg.Shellout("git", "checkout", "-b", branchName)
	r.Output()

	r = pkg.Shellout("git", "pull", "origin", branchName)
	r.Output()
}

func main() {
	var commitArgs string
	r := pkg.Shellout("git", "add", ".")
	r.Output()

	commandLine := pkg.LoadArgs()
	if commandLine != "" {
		commitArgs = commandLine
		log.Println("commit message:", commitArgs)
	} else {
		commitArgs = pkg.ReadFromCommand()
	}

	r = pkg.Shellout("git", "commit", "-m", commitArgs)
	r.Output()

	r = pkg.Shellout("git", "push")
	r.Output()
}
