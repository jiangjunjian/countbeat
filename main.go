package main

import (
	"os"

	"github.com/jiangjunjian/countbeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
