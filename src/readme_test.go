package pike

import (
	"testing"

	"github.com/urfave/cli/v2"
)

func TestReadme(t *testing.T) {
	type args struct {
		dirName    string
		output     string
		init       bool
		autoAppend bool
		exclude    *cli.StringSlice
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"missing", args{"../testdata/readme/missing", "json", true, false, nil}, true},
		{"missing-tf", args{"../testdata/readme/missing", "tf", true, false, nil}, true},
		{"exists", args{"../testdata/readme/exists", "tf", true, false, nil}, false},
		{"exists-notags", args{"../testdata/readme/exists-notags", "tf", true, false, nil}, true},
		{"empty", args{"../testdata/readme/empty", "tf", true, false, nil}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Readme(tt.args.dirName, tt.args.output, tt.args.init, tt.args.autoAppend, tt.args.exclude); (err != nil) != tt.wantErr {
				t.Errorf("Readme() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
