package main

import (
	"fmt"
	"os"

	"github.com/gofulljs/gbook/cmds/app"
)

func main() {
	if err := app.InitApp().Run(os.Args); err != nil {
		if app.LogDetail {
			fmt.Printf("%+v\n", err)
		} else {
			fmt.Println(err)
		}
	}
}
