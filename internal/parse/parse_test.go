package parse

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestGetGoFiles(t *testing.T) {
	t.Parallel()

	type args struct {
		path      string
		extension string
	}

	wanted := []string{filepath.Join("testdata", "test.go")}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "Pass", args: args{path: "./testdata", extension: "go"}, want: wanted},
		{name: "None", args: args{path: "../../src/mapping", extension: "go"}},
		{
			name: "Valid path",
			args: args{
				path:      "./testdata",
				extension: "markdown",
			},
			wantErr: false,
		},
		{
			name: "Invalid path",
			args: args{
				path:      "/nonexistent",
				extension: "markdown",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := getGoFiles(tt.args.path, tt.args.extension)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetGoFiles() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGoFiles() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetKeys(t *testing.T) {
	t.Parallel()

	type args struct {
		m map[string]struct{}
	}

	sample := map[string]struct{}{"first": {}}
	nothing := map[string]struct{}{}
	var bumpkis []string

	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "pass", args: args{sample}, want: []string{"first"}},
		{name: "nil", args: args{nothing}},
		{
			name: "Non-empty map",
			args: args{m: map[string]struct{}{"key1": {}, "key2": {}}},
			want: []string{"key1", "key2"},
		},
		{
			name: "Empty map",
			args: args{m: map[string]struct{}{}},
			want: bumpkis,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := getKeys(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMatches(t *testing.T) {
	t.Parallel()

	var empty []string

	type args struct {
		source    string
		match     string
		extension string
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "pass", wantErr: true},
		{name: "go", args: args{source: "./testdata", match: "(aws_.*?)", extension: "go"}, want: []string{"aws_"}},
		{
			name: "Valid pattern",
			args: args{
				source:    "./testdata",
				match:     `resource "(test_.*?)"`,
				extension: "markdown",
			},
			want:    empty,
			wantErr: false,
		},
		{
			name: "Invalid regex pattern",
			args: args{
				source:    "./testdata",
				match:     "[",
				extension: "markdown",
			},
			want:    empty,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := getMatches(tt.args.source, tt.args.match, tt.args.extension)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetMatches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMatches() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestParse covers the input-validation contract of the public Parse
// entrypoint: a missing name must error out before any schema or docs
// work is attempted.
//
// End-to-end provider parsing used to live here too, via live git clones
// of terraform-provider-{aws,azurerm,google}. Those subtests were dropped
// after the schema-based rewrite because (a) Parse() now takes the schema
// path first and ignores the codebase on success, so the clone was pure
// overhead, and (b) the pass/fail signal didn't tell us which path ran
// — environment-dependent coverage is worse than none. The schema path
// is covered by schema_test.go; the docs path is covered below via
// parseFromDocs directly.
func TestParse(t *testing.T) {
	t.Parallel()

	if err := Parse("", ""); err == nil {
		t.Error("Parse(\"\", \"\") expected an error, got nil")
	}
}

// TestParseFromDocs_EmptyCodebase pins the input-validation guard on the
// docs fallback path. Previously this was reached via Parse("", "azure"),
// but Parse now takes the schema path first — so to exercise this guard
// deterministically we call parseFromDocs directly.
func TestParseFromDocs_EmptyCodebase(t *testing.T) {
	t.Parallel()

	if _, err := parseFromDocs("", "azurerm"); err == nil {
		t.Error("parseFromDocs(\"\", ...) expected an error, got nil")
	}
}
