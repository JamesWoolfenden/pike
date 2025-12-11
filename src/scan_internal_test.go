package pike

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_emptyIACError_Error(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"emptyIACError", "no IAC found"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &emptyIACError{}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_makePolicyError_Error(t *testing.T) {
	type fields struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"policy error", fields{err: errors.New("invalid policy")}, "failed to make policy invalid policy"},
		{"nil error", fields{err: nil}, "failed to make policy <nil>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &makePolicyError{
				err: tt.fields.err,
			}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_emptyScanLocationError_Error(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"emptyScanLocationError", "no scan location"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &emptyScanLocationError{}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_makeDirectoryError_Error(t *testing.T) {
	type fields struct {
		directory string
		err       error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"directory error", fields{directory: "/tmp/test", err: errors.New("permission denied")}, "failed to make directory /tmp/test permission denied"},
		{"empty directory", fields{directory: "", err: errors.New("no path")}, "failed to make directory  no path"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &makeDirectoryError{
				directory: tt.fields.directory,
				err:       tt.fields.err,
			}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_locateTerraformError_Error(t *testing.T) {
	type fields struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"not found", fields{err: errors.New("not in PATH")}, "failed to find Terraform not in PATH"},
		{"nil error", fields{err: nil}, "failed to find Terraform <nil>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &locateTerraformError{
				err: tt.fields.err,
			}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_terraformExecError_Error(t *testing.T) {
	type fields struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"exec error", fields{err: errors.New("command failed")}, "Terraform execution error command failed"},
		{"nil error", fields{err: nil}, "Terraform execution error <nil>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &terraformExecError{
				err: tt.fields.err,
			}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_terraformInitError_Error(t *testing.T) {
	type fields struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"init error", fields{err: errors.New("init failed")}, "Terraform init error init failed"},
		{"nil error", fields{err: nil}, "Terraform init error <nil>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &terraformInitError{
				err: tt.fields.err,
			}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_readDirectoryError_Error(t *testing.T) {
	type fields struct {
		directory string
		err       error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"read error", fields{directory: "/tmp/test", err: errors.New("permission denied")}, "failed to read directory /tmp/test permission denied"},
		{"empty directory", fields{directory: "", err: errors.New("no path")}, "failed to read directory  no path"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &readDirectoryError{
				directory: tt.fields.directory,
				err:       tt.fields.err,
			}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_absolutePathError_Error(t *testing.T) {
	type fields struct {
		directory string
		err       error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"path error", fields{directory: "relative/path", err: errors.New("invalid path")}, "failed to get absolute path relative/path invalid path"},
		{"empty directory", fields{directory: "", err: errors.New("empty")}, "failed to get absolute path  empty"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &absolutePathError{
				directory: tt.fields.directory,
				err:       tt.fields.err,
			}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_getTFError_Error(t *testing.T) {
	type fields struct {
		directory string
		err       error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"tf error", fields{directory: "/tmp/test", err: errors.New("no tf files")}, "failed to get Terraform templates /tmp/test no tf files"},
		{"empty directory", fields{directory: "", err: errors.New("empty")}, "failed to get Terraform templates  empty"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &getTFError{
				directory: tt.fields.directory,
				err:       tt.fields.err,
			}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}

func Test_getPolicyError_Error(t *testing.T) {
	type fields struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"policy error", fields{err: errors.New("invalid policy")}, "failed to get policy invalid policy"},
		{"nil error", fields{err: nil}, "failed to get policy <nil>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &getPolicyError{
				err: tt.fields.err,
			}
			assert.Equalf(t, tt.want, m.Error(), "Error()")
		})
	}
}
