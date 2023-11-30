package main

import (
	"fmt"
	"os"
)

const (
	Name     = "alpen"
	Version  = "0.0.11"
	Revision = "HEAD"
)

func main() {
	if err := newApp().run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
