package main

import (
	"log"
	"os"

	"github.com/svtter/commit/pkg"
	"github.com/urfave/cli/v2"
)

// change branch and pull from master
func changePull(branchName string) {
	r := pkg.Shellout("git", "checkout", "-b", branchName)
	r.Output()

	r = pkg.Shellout("git", "pull", "origin", branchName)
	r.Output()
}

// handle args from os.Args
func handleArgs() error {
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
	return nil
}

func main() {
	app := &cli.App{
		Name:  "commit",
		Usage: "Make git convenient!",
		Action: func(c *cli.Context) error {
			err := handleArgs()
			return err
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
