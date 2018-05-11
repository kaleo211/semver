package main

import (
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "semver"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "increment",
			Value: "patch",
			Destination: "",
		}
	}
}
