package app

import (
	"github.com/gofulljs/gbook/cmds/sync"
	"github.com/urfave/cli/v2"
)

func InitApp() *cli.App {
	return &cli.App{
		Name:    AppName,
		Usage:   "uniswap tick update",
		Version: "v1.0.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "chain-uri",
				//Required: true,
				Aliases: []string{"u"},
				EnvVars: []string{"CHAIN_URI"},
			},
		},
		Action: func(cctx *cli.Context) error {
			return nil
		},
		Commands: []*cli.Command{
			sync.Run,
		},
	}
}
