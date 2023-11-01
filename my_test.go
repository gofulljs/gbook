package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/gofulljs/gbook/util"
	"github.com/stretchr/testify/assert"
)

// 测试环境变量生效
func Test00(t *testing.T) {
	path := os.Getenv("PATH")
	fmt.Printf("path:%s\n", path)

	util.Active("C:\\Users\\huo\\AppData\\Local\\nvs\\node\\10.24.1\\x64")
	fmt.Println()
	fmt.Printf("path:%s\n", os.Getenv("PATH"))
	cmd := exec.Command("node", "-v")
	output, err := cmd.Output()
	assert.NoError(t, err)
	fmt.Printf("output:%s\n", output)
}
