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
	switch strings.ToLower(ctx.String("type")) {
	case "http":
		s := echo.HttpD{}
		s.Start(ctx.String("port"))
	case "tcp":
		s := echo.TCP{}
		s.Start(ctx.String("port"))
	default:
		s := echo.UDP{}
		s.Start(ctx.String("port"))
	}

	return nil
}
