package main

import (
	"os"

	"github.com/joseph0x45/argus/internal/cli"
	"github.com/joseph0x45/goutils"
)

func main() {
  goutils.SetAppName("argus")
	os.Exit(cli.DispatchCommands(os.Args))
}
