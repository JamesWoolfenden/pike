package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	//nolint:goimports
	pike "github.com/jameswoolfenden/pike/src" //nolint:goimports
	"github.com/urfave/cli/v2"
)

func main() {
	var directory string

	app := &cli.App{
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
			},
			&cli.StringFlag{
				Name:        "directory",
				Aliases:     []string{"D"},
				Usage:       "Directory to scan",
				Destination: &directory,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "scan",
				Aliases: []string{"s"},
				Usage:   "scan",
				Action: func(*cli.Context) error {
					return pike.Scan(directory)
				},
			},
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "version",
				Action: func(*cli.Context) error {
					fmt.Println(pike.Version)
					return nil
				},
			},
		},
		Name:     "pike",
		Usage:    "Generate IAM policy from your IAC code",
		Compiled: time.Time{},
		Authors:  []*cli.Author{{Name: "James Woolfenden", Email: "support@bridgecrew.io"}},
		Version:  pike.Version,
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
