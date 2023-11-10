package main

import (
	"fmt"
	"os"
)

const (
	Name     = "alpen"
	Version  = "0.0.7"
	Revision = "HEAD"
)

func main() {
	app := NewApp()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
