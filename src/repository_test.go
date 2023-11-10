package pike_test

import (
	"testing"

	pike "github.com/jameswoolfenden/pike/src"
)

func TestRepository(t *testing.T) {
	t.Parallel()

	type args struct {
		repository  string
		destination string
		directory   string
		output      string
		init        bool
		write       bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "pass",
			args: args{
				repository:  "https://github.com/JamesWoolfenden/terraform-aws-codebuild",
				destination: "../../terraform-gcp-redis",
				directory:   ".",
				output:      "terraform",
				init:        true,
				write:       true,
			},
		},
		{
			name: "nothing",
			args: args{
				repository:  "https://github.com/JamesWoolfenden/pike",
				destination: "../../placeholder",
				directory:   ".",
				output:      "terraform",
				init:        true,
				write:       true,
			}, wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := pike.Repository(tt.args.repository, tt.args.destination, tt.args.directory, tt.args.output, tt.args.init, tt.args.write, false); (err != nil) != tt.wantErr {
				t.Errorf("Repository() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
