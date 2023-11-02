package ctest

import (
	"fmt"

	"github.com/gofulljs/gbook/global"
	"github.com/urfave/cli/v2"
)

var Run = &cli.Command{
	Name:   global.CmdTest,
	Usage:  "for test",
	Hidden: true,
	Action: func(cctx *cli.Context) error {

		fmt.Println("test")

		return nil
	},
}
