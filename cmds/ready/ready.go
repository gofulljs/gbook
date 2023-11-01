package ready

import (
	"bytes"
	"os/exec"
	"strconv"

	"github.com/gofulljs/gbook/cmds/cmdutil"
	"github.com/gofulljs/gbook/global"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var Run = &cli.Command{
	Name:  global.CmdReady,
	Usage: "check env is ready",
	Action: func(cctx *cli.Context) error {

		cmdutil.SetNodePath(cctx)

		// check node version
		err := checkNodeVersion()
		if err != nil {
			return err
		}

		// check gitbook-cli
		err = checkGitbookCli()
		if err != nil {
			return err
		}
		return nil
	},
}

func checkNodeVersion() error {
	command := exec.Command("node", "-v")
	output, err := command.Output()
	if err != nil {
		return xerrors.Errorf("%w", err)
	}
	output = bytes.TrimPrefix(output, []byte{'v'})
	arr := bytes.Split(output, []byte{'.'})
	version, err := strconv.Atoi(string(arr[0]))
	if err != nil {
		return xerrors.Errorf("%w", err)
	}

	if version > 10 {
		return xerrors.Errorf("err: node version %v is greater than 10", version)
	}

	return nil
}

func checkGitbookCli() error {
	command := exec.Command("gitbook", "ls")
	output, err := command.Output()
	if err != nil {
		return xerrors.Errorf("%w", err)
	}
	if bytes.Contains(output, []byte("There is no versions installed")) {
		return xerrors.New("err:" + string(output))
	}
	return nil
}
