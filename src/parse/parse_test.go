package parse

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestGetGoFiles(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetGoFiles(tt.args.path, tt.args.extension)
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

	myKeys := []string{"first"}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "pass", args: args{sample}, want: myKeys},
		{name: "nil", args: args{nothing}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := GetKeys(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMatches(t *testing.T) {
	t.Parallel()

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
		{name: "pass"},
		{name: "go", args: args{source: "./testdata", match: "(aws_.*?)", extension: "go"}, want: []string{"aws_"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := GetMatches(tt.args.source, tt.args.match, tt.args.extension)
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

//func TestParse(t *testing.T) {
//	t.Parallel()
//
//	_, err := git.PlainClone("./terraform-provider-aws", false, &git.CloneOptions{
//		URL:      "https://github.com/hashicorp/terraform-provider-aws",
//		Progress: os.Stdout,
//		Depth:    1,
//	})
//
//	_, err = git.PlainClone("./terraform-provider-azurerm", false, &git.CloneOptions{
//		URL:      "https://github.com/hashicorp/terraform-provider-azurerm",
//		Progress: os.Stdout,
//		Depth:    1,
//	})
//
//	_, err = git.PlainClone("./terraform-provider-google", false, &git.CloneOptions{
//		URL:      "https://github.com/hashicorp/terraform-provider-google",
//		Progress: os.Stdout,
//		Depth:    1,
//	})
//
//	type args struct {
//		codebase string
//		name     string
//	}
//
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//	}{
//		{name: "aws", args: args{codebase: "./terraform-provider-aws", name: "aws"}},
//		{name: "azure", args: args{codebase: "./terraform-provider-azurerm", name: "azurerm"}},
//		{name: "google", args: args{codebase: "./terraform-provider-google", name: "google"}},
//	}
//
//	for _, tt := range tests {
//		tt := tt
//		t.Run(tt.name, func(t *testing.T) {
//			t.Parallel()
//			if err := Parse(tt.args.codebase, tt.args.name); (err != nil) != tt.wantErr {
//				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

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
		{name: "pass",
			args: args{s: "aws_ami", m: myMap, a: []string{"aws_s3_bucket"}},
			want: []string{"aws_s3_bucket", "aws_ami"}, want1: wantMap},
		{name: "duplicate",
			args: args{s: "aws_ami", m: wantMap, a: []string{"aws_s3_bucket", "aws_ami"}},
			want: []string{"aws_s3_bucket", "aws_ami"}, want1: wantMap},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
