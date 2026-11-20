package actions

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func GreetCommand() cli.Command {
	return cli.Command{
		Name:  "greet",
		Usage: "Greet someone",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Value:   "world",
				Usage:   "name of the person to greet",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Printf("Hello, %s!\n", c.String("name"))
			return nil
		},
	}
}
