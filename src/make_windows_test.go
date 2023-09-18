//go:build auth
// +build auth

package pike_test

import (
	"testing"

	pike "github.com/jameswoolfenden/pike/src"
)

func TestMakeWindows(t *testing.T) {
	t.Parallel()

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
			t.Parallel()
			if _, err := pike.Make(tt.args.directory); (err != nil) != tt.wantErr {
				t.Errorf("Make() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
