package main

import (
	"os"

	"github.com/gofulljs/gbook/app"
)

func main() {
	if err := app.InitApp().Run(os.Args); err != nil {
		panic(err)
	}
}
