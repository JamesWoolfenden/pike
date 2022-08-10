package pike

import (
	"errors"
	"os"
)

func Readme(dirname string, output string) error {

	file := dirname + "/README.md"

	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		return err
	}

	Policy, err := MakePolicy(dirname, output)

	markdown := "\nThe Policy required is \n```json\n" + Policy + "\n```\n"
	err = ReplaceSection(file, markdown)

	return err
}
