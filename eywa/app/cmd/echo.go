package cmd

import (
	"eywa/echo"
	"eywa/utils"
	"strings"

	"github.com/urfave/cli/v2"
)

var Echo = &cli.Command{
	Name:        "echo",
	Usage:       "Start to serve echo",
	Description: "Start to serve echo",
	Action:      ExecuteEcho,
	Flags: []cli.Flag{
		utils.StringFlag("type", "http", "`type` of echo server protocol (http|tcp|udp)"),
		utils.StringFlag("port", "8080", "`port` of server"),
	},
}

func ExecuteEcho(ctx *cli.Context) error {
	var r echo.Echoer

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
