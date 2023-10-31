package app

import (
	"github.com/gofulljs/gbook/cmds/sync2"
	"github.com/gofulljs/gbook/global"
	"github.com/urfave/cli/v2"
)

func InitApp() *cli.App {
	return &cli.App{
		Name:    global.AppName,
		Usage:   "uniswap tick update",
		Version: "v1.0.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "bookVersion",
				Aliases: []string{"bv"},
				Value:   global.BOOK_VERSION,
				EnvVars: []string{"BOOK_VERSION"},
			},
			&cli.StringFlag{
				Name:    "bookHome",
				Aliases: []string{"bh"},
				Usage:   "gitbook path, default is $HOME/.gitbook/versions/",
			},
			&cli.StringFlag{
				Name:    "nodePath",
				Usage:   "nodejs home, if not specified, use current node",
				EnvVars: []string{"BOOK_NODE_HOME"},
			},
		},
		Action: func(cctx *cli.Context) error {
			return nil
		},
		Commands: []*cli.Command{
			sync2.Run,
		},
	}
}
