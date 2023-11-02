package installtest

import (
	"os"
	"testing"

	"github.com/gofulljs/gbook/cmds/app"
	"github.com/gofulljs/gbook/global"
	"github.com/stretchr/testify/assert"
)

func TestInstall(t *testing.T) {
	os.Setenv("BOOK_NODE_HOME", "C:\\Users\\huo\\AppData\\Local\\nvs\\node\\10.24.1\\x64")
	app := app.InitApp()
	err := app.Run([]string{global.AppName, global.CmdInstall})
	assert.NoError(t, err)
}

func TestInstall_one(t *testing.T) {
	os.Setenv("BOOK_NODE_HOME", "C:\\Users\\huo\\AppData\\Local\\nvs\\node\\10.24.1\\x64")
	app := app.InitApp()
	err := app.Run([]string{global.AppName, global.CmdInstall, "3-ba"})
	assert.NoError(t, err)
}
