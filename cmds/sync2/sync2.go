package sync2

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofulljs/gbook/cmdutil"
	"github.com/gofulljs/gbook/global"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var Run = &cli.Command{
	Name:  "sync2",
	Usage: "sync2 gitbook, 不包含node_modules, suggest",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "source",
			Usage: "gitbook数据源",
			Value: "https://github.com/gofulljs/gitbook/archive/refs/tags/3.2.3.tar.gz",
		},
		&cli.StringFlag{
			Name:  "proxy1",
			Usage: fmt.Sprintf(`自定义加速源(前缀+source), 不传采用以下\n%v`, global.Proxy1s),
		},
		&cli.StringFlag{
			Name:  "proxy2",
			Usage: `自定义加速源(替换https://github.com前缀)`,
		},
	},
	Action: func(cctx *cli.Context) error {
		bookVersion := cctx.String("bookVersion")

		var err error
		bookHome := cctx.String("bookHome")
		if bookHome == "" {
			bookHome, err = getRootPath()
			if err != nil {
				return err
			}
		}

		bookVersionPath := filepath.Join(bookHome, bookVersion)

		if checkGitbookIsExist(bookVersionPath) {
			return nil
		}

		err = MustDownloadGitbook(bookHome, cctx.String("proxy1"), cctx.String("proxy2"), cctx.String("source"))
		if err != nil {
			return err
		}

		npmCmd := cmdutil.GetNpmCmd(cctx)
		return nodeInstall(bookVersionPath, npmCmd)
	},
}

// getRootPath get gitbook store path
func getRootPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", xerrors.Errorf("%w", err)
	}
	return filepath.Join(homeDir, ".gitbook", "versions"), nil
}

// checkGitbookIsExist check gitbook is in $HOME dir exist
func checkGitbookIsExist(bookVersionPath string) bool {
	bookJsFullPath := filepath.Join(bookVersionPath, "bin", "gitbook.js")
	_, err := os.Stat(bookJsFullPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}

	nodeModulesPath := filepath.Join(bookVersionPath, "node_modules")
	_, err = os.Stat(nodeModulesPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}

	return true
}
