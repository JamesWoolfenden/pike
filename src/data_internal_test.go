package pike

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fileStringEmptyError_Error(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Invoke", "no file provided"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &fileStringEmptyError{}
			assert.Equalf(t, tt.want, e.Error(), "Error()")
		})
	}
}
