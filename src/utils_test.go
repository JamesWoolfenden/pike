package pike

import (
	"errors"
	"os"
	"reflect"
	"strings"
	"testing"
)

func Test_randSeq(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"correct", args{8}, 8},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := RandSeq(tt.args.n); len(got) != tt.want {
				t.Errorf("RandSeq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceSection(t *testing.T) {
	t.Parallel()

	type args struct {
		source  string
		middle  string
		autoadd bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"found",
			args{
				"testdata/test-readme.md",
				"Any thing in here",
				false,
			},
			false,
		},
		{
			"file not found",
			args{
				"testdata/bogus-readme.md",
				"Any thing in here",
				false,
			},
			true,
		},
		{
			"empty file",
			args{
				"testdata/test-readme-empty.md",
				"Any thing in here",
				false,
			},
			true,
		},
		{
			"wrong format",
			args{
				"testdata/test-readme-wrong.md",
				"Any thing in here",
				false,
			},
			true,
		},
		{
			"half man half biscuit",
			args{
				"testdata/biscuit.md",
				"Any thing in here",
				false,
			},
			true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if err := ReplaceSection(tt.args.source, tt.args.middle, tt.args.autoadd); (err != nil) != tt.wantErr {
				t.Errorf("ReplaceSection() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fileExists(t *testing.T) {
	t.Parallel()

	type args struct {
		filename string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"found", args{"testdata/test-readme-empty.md"}, true},
		{"bogus", args{"testdata/bogus-readme-empty.md"}, false},
		{"half-man half-biscuit", args{"testdata/biscuit.md"}, true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := FileExists(tt.args.filename); got != tt.want {
				t.Errorf("FileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandSeq(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name    string
		args    args
		notWant string
	}{
		{"correct", args{8}, "jslbVKIr"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := RandSeq(tt.args.n); got == tt.notWant {
				t.Errorf("RandSeq() = %v, want %v", got, tt.notWant)
			}
		})
	}
}

func TestAlmostEqual(t *testing.T) {
	t.Parallel()

	type args struct {
		a float64
		b float64
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "exact equal",
			args: args{
				a: 1.0,
				b: 1.0,
			},
			want: true,
		},
		{
			name: "not equal beyond threshold",
			args: args{
				a: 1.1,
				b: 1.0,
			},
			want: false,
		},
		{
			name: "negative numbers equal",
			args: args{
				a: -1.0,
				b: -1.0,
			},
			want: true,
		},
		{
			name: "zero values",
			args: args{
				a: 0.0,
				b: 0.0,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := AlmostEqual(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("AlmostEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvVariableNotSetError_Error(t *testing.T) {
	type fields struct {
		Key string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"fail", fields{"key"}, "environment variable key not set"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EnvVariableNotSetError{
				Key: tt.fields.Key,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getEnv(t *testing.T) {
	type args struct {
		key string
	}

	err := os.Setenv("fortest", "value")
	if err != nil {
		return
	}

	tests := []struct {
		name    string
		args    args
		want    *string
		wantErr bool
	}{
		{"fail", args{"key"}, nil, true},
		{"pass", args{"fortest"}, &[]string{"value"}[0], false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetEnv(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEnv() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteFileError_Error(t *testing.T) {
	tests := []struct {
		name     string
		file     string
		err      error
		expected string
	}{
		{
			name:     "normal file and error",
			file:     "/path/to/file.txt",
			err:      errors.New("permission denied"),
			expected: "failed to write file /path/to/file.txt permission denied",
		},
		{
			name:     "empty file path",
			file:     "",
			err:      errors.New("some error"),
			expected: "failed to write file  some error",
		},
		{
			name:     "nil error",
			file:     "/path/to/file.txt",
			err:      nil,
			expected: "failed to write file /path/to/file.txt <nil>",
		},
		{
			name:     "file path with spaces",
			file:     "/path/to/file with spaces.txt",
			err:      errors.New("write failed"),
			expected: "failed to write file /path/to/file with spaces.txt write failed",
		},
		{
			name:     "file path with special characters",
			file:     "/path/to/file-name_123.txt",
			err:      errors.New("disk full"),
			expected: "failed to write file /path/to/file-name_123.txt disk full",
		},
		{
			name:     "long file path",
			file:     strings.Repeat("/very/long/path", 10) + "/file.txt",
			err:      errors.New("timeout"),
			expected: "failed to write file " + strings.Repeat("/very/long/path", 10) + "/file.txt timeout",
		},
		{
			name:     "error with newlines",
			file:     "/path/to/file.txt",
			err:      errors.New("error\nwith\nnewlines"),
			expected: "failed to write file /path/to/file.txt error\nwith\nnewlines",
		},
		{
			name:     "both file and error empty/nil",
			file:     "",
			err:      nil,
			expected: "failed to write file  <nil>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &writeFileError{
				file: tt.file,
				err:  tt.err,
			}

			result := e.Error()

			if result != tt.expected {
				t.Errorf("writeFileError.Error() = %q, expected %q", result, tt.expected)
			}
		})
	}
}

func TestCustomErrors(t *testing.T) {
	t.Run("readFileError", func(t *testing.T) {
		err := &readFileError{file: "test.txt", err: errors.New("permission denied")}
		expected := "failed to read file test.txt permission denied"
		if err.Error() != expected {
			t.Errorf("Expected: %s, got: %s", expected, err.Error())
		}
	})

	t.Run("delimiterMismatchError", func(t *testing.T) {
		err := &delimiterMismatchError{}
		expected := "pike delimiters mismatch in Readme"
		if err.Error() != expected {
			t.Errorf("Expected: %s, got: %s", expected, err.Error())
		}
	})

	t.Run("delimiterHooksMissingError", func(t *testing.T) {
		err := &delimiterHooksMissingError{}
		expected := "pike hooks delimiter missing in Readme,  consider using the flag -auto"
		if err.Error() != expected {
			t.Errorf("Expected: %s, got: %s", expected, err.Error())
		}
	})

	t.Run("writeFileError", func(t *testing.T) {
		err := &writeFileError{file: "output.txt", err: errors.New("disk full")}
		expected := "failed to write file output.txt disk full"
		if err.Error() != expected {
			t.Errorf("Expected: %s, got: %s", expected, err.Error())
		}
	})
}
