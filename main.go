package main

import (
	"github.com/svtter/commit/pkg"
	"log"
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
	if !pkg.CheckPrefix(commitArgs) {
		log.Println("commit message is not allowed. Please input with fea/fix/docs/style/refactor/test/chore.")
		return
	}

	errout, out, err = pkg.Shellout("git", "commit", "-m", commitArgs)
	pkg.Output(out, errout, err)

	errout, out, err = pkg.Shellout("git", "push")
	pkg.Output(out, errout, err)
}

