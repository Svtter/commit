package main

import (
	"fmt"
	"log"
	"os"

	"github.com/svtter/commit/pkg"
	"github.com/urfave/cli/v2"
)

func newMain() cli.App {
	var message string

	app := &cli.App{
		Name:  "commit",
		Usage: "commit a git pipeline.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "message",
				Usage:       "message for current commit.",
				Destination: &message,
			},
			&cli.BoolFlag{
				Name:    "feat",
				Aliases: []string{"f"},
				Usage:   "Use use flag if commit is a feat",
			},
		},
		Action: func(c *cli.Context) error {
			// fmt.Printf("%+v\n", c.Args())
			message = c.Args().First()
			fmt.Printf("%+v\n", message)
			return nil
		},
	}
	return *app
}

func oldMain() {
	var commitArgs string

	// git add command
	errout, out, err := pkg.Shellout("git", "add", ".")
	pkg.Output(out, errout, err)

	// parse command arguments
	commandLine := pkg.LoadArgs()
	if commandLine != "" {
		commitArgs = commandLine
		log.Println("commit message:", commitArgs)
	} else {
		commitArgs = pkg.ReadFromCommand()
	}

	// check the commit message
	if !pkg.CheckPrefix(commitArgs) {
		log.Println("commit message is not allowed. Please input with fea/fix/docs/style/refactor/test/chore.")
		return
	}

	// make a git commit command
	errout, out, err = pkg.Shellout("git", "commit", "-m", commitArgs)
	pkg.Output(out, errout, err)

	// make a git push command
	errout, out, err = pkg.Shellout("git", "push")
	pkg.Output(out, errout, err)

}

func main() {
	app := newMain()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
