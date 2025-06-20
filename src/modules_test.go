package pike

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReadModuleJson_ValidInput(t *testing.T) {
	jsonInput := `{
		"Modules": [
			{
				"Key": "module1",
				"Source": "github.com/example/module1",
				"Version": "1.0.0",
				"Dir": "modules/module1"
			},
			{
				"Key": "module2",
				"Source": "github.com/example/module2",
				"Dir": "modules/module2"
			}
		]
	}`

	reader := strings.NewReader(jsonInput)
	result, err := ReadModuleJson(reader)

	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if len(result) != 2 {
		t.Fatalf("Expected 2 modules, got: %d", len(result))
	}

	// Test module1
	module1, exists := result["module1"]
	if !exists {
		t.Fatal("Expected module1 to exist")
	}
	if module1.Key != "module1" {
		t.Errorf("Expected Key 'module1', got: %s", module1.Key)
	}
	if module1.SourceAddr != "github.com/example/module1" {
		t.Errorf("Expected SourceAddr 'github.com/example/module1', got: %s", module1.SourceAddr)
	}
	if module1.VersionStr != "1.0.0" {
		t.Errorf("Expected VersionStr '1.0.0', got: %s", module1.VersionStr)
	}
	if module1.Version == nil {
		t.Error("Expected Version to be parsed")
	} else if module1.Version.String() != "1.0.0" {
		t.Errorf("Expected Version '1.0.0', got: %s", module1.Version.String())
	}

	// Test module2 (no version)
	module2, exists := result["module2"]
	if !exists {
		t.Fatal("Expected module2 to exist")
	}
	if module2.Version != nil {
		t.Error("Expected Version to be nil for module without version")
	}
}

func TestReadModuleJson_EmptyInput(t *testing.T) {
	reader := strings.NewReader("")
	result, err := ReadModuleJson(reader)

	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if len(result) != 0 {
		t.Fatalf("Expected empty ModuleJson, got: %d modules", len(result))
	}
}

func TestReadModuleJson_InvalidJSON(t *testing.T) {
	reader := strings.NewReader(`{"invalid": json}`)
	_, err := ReadModuleJson(reader)

	if err == nil {
		t.Fatal("Expected error for invalid JSON")
	}

	if !strings.Contains(err.Error(), "invalid character 'j' looking for beginning of value") {
		t.Errorf("Expected unmarshalling error, got: %v", err)
	}
}

func TestReadModuleJson_DuplicateKeys(t *testing.T) {
	jsonInput := `{
		"Modules": [
			{
				"Key": "duplicate",
				"Source": "github.com/example/module1",
				"Dir": "modules/module1"
			},
			{
				"Key": "duplicate",
				"Source": "github.com/example/module2",
				"Dir": "modules/module2"
			}
		]
	}`

	reader := strings.NewReader(jsonInput)
	_, err := ReadModuleJson(reader)

	if err == nil {
		t.Fatal("Expected error for duplicate keys")
	}

	if !strings.Contains(err.Error(), "snapshot file contains two records for path duplicate") {
		t.Errorf("Expected duplicate key error, got: %v", err)
	}
}

func TestReadModuleJson_InvalidVersion(t *testing.T) {
	jsonInput := `{
		"Modules": [
			{
				"Key": "module1",
				"Source": "github.com/example/module1",
				"Version": "invalid-version",
				"Dir": "modules/module1"
			}
		]
	}`

	reader := strings.NewReader(jsonInput)
	_, err := ReadModuleJson(reader)

	if err == nil {
		t.Fatal("Expected error for invalid version")
	}

	if !strings.Contains(err.Error(), "invalid version") {
		t.Errorf("Expected invalid version error, got: %v", err)
	}
}

