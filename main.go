package main

import (
	"fmt"
	"os"
)

const Name = "alpen"

func main() {
	if err := newApp().cli.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("%s: %w", Name, err))
		os.Exit(1)
	}
}
