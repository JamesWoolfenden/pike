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
		// TODO: Add test cases.
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
