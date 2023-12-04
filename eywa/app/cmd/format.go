package cmd

import (
	"eywa/format"
	"eywa/logger"
	"eywa/utils"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v2"
)

var Format = &cli.Command{
	Name:        "format",
	Usage:       "Start to format",
	Description: "Start to format",
	Action:      executeFormat,
	Flags: []cli.Flag{
		utils.StringFlag("type", "json", "`type` of action (json|xml)"),
	},
}

func executeFormat(ctx *cli.Context) error {
	logger.Init(ctx)
	log.Debugf("type:%s", ctx.String("type"))

	if ctx.NArg() < 1 {
		log.Errorf("ctx.NArg() is less than 1(%d)", ctx.NArg())
		return nil
	}

	data := ctx.Args().Get(0)
	switch strings.ToLower(ctx.String("type")) {
	case "json":
		log.Infof("\n%s", format.Json(data))
	default:
		log.Infof("\n%s", format.Json(data))
	}

	return nil
}
