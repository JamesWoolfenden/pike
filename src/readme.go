package pike

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/rs/zerolog/log"
)

type replaceSectionError struct {
	err error
}

type fileDoesNotExistError struct {
	file string
	err  error
}

func (e fileDoesNotExistError) Error() string {
	return fmt.Sprintf("file %s does not exist %v", e.file, e.err)
}

func (m *replaceSectionError) Error() string {
	return fmt.Sprintf("failed to replace section %v", m.err)
}

// Readme Updates a README.md file.
func Readme(dirName string, output string, init bool, autoAppend bool) error {
	file := path.Join(dirName, "README.md")

	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		return &fileDoesNotExistError{file, err}
	}

	OutPolicy, err := MakePolicy(dirName, nil, init, false, "", "")
	if err != nil {
		log.Info().Msg("failed to make policy")

		return &makePolicyError{err}
	}

	var markdown string

	switch strings.ToLower(output) {
	case terraform:
		markdown = "\nThe Terraform resource required is:\n\n```golang\n" + OutPolicy.AsString(output) + "\n```\n"
	case "json":
		markdown = "\nThe Policy required is:\n\n```json\n" + OutPolicy.AsString(output) + "\n```\n"
	default:
		return &tfPolicyFormatError{}
	}

	err = ReplaceSection(file, markdown, autoAppend)
	if err != nil {
		return &replaceSectionError{err}
	}

	log.Info().Msg("readme updated")

	return err
}

type tfPolicyFormatError struct{}

func (m *tfPolicyFormatError) Error() string {
	return "output formats are Terraform and JSON"
}
