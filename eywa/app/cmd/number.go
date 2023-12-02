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
	Usage:       "Start to number",
	Description: "Start to number",
	Action:      executeNumber,
	Flags: []cli.Flag{
		utils.StringFlag("type", "bit", "`type` of number bit"),
		utils.StringFlag("num", "100", "data of `num`"),
	},
}

func executeNumber(ctx *cli.Context) error {
	logger.Init(ctx)
	log.Debugf("type:%s", ctx.String("type"))

	n, _ := strconv.ParseUint(ctx.String("num"), 10, 64)
	switch strings.ToLower(ctx.String("type")) {
	case "bit":
		log.Infof("%d", number.BitCount(n))
	default:
		log.Infof("%d", number.BitCount(n))
	}

	return nil
}
