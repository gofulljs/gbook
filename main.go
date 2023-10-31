package main

import (
	"fmt"
	"os"

	"github.com/gofulljs/gbook/app"
)

func main() {
	if err := app.InitApp().Run(os.Args); err != nil {
		fmt.Printf("%+v\n", err)
	}
}
