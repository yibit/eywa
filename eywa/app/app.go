package app

import (
	"eywa/app/cmd"

	"github.com/urfave/cli/v2"
)

func Run(version string, args []string) {
	app := cli.NewApp()
	app.Name = "eywa"
	app.Version = version
	app.Usage = "a toolkit for hacking"
	app.Commands = []*cli.Command{cmd.Differ, cmd.Text, cmd.Crypto, cmd.Encoding}

	_ = app.Run(args)
}
