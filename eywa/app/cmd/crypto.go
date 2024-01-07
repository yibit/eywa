package cmd

import (
	"encoding/hex"
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

	switch strings.ToLower(ctx.String("type")) {
	case "md5":
		log.Infof("%s", crypto.MD5(ctx.Args().Get(0)))
	case "aes":
		if ctx.NArg() != 4 {
			log.Errorf("ctx.NArg() != 4")
			return nil
		}
		log.Infof("%s", hex.EncodeToString(crypto.EncryptAES(ctx.Args().Get(0), ctx.Args().Get(1), []byte(ctx.Args().Get(2)), ctx.Args().Get(3))))
	default:
		if ctx.NArg() != 2 {
			log.Errorf("ctx.NArg() != 2")
			return nil
		}
		data, err := crypto.EncryptDES([]byte(ctx.Args().Get(0)), []byte(ctx.Args().Get(1)))
		if err == nil {
			log.Infof("%s", hex.EncodeToString(data))
		}
	}

	return nil
}
