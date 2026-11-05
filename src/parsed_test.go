package pike_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	pike "github.com/jameswoolfenden/pike/src"
)

func readTF(t *testing.T, dir string) map[string][]byte {
	t.Helper()

	files, err := pike.GetTF(dir)
	if err != nil {
		t.Fatal(err)
	}

	sources := map[string][]byte{}

	for _, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}

		sources[f] = b
	}

	return sources
}

// TestParsedDirMatchesDisk checks the caller-supplied-content path gives the
// same permissions and findings as the path-based entry points.
func TestParsedDirMatchesDisk(t *testing.T) {
	t.Parallel()

	dir, err := filepath.Abs("testdata/audit/aws")
	if err != nil {
		t.Fatal(err)
	}

	wantBag, err := pike.MakePermissionBag(dir, nil, false, "", false)
	if err != nil {
		t.Fatal(err)
	}

	wantFindings, err := pike.CollectAuditFindings(dir)
	if err != nil {
		t.Fatal(err)
	}

	parsed := pike.ParseSources(dir, readTF(t, dir))

	gotBag, gotFindings, err := parsed.Analyse("", false)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(wantBag, gotBag) {
		t.Errorf("permission bag differs:\nwant %+v\ngot  %+v", wantBag, gotBag)
	}

	if !reflect.DeepEqual(wantFindings, gotFindings) {
		t.Errorf("findings differ:\nwant %+v\ngot  %+v", wantFindings, gotFindings)
	}
}

func TestParseSourcesRecordsErrors(t *testing.T) {
	t.Parallel()

	parsed := pike.ParseSources(t.TempDir(), map[string][]byte{"bad.tf": []byte(`resource "x" {`)})

	if len(parsed.Errors) != 1 || len(parsed.Bodies) != 0 {
		t.Fatalf("expected one parse error, got errors=%v bodies=%v", parsed.Errors, parsed.Bodies)
	}

	if _, err := parsed.PermissionBag("", false); err == nil {
		t.Error("expected an error when every file fails to parse")
	}
}
