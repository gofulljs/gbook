package cmdutil

import (
	"github.com/gofulljs/gbook/global"
	"github.com/gofulljs/gbook/util"
	"github.com/urfave/cli/v2"
)

func Check(cctx *cli.Context) error {
	err := cctx.App.Run([]string{global.AppName, global.CmdReady})
	if err != nil {
		return err
	}

	err = cctx.App.Run([]string{global.AppName, global.CmdSync2})
	if err != nil {
		return err
	}

	return nil
}

func SetNodePath(cctx *cli.Context) {
	nodePath := cctx.String("nodePath")
	if nodePath != "" {
		util.Active(nodePath)
	}
}

// Plugin: get full plugin name
func Plugin(plugin string) string {
	return global.PLUGIN_PREFIX + plugin
}
