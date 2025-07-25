//go:build auth
// +build auth

package pike_test

import (
	"testing"

	pike "github.com/jameswoolfenden/pike/src"
)

func TestInvokeGithubDispatchEvent(t *testing.T) {
	t.Parallel()

	type args struct {
		repository       string
		workflowFileName string
		branch           string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "wrong branch", args: args{repository: "https://github.com/jameswoolfenden/pike", workflowFileName: "resources.yml", branch: "main"}, wantErr: true},
		{name: "branch", args: args{repository: "https://github.com/jameswoolfenden/pike", workflowFileName: "resources.yml", branch: "master"}, wantErr: false},
		{name: "guff", args: args{repository: "github.guff/jameswoolfenden/pike", workflowFileName: "resources", branch: "main"}, wantErr: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := pike.InvokeGithubDispatchEvent(tt.args.repository, tt.args.workflowFileName, tt.args.branch); (err != nil) != tt.wantErr {
				t.Errorf("InvokeGithubDispatchEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
