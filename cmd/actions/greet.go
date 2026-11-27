package actions

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func GreetCommand() *cli.Command {
	return &cli.Command{
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
		Action: func(ctx context.Context, c *cli.Command) error {
			fmt.Printf("Hello, %s!\n", c.String("name"))
			return nil
		},
	}
}
