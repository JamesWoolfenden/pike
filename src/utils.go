package pike

import (
	"bytes"
	"errors"
	"math/rand"
	"os"
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
func ReplaceSection(source string, middle string) error {

	const start = "<!-- BEGINNING OF PRE-COMMIT-PIKE DOCS HOOK -->"
	const stop = "<!-- END OF PRE-COMMIT-PIKE DOCS HOOK -->"

	dat, err := os.ReadFile(source)
	if (err) != nil {
		return err
	}

	file := string(dat)

	if !strings.Contains(file, start) || !strings.Contains(file, stop) {
		return errors.New("Pike delimiters not found in Readme")
	}

	section1 := (strings.Split(file, start)[0]) + start
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
