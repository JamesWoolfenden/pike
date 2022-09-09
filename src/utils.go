package pike

import (
	"bytes"
	"errors"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// ReplaceSection find a section in a readme and replaces the section
func ReplaceSection(source string, middle string, autoadd bool) error {

	const start = "<!-- BEGINNING OF PRE-COMMIT-PIKE DOCS HOOK -->"
	const stop = "<!-- END OF PRE-COMMIT-PIKE DOCS HOOK -->"

	newSource, _ := filepath.Abs(source)
	//log.Print(newSource)
	dat, err := os.ReadFile(newSource)
	if (err) != nil {
		return err
	}

	file := string(dat)

	if !strings.Contains(file, start) || !strings.Contains(file, stop) {
		// add to new readme files
		if !strings.Contains(file, start) && !strings.Contains(file, stop) {
			if autoadd {
				file = file + "\n\n" + start + stop
			} else {
				return errors.New("missing both hooks in Readme, consider using the flag -autoappend")
			}
		} else {
			return errors.New("pike delimiters mismatch in Readme")
		}

	}

	section1 := (strings.Split(file, start)[0]) + start
	if strings.Contains(section1, stop) {
		return errors.New("pike delimiters mismatch in Readme")
	}
	section2 := stop + (strings.Split(file, stop)[1])

	var Output bytes.Buffer
	Output.WriteString(section1)
	Output.WriteString(middle)
	Output.WriteString(section2)

	err = os.WriteFile(source, Output.Bytes(), 0644)
	if (err) != nil {
		return err
	}

	return err
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
