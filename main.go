package main

import (
	"fmt"
	"github.com/jameswoolfenden/pike/src"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

func main() {
	var language string
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
					return scan(directory)
				},
			},
		},
		Name:  "pike",
		Usage: "Generate IAM policy from your IAC code",
		Action: func(cCtx *cli.Context) error {
			name := "someone"
			if cCtx.NArg() > 0 {
				name = cCtx.Args().Get(0)
			}
			if language == "spanish" {
				fmt.Println("Hola", name)
			} else {
				fmt.Println("Hello", name)
			}
			return nil
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func scan(dirname string) error {
	files, err := ioutil.ReadDir(dirname)

	if err != nil {
		return err
	}

	var results []template

	for _, file := range files {

		resources, filename, code := pike.GetResources(file, dirname)
		for _, resource := range resources {
			myAPI := pike.GetAPI(resource)
			result := template{resource, myAPI, filename, code}
			results = append(results, result)
		}
	}

	log.Print(results)
	return nil
}

type template struct {
	resource string
	api      string
	filename string
	code     string
}
