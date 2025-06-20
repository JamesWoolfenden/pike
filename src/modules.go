package pike

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/hashicorp/go-version"
)

const (
	ManifestSnapshotFilename = "modules.json"
)

// Record represents some metadata about an installed module, as part
// of a module JSON.
type Record struct {
	Key        string           `json:"Key"`
	SourceAddr string           `json:"Source"`
	Version    *version.Version `json:"-"`
	VersionStr string           `json:"Version,omitempty"`
	Dir        string           `json:"Dir"`
}

type ModuleJson map[string]Record

type modulesJson struct {
	Records []Record `json:"Modules"`
}

type invalidVersionError struct {
	err     error
	key     string
	version string
}

func (m *invalidVersionError) Error() string {
	return fmt.Sprintf("invalid version %q for %s: %s", m.version, m.key, m.err)
}

func ReadModuleJson(r io.Reader) (ModuleJson, error) {
	src, err := io.ReadAll(r)

	if err != nil {
		return nil, err
	}

	if len(src) == 0 {
		return make(ModuleJson), nil
	}

	var read modulesJson
	err = json.Unmarshal(src, &read)

	if err != nil {
		return nil, &unmarshallJSONError{err, ""}
	}

	newModuleJson := make(ModuleJson)
	for _, record := range read.Records {
		if record.VersionStr != "" {
			record.Version, err = version.NewVersion(record.VersionStr)
			if err != nil {
				return nil, &invalidVersionError{err, record.Key, record.VersionStr}
			}
		}
		// Ensure Windows is using the proper modules path format after
		// reading the module's manifest Dir records
		record.Dir = filepath.FromSlash(record.Dir)

		if _, exists := newModuleJson[record.Key]; exists {
			return nil, fmt.Errorf("snapshot file contains two records for path %s", record.Key)
		}
		newModuleJson[record.Key] = record
	}
	return newModuleJson, nil
}

func ReadModuleJsonForDir(dir string) (ModuleJson, error) {
	fn := filepath.Join(dir, ManifestSnapshotFilename)
	r, err := os.Open(fn)
	if err != nil {
		if os.IsNotExist(err) {
			return make(ModuleJson), nil
		}
		return nil, err
	}
	defer r.Close()
	return ReadModuleJson(r)
}

func ReturnLocalAddrFromSource(source string, listModules ModuleJson) string {
	for _, module := range listModules {
		if module.SourceAddr == source {
			return module.Dir
		}
	}
	return ""
}
