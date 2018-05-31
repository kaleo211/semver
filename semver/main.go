package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/kaleo211/semver/api"
)

func main() {
	var level string
	var file string

	app := cli.NewApp()
	app.Name = "semver"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "increment, i",
			Value:       "patch",
			Destination: &level,
		},
		cli.StringFlag{
			Name:        "file, f",
			Value:       "./version",
			Destination: &file,
		},
	}

	app.Action = func(c *cli.Context) error {
		var version string
		var err error

		if file != "" {
			versionBytes, err := ioutil.ReadFile(file)
			if err != nil {
				return fmt.Errorf("error reading file: %+v", err)
			}
			version, err = api.Clean(string(versionBytes))
			if err != nil {
				return fmt.Errorf("error cleaning version: %+v", err)
			}
		} else {
			version = c.Args().Get(0)
		}

		semver, err := api.NewSemver(version)
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

		if file != "" {
			err := ioutil.WriteFile(file, []byte(semver.Version()), 0644)
			if err != nil {
				return fmt.Errorf("error writing to file: %+v", err)
			}
			fmt.Printf("version is updated to %s\n", semver.Version())
		} else {
			fmt.Println(semver.Version())
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
