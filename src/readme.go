package pike

import (
	"errors"
	"os"
)

// Readme Updates a README.md file
func Readme(dirname string, output string) error {

	file := dirname + "/README.md"

	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		return err
	}

	Policy, err2 := MakePolicy(dirname, output, "")
	if err2 != nil {
		return err2
	}
	markdown := "\nThe Policy required is:\n\n```json\n" + Policy + "\n```\n"
	err := ReplaceSection(file, markdown)

	return err
}
