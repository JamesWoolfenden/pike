package pike

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	version "github.com/hashicorp/go-version"
)

const (
	ManifestSnapshotFilename = "modules.json"
)

// Record represents some metadata about an installed module, as part
// of a modules json.
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
		return nil, fmt.Errorf("error unmarshalling snapshot: %v", err)
	}
	new := make(ModuleJson)
	for _, record := range read.Records {
		if record.VersionStr != "" {
			record.Version, err = version.NewVersion(record.VersionStr)
			if err != nil {
				return nil, fmt.Errorf("invalid version %q for %s: %s", record.VersionStr, record.Key, err)
			}
		}
		// Ensure Windows is using the proper modules path format after
		// reading the modules manifest Dir records
		record.Dir = filepath.FromSlash(record.Dir)

		if _, exists := new[record.Key]; exists {
			return nil, fmt.Errorf("snapshot file contains two records for path %s", record.Key)
		}
		new[record.Key] = record
	}
	return new, nil
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
