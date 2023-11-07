package ready

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
	"time"

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
		if strings.Contains(err.Error(), "executable file not found in") {
			return nodeInstallGitbookCli()
		}
		return xerrors.Errorf("%w", err)
	}
	if bytes.Contains(output, []byte("No such file or directory")) {
		fmt.Println("gitbook-cli is not install, will install")
		return nodeInstallGitbookCli()
	}
	return nil
}

func nodeInstallGitbookCli() (err error) {

	cmd := exec.Command("npm", "install", "-g", "gitbook-cli")

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
