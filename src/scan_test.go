package pike

import (
	"io/fs"
	"reflect"
	"testing"
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Scan(tt.args.dirname); (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetTF(t *testing.T) {
	type args struct {
		dirname string
	}
	tests := []struct {
		name    string
		args    args
		want    []fs.FileInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTF(tt.args.dirname)
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
