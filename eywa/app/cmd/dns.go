package cmd

import (
	"fmt"
	"log"

	"eywa/dns"
	"eywa/utils"

	"github.com/urfave/cli/v2"
)

var Dns = &cli.Command{
	Name:        "dns",
	Usage:       "Start to run dns",
	Description: "Start to run dns",
	Action:      ExecuteDns,
	Flags: []cli.Flag{
		utils.StringFlag("domain", "d", "www.github.com", "name of `domain`"),
		utils.StringFlag("server", "s", "8.8.8.8:53", "host of dns `server`"),
		utils.IntFlag("timeout", "t", 5, "`timeout` of dns query"),
		utils.BoolFlag("fqdn", "f", false, "is it `fqdn` query"),
	},
}

func ExecuteDns(ctx *cli.Context) error {
	switch ctx.Bool("fqdn") {
	case true:
		return fqdnQuery(ctx)
	default:
		return dnsQuery(ctx)
	}
}

func dnsQuery(ctx *cli.Context) error {
	fmt.Printf("--- dns --- %s: %s\n", ctx.String("server"), ctx.String("domain"))

	cname, addrs, err := dns.Dns(ctx.String("server"), ctx.String("domain"), ctx.Int("timeout"))
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	fmt.Printf("cname: %s\n", cname)
	for _, addr := range addrs {
		fmt.Printf("addr: %s\n", addr)
	}

	return nil
}

func fqdnQuery(ctx *cli.Context) error {
	fmt.Printf("--- fqdn --- %s: %s\n", ctx.String("server"), ctx.String("domain"))

	addrs, err := dns.Fqdn(ctx.String("server"), ctx.String("domain"), ctx.Int("timeout"))
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	for _, addr := range addrs {
		fmt.Printf("addr: %s\n", addr)
	}

	return nil
}
