package cmd

import (
	"strings"

	"eywa/curl"
	"eywa/logger"
	"eywa/utils"

	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v2"
)

var Curl = &cli.Command{
	Name:        "curl",
	Usage:       "Simple style of curl",
	Description: "Simple style of curl",
	Action:      executeCurl,
	Flags: []cli.Flag{
		utils.StringFlag("method", "X", "GET", "type of `method` (POST|GET|PUT|DELETE|OPTIONS)"),
		utils.StringSliceFlag("header", "H", []string{"Content-Type:application/json;charset=UTF-8"}, "fields of `headers`"),
		utils.StringFlag("curl", "c", "", "curl of `object`"),
		utils.StringFlag("data", "d", "", "data of `body`"),
	},
}

func executeCurl(ctx *cli.Context) error {
	logger.Init(ctx)
	log.Debugf("method:%s", ctx.String("method"))

	switch strings.ToUpper(ctx.String("method")) {
	case "POST":
		fallthrough
	case "GET":
		fallthrough
	case "PUT":
		fallthrough
	case "DELETE":
		curl.InvokeYService(ctx.String("method"), ctx.String("curl"), ctx.StringSlice("header"), ctx.String("data"))
	default:
		log.Errorf("%s not supported!", ctx.String("method"))
	}

	return nil
}
