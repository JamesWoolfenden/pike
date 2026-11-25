package pike

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGCPPolicyRejectsInjection(t *testing.T) {
	t.Setenv("GOOGLE_CLOUD_PROJECT", "p\"\n  injected_attr = \"pwned")

	if _, err := GCPPolicy([]string{"storage.buckets.create"}, "", t.TempDir()); err == nil {
		t.Fatal("expected error for malformed project id")
	}

	t.Setenv("GOOGLE_CLOUD_PROJECT", "valid-project-1")

	if _, err := GCPPolicy([]string{"storage.buckets.create"}, "x\"\ny", t.TempDir()); err == nil {
		t.Fatal("expected error for malformed policy name")
	}
}

func TestProviderProjectInjectionIgnored(t *testing.T) {
	dir := t.TempDir()
	tf := "provider \"google\" {\n  project = \"p\\\"\\n  injected_attr = \\\"pwned\"\n}\n"

	if err := os.WriteFile(filepath.Join(dir, "main.tf"), []byte(tf), 0o600); err != nil {
		t.Fatal(err)
	}

	if got := getProjectFromTerraformProvider(dir); got != "" {
		t.Fatalf("expected malformed project to be ignored, got %q", got)
	}
}

func TestEnsurePikeDirRefusesSymlink(t *testing.T) {
	scan := t.TempDir()
	victim := t.TempDir()

	if err := os.Symlink(victim, filepath.Join(scan, ".pike")); err != nil {
		t.Skip("symlinks unsupported")
	}

	if _, err := ensurePikeDir(scan); err == nil {
		t.Fatal("expected symlinked .pike to be refused")
	}

	if err := writeTwoRoleFile("x", scan, ""); err == nil {
		t.Fatal("expected write to be refused")
	}

	if entries, _ := os.ReadDir(victim); len(entries) != 0 {
		t.Fatalf("victim directory was written to: %v", entries)
	}
}

func TestWriteFileNoSymlink(t *testing.T) {
	dir := t.TempDir()
	target := filepath.Join(dir, "target")
	link := filepath.Join(dir, "link")

	if err := os.WriteFile(target, []byte("orig"), 0o600); err != nil {
		t.Fatal(err)
	}

	if err := os.Symlink(target, link); err != nil {
		t.Skip("symlinks unsupported")
	}

	if err := writeFileNoSymlink(link, []byte("new")); err == nil {
		t.Fatal("expected symlink target to be refused")
	}

	if b, _ := os.ReadFile(target); string(b) != "orig" {
		t.Fatalf("target modified: %s", b)
	}
}

func TestReadConfinedFile(t *testing.T) {
	root := t.TempDir()
	base := filepath.Join(root, "scan")
	outside := filepath.Join(root, "secret.json")

	if err := os.Mkdir(base, 0o750); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(outside, []byte("{}"), 0o600); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(filepath.Join(base, "ok.json"), []byte("{}"), 0o600); err != nil {
		t.Fatal(err)
	}

	if _, err := readConfinedFile(base, "ok.json"); err != nil {
		t.Fatalf("in-tree file should be readable: %v", err)
	}

	if _, err := readConfinedFile(base, "../secret.json"); err == nil {
		t.Fatal("expected .. traversal to be refused")
	}

	if err := os.Symlink(outside, filepath.Join(base, "link.json")); err == nil {
		if _, err := readConfinedFile(base, "link.json"); err == nil {
			t.Fatal("expected symlink escape to be refused")
		}
	}

	if err := os.WriteFile(filepath.Join(base, "big.json"), []byte(strings.Repeat("a", maxPolicyFileSize+1)), 0o600); err != nil {
		t.Fatal(err)
	}

	if _, err := readConfinedFile(base, "big.json"); err == nil {
		t.Fatal("expected oversized file to be refused")
	}

	if _, err := readConfinedFile(base, "."); err == nil {
		t.Fatal("expected directory to be refused")
	}
}
