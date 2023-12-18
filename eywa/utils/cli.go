package utils

import (
	"github.com/urfave/cli/v2"
)

func StringFlag(name, value, usage string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:    name,
		Aliases: Alias(name),
		Value:   value,
		Usage:   usage,
	}
}

func IntFlag(name string, value int, usage string) *cli.IntFlag {
	return &cli.IntFlag{
		Name:    name,
		Aliases: Alias(name),
		Value:   value,
		Usage:   usage,
	}
}

func BoolFlag(name string, value bool, usage string) *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:    name,
		Aliases: Alias(name),
		Value:   value,
		Usage:   usage,
	}
}

func Alias(name string) []string {
	return []string{string(name[0])}
}
