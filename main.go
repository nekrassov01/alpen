package main

import (
	"fmt"
	"os"
)

func main() {
	if err := newApp().run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
