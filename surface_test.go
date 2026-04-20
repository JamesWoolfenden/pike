// Surface tests lock down the externally-visible behaviour of the pike
// binary and the pike Go package so the upcoming src/ -> internal/ refactor
// cannot silently change either. They are deliberately blunt:
//
//   - TestHelpGolden:    builds the current binary and diffs `pike --help` output
//     against testdata/help.golden. This catches accidental
//     subcommand drops, flag renames, and usage-string drift.
//
//   - TestAPISurfaceGolden: snapshots the exported identifiers in `package pike`
//     (names, signatures, const values) into
//     testdata/api_surface.txt. While we carve pike up
//     into internal/ packages the root `pike` package will
//     be re-exporting via type aliases; this test fails
//     loudly if an alias is forgotten or a type shape
//     drifts.
//
// Run `go test -run '^TestHelpGolden$|^TestAPISurfaceGolden$' -update` to
// regenerate the goldens after an intentional change, then diff and commit.
//
// These tests do NOT run during day-to-day `go test ./...` on machines that
// lack `go`/`pike` in PATH after a build; they gate on t.Skip when the
// prerequisites are missing so local developer workflow is never blocked.
package main_test

import (
	"bytes"
	"flag"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

var updateGoldens = flag.Bool("update", false, "rewrite testdata/*.golden files from observed output")

const (
	helpGoldenPath       = "testdata/help.golden"
	apiSurfaceGoldenPath = "testdata/api_surface.txt"
)

// TestHelpGolden builds the pike binary via `go run .` and diffs --help output.
//
// We `go run` instead of `go build` + exec because (a) it's a single command,
// (b) it inherits GOCACHE so repeat runs are fast, and (c) it avoids having to
// clean up a temp binary.
func TestHelpGolden(t *testing.T) {
	if _, err := exec.LookPath("go"); err != nil {
		t.Skip("go not on PATH; run this test inside a Go toolchain environment")
	}

	// Windows `go run` invocations are sometimes slower than CI's per-test
	// timeout defaults, but this test doesn't need any nondeterministic setup
	// so we let the default test timeout apply.
	cmd := exec.Command("go", "run", ".", "--help")

	// `urfave/cli` writes help to stdout, but some branches (e.g. unknown
	// command) emit to stderr. We capture both so we don't accidentally
	// miss drift into the other stream.
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		// urfave/cli returns 0 from --help; non-zero here is a real failure,
		// but include the captured output in the message so CI logs are useful.
		t.Fatalf("go run . --help: %v\noutput:\n%s", err, out.String())
	}

	got := normalizeHelp(out.String())

	if *updateGoldens {
		if err := os.MkdirAll(filepath.Dir(helpGoldenPath), 0o755); err != nil {
			t.Fatalf("mkdir testdata: %v", err)
		}
		if err := os.WriteFile(helpGoldenPath, []byte(got), 0o600); err != nil {
			t.Fatalf("write golden: %v", err)
		}
		t.Logf("wrote %s (%d bytes)", helpGoldenPath, len(got))
		return
	}

	want, err := os.ReadFile(helpGoldenPath)
	if err != nil {
		if os.IsNotExist(err) {
			t.Fatalf("%s does not exist; run `go test -run TestHelpGolden -update` to create it", helpGoldenPath)
		}
		t.Fatalf("read golden: %v", err)
	}
	want = bytes.ReplaceAll(want, []byte("\r\n"), []byte("\n"))

	if got != string(want) {
		t.Errorf("pike --help output drifted from %s.\nRun `go test -run TestHelpGolden -update` to accept the change and review the diff before committing.\n\n--- got ---\n%s\n--- want ---\n%s", helpGoldenPath, got, string(want))
	}
}

// normalizeHelp strips volatile bits from --help output so the golden is
// stable across environments. Today that's just the binary version (set to
// pike.Version which is bumped per release) and any trailing whitespace.
func normalizeHelp(s string) string {
	var keep []string
	for _, line := range strings.Split(s, "\n") {
		trimmed := strings.TrimRight(line, " \t\r")
		// Skip the VERSION: block entirely - it's release-gated and would
		// break the golden every tag.
		if strings.HasPrefix(strings.TrimSpace(trimmed), "VERSION:") {
			keep = append(keep, "VERSION: <redacted-for-golden>")
			continue
		}
		keep = append(keep, trimmed)
	}
	// Collapse trailing blank lines.
	for len(keep) > 0 && keep[len(keep)-1] == "" {
		keep = keep[:len(keep)-1]
	}
	return strings.Join(keep, "\n") + "\n"
}

