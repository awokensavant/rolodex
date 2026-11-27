package actions

import "github.com/urfave/cli/v3"

func Commands() []*cli.Command {
	return []*cli.Command{
		GreetCommand(),
	}
}
