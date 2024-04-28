package utils

import (
	"github.com/urfave/cli/v2"
)

func StringFlag(name, aliases, value, usage string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:    name,
		Aliases: Alias(aliases),
		Value:   value,
		Usage:   usage,
	}
}

func StringSliceFlag(name, aliases string, values []string, usage string) *cli.StringSliceFlag {
	return &cli.StringSliceFlag{
		Name:    name,
		Aliases: Alias(aliases),
		Value:   cli.NewStringSlice(values...),
		Usage:   usage,
	}
}

func IntFlag(name, aliases string, value int, usage string) *cli.IntFlag {
	return &cli.IntFlag{
		Name:    name,
		Aliases: Alias(aliases),
		Value:   value,
		Usage:   usage,
	}
}

func Int64Flag(name, aliases string, value int64, usage string) *cli.Int64Flag {
	return &cli.Int64Flag{
		Name:    name,
		Aliases: Alias(aliases),
		Value:   value,
		Usage:   usage,
	}
}

func Uint64Flag(name, aliases string, value uint64, usage string) *cli.Uint64Flag {
	return &cli.Uint64Flag{
		Name:    name,
		Aliases: Alias(aliases),
		Value:   value,
		Usage:   usage,
	}
}

func IntSliceFlag(name, aliases string, value int, usage string) *cli.IntSliceFlag {
	return &cli.IntSliceFlag{
		Name:    name,
		Aliases: Alias(aliases),
		Value:   cli.NewIntSlice(value),
		Usage:   usage,
	}
}

func Int64SliceFlag(name, aliases string, value int64, usage string) *cli.Int64SliceFlag {
	return &cli.Int64SliceFlag{
		Name:    name,
		Aliases: Alias(aliases),
		Value:   cli.NewInt64Slice(value),
		Usage:   usage,
	}
}

func BoolFlag(name, aliases string, value bool, usage string) *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:    name,
		Aliases: Alias(aliases),
		Value:   value,
		Usage:   usage,
	}
}

func Alias(name string) []string {
	return []string{name}
}
