package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "semver"

	var version string
	var level string

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "increment, i",
			Value:       "patch",
			Destination: &level,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
