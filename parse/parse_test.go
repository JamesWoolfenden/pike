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
		{name: "None", args: args{path: "../src/mapping", extension: "go"}},
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
		m map[string]bool
	}

	sample := map[string]bool{
		"first": true,
	}

	nothing := map[string]bool{}
	var bumpkis []string

	myKeys := []string{"first"}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "pass", args: args{sample}, want: myKeys},
		{name: "nil", args: args{nothing}},

		{
			name: "Non-empty map", args: args{
				m: map[string]bool{
					"key1": true,
					"key2": true,
				},
			},
			want: []string{"key1", "key2"},
		},
		{
			name: "Empty map",
			args: args{m: map[string]bool{}},
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

func Test_add(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
		m map[string]bool
		a []string
	}

	myMap := map[string]bool{
		"aws_s3_bucket": true,
	}

	wantMap := map[string]bool{
		"aws_ami":       true,
		"aws_s3_bucket": true,
	}

	tests := []struct {
		name  string
		args  args
		want  []string
		want1 map[string]bool
	}{
		{
			name: "pass",
			args: args{s: "aws_ami", m: myMap, a: []string{"aws_s3_bucket"}},
			want: []string{"aws_s3_bucket", "aws_ami"}, want1: wantMap,
		},
		{
			name: "duplicate",
			args: args{s: "aws_ami", m: wantMap, a: []string{"aws_s3_bucket", "aws_ami"}},
			want: []string{"aws_s3_bucket", "aws_ami"}, want1: wantMap,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, got1 := add(tt.args.s, tt.args.m, tt.args.a)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("add() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("add() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name      string
		s         string
		m         map[string]bool
		a         []string
		wantSlice []string
		wantMap   map[string]bool
	}{
		{
			name:      "New element",
			s:         "test",
			m:         map[string]bool{},
			a:         []string{},
			wantSlice: []string{"test"},
			wantMap:   map[string]bool{"test": true},
		},
		{
			name:      "Duplicate element",
			s:         "test",
			m:         map[string]bool{"test": true},
			a:         []string{"test"},
			wantSlice: []string{"test"},
			wantMap:   map[string]bool{"test": true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSlice, gotMap := add(tt.s, tt.m, tt.a)

			if !reflect.DeepEqual(gotSlice, tt.wantSlice) {
				t.Errorf("add() gotSlice = %v, want %v", gotSlice, tt.wantSlice)
			}

			if !reflect.DeepEqual(gotMap, tt.wantMap) {
				t.Errorf("add() gotMap = %v, want %v", gotMap, tt.wantMap)
			}
		})
	}
}
