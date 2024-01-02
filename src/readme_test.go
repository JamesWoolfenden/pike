package pike_test

import (
	"testing"

	pike "github.com/jameswoolfenden/pike/src"
)

func TestReadme(t *testing.T) {
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
		{name: "missing",
			args:    args{dirName: "../testdata/readme/missing", output: "json", init: true, autoAppend: false},
			wantErr: true},
		{name: "missing-tf",
			args:    args{dirName: "../testdata/readme/missing", output: "terraform", init: true, autoAppend: false},
			wantErr: true},
		{name: "exists", args: args{dirName: "../testdata/readme/exists", output: "terraform", init: true, autoAppend: false}},
		{name: "exists-json",
			args: args{dirName: "../testdata/readme/exists", output: "json", init: true, autoAppend: false}},
		{name: "wrong output",
			args:    args{dirName: "../testdata/readme/exists", output: "cdk", init: true, autoAppend: false},
			wantErr: true},
		{name: "exists-noTags",
			args:    args{dirName: "../testdata/readme/exists-notags", output: "terraform", init: true},
			wantErr: true},
		{name: "empty", args: args{dirName: "../testdata/readme/empty", output: "terraform", init: true},
			wantErr: true},
		{name: "exists-noInit", args: args{dirName: "../testdata/readme/exists", output: "terraform"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := pike.Readme(tt.args.dirName, tt.args.output, tt.args.init, tt.args.autoAppend); (err != nil) != tt.wantErr {
				t.Errorf("Readme() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
