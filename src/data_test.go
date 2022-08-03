package pike

import (
	//"io/fs"
	"reflect"
	"testing"
)

// func TestGetResources(t *testing.T) {
// 	type args struct {
// 		file    fs.FileInfo
// 		dirname string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want []Resource
// 	}{
// 		{"missing", args{nil, "./terraform/"}, []Resource{}},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := GetResources(tt.args.file, tt.args.dirname); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("GetResources() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestGetProvider(t *testing.T) {
	type args struct {
		resource string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"pass", args{"aws_something.tf"}, "aws"},
		{"nothing", args{"something.tf"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetProvider(tt.args.resource); got != tt.want {
				t.Errorf("GetProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPermission(t *testing.T) {
	type args struct {
		result template
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := GetPermission(tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPermission() = %v, want %v", got, tt.want)
			}
		})
	}
}
