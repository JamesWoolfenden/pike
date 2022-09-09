package pike

import (
	_ "embed"
	"reflect"
	"testing"
)

func TestNewAWSPolicy(t *testing.T) {
	type args struct {
		Actions []string
	}
	tests := []struct {
		name string
		args args
		want Policy
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAWSPolicy(tt.args.Actions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAWSPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPolicy(t *testing.T) {
	type args struct {
		actions Sorted
		output  string
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
			got, err := GetPolicy(tt.args.actions, tt.args.output)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPolicy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAWSPolicy(t *testing.T) {
	type args struct {
		Permissions []string
		output      string
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
			got, err := AWSPolicy(tt.args.Permissions, tt.args.output)
			if (err != nil) != tt.wantErr {
				t.Errorf("AWSPolicy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AWSPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unique(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unique(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unique() = %v, want %v", got, tt.want)
			}
		})
	}
}
