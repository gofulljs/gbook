package cmdutil

import (
	"os"
	"path/filepath"

	"github.com/gofulljs/gbook/global"
	"github.com/gofulljs/gbook/util"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
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

func GetBookVars(cctx *cli.Context) (bookVersion, bookHome, bookVersionPath string, err error) {
	bookVersion = cctx.String("bookVersion")

	bookHome = cctx.String("bookHome")
	if bookHome == "" {
		bookHome, err = getRootPath()
		if err != nil {
			return "", "", "", err
		}
	}

	bookVersionPath = filepath.Join(bookHome, bookVersion)

	return bookVersion, bookHome, bookVersionPath, nil
}

// Plugin: get full plugin name
func Plugin(plugin string) string {
	return global.PLUGIN_PREFIX + plugin
}

// getRootPath get gitbook store path
func getRootPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", xerrors.Errorf("%w", err)
	}
	return filepath.Join(homeDir, ".gitbook", "versions"), nil
}
