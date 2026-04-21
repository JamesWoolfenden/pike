package pike

import (
	"embed"
	"io/fs"
	"path/filepath"
	"strings"
)

//go:embed mapping/aws
var awsFS embed.FS

//go:embed mapping/google
var gcpFS embed.FS

//go:embed mapping/azurerm
var azureFS embed.FS

var (
	awsResourceMap   map[string][]byte
	awsDataMap       map[string][]byte
	gcpResourceMap   map[string][]byte
	gcpDataMap       map[string][]byte
	azureResourceMap map[string][]byte
	azureDataMap     map[string][]byte
)

func init() {
	awsResourceMap = buildMappingMap(awsFS, "mapping/aws/resource")
	awsDataMap = buildMappingMap(awsFS, "mapping/aws/data")
	gcpResourceMap = buildMappingMap(gcpFS, "mapping/google/resource")
	gcpDataMap = buildMappingMap(gcpFS, "mapping/google/data")
	azureResourceMap = buildMappingMap(azureFS, "mapping/azurerm/resource")
	azureDataMap = buildMappingMap(azureFS, "mapping/azurerm/data")
}

// buildMappingMap walks an embedded FS subtree and returns a map from
// resource name (JSON filename without extension) to raw JSON content.
// Adding a new resource only requires dropping a JSON file in the right
// directory — no Go file needs updating.
func buildMappingMap(fsys embed.FS, root string) map[string][]byte {
	m := make(map[string][]byte)

	_ = fs.WalkDir(fsys, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}

		data, err := fsys.ReadFile(path)
		if err != nil {
			return nil
		}

		name := strings.TrimSuffix(filepath.Base(path), ".json")
		m[name] = data

		return nil
	})

	return m
}
