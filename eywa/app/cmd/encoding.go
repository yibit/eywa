package cmd

import (
	"eywa/encoding"
	"eywa/logger"
	"eywa/utils"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v2"
)

var Encoding = &cli.Command{
	Name:        "encoding",
	Usage:       "Start to encoding",
	Description: "Start to encoding",
	Action:      executeEncoding,
	Flags: []cli.Flag{
		utils.StringFlag("type", "base64", "`type` of crypto (base64|uuid|random)"),
		utils.StringFlag("mode", "decoding", "`type` of crypto (encoding|decoding)"),
		utils.StringFlag("num", "8", "`num` of length"),
	},
}

func executeEncoding(ctx *cli.Context) error {
	logger.Init(ctx)
	log.Debugf("type:%s", ctx.String("type"))

	if ctx.NArg() < 1 && ctx.String("type") != "uuid" && ctx.String("type") != "random" {
		log.Errorf("ctx.NArg() is less than 2(%d)", ctx.NArg())
		return nil
	}

	data := ctx.Args().Get(0)
	switch strings.ToLower(ctx.String("type")) {
	case "base64":
		if ctx.String("mode") == "encoding" {
			log.Infof("%s", encoding.EncodeToString(data))
		} else {
			log.Infof("%s", encoding.DecodeString(data))
		}
	case "uuid":
		log.Infof("%s", utils.UUID())
	case "random":
		n, _ := strconv.ParseUint(ctx.String("num"), 10, 64)
		log.Infof("%s", utils.RandomString(int(n)))
	default:
		log.Infof("%s", encoding.DecodeString(data))
	}

	return nil
}
