package main

import (
	"os"

	"github.com/codegangsta/cli"
)

const version string = "0.1.0"

func main() {
	app := cli.NewApp()
	app.Name = "faw"
	app.Version = version
	app.Usage = "Font Awesome Workflow for Alfred"
	app.Author = "ruedap"
	app.Email = "ruedap@ruedap.com"
	app.Commands = commands
	app.Run(os.Args)
}
