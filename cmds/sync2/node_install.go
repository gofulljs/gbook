package sync2

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"

	"golang.org/x/xerrors"
)

func nodeInstall(bookVersionPath string) (err error) {

	cmd := exec.Command("npm", "install")

	cmd.Dir = bookVersionPath

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

func nodeInstallBk(bookVersionPath, nodePath string) (output []byte, err error) {
	npmCmd := "npm"
	if nodePath != "" {
		nodePath = strings.TrimSuffix(nodePath, "/")
		npmCmd = nodePath + "/" + npmCmd
	}
	command := exec.Command(npmCmd, "install", bookVersionPath, "-C", bookVersionPath)

	output, err = command.Output()
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	return output, nil
}
