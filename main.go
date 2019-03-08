package main

import (
	"os"

	"github.com/maxverro/stockbeat/cmd"

	_ "github.com/maxverro/stockbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
