package pike

import "testing"

func TestReadme(t *testing.T) {
	type args struct {
		dirName    string
		output     string
		init       bool
		autoAppend bool
		exclude    []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Readme(tt.args.dirName, tt.args.output, tt.args.init, tt.args.autoAppend, tt.args.exclude); (err != nil) != tt.wantErr {
				t.Errorf("Readme() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
