package main

import (
	"os"

	rolodex "github.com/awokensavant/rolodex"
	actions "github.com/awokensavant/rolodex/cmd/actions"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{}
	app.Name = "rolodex"
	app.Usage = "usage"
	app.Version = rolodex.GetVersion()
	app.Commands = actions.Commands()

	if errRun := app.Run(os.Args); errRun != nil {
		panic(errRun)
	}
}

// func main() {

// 	app := &cli.App{
// 		Name:     "rolodex",
// 		Usage:    "A sample command-line application",
// 		Version:  "0.1.0",
// 		Commands: []*cli.Command{},
// 	}

// 	if err := app.Run(os.Args); err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		os.Exit(1)
// 	}
// }
