package pike

import "testing"

func TestInvokeGithubDispatchEvent(t *testing.T) {
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
		t.Run(tt.name, func(t *testing.T) {
			if err := InvokeGithubDispatchEvent(tt.args.repository, tt.args.workflowFileName, tt.args.branch); (err != nil) != tt.wantErr {
				t.Errorf("InvokeGithubDispatchEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
