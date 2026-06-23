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
func Readme(dirName string, output string, init bool, autoAppend bool, legacy bool) error {
	file := path.Join(dirName, "README.md")

	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		log.Info().Str("file", file).Msg("no README.md found, skipping")
		return nil
	}

	OutPolicy, err := MakePolicy(dirName, nil, init, false, "", "", false)
	if err != nil {
		var emptyIAC *emptyIACError
		var emptyPerms *emptyPermissionsError
		if errors.As(err, &emptyIAC) || errors.As(err, &emptyPerms) {
			log.Info().Str("dir", dirName).Msg("no IAM permissions found, readme unchanged")
			return nil
		}
		log.Info().Msg("failed to make policy")
		return &makePolicyError{err}
	}

	var rendered string
	if legacy {
		rendered = OutPolicy.AsString(output)
	} else {
		rendered = OutPolicy.AsTwoRoleString(output)
	}

	var markdown string

	switch strings.ToLower(output) {
	case terraform:
		markdown = "\nThe Terraform resource required is:\n\n```golang\n" + rendered + "\n```\n"
	case "json":
		markdown = "\nThe Policy required is:\n\n```json\n" + rendered + "\n```\n"
	default:
		return &tfPolicyFormatError{}
	}

	err = ReplaceSection(file, markdown, autoAppend)
	if err != nil {
		return &replaceSectionError{err}
	}

	log.Info().Msg("readme updated")

	return nil
}

type tfPolicyFormatError struct{}

func (m *tfPolicyFormatError) Error() string {
	return "output formats are Terraform and JSON"
}
