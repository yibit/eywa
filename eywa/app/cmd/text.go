package cmd

import (
	"eywa/logger"
	"eywa/text"
	"eywa/utils"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v2"
)

var Text = &cli.Command{
	Name:        "text",
	Usage:       "Start to convert text",
	Description: "Start to convert text",
	Action:      executeText,
	Flags: []cli.Flag{
		utils.StringFlag("type", "title", "`type` of action (upper|lower|title)"),
	},
}

func executeText(ctx *cli.Context) error {
	logger.Init(ctx)
	log.Debugf("type:%s", ctx.String("type"))

	if ctx.NArg() < 1 {
		log.Errorf("ctx.NArg() is less than 1(%d)", ctx.NArg())
		return nil
	}

	data := ctx.Args().Get(0)
	switch strings.ToLower(ctx.String("type")) {
	case "upper":
		log.Infof("%s", text.Upper(data))
	case "lower":
		log.Infof("%s", text.Lower(data))
	default:
		log.Infof("%s", text.Title(data))
	}

	return nil
}
