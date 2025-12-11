package pike

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_emptyRegionError_Error(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"emptyRegionError", "region cannot be empty"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := emptyRegionError{}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_iamRoleEmptyError_Error(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"iamRoleEmptyError", "iamRole cannot be empty"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := iamRoleEmptyError{}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_getAWSCredentialsError_Error(t *testing.T) {
	type fields struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"credential error", fields{err: errors.New("invalid credentials")}, "failed to get AWS credentials: invalid credentials"},
		{"access denied", fields{err: errors.New("access denied")}, "failed to get AWS credentials: access denied"},
		{"nil error", fields{err: nil}, "failed to get AWS credentials: <nil>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := getAWSCredentialsError{
				err: tt.fields.err,
			}
			assert.Equalf(t, tt.want, e.Error(), "Error()")
		})
	}
}
