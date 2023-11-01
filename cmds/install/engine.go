package install

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/gofulljs/gbook/cmds/cmdutil"
	"github.com/gofulljs/gbook/global"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/xerrors"
)

func getAlreadyPlugins() (mAlreadyPlugin map[string]struct{}, err error) {
	cmd := exec.Command("npm", "list", "--json")
	output, err := cmd.Output()
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	return parseAlreadyPlugins(output), nil
}

// getBookJsonPlugins
func getBookJsonPlugins() (mPlugin map[string]struct{}, err error) {

	content, err := os.ReadFile(global.BOOK_JSON_FILE)
	if err != nil {
		if err == os.ErrNotExist {
			return nil, nil
		}
		return nil, err
	}

	var plugins []string
	jsoniter.Get(content, "plugins").ToVal(&plugins)

	mPlugin = make(map[string]struct{})
	for _, plugin := range plugins {
		if plugin[0] == '-' {
			continue
		}
		mPlugin[cmdutil.Plugin(plugin)] = struct{}{}
	}

	return mPlugin, nil
}

func getNeedPlugins() (mNeed map[string]struct{}, err error) {
	mAlready, err := getAlreadyPlugins()
	if err != nil {
		return nil, err
	}

	mAll, err := getBookJsonPlugins()
	if err != nil {
		return nil, err
	}

	mNeed = make(map[string]struct{})
	for k := range mAll {
		if _, ok := mAlready[k]; !ok {
			mNeed[k] = struct{}{}
		}
	}

	return mNeed, nil
}

func installPlugins(bookVersion string) (mNeed map[string]struct{}, err error) {
	mNeed, err = getNeedPlugins()
	if err != nil {
		return nil, err
	}

	for k := range mNeed {
		err = installSinglePlugin(k, bookVersion)
		if err != nil {
			return nil, err
		}
	}

	return mNeed, nil
}

// getValidePluginName return plugin with version
func getValidePluginName(plugin, bookVersion string) (string, error) {
	cmd := exec.Command("npm", "view", plugin+"@*", "engines.gitbook")
	output, err := cmd.Output()
	if err != nil {
		return "", xerrors.Errorf("%w", err)
	}

	return parseValidPluginName(plugin, output, bookVersion)
}

func installSinglePlugin(plugin, bookVersion string) (err error) {

	pluginName, err := getValidePluginName(plugin, bookVersion)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}
	cmd := exec.Command("npm", "install", pluginName)

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
