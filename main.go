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
	var branch string
	var directory string
	var file string
	var output string
	var arn string
	var wait int
	var init bool
	var autoAppend bool
	var write bool
	var repository string
	var region string
	var workflow string

	var app = &cli.App{
		EnableBashCompletion: true,
		Flags:                []cli.Flag{},
		Commands: []*cli.Command{
			{
				Name:    "make",
				Aliases: []string{"m"},
				Usage:   "make the policy/role required for this IAC to deploy",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "directory",
						Aliases:     []string{"d"},
						Usage:       "Directory to scan (defaults to .)",
						Value:       ".",
						Destination: &directory,
					},
				},
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
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "directory",
						Aliases:     []string{"d"},
						Usage:       "Directory to scan (defaults to .)",
						Value:       ".",
						Destination: &directory,
					},
					&cli.StringFlag{
						Name:        "region",
						Aliases:     []string{"g"},
						Usage:       "The region",
						DefaultText: "eu-west-2",
						Destination: &region,
					},
				},
				Action: func(*cli.Context) error {
					return pike.Apply(directory, region)
				},
			},
			{
				Name:    "remote",
				Aliases: []string{"m"},
				Usage:   "Create/Update the Policy and set credentials/secret for Github Action",
				Flags: []cli.Flag{
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
						Usage:       "The github repository and owner (to set secrets) e.g. jameswoolfenden/terraform-aws-s3",
						Destination: &repository,
						Required:    true,
					},
					&cli.StringFlag{
						Name:        "region",
						Aliases:     []string{"g"},
						Usage:       "The region",
						DefaultText: "eu-west-2",
						Destination: &region,
					},
				},
				Action: func(*cli.Context) error {
					return pike.Remote(directory, repository, region)
				},
			},
			{
				Name:    "scan",
				Aliases: []string{"s"},
				Usage:   "scan a directory for IAM code",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "directory",
						Aliases:     []string{"d"},
						Usage:       "Directory to scan (defaults to .)",
						Value:       ".",
						Destination: &directory,
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
						Name:        "file",
						Aliases:     []string{"f"},
						Usage:       "File to scan",
						Destination: &file,
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
				},
				Action: func(*cli.Context) error {
					return pike.Scan(directory, output, &file, init, write)
				},
			},
			{
				Name:    "compare",
				Aliases: []string{"c"},
				Usage:   "policy comparison of deployed versus IAC",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "directory",
						Aliases:     []string{"d"},
						Usage:       "Directory to scan (defaults to .)",
						Value:       ".",
						Destination: &directory,
					},
					&cli.StringFlag{
						Name:        "arn",
						Aliases:     []string{"a"},
						Usage:       "Policy identifier e.g. arn",
						Value:       "arn:aws:iam::680235478471:policy/basic",
						Destination: &arn,
						EnvVars:     []string{"ARN"},
					},
					&cli.BoolFlag{
						Name:        "init",
						Aliases:     []string{"i"},
						Usage:       "Run Terraform init to download modules",
						Destination: &init,
					},
				},
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
				Flags: []cli.Flag{
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
				Action: func(*cli.Context) error {
					return pike.Watch(arn, wait)
				},
			},
			{
				Name:    "readme",
				Aliases: []string{"r"},
				Usage:   "Looks in dir for a README.md and updates it with the Policy required to build the code",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "directory",
						Aliases:     []string{"d"},
						Usage:       "Directory to scan (defaults to .)",
						Value:       ".",
						Destination: &directory,
					},
					&cli.StringFlag{
						Name:        "output",
						Aliases:     []string{"o"},
						Usage:       "Output types e.g. `json` terraform",
						Value:       "terraform",
						Destination: &output,
						EnvVars:     []string{"OUTPUT"},
					},
					&cli.BoolFlag{
						Name:        "init",
						Aliases:     []string{"i"},
						Usage:       "Run Terraform init to download modules",
						Destination: &init,
					},
					&cli.BoolFlag{
						Name:        "auto append",
						Aliases:     []string{"A"},
						Usage:       "Automatically adds policy section to the end of Readme",
						Destination: &autoAppend,
					},
				},
				Action: func(*cli.Context) error {
					return pike.Readme(directory, output, init, autoAppend)
				},
			},
			{
				Name:    "invoke",
				Aliases: []string{"i"},
				Usage:   "Triggers a gitHub action specified with the workflow flag",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "repository",
						Aliases:     []string{"r"},
						Usage:       "The github repository and owner (to set secrets) e.g. jameswoolfenden/pike   ",
						Destination: &repository,
					},
					&cli.StringFlag{
						Name:        "workflow",
						Usage:       "Github action workflows filename",
						Value:       "main.yml",
						Destination: &workflow,
					},
					&cli.StringFlag{
						Name:        "branch",
						Aliases:     []string{"b"},
						Value:       "main",
						Destination: &branch,
						Usage:       "The branch to use when invoke is called",
					},
				},
				Action: func(*cli.Context) error {
					return pike.InvokeGithubDispatchEvent(repository, workflow, branch)
				},
			},
			{
				Name:      "version",
				Aliases:   []string{"v"},
				Usage:     "Outputs the application version",
				UsageText: "pike version",
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
