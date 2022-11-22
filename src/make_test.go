//go:build auth

package pike

import (
	"testing"
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
		{"linux-random", args{"testdata/scan/examples/random"}, true},
		{"linux-full", args{"testdata/scan/examples/simple"}, false},
		{"linux-fail", args{"../modules/aws/terraform-aws-budget/rubbish"}, true},
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
		wantErr bool
	}{
		{name: "fail", args: args{"asdasd"}, wantErr: true},
		{name: "pass-random", args: args{"testdata/scan/examples/random"}, wantErr: false},
		{name: "pass", args: args{"testdata/scan/examples/simple"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tfApply(tt.args.policyPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("tfApply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestApply(t *testing.T) {
	type args struct {
		target string
		region string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"pass", args{"testdata/scan/examples/simple", "eu-west-2"}, false},
		{"pass", args{"testdata/scan/examples/balls", "eu-west-2"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Apply(tt.args.target, tt.args.region); (err != nil) != tt.wantErr {
				t.Errorf("Apply() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}
