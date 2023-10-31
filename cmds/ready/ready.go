package ready

import (
	"github.com/urfave/cli/v2"
)

var Run = &cli.Command{
	Name:  "ready",
	Usage: "env is ready",
	Action: func(cctx *cli.Context) error {
		// check node version
		// checkNodeVersion()

		// check gitbook-cli
		return nil
	},
}
