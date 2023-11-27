package pike

import (
	"errors"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

// Readme Updates a README.md file.
func Readme(dirName string, output string, init bool, autoAppend bool) error {
	file := dirName + "/README.md"

	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		return err
	}

	OutPolicy, err2 := MakePolicy(dirName, nil, init, false)
	if err2 != nil {
		log.Print("failed to make policy")

		return err2
	}

	var markdown string

	switch strings.ToLower(output) {
	case terraform:
		markdown = "\nThe Terraform resource required is:\n\n```golang\n" + OutPolicy.AsString(output) + "\n```\n"
	case "json":
		markdown = "\nThe Policy required is:\n\n```json\n" + OutPolicy.AsString(output) + "\n```\n"
	default:
		return errors.New("output formats are terraform or json")
	}

	err := ReplaceSection(file, markdown, autoAppend)

	log.Print("readme updated")

	return err
}
