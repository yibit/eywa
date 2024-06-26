package cmd

import (
	"strings"

	"eywa/echo"
	"eywa/utils"

	"github.com/urfave/cli/v2"
)

var Echo = &cli.Command{
	Name:        "echo",
	Usage:       "Start to serve echo",
	Description: "Start to serve echo",
	Action:      ExecuteEcho,
	Flags: []cli.Flag{
		utils.StringFlag("type", "t", "http", "`type` of echo server protocol (http|tcp|udp)"),
		utils.StringFlag("port", "p", "1037", "`port` of server"),
	},
}

func ExecuteEcho(ctx *cli.Context) error {
	var r echo.Echo

	switch strings.ToLower(ctx.String("type")) {
	case "http":
		r = echo.HttpD{}
	case "tcp":
		r = echo.TCP{}
	default:
		r = echo.UDP{}
	}

	r.Start(ctx.String("port"))

	return nil
}
