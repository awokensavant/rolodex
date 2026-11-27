package actions

import (
	"context"
	"fmt"

	"github.com/awokensavant/rolodex/domain"
	"github.com/urfave/cli/v3"
)

func VersionCommand() *cli.Command {
	return &cli.Command{
		Name:  "version",
		Usage: "Print the version of the application",
		Action: func(ctx context.Context, c *cli.Command) error {
			fmt.Printf("Version: %s\n", domain.Version4.String())
			return nil
		},
	}
}
