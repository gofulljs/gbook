package readytest

import (
	"testing"

	"github.com/gofulljs/gbook/app"
	"github.com/gofulljs/gbook/global"
	"github.com/stretchr/testify/assert"
)

func TestSync(t *testing.T) {
	app := app.InitApp()
	err := app.Run([]string{global.AppName, "ready"})
	assert.NoError(t, err)
}