func TestReadModuleJson_WindowsPathHandling(t *testing.T) {
	jsonInput := `{
		"Modules": [
			{
				"Key": "module1",
				"Source": "github.com/example/module1",
				"Dir": "modules/subdir/module1"
			}
		]
	}`

	reader := strings.NewReader(jsonInput)
	result, err := ReadModuleJson(reader)

	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	module1 := result["module1"]
	expectedDir := filepath.FromSlash("modules/subdir/module1")
	if module1.Dir != expectedDir {
		t.Errorf("Expected Dir '%s', got: '%s'", expectedDir, module1.Dir)
	}
}

func TestReadModuleJsonForDir_ValidFile(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "pike-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a modules.json file
	jsonContent := `{
		"Modules": [
			{
				"Key": "test-module",
				"Source": "github.com/example/test",
				"Dir": "modules/test"
			}
		]
	}`

	modulesFile := filepath.Join(tempDir, ManifestSnapshotFilename)
	err = os.WriteFile(modulesFile, []byte(jsonContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write modules file: %v", err)
	}

	result, err := ReadModuleJsonForDir(tempDir)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if len(result) != 1 {
		t.Fatalf("Expected 1 module, got: %d", len(result))
	}

	module, exists := result["test-module"]
	if !exists {
		t.Fatal("Expected test-module to exist")
	}
	if module.SourceAddr != "github.com/example/test" {
		t.Errorf("Expected SourceAddr 'github.com/example/test', got: %s", module.SourceAddr)
	}
}

func TestReadModuleJsonForDir_NonExistentFile(t *testing.T) {
	// Create a temporary directory without modules.json
	tempDir, err := os.MkdirTemp("", "pike-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	result, err := ReadModuleJsonForDir(tempDir)
	if err != nil {
		t.Fatalf("Expected no error for non-existent file, got: %v", err)
	}

	if len(result) != 0 {
		t.Fatalf("Expected empty ModuleJson, got: %d modules", len(result))
	}
}

func TestReadModuleJsonForDir_NonExistentDirectory(t *testing.T) {
	result, err := ReadModuleJsonForDir("/non/existent/directory")
	if err != nil {
		t.Fatalf("Expected no error for non-existent directory, got: %v", err)
	}

	if len(result) != 0 {
		t.Fatalf("Expected empty ModuleJson, got: %d modules", len(result))
	}
}

func TestReturnLocalAddrFromSource_Found(t *testing.T) {
	modules := ModuleJson{
		"module1": Record{
			Key:        "module1",
			SourceAddr: "github.com/example/module1",
			Dir:        "/path/to/module1",
		},
		"module2": Record{
			Key:        "module2",
			SourceAddr: "github.com/example/module2",
			Dir:        "/path/to/module2",
		},
	}

	result := ReturnLocalAddrFromSource("github.com/example/module2", modules)
	expected := "/path/to/module2"

	if result != expected {
		t.Errorf("Expected '%s', got: '%s'", expected, result)
	}
}

func TestReturnLocalAddrFromSource_NotFound(t *testing.T) {
	modules := ModuleJson{
		"module1": Record{
			Key:        "module1",
			SourceAddr: "github.com/example/module1",
			Dir:        "/path/to/module1",
		},
	}

	result := ReturnLocalAddrFromSource("github.com/example/nonexistent", modules)

	if result != "" {
		t.Errorf("Expected empty string, got: '%s'", result)
	}
}

func TestReturnLocalAddrFromSource_EmptyModules(t *testing.T) {
	modules := make(ModuleJson)

	result := ReturnLocalAddrFromSource("github.com/example/any", modules)

	if result != "" {
		t.Errorf("Expected empty string, got: '%s'", result)
	}
}

func TestReturnLocalAddrFromSource_EmptySource(t *testing.T) {
	modules := ModuleJson{
		"module1": Record{
			Key:        "module1",
			SourceAddr: "github.com/example/module1",
			Dir:        "/path/to/module1",
		},
	}

	result := ReturnLocalAddrFromSource("", modules)

	if result != "" {
		t.Errorf("Expected empty string, got: '%s'", result)
	}
}
