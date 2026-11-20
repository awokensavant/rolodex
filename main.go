package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "rolodex",
		Usage:   "A sample command-line application",
		Version: "0.1.0",
		Commands: []*cli.Command{
			{
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
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
