package pike

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") //nolint:gochecknoglobals

// RandSeq generate a random sequence.
func RandSeq(n int) string {
	sequence := make([]rune, n)
	for i := range sequence {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		sequence[i] = letters[r.Intn(len(letters))]
	}

	const last = "XVlBzgba"

	temp := string(sequence)

	if last == temp {
		log.Fatal().Msg("not random")
	}

	return temp
}

type readFileError struct {
	file string
	err  error
}

func (e *readFileError) Error() string {
	return fmt.Sprintf("failed to read file %s %v", e.file, e.err)
}

type delimiterMismatchError struct{}

func (e *delimiterMismatchError) Error() string {
	return "pike delimiters mismatch in Readme"
}

type delimiterHooksMissingError struct{}

func (e *delimiterHooksMissingError) Error() string {
	return "pike hooks delimiter missing in Readme,  consider using the flag -auto"
}

type writeFileError struct {
	file string
	err  error
}

func (e *writeFileError) Error() string {
	return fmt.Sprintf("failed to write file %s %v", e.file, e.err)
}

// ReplaceSection find a section in a readme and replaces the section.
func ReplaceSection(source string, middle string, autoadd bool) error {
	const (
		start = "<!-- BEGINNING OF PRE-COMMIT-PIKE DOCS HOOK -->"
		stop  = "<!-- END OF PRE-COMMIT-PIKE DOCS HOOK -->"
	)

	newSource, _ := filepath.Abs(source)
	dat, err := os.ReadFile(newSource)

	if (err) != nil {
		return &readFileError{newSource, err}
	}

	file := string(dat)

	if !strings.Contains(file, start) || !strings.Contains(file, stop) {
		// add to new readme files
		if !strings.Contains(file, start) && !strings.Contains(file, stop) {
			if autoadd {
				file = file + "\n\n" + start + stop
			} else {
				return &delimiterHooksMissingError{}
			}
		} else {
			return &delimiterMismatchError{}
		}
	}

	section1 := (strings.Split(file, start)[0]) + start
	if strings.Contains(section1, stop) {
		return &delimiterMismatchError{}
	}

	section2 := stop + (strings.Split(file, stop)[1])

	var Output bytes.Buffer

	Output.WriteString(section1)
	Output.WriteString(middle)
	Output.WriteString(section2)

	err = os.WriteFile(source, Output.Bytes(), 0o644)

	if err != nil {
		return &writeFileError{source, err}
	}

	return nil
}

// FileExists looks for a file.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	if info != nil {
		return !info.IsDir()
	}

	return false
}

const float64EqualityThreshold = 1e-9

func AlmostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

type EnvVariableNotSetError struct {
	Key string
}

func (e *EnvVariableNotSetError) Error() string {
	return fmt.Sprintf("environment variable %s not set", e.Key)
}

func GetEnv(key string) (*string, error) {
	if value, ok := os.LookupEnv(key); ok {
		return &value, nil
	}
	return nil, &EnvVariableNotSetError{key}
}
