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
		utils.StringFlag("type", "base64", "`type` of encoding (number|uuid|random|hex)"),
		utils.StringFlag("mode", "decoding", "`mode` of encoding (encoding|decoding)"),
		utils.StringFlag("num", "8", "`length` of string"),
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
	case "hex":
		if ctx.String("mode") == "encoding" {
			log.Infof("%s", encoding.HexEncode(data))
		} else {
			log.Infof("%s", encoding.HexDecode(data))
		}
	case "random":
		n, _ := strconv.ParseUint(ctx.String("num"), 10, 64)
		log.Infof("%s", utils.RandomString(int(n)))
	case "uuid":
		log.Infof("%s", utils.UUID())
	default:
		log.Infof("%s", encoding.DecodeString(data))
	}

	return nil
}
