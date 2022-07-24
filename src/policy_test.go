package pike

import (
	"reflect"
	"testing"
)

func TestNewPolicy(t *testing.T) {
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
			if got := NewPolicy(tt.args.Actions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPolicy(t *testing.T) {
	type args struct {
		actions []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetPolicy(tt.args.actions)
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
