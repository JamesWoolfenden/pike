//go:build auth
// +build auth

package pike

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-exec/tfexec"
)

func TestMake(t *testing.T) {
	type args struct {
		directory string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"windows-full", args{"testdata\\scan\\examples\\simple"}, false},
		{"windows-fail", args{"e:\\code\\modules\\aws\\terraform-aws-budget\\rubbish"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := Make(tt.args.directory); (err != nil) != tt.wantErr {
				t.Errorf("Make() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_tfApply(t *testing.T) {
	type args struct {
		policyPath string
	}
	tests := []struct {
		name    string
		args    args
		want    *tfexec.Terraform
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tfApply(tt.args.policyPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("tfApply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tfApply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApply(t *testing.T) {
	type args struct {
		target string
	}
	tests := []struct {
		name string
		args args
	}{
		{"pass", args{"testdata/scan/examples/simple"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Apply(tt.args.target)
		})
	}
}
