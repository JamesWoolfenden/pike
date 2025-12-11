package pike

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parameterNilError_Error(t *testing.T) {
	type fields struct {
		parameter string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"permissionMap nil", fields{parameter: "permissionMap"}, "permissionMap was nil"},
		{"empty parameter", fields{parameter: ""}, " was nil"},
		{"action parameter", fields{parameter: "action"}, "action was nil"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &parameterNilError{
				parameter: tt.fields.parameter,
			}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}
