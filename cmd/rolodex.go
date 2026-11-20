package main

import (
	"os"

	"github.com/awokensavant/rolodex"
	"github.com/awokensavant/rolodex/cmd/actions"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{}
	app.Name = "rolodex"
	app.Usage = "usage"
	app.Version = rolodex.GetVersion()
	app.Commands = actions.Commands()

	if errRun := app.Run(os.Args); errRun != nil {
		panic(errRun)
	}
}
