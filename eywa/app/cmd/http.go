package cmd

import (
	"eywa/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/urfave/cli/v2"
)

var HTTP = &cli.Command{
	Name:        "http",
	Usage:       "Show status of HTTP code",
	Description: "Show status of HTTP code",
	Action:      ExecuteHTTP,
	Flags: []cli.Flag{
		utils.IntFlag("code", 200, "`code` of status"),
		utils.BoolFlag("list", false, "`list` all of status"),
	},
}

func ExecuteHTTP(ctx *cli.Context) error {
	statusCode := []int{100, 103, 200, 208, 226, 226, 300, 308, 400, 418, 421, 426, 428, 429, 431, 431, 451, 451, 500, 508, 510, 511}

	var str string
	switch ctx.Bool("list") {
	case true:
		for i := 0; i < len(statusCode); i += 2 {
			for k := statusCode[i]; k <= statusCode[i+1]; k++ {
				str += strings.TrimSpace(fmt.Sprintf("%d %s", k, http.StatusText(k))) + "\n"
			}
			str += "\n"
		}
	default:
		str += strings.TrimSpace(fmt.Sprintf("%d %s", ctx.Int("code"), http.StatusText(ctx.Int("code")))) + "\n"
	}

	fmt.Printf("%s", str)

	return nil
}
