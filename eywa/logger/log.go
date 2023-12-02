package logger

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v2"
)

const TimeFormat = "2006/01/02 15:04:05"

func Init(ctx *cli.Context) {
	logger := log.NewWithOptions(os.Stdout, log.Options{
		Level:           log.ParseLevel(ctx.String("level")),
		ReportCaller:    ctx.Bool("report"),
		ReportTimestamp: false,
		Prefix:          "Eywa 🍪 ",
	})

	log.SetDefault(logger)
}
