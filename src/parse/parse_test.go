package parse

import (
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/go-git/go-git/v5"
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
		{name: "None", args: args{path: "../mapping", extension: "go"}},
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

func setup(cloud string) {
	log.Println("setup")

	switch cloud {
	case "aws":
		_, _ = git.PlainClone("./terraform-provider-aws", false, &git.CloneOptions{
			URL:      "https://github.com/hashicorp/terraform-provider-aws",
			Progress: os.Stdout,
			Depth:    1,
		})
	case "azurerm":
		_, _ = git.PlainClone("./terraform-provider-azurerm", false, &git.CloneOptions{
			URL:      "https://github.com/hashicorp/terraform-provider-azurerm",
			Progress: os.Stdout,
			Depth:    1,
		})
	case "google":
		_, _ = git.PlainClone("./terraform-provider-google", false, &git.CloneOptions{
			URL:      "https://github.com/hashicorp/terraform-provider-google",
			Progress: os.Stdout,
			Depth:    1,
		})
	}
}

func teardown(cloud string) {
	log.Println("teardown")
	switch cloud {
	case "aws":
		_ = os.RemoveAll("./terraform-provider-aws")
	case "azurerm":
		_ = os.RemoveAll("./terraform-provider-azurerm")
	case "google":
		_ = os.RemoveAll("./terraform-provider-google")
	}
}

func TestParse(t *testing.T) {
	type args struct {
		codebase string
		name     string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "aws", args: args{codebase: "./terraform-provider-aws", name: "aws"}},
		{name: "azure", args: args{codebase: "./terraform-provider-azurerm", name: "azurerm"}},
		{name: "google", args: args{codebase: "./terraform-provider-google", name: "google"}},
		{
			name:    "Empty codebase",
			args:    args{codebase: "", name: "azure"},
			wantErr: true,
		},
	}

	for _, tt := range tests {

		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			setup(tt.args.name)

			if err := Parse(tt.args.codebase, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}

			teardown(tt.args.name)
		})
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
