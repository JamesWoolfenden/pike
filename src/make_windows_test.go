//go:build auth
// +build auth

package pike

import (
	"testing"
)

func TestMakeWindows(t *testing.T) {
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
