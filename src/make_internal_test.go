package pike

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_arnNotFoundInStateError_Error(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Invoke", "no arn found in state"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &arnNotFoundInStateError{}
			assert.Equalf(t, tt.want, e.Error(), "Error()")
		})
	}
}

func Test_castToStringError_Error(t *testing.T) {
	type fields struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Invoke", fields{"test"}, "cannot convert test to a string"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &castToStringError{
				value: tt.fields.value,
			}
			assert.Equalf(t, tt.want, e.Error(), "Error()")
		})
	}
}
