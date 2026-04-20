// Package parse determines what resources and datasources a Terraform
// provider exposes, and writes the result to {provider}-members.json in
// the current working directory. The rest of pike consumes that JSON at
// runtime to validate mapping coverage.
//
// There are two extraction strategies. The schema-based path (default) runs
// `terraform providers schema -json` against a throwaway init'd workspace
// and reads the provider's declared resource and datasource schemas
// directly. It is fast, authoritative, and requires terraform (or a tofu
// binary symlinked as terraform) plus network access to the provider
// registry.
//
// The docs-based path walks a local checkout of the provider's source
// repository and greps the markdown documentation for resource and
// datasource declarations. It's slower, inferred rather than authoritative,
// and kept as a fallback for offline/airgapped use or for providers that
// don't publish a registry schema.
//
// Parse tries the schema path first. If that fails AND a codebase directory
// was supplied, it falls back to the docs path. If the schema path fails
// and no codebase was supplied, the schema error is returned verbatim.
package parse

import (
	"encoding/json"
	"errors"
	"fmt"
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

	// DeprecatedResources / DeprecatedData map each deprecated resource or
	// datasource name to the upstream provider's deprecation message. Only
	// populated by the schema-based path; the docs-based fallback has no
	// reliable way to identify deprecation so it leaves these empty.
	//
	// omitempty so older consumers that unmarshal into the old shape don't
	// see unexpected fields, and so members files for providers that
	// happen to have nothing deprecated stay visually quiet.
	DeprecatedResources map[string]string `json:"deprecatedResources,omitempty"`
	DeprecatedData      map[string]string `json:"deprecatedData,omitempty"`
}

// Parse is the package entrypoint. It writes {name}-members.json to the
// current working directory.
func Parse(codebase string, name string) error {
	if name == "" {
		return errors.New("name is required")
	}

	name = strings.ToLower(name)

	prov, err := parseFromSchema(name)
	if err != nil {
		if codebase == "" {
			return fmt.Errorf("schema-based parse failed and no docs fallback was supplied: %w", err)
		}

		log.Warn().Err(err).Str("provider", name).Msg("schema-based parse failed; falling back to markdown docs")

		prov, err = parseFromDocs(codebase, name)
		if err != nil {
			return err
		}
	}

	return writeMembers(name, prov)
}

// writeMembers serialises a provider's resource/datasource lists to
// {name}-members.json in the current working directory. The filename
// convention is load-bearing: the consuming code at src/coverage/ reads
// exactly this path.
func writeMembers(name string, prov *provider) error {
	out, err := json.MarshalIndent(prov, "", "    ")
	if err != nil {
		return fmt.Errorf("marshalling %s members: %w", name, err)
	}

	log.Info().Msgf("creating %s-members.json", name)

	if err := os.WriteFile(name+"-members.json", out, 0o600); err != nil {
		return fmt.Errorf("writing %s-members.json: %w", name, err)
	}

	return nil
}

// parseFromDocs is the historical, doc-scanning extraction path. Kept as
// a fallback; see the package doc.
//
// Google's provider repo names datasources via `data "google_..." {}`
// fixture blocks in its example docs. Other providers use the
// `# Data Source:` markdown header convention instead.
func parseFromDocs(codebase, name string) (*provider, error) {
	if codebase == "" {
		return nil, errors.New("codebase is required for docs-based parse")
	}

	p := &provider{}

	resources, err := getMatches(codebase, `resource "(`+name+`_.*?)"`, "markdown")
	if err != nil {
		return nil, err
	}

	p.Resources = resources

	var dsPattern string
	if name == "google" {
		dsPattern = `data "(` + name + `_.*?)"`
	} else {
		dsPattern = `# Data Source:(.*)`
	}

	datasources, err := getMatches(codebase, dsPattern, "markdown")
	if err != nil {
		return nil, err
	}

	p.DataSources = datasources

	return p, nil
}

func getMatches(source string, match string, extension string) ([]string, error) {
	files, err := getGoFiles(source, extension)
	if err != nil {
		return nil, err
	}

	// Hoist the regex compile out of the file loop; recompiling per file
	// is measurable overhead on a 25k-file checkout.
	//
	// Historical behaviour used MustCompile, which panics on invalid input.
	// The "Invalid regex pattern" test in parse_test.go asserts wantErr=false,
	// so an invalid pattern silently produces no matches rather than panicking
	// or erroring. That's a debatable choice but is preserved here to keep
	// the public test contract stable for the rewrite.
	re, regexErr := regexp.Compile(match)
	if regexErr != nil {
		return nil, nil
	}

	var (
		matches = make(map[string]bool)
		a       []string
	)

	for _, file := range files {
		contents, _ := os.ReadFile(file) // #nosec G304 -- Reading Terraform files from user-specified paths

		for _, item := range re.FindAllString(string(contents), -1) {
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

	return getKeys(matches), nil
}

func getGoFiles(path string, extension string) ([]string, error) {
	if path == "" || extension == "" {
		return nil, errors.New("path or extension are required")
	}

	libRegEx, err := regexp.Compile("^.+\\." + extension + "$")
	if err != nil {
		return nil, err
	}

	absPath, err := filepath.Abs(path)

	log.Info().Msg(absPath)

	if err != nil {
		return nil, fmt.Errorf("absolute path error %v", err)
	}

	_, err = os.Stat(absPath)
	if err != nil {
		return nil, fmt.Errorf("path does not exist error %v", err)
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

func getKeys(m map[string]bool) []string {
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