// TestAPISurfaceGolden parses src/*.go and extracts the exported identifier
// set, then diffs that against testdata/api_surface.txt.
//
// The surface we care about during the refactor:
//   - package-level func signatures
//   - exported type names and their kinds (struct, interface, alias)
//   - exported const names and values
//   - exported var names
//
// We intentionally DO NOT capture struct field lists - those drift for
// internal reasons (unexported field added, tag tweak) and would generate
// noise. The Stage 3+ alias-shim work relies on type identity being
// preserved, which the type name + kind check catches.
func TestAPISurfaceGolden(t *testing.T) {
	fset := token.NewFileSet()
	// Parse every .go file in src/ at the top level (not subdirs) - that's
	// package pike today.
	entries, err := os.ReadDir("src")
	if err != nil {
		if os.IsNotExist(err) {
			t.Skip("src/ not present (post-Stage-7?); this test is deliberately src-specific")
		}
		t.Fatalf("read src/: %v", err)
	}

	var lines []string
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") || strings.HasSuffix(e.Name(), "_test.go") {
			continue
		}
		path := filepath.Join("src", e.Name())
		file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			t.Fatalf("parse %s: %v", path, err)
		}
		for _, d := range file.Decls {
			switch decl := d.(type) {
			case *ast.FuncDecl:
				if !decl.Name.IsExported() || decl.Recv != nil {
					continue
				}
				lines = append(lines, "func "+decl.Name.Name+" "+formatFuncSig(decl.Type))
			case *ast.GenDecl:
				for _, spec := range decl.Specs {
					switch s := spec.(type) {
					case *ast.TypeSpec:
						if s.Name.IsExported() {
							lines = append(lines, "type "+s.Name.Name+" "+typeKind(s))
						}
					case *ast.ValueSpec:
						for i, n := range s.Names {
							if !n.IsExported() {
								continue
							}
							tok := "var"
							if decl.Tok == token.CONST {
								tok = "const"
							}
							val := ""
							if i < len(s.Values) {
								val = " = " + exprString(s.Values[i])
							}
							lines = append(lines, tok+" "+n.Name+val)
						}
					}
				}
			}
		}
	}
	sort.Strings(lines)
	got := strings.Join(lines, "\n") + "\n"

	if *updateGoldens {
		if err := os.MkdirAll(filepath.Dir(apiSurfaceGoldenPath), 0o755); err != nil {
			t.Fatalf("mkdir testdata: %v", err)
		}
		if err := os.WriteFile(apiSurfaceGoldenPath, []byte(got), 0o600); err != nil {
			t.Fatalf("write golden: %v", err)
		}
		t.Logf("wrote %s (%d lines, %d bytes)", apiSurfaceGoldenPath, len(lines), len(got))
		return
	}

	want, err := os.ReadFile(apiSurfaceGoldenPath)
	if err != nil {
		if os.IsNotExist(err) {
			t.Fatalf("%s does not exist; run `go test -run TestAPISurfaceGolden -update` to create it", apiSurfaceGoldenPath)
		}
		t.Fatalf("read golden: %v", err)
	}
	want = bytes.ReplaceAll(want, []byte("\r\n"), []byte("\n"))

	if got != string(want) {
		t.Errorf("exported pike surface drifted from %s.\nRun `go test -run TestAPISurfaceGolden -update` to accept the change and review the diff before committing.\n\n--- got ---\n%s\n--- want ---\n%s", apiSurfaceGoldenPath, got, string(want))
	}
}

// formatFuncSig produces a stable printed form of a function signature,
// using default formatting from go/printer would be ideal but go/printer
// needs a token.FileSet which would make this function noisier; for the
// refactor's purposes "same names in order" is enough.
func formatFuncSig(ft *ast.FuncType) string {
	var b strings.Builder
	b.WriteString("(")
	if ft.Params != nil {
		for i, field := range ft.Params.List {
			if i > 0 {
				b.WriteString(", ")
			}
			for j, name := range field.Names {
				if j > 0 {
					b.WriteString(", ")
				}
				b.WriteString(name.Name)
				b.WriteString(" ")
			}
			b.WriteString(exprString(field.Type))
		}
	}
	b.WriteString(")")
	if ft.Results != nil {
		b.WriteString(" ")
		if len(ft.Results.List) > 1 || (len(ft.Results.List) == 1 && len(ft.Results.List[0].Names) > 0) {
			b.WriteString("(")
		}
		for i, field := range ft.Results.List {
			if i > 0 {
				b.WriteString(", ")
			}
			b.WriteString(exprString(field.Type))
		}
		if len(ft.Results.List) > 1 || (len(ft.Results.List) == 1 && len(ft.Results.List[0].Names) > 0) {
			b.WriteString(")")
		}
	}
	return b.String()
}

// typeKind returns a short string categorising a TypeSpec, e.g. "struct",
// "interface", "alias <underlying>", or "defined <underlying>".
func typeKind(t *ast.TypeSpec) string {
	if t.Assign.IsValid() {
		return "alias " + exprString(t.Type)
	}
	switch underlying := t.Type.(type) {
	case *ast.StructType:
		return "struct"
	case *ast.InterfaceType:
		return "interface"
	default:
		return "defined " + exprString(underlying)
	}
}

// exprString renders a type or simple-value expression as Go source-level
// text. We walk the AST directly rather than calling printer.Fprint so the
// output is deterministic across go releases (printer format occasionally
// tweaks spacing).
func exprString(e ast.Expr) string {
	switch x := e.(type) {
	case *ast.Ident:
		return x.Name
	case *ast.SelectorExpr:
		return exprString(x.X) + "." + x.Sel.Name
	case *ast.StarExpr:
		return "*" + exprString(x.X)
	case *ast.ArrayType:
		if x.Len == nil {
			return "[]" + exprString(x.Elt)
		}
		return "[" + exprString(x.Len) + "]" + exprString(x.Elt)
	case *ast.MapType:
		return "map[" + exprString(x.Key) + "]" + exprString(x.Value)
	case *ast.ChanType:
		return "chan " + exprString(x.Value)
	case *ast.Ellipsis:
		return "..." + exprString(x.Elt)
	case *ast.FuncType:
		return "func" + formatFuncSig(x)
	case *ast.BasicLit:
		return x.Value
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.StructType:
		return "struct{...}"
	case *ast.UnaryExpr:
		return x.Op.String() + exprString(x.X)
	case *ast.BinaryExpr:
		return exprString(x.X) + " " + x.Op.String() + " " + exprString(x.Y)
	case *ast.CallExpr:
		return exprString(x.Fun) + "(...)"
	}
	return "?"
}
