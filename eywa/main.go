package main

import (
	"os"

	"eywa/app"
)

var Version string

func main() {
	app.Run(Version, os.Args)
}
