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
	var file string
	var output string
	var arn string
	var wait int
	var init bool
	var autoAppend bool
	var write bool
	var repository string
	var owner string

	var app = &cli.App{
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
			},
			&cli.BoolFlag{
				Name:        "init",
				Aliases:     []string{"i"},
				Usage:       "Run Terraform init to download modules",
				Destination: &init,
			},
			&cli.BoolFlag{
				Name:        "write",
				Aliases:     []string{"w"},
				Usage:       "Write the policy output to a file at .pike",
				Destination: &write,
			},
			&cli.BoolFlag{
				Name:        "auto append",
				Aliases:     []string{"A"},
				Usage:       "Automatically adds policy section to the end of Readme",
				Destination: &autoAppend,
			},
			&cli.StringFlag{
				Name:        "directory",
				Aliases:     []string{"d"},
				Usage:       "Directory to scan (defaults to .)",
				Value:       ".",
				Destination: &directory,
			},
			&cli.StringFlag{
				Name:        "repository",
				Aliases:     []string{"r"},
				Usage:       "The github repository (to set secrets)",
				Destination: &repository,
			},
			&cli.StringFlag{
				Name:        "file",
				Aliases:     []string{"f"},
				Usage:       "File to scan",
				Destination: &file,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Usage:       "Output types e.g. `json` terraform",
				Value:       "terraform",
				Destination: &output,
				EnvVars:     []string{"OUTPUT"},
			},
			&cli.StringFlag{
				Name:        "owner",
				Aliases:     []string{"n"},
				Usage:       "The Owner of the Github repo",
				Destination: &owner,
			},
			&cli.StringFlag{
				Name:        "arn",
				Aliases:     []string{"a"},
				Usage:       "Policy identifier e.g. arn",
				Value:       "arn:aws:iam::680235478471:policy/basic",
				Destination: &arn,
				EnvVars:     []string{"ARN"},
			},
			&cli.IntFlag{
				Name:        "wait",
				Aliases:     []string{"W"},
				Value:       100,
				Usage:       "Time to wait for policy change (in tenths of seconds)",
				Destination: &wait,
				EnvVars:     []string{"WAIT"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "make",
				Aliases: []string{"m"},
				Usage:   "make the policy/role required for this IAC to deploy",
				Action: func(*cli.Context) error {
					arn, err := pike.Make(directory)
					log.Print(*arn)
					return err
				},
			},
			{
				Name:    "apply",
				Aliases: []string{"a"},
				Usage:   "Create a policy and use it to instantiate the IAC",
				Action: func(*cli.Context) error {
					return pike.Apply(directory)
				},
			},
			{
				Name:    "remote",
				Aliases: []string{"r"},
				Usage:   "Create/Update the Policy and set credentials/secret for Github Action",

				Action: func(*cli.Context) error {
					if owner == "" {
						log.Fatal("owner details required")
					}
					if repository == "" {
						log.Fatal("repository must be set for Remote")
					}
					return pike.Remote(directory, owner, repository)
				},
			},
			{
				Name:    "scan",
				Aliases: []string{"s"},
				Usage:   "scan a directory for IAM code",
				Action: func(*cli.Context) error {
					return pike.Scan(directory, output, file, init, write)
				},
			},
			{
				Name:    "compare",
				Aliases: []string{"c"},
				Usage:   "policy comparison of deployed versus IAC",
				Action: func(*cli.Context) error {
					theSame, err := pike.Compare(directory, arn, init)
					log.Print("The same: ", theSame)
					return err
				},
			},
			{
				Name:    "watch",
				Aliases: []string{"w"},
				Usage:   "Waits for policy update",
				Action: func(*cli.Context) error {
					return pike.Watch(arn, wait)
				},
			},
			{
				Name:    "readme",
				Aliases: []string{"r"},
				Usage:   "Looks in dir for a README.md and updates it with the Policy required to build the code",
				Action: func(*cli.Context) error {
					return pike.Readme(directory, output, init, autoAppend)
				},
			},
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Outputs the application version",
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
