package cmd

import (
	"io"
	"os"
	"strings"

	"eywa/differ"
	"eywa/logger"
	"eywa/utils"

	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v2"
)

var Differ = &cli.Command{
	Name:        "diff",
	Usage:       "Start to diff files",
	Description: "Start to diff files",
	Action:      executeDiff,
	Flags: []cli.Flag{
		utils.StringFlag("type", "t", "json", "`type` of file to diff in (json|yaml|text|xml)"),
		utils.StringFlag("level", "l", "info", "`level` of logger in (debug|info|warn|error)"),
	},
}

func executeDiff(ctx *cli.Context) error {
	logger.Init(ctx)
	log.Debugf("type:%s", ctx.String("type"))

	if ctx.NArg() < 2 {
		log.Errorf("ctx.NArg() is less than 2(%d)", ctx.NArg())
		return nil
	}

	x, err := readData(ctx.Args().Get(0))
	if err != nil {
		return err
	}
	y, err := readData(ctx.Args().Get(1))
	if err != nil {
		return err
	}

	differ.Diff(strings.ToLower(ctx.String("type")), x, y)

	return nil
}

func readData(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return data, nil
}
