package pike

import (
	"errors"
	"os"

	"github.com/urfave/cli/v2"
)

// Readme Updates a README.md file
func Readme(dirName string, output string, init bool, autoAppend bool, excludes *cli.StringSlice) error {

	file := dirName + "/README.md"

	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		return err
	}

	Policy, err2 := MakePolicy(dirName, output, "", init, excludes)
	if err2 != nil {
		return err2
	}

	var markdown string
	if init {
		markdown = "\nThe Terraform resource required is:\n\n```golang\n" + Policy + "\n```\n"
	} else {
		markdown = "\nThe Policy required is:\n\n```json\n" + Policy + "\n```\n"
	}
	err := ReplaceSection(file, markdown, autoAppend)

	return err
}
