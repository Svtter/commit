package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/svtter/commit/pkg"
	"github.com/urfave/cli/v2"
)

// v2 version of commit
func MainV2() cli.App {
	var commitMessage string
	var isFeature bool

	app := &cli.App{
		Name:  "commit",
		Usage: "commit a git pipeline.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "message",
				Usage:       "message for current commit.",
				Destination: &commitMessage,
			},
			&cli.BoolFlag{
				Name:        "feat",
				Aliases:     []string{"f"},
				Usage:       "Use use flag if commit is a feat",
				Destination: &isFeature,
			},
		},
		Action: func(c *cli.Context) error {
			// fmt.Printf("%+v\n", c.Args())
			commitMessage = c.Args().First()
			if isFeature {
				commitMessage = "feat: " + commitMessage
			}
			fmt.Printf("%+v\n", commitMessage)

			// check the commit message
			if !pkg.CheckPrefix(commitMessage) {
				errMessage := "commit message is not allowed. Please input with fea/fix/docs/style/refactor/test/chore"
				log.Println(errMessage)
				return errors.New(errMessage)
			}

			pkg.CommitPipeline(commitMessage)
			return nil
		},
	}
	return *app
}

// v1 version of commit
func MainV1() {
	var commitArgs string

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
	pkg.CommitPipeline(commitArgs)
}

func main() {
	// MainV1()
	app := MainV2()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
