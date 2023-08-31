// The main package is the command line runner
package main

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/oalders/is/types"
)

func main() {
	var API struct {
		Debug   bool             `help:"turn on debugging statements"`
		OS      OSCmd            `cmd:"" help:"Check OS attributes. e.g. \"is os name eq darwin\""`
		CLI     CLICmd           `cmd:"" help:"Check cli version. e.g. \"is cli version tmux gte 3\""`
		Known   KnownCmd         `cmd:""`
		There   ThereCmd         `cmd:"" help:"Check if command exists. e.g. \"is there git\""`
		User    UserCmd          `cmd:"" help:"Info about current user. e.g. \"is user sudoer\""`
		Version kong.VersionFlag `help:"Print version to screen"`
	}

	ctx := kong.Parse(&API,
		kong.Vars{
			"version": "0.1.1",
		})
	runContext := types.Context{Debug: API.Debug}
	err := ctx.Run(&runContext)
	ctx.FatalIfErrorf(err)

	if runContext.Success {
		os.Exit(0)
	}
	os.Exit(1)
}
