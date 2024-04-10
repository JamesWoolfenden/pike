package parse

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"
	"golang.org/x/exp/slices"
)

type provider struct {
	Resources   []string `json:"resources"`
	DataSources []string `json:"dataSources"`
}

func Parse(codebase string, name string) error {
	var err error

	var jsonOut []byte

	myProvider := provider{}

	name = strings.ToLower(name)

	switch name {
	case "google":
		{
			match := `resource "(` + name + `_.*?)"`
			myProvider.Resources, err = GetMatches(codebase, match, "markdown")
			if err != nil {
				return err
			}

			myProvider.DataSources, err = GetMatches(codebase, `data "(`+name+`_.*?)"`, "markdown")
			if err != nil {
				return err
			}
		}
	default:
		match := `resource "(` + name + `_.*?)"`
		myProvider.Resources, err = GetMatches(codebase, match, "markdown")

		if err != nil {
			return err
		}

		myProvider.DataSources, err = GetMatches(codebase, `# Data Source:(.*)`, "markdown")
		if err != nil {
			return err
		}
	}

	jsonOut, err = json.MarshalIndent(myProvider, "", "    ")

	if err != nil {
		return err
	}

	log.Info().Msgf("creating %s-members.json", name)
	err = os.WriteFile(name+"-members.json", jsonOut, 0o700)

	if err != nil {
		return err
	}

	return nil
}

func GetMatches(source string, match string, extension string) ([]string, error) {
	files, err := GetGoFiles(source, extension)
	if err != nil {
		return nil, err
	}

	var (
		matches = make(map[string]bool)
		a       []string
	)

	for _, file := range files {
		contents, _ := os.ReadFile(file)

		re := regexp.MustCompile(match)
		match := re.FindAllString(string(contents), -1)

		for _, item := range match {
			if strings.Contains(item, "%s") {
				continue
			}

			matched := strings.TrimSpace(strings.ReplaceAll(item, "\"", ""))
			matched = strings.TrimSpace(strings.ReplaceAll(matched, "# Data Source: ", ""))
			matched = strings.TrimSpace(strings.ReplaceAll(matched, "data ", ""))
			matched = strings.TrimSpace(strings.ReplaceAll(matched, "resource ", ""))
			matched = strings.TrimSpace(strings.ReplaceAll(matched, "`", ""))
			a, matches = add(matched, matches, a)
		}
	}

	keys := GetKeys(matches)

	return keys, nil
}

func GetGoFiles(path string, extension string) ([]string, error) {
	libRegEx, err := regexp.Compile("^.+\\." + extension + "$")
	if err != nil {
		return nil, err
	}

	absPath, err := filepath.Abs(path)

	log.Info().Msgf(absPath)

	if err != nil {
		return nil, err
	}

	var files []string

	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err == nil && libRegEx.MatchString(info.Name()) {
			if strings.Contains(path, "_test") || strings.Contains(path, ".ci") || info.IsDir() {
			} else {
				files = append(files, path)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func GetKeys(m map[string]bool) []string {
	var keys []string

	for k := range m {
		keys = append(keys, k)
	}

	slices.Sort(keys)

	return keys
}

func add(s string, m map[string]bool, a []string) ([]string, map[string]bool) {
	if m[s] {
		return a, m
	}

	a = append(a, s)

	m[s] = true

	return a, m
}
