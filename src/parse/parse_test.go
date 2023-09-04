package parse

import (
	"reflect"
	"testing"
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

	type args struct {
		codebase string
		name     string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		//{"aws", args{"e:/code/terraform-provider-aws", "aws"}, false},
		//{"azure", args{"e:/code/terraform-provider-azurerm", "azurerm"}, false},
		//{"google", args{"e:/code/terraform-provider-google", "google"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Parse(tt.args.codebase, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_add(t *testing.T) {
	type args struct {
		s string
		m map[string]bool
		a []string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 map[string]bool
	}{
		// TODO: Add test cases.
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
