package app

import (
	"fmt"
	"io"
	"os/exec"
	"time"

	"github.com/gofulljs/gbook/cmds/cmdutil"
	"github.com/gofulljs/gbook/cmds/ctest"
	"github.com/gofulljs/gbook/cmds/install"
	"github.com/gofulljs/gbook/cmds/ready"
	"github.com/gofulljs/gbook/cmds/sync"
	"github.com/gofulljs/gbook/cmds/sync2"
	"github.com/gofulljs/gbook/global"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var LogDetail = false

func InitApp() *cli.App {
	return &cli.App{
		Name:    global.AppName,
		Usage:   "uniswap tick update, other command will forward gitbook *",
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
			&cli.BoolFlag{
				Name:        "logDetail",
				Usage:       "print log Detail for development",
				Hidden:      true,
				Aliases:     []string{"ld"},
				Destination: &LogDetail,
			},
		},
		Action: func(cctx *cli.Context) error {

			//check
			err := cmdutil.Check(cctx)
			if err != nil {
				return err
			}

			args := cctx.Args().Slice()
			if args[0] == "build" || args[0] == "serve" {
				// 会先install
				err = cctx.App.Run([]string{global.AppName, global.CmdInstall})
				if err != nil {
					return err
				}
			}
			fmt.Print("forward:\n", "gitbook ")
			for _, v := range args {
				fmt.Print(v + " ")
			}
			fmt.Println()

			return gitbookForward(args)
		},
		Commands: []*cli.Command{
			ready.Run,
			sync.Run,
			sync2.Run,
			install.Run,
			ctest.Run,
		},
	}
}

func gitbookForward(args []string) (err error) {

	cmd := exec.Command("gitbook", args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return xerrors.Errorf("%w", err)
	}
	cmd.Stderr = cmd.Stdout

	if err = cmd.Start(); err != nil {
		return xerrors.Errorf("%w", err)
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 4096)
		n, err := stdout.Read(tmp)
		fmt.Print(string(tmp[:n]))
		if err != nil {
			if err == io.EOF {
				break
			}
			return xerrors.Errorf("%w", err)
		}
		time.Sleep(100 * time.Millisecond)
	}
	if err = cmd.Wait(); err != nil {
		return xerrors.Errorf("%w", err)
	}

	return nil
}
