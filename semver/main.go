package main

import (
	"errors"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/kaleo211/semver/api"
)

func main() {
	var level string

	app := cli.NewApp()
	app.Name = "semver"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "increment, i",
			Value:       "patch",
			Destination: &level,
		},
	}

	app.Action = func(c *cli.Context) error {
		semver, err := api.NewSemver(c.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}

		switch level {
		case "patch":
			semver.IncPatch()
		case "minor":
			semver.IncMinor()
		case "major":
			semver.IncMajor()
		default:
			return errors.New("release type is not valide")
		}

		fmt.Println(semver.Version())
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
