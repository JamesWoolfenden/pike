package parse

import (
	"os"
	"reflect"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/rs/zerolog/log"
)

func TestGetGoFiles(t *testing.T) {
	type args struct {
		path      string
		extension string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
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
	type args struct {
		m map[string]bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetKeys(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMatches(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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

func TestParse(t *testing.T) {
	t.Parallel()

	_, err := git.PlainClone("./terraform-provider-aws", false, &git.CloneOptions{
		URL:      "https://github.com/hashicorp/terraform-provider-aws",
		Progress: os.Stdout,
		Depth:    1,
	})
	log.Print(err)

	_, err = git.PlainClone("./terraform-provider-azurerm", false, &git.CloneOptions{
		URL:      "https://github.com/hashicorp/terraform-provider-azurerm",
		Progress: os.Stdout,
		Depth:    1,
	})

	log.Print(err)
	_, err = git.PlainClone("./terraform-provider-google", false, &git.CloneOptions{
		URL:      "https://github.com/hashicorp/terraform-provider-google",
		Progress: os.Stdout,
		Depth:    1,
	})
	log.Print(err)

	type args struct {
		codebase string
		name     string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"aws", args{"./terraform-provider-aws", "aws"}, false},
		{"azure", args{"./terraform-provider-azurerm", "azurerm"}, false},
		{"google", args{"./terraform-provider-google", "google"}, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := Parse(tt.args.codebase, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
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
		{"pass",
			args{"aws_ami", myMap, []string{"aws_s3_bucket"}},
			[]string{"aws_s3_bucket", "aws_ami"}, wantMap},
		{"duplicate",
			args{"aws_ami", wantMap, []string{"aws_s3_bucket", "aws_ami"}},
			[]string{"aws_s3_bucket", "aws_ami"}, wantMap},
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
