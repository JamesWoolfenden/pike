package pike_test

import (
	"testing"

	pike "github.com/jameswoolfenden/pike/src"
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
			if got := pike.RandSeq(tt.args.n); len(got) != tt.want {
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
			if err := pike.ReplaceSection(tt.args.source, tt.args.middle, tt.args.autoadd); (err != nil) != tt.wantErr {
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
			if got := pike.FileExists(tt.args.filename); got != tt.want {
				t.Errorf("FileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
