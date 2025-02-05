package cmd

import (
	"strings"

	"eywa/logger"
	"eywa/text"
	"eywa/utils"

	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v2"
)

var Text = &cli.Command{
	Name:        "text",
	Usage:       "Start to convert text",
	Description: "Start to convert text",
	Action:      executeText,
	Flags: []cli.Flag{
		utils.StringFlag("type", "t", "title", "`type` of action (upper|lower|title|sub|pinyin|s2t)"),
	},
}

func executeText(ctx *cli.Context) error {
	logger.Init(ctx)
	log.Debugf("type: %s", ctx.String("type"))

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
	case "sub":
		log.Infof("%s %s %s", data, ctx.Args().Get(1), ctx.Args().Get(2))
		if ctx.NArg() < 3 {
			log.Errorf("ctx.NArg() is less than 3(%d)", ctx.NArg())
			return nil
		}
		log.Infof("%s", strings.ReplaceAll(data, ctx.Args().Get(1), ctx.Args().Get(2)))
	case "pinyin":
		log.Infof("%s\nINFO %s: %s", text.Pinyin(data), log.Default().GetPrefix(), data)
	case "s2t":
		log.Infof("%s", text.S2T(data))
	default:
		log.Infof("%s", text.Title(data))
	}

	return nil
}
