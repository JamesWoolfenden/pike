package pike

import (
	"reflect"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestScan(t *testing.T) {
	type args struct {
		dirname string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"aws", args{"../terraform/aws/backup"}, false},
		{"gcp", args{"../terraform/gcp/backup"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Scan(tt.args.dirname, "json", "", false, nil); (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetTF(t *testing.T) {
	type args struct {
		dirname  string
		recurse  bool
		excludes *cli.StringSlice
	}

	excludes := cli.NewStringSlice("simple")
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"first", args{"../testdata/scan/examples/simple", false, nil}, []string{"../testdata/scan/examples/simple/aws_s3_bucket.pike.tf"}, false},
		{"empty", args{"../testdata/scan", false, nil}, nil, false},
		{"recurse", args{"../testdata/scan", true, nil}, []string{"../testdata/scan/examples/simple/aws_s3_bucket.pike.tf"}, false},
		{"recurse-exclude", args{"../testdata/scan", true, excludes}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTF(tt.args.dirname, tt.args.recurse, tt.args.excludes)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringInSlice(t *testing.T) {
	type args struct {
		a    string
		list []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"pass", args{"a", []string{"a", "b", "c"}}, true},
		{"fail", args{"d", []string{"a", "b", "c"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringInSlice(tt.args.a, tt.args.list); got != tt.want {
				t.Errorf("stringInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInit(t *testing.T) {
	type args struct {
		dirName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Init(tt.args.dirName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Init() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakePolicy(t *testing.T) {
	type args struct {
		dirName  string
		file     string
		init     bool
		excludes *cli.StringSlice
	}
	tests := []struct {
		name    string
		args    args
		want    OutputPolicy
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MakePolicy(tt.args.dirName, tt.args.file, tt.args.init, tt.args.excludes)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakePolicy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakePolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetHCLType(t *testing.T) {
	type args struct {
		resourceName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHCLType(tt.args.resourceName); got != tt.want {
				t.Errorf("GetHCLType() = %v, want %v", got, tt.want)
			}
		})
	}
}
