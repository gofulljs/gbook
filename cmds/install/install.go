package install

import (
	"fmt"

	"github.com/gofulljs/gbook/cmds/cmdutil"
	"github.com/gofulljs/gbook/global"
	"github.com/gofulljs/gbook/util"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var Run = &cli.Command{
	Name:    global.CmdInstall,
	Aliases: []string{"i"},
	Usage:   "install plugin\n `install`: install all plugins from gitbook\t\n `install [plugins...]`: install plugin you want, eg: `gbook install code ga`",
	Action: func(cctx *cli.Context) error {

		//book.json is Exist?
		isExist, err := util.GetFileExist("book.json")
		if err != nil {
			return xerrors.Errorf("%w", err)
		}
		if !isExist {
			return nil
		}

		//check
		err = cmdutil.Check(cctx)
		if err != nil {
			return err
		}

		bookVersion := cctx.String("bookVersion")

		plugins := cctx.Args().Slice()
		if len(plugins) == 0 {
			mNeed, err := installPlugins(bookVersion)
			if err != nil {
				return err
			}
			if len(mNeed) == 0 {
				fmt.Println("no plugin need to install")
			}
		}

		for _, plugin := range plugins {
			err = installSinglePlugin(cmdutil.Plugin(plugin), bookVersion)
			if err != nil {
				return err
			}
		}

		// 清空lock文件
		return util.DeleteFileIfExist("package-lock.json")
	},
}
