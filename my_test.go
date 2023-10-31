package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test00(t *testing.T) {
	path := os.Getenv("PATH")
	fmt.Printf("path:%s\n", path)
	os.Setenv("PATH", path+";"+"E:\\Scoop\\apps\\nvs\\current\\nodejs\\node\\10.24.1\\x64")

	cmd := exec.Command("node", "-v")
	output, err := cmd.Output()
	assert.NoError(t, err)
	fmt.Printf("output:%s\n", output)
}
