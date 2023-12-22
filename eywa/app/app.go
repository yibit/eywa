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
	app.Commands = []*cli.Command{cmd.Crypto, cmd.Differ, cmd.Dns,
		cmd.Encoding, cmd.Echo, cmd.Format, cmd.Number, cmd.Text}

	_ = app.Run(args)
}
