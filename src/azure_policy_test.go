package pike

import (
	_ "embed"
	"testing"
)

func TestAZUREPolicy(t *testing.T) {
	type args struct {
		permissions []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AZUREPolicy(tt.args.permissions)
			if (err != nil) != tt.wantErr {
				t.Errorf("AZUREPolicy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AZUREPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}
