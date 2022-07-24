package pike

import (
	"reflect"
	"testing"
)

func TestGetAWSPermissions(t *testing.T) {

	data:=make([]inferface, 3)
	data[0]="ec2:AuthorizeSecurityGroupIngress"
	data[1]="ec2:AuthorizeSecurityGroupEgress"
        data[2]="ec2:CreateSecurityGroup"

    

	type args struct {
		result template
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{"First",args{},data}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAWSPermissions(tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAWSPermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAttributes(t *testing.T) {
	type args struct {
		result template
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAttributes(tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAttributes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		s []string
		e string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"pass",args{[]string{"a","b","c"},"c"},true},
		{"fail",args{[]string{"a","b","c"},"d"},false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPermissionMap(t *testing.T) {
	type args struct {
		raw        []byte
		attributes []string
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
			if got := GetPermissionMap(tt.args.raw, tt.args.attributes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPermissionMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
