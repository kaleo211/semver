package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/kaleo211/semver/api/naive"
)

func main() {

	app := cli.NewApp()
	app.Name = "semver"

	var level string

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "increment, i",
			Value:       "patch",
			Destination: &level,
		},
	}

	app.Action = func(c *cli.Context) error {
		semver, err := naive.NewSemver(c.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}

		semver.Increment(level)
		fmt.Println(semver.Version())
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
