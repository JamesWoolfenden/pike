package pike

import (
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
