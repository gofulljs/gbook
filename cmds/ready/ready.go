package ready

import (
	"bytes"
	"os/exec"
	"strconv"

	"github.com/gofulljs/gbook/cmdutil"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var Run = &cli.Command{
	Name:  "ready",
	Usage: "check env is ready",
	Action: func(cctx *cli.Context) error {

		nodeCmd := cmdutil.GetNodeCmd(cctx)

		// check node version
		err := checkNodeVersion(nodeCmd)
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

func checkNodeVersion(nodeCmd string) error {
	command := exec.Command(nodeCmd, "-v")
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
		return xerrors.Errorf("node version %v is greater than 10", version)
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
		return xerrors.New(string(output))
	}
	return nil
}
