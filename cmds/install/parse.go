package install

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strings"

	"github.com/gofulljs/gbook/cmds/cmdutil"
	"github.com/gofulljs/gbook/global"
	"github.com/hashicorp/go-version"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/xerrors"
)

var (
	errInvalidBookJson = errors.New("invalid book.json")
)

// parseBookJsonPlugins 解析book.json中的所有plugin
func parseBookJsonPlugins(content []byte) (mPlugin map[string]struct{}, err error) {

	defer func() {
		er := recover()
		if er != nil {
			err = xerrors.Errorf("%w", errInvalidBookJson)
		}
	}()

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

// parseAlreadyPlugins 解析已存在的plugin
func parseAlreadyPlugins(output []byte) (mAlreadyPlugin map[string]struct{}) {
	keys := jsoniter.Get(output, "dependencies").Keys()
	if len(keys) == 0 {
		return nil
	}

	mAlreadyPlugin = make(map[string]struct{})
	for _, key := range keys {
		if !strings.HasPrefix(key, global.PLUGIN_PREFIX) {
			continue
		}
		mAlreadyPlugin[key] = struct{}{}
	}

	return mAlreadyPlugin
}

func parseValidPluginName(pluginNoVersion string, output []byte, bookVersion string) (string, error) {

	buf := bytes.NewBuffer(output)

	r := bufio.NewReader(buf)

	maxVersion, err := version.NewVersion("0.0.0")
	if err != nil {
		return "", xerrors.Errorf("install %v err: %w", pluginNoVersion, err)
	}
	nameWithMaxVersion := pluginNoVersion

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", xerrors.Errorf("install %v err: %w", pluginNoVersion, err)
		}
		arr := strings.Split(string(line), " ")
		if len(arr) != 2 {
			continue
		}
		needBookVersion := strings.Trim(arr[1], "'")

		if needBookVersion != "*" {

			splitIndex := strings.Index(needBookVersion, "-")
			if splitIndex != -1 {
				afterStr := needBookVersion[splitIndex:]
				needBookVersion = strings.ReplaceAll(needBookVersion[:splitIndex], "x", "0") + afterStr
			} else {
				needBookVersion = strings.ReplaceAll(needBookVersion, "x", "0")
			}

			// 判断是否符合规定
			constraints, err := version.NewConstraint(needBookVersion)
			if err != nil {
				return "", xerrors.Errorf("install %v err: %w", pluginNoVersion, err)
			}
			bv, err := version.NewVersion(bookVersion)
			if err != nil {
				return "", xerrors.Errorf("install %v err: %w", pluginNoVersion, err)
			}
			// 校验不通过
			if !constraints.Check(bv) {
				continue
			}
		}
		name := arr[0]
		i := strings.Index(name, "@")
		version, err := version.NewVersion(name[i+1:])
		if err != nil {
			return "", xerrors.Errorf("install %v err: %w", pluginNoVersion, err)
		}

		if version.GreaterThan(maxVersion) {
			maxVersion = version
			nameWithMaxVersion = name
		}
	}

	return nameWithMaxVersion, nil
}
