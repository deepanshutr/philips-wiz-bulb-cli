package main

import (
	"fmt"
	"os"

	"github.com/deepanshutr/philips-wiz-bulb-cli/internal/cli"
)

func main() {
	if err := cli.NewRoot().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
