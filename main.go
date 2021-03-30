package main

import (
	"log"
	"strings"

	"github.com/svtter/commit/pkg"
)

var (
	starts = map[string]string {
		"feat": "new feature",
		"fix": "fix bug",
		"docs": "documentation",
		"style": "code style without business change",
		"refactor": "refactor code",
		"test": "test code",
		"chore": "change for build tools",
	}
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
	if !checkPrefix(commitArgs) {
		log.Println("commit message is not allowed. Please input with fea/fix/docs/style/refactor/test/chore.")
	}

	errout, out, err = pkg.Shellout("git", "commit", "-m", commitArgs)
	pkg.Output(out, errout, err)

	errout, out, err = pkg.Shellout("git", "push")
	pkg.Output(out, errout, err)
}

func checkPrefix(args string) bool {
	for key, _ := range starts {
		if strings.HasPrefix(key, args) {
			return true
		}
	}
	return false
}
