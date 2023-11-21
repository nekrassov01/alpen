package main

import (
	"fmt"
	"os"
)

const (
	Name     = "alpen"
	Version  = "0.0.9"
	Revision = "HEAD"
)

func main() {
	app := newApp()
	if err := app.run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
