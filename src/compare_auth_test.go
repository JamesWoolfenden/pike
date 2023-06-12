//go:build auth

package pike

import (
	"testing"
)

func TestCompare(t *testing.T) {
	type args struct {
		directory string
		arn       string
		init      bool
	}

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"fail", args{"../testdata/scan/examples/simple", "arn:aws:iam::680235478471:policy/basic", true}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Compare(tt.args.directory, tt.args.arn, tt.args.init)
			if (err != nil) != tt.wantErr {
				t.Errorf("Compare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
