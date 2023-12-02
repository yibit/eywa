package main

import (
	"eywa/app"
	"os"
)

var Version string

func main() {
	app.Run(Version, os.Args)
}
