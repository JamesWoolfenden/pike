package main

import (
	"github.com/jameswoolfenden/pike/src" //nolint:goimports
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
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
		},
		Name:  "pike",
		Usage: "Generate IAM policy from your IAC code",
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
