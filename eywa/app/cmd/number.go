package cmd

import (
	"eywa/logger"
	"eywa/number"
	"eywa/utils"
	"math/bits"
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
		utils.StringFlag("type", "bits", "`type` of action (ones|bits)"),
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
	case "ones":
		log.Infof("%d", number.OnesCount(n))
	case "bits":
		log.Infof("%s", number.Bits(n))
	default:
		log.Infof("%b:%d", n, bits.OnesCount(uint(n)))
	}

	return nil
}
