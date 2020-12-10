package main

import (
	"log"

	"github.com/svtter/commit/pkg"
)

func main() {
	var commitArgs string
	errout, out, err := pkg.Shellout("git", "add", ".")
	pkg.Output(out, errout, err)

	commandLine := pkg.LoadArgs()
	if commandLine != "" {
		commitArgs = commandLine
		log.Println("commit message:", commitArgs)
	} else {
		commitArgs = pkg.ReadFromCommand()
	}

	errout, out, err = pkg.Shellout("git", "commit", "-m", commitArgs)
	pkg.Output(out, errout, err)

	errout, out, err = pkg.Shellout("git", "push")
	pkg.Output(out, errout, err)
}
