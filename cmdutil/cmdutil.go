package cmdutil

import (
	"strings"

	"github.com/urfave/cli/v2"
)

func GetNpmCmd(cctx *cli.Context) string {
	nodePath := cctx.String("nodePath")
	npmCmd := "npm"
	if nodePath != "" {
		nodePath = strings.TrimSuffix(nodePath, "/")
		npmCmd = nodePath + "/" + npmCmd
	}
	return npmCmd
}

func GetNodeCmd(cctx *cli.Context) string {
	nodePath := cctx.String("nodePath")
	nodeCmd := "npm"
	if nodePath != "" {
		nodePath = strings.TrimSuffix(nodePath, "/")
		nodeCmd = nodePath + "/" + nodeCmd
	}
	return nodeCmd
}
