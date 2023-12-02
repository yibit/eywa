package cmd

import (
	"eywa/crypto"
	"eywa/logger"
	"eywa/utils"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v2"
)

var Crypto = &cli.Command{
	Name:        "crypto",
	Usage:       "Start to crypto",
	Description: "Start to crypto",
	Action:      executeCrypto,
	Flags: []cli.Flag{
		utils.StringFlag("type", "md5", "`type` of crypto (aes|des|md5)"),
	},
}

func executeCrypto(ctx *cli.Context) error {
	logger.Init(ctx)
	log.Debugf("type:%s", ctx.String("type"))

	if ctx.NArg() < 1 {
		log.Errorf("ctx.NArg() is less than 2(%d)", ctx.NArg())
		return nil
	}

	data := ctx.Args().Get(0)
	switch strings.ToLower(ctx.String("type")) {
	case "md5":
		log.Infof("%s", crypto.MD5(data))
	default:
		log.Infof("%s", crypto.MD5(data))
	}

	return nil
}
