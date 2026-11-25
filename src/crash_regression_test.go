package pike

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func writeTF(t *testing.T, content string) string {
	t.Helper()

	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "main.tf"), []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}

	return dir
}

// #216: a module whose source points at its own directory must terminate.
func TestSelfReferencingModuleTerminates(t *testing.T) {
	dir := writeTF(t, "module \"self\" {\n  source = \"./\"\n}\n")

	done := make(chan struct{})

	go func() {
		defer close(done)

		files, _ := GetTF(dir)
		for _, f := range files {
			_, _ = GetResources(f, dir)
		}
	}()

	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Fatal("scan did not terminate on a self-referencing module")
	}
}

// #217: a non-string interpolation in a module source must not panic.
func TestModuleSourceNonStringInterpolation(t *testing.T) {
	dir := writeTF(t, "module \"m\" {\n  source = \"a${1}b\"\n}\n")

	body, err := GetResourceBlocks(filepath.Join(dir, "main.tf"))
	if err != nil {
		t.Fatal(err)
	}

	for _, b := range body.Blocks {
		if got := GetModulePath(b); got != "" {
			t.Fatalf("expected empty path, got %q", got)
		}
	}
}

// #218: a non-string provider project must not panic.
func TestProviderProjectNonString(t *testing.T) {
	for _, v := range []string{"123", "null", "true"} {
		dir := writeTF(t, "provider \"google\" {\n  project = "+v+"\n}\n")

		if got := getProjectFromTerraformProvider(dir); got != "" {
			t.Fatalf("project = %s: expected empty, got %q", v, got)
		}
	}
}
