package main

import (
	"context"
	"os"

	"github.com/awokensavant/rolodex"
	"github.com/awokensavant/rolodex/cmd/actions"
	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{}
	app.Name = "rolodex"
	app.Usage = "usage"
	app.Version = rolodex.GetVersion()
	app.Commands = actions.Commands()

	app.Run(context.Background(), os.Args)
}
