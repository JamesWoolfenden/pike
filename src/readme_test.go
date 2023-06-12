package pike_test

import (
	"testing"

	pike "github.com/jameswoolfenden/pike/src"
)

func TestReadme(t *testing.T) {
	t.Parallel()

	type args struct {
		dirName    string
		output     string
		init       bool
		autoAppend bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"missing", args{"../testdata/readme/missing", "json", true, false}, true},
		{"missing-tf", args{"../testdata/readme/missing", "terraform", true, false}, true},
		{"exists", args{"../testdata/readme/exists", "terraform", true, false}, false},
		{"exists-json", args{"../testdata/readme/exists", "json", true, false}, false},
		{"wrong output", args{"../testdata/readme/exists", "cdk", true, false}, true},
		{"exists-notags", args{"../testdata/readme/exists-notags", "terraform", true, false}, true},
		{"empty", args{"../testdata/readme/empty", "terraform", true, false}, true},
		{"exists-noinit", args{"../testdata/readme/exists", "terraform", false, false}, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := pike.Readme(tt.args.dirName, tt.args.output, tt.args.init, tt.args.autoAppend); (err != nil) != tt.wantErr {
				t.Errorf("Readme() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
