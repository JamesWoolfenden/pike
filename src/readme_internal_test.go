package pike

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fileDoesNotExistError_Error(t *testing.T) {
	type fields struct {
		file string
		err  error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"file not found", fields{file: "/tmp/README.md", err: errors.New("no such file")}, "file /tmp/README.md does not exist no such file"},
		{"empty file", fields{file: "", err: errors.New("empty path")}, "file  does not exist empty path"},
		{"nil error", fields{file: "README.md", err: nil}, "file README.md does not exist <nil>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := fileDoesNotExistError{
				file: tt.fields.file,
				err:  tt.fields.err,
			}
			assert.Equalf(t, tt.want, e.Error(), "Error()")
		})
	}
}

func Test_replaceSectionError_Error(t *testing.T) {
	type fields struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"replace error", fields{err: errors.New("section not found")}, "failed to replace section section not found"},
		{"nil error", fields{err: nil}, "failed to replace section <nil>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &replaceSectionError{
				err: tt.fields.err,
			}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_tfPolicyFormatError_Error(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"tfPolicyFormatError", "output formats are Terraform and JSON"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &tfPolicyFormatError{}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}
