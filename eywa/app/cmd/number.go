package cmd

import (
	"eywa/logger"
	"eywa/number"
	"eywa/utils"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v2"
)

var Number = &cli.Command{
	Name:        "number",
	Usage:       "Start to convert number",
	Description: "Start to convert number",
	Action:      executeNumber,
	Flags: []cli.Flag{
		utils.StringFlag("type", "bit", "`type` of action (bit)"),
		utils.StringFlag("num", "100", "then `number`"),
	},
}

func executeNumber(ctx *cli.Context) error {
	logger.Init(ctx)
	log.Debugf("type:%s", ctx.String("type"))

	n, err := strconv.ParseUint(ctx.String("num"), 10, 64)
	if err != nil {
		return nil
	}
	switch strings.ToLower(ctx.String("type")) {
	case "bit":
		log.Infof("%d", number.BitCount(n))
	default:
		log.Infof("%d", number.BitCount(n))
	}

	return nil
}
