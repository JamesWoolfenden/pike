//go:build auth
// +build auth

package pike

import (
	"reflect"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/service/sts"
)

func Test_getAWSCredentials(t *testing.T) {
	type args struct {
		IAMRole string
	}
	arn := "arn:aws:sts::680235478471:assumed-role/terraform_pike_20220924074025950900000002/testAssumeRoleSession"
	assumeroleid := "AROAZ4YJRVXDRTVF6YDQI:testAssumeRoleSession"
	find := &sts.AssumedRoleUser{Arn: &arn, AssumedRoleId: &assumeroleid}

	tests := []struct {
		name    string
		args    args
		want    *sts.AssumedRoleUser
		wantErr bool
	}{
		{"pass", args{"arn:aws:iam::680235478471:role/terraform_pike_20220924074025950900000002"}, find, false},
		{"denied", args{"arn:aws:iam::123456789012:role/demo"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getAWSCredentials(tt.args.IAMRole, "eu-west-2")
			if (err != nil) != tt.wantErr {
				t.Errorf("getAWSCredentails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want == nil {
				if got != nil {
					t.Errorf("getAWSCredentails() = %v, want %v", got, tt.want)
				}
				return
			}
			if !reflect.DeepEqual(got.AssumedRoleUser, tt.want) {
				t.Errorf("getAWSCredentails() = %v, want %v", got.AssumedRoleUser, tt.want)
			}
		})
	}
}

func Test_setAWSAuth(t *testing.T) {
	type args struct {
		iamRole string
		Region  string
	}

	arghh := "User: arn:aws:iam::680235478471:user/jameswoolfenden is not authorized to perform: sts:AssumeRole on " +
		"resource: arn:aws:iam::123456789012:role/demo"
	// myErr := awserr.NewRequestFailure(arghh, 403, "")
	tests := []struct {
		name string
		args args
		want *string
	}{
		{"pass", args{"arn:aws:iam::680235478471:role/terraform_pike_20221003170618461900000002", "eu-west-2"}, nil},
		{"denied", args{"arn:aws:iam::123456789012:role/demo", "eu-west-2"}, &arghh},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// could inherit temp set from before
			unSetAWSAuth()
			got := setAWSAuth(tt.args.iamRole, tt.args.Region)

			if tt.want == nil && got != nil {
				t.Errorf("setAWSAuth() = %v, want %v", got, tt.want)
				return
			}

			if !(tt.want == nil && got == nil) {
				wanted := tt.want
				gotten := got.Error()

				if !strings.Contains(gotten, *wanted) {
					t.Errorf("setAWSAuth() = %v, want %v", gotten, *wanted)
					return
				}
			}
		})
	}
}

func Test_unSetAWSAuth(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			unSetAWSAuth()
		})
	}
}

func Test_getAWSCredentials1(t *testing.T) {
	type args struct {
		IAMRole string
		region  string
	}
	tests := []struct {
		name    string
		args    args
		want    *sts.AssumeRoleOutput
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getAWSCredentials(tt.args.IAMRole, tt.args.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAWSCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAWSCredentials() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setAWSAuth1(t *testing.T) {
	type args struct {
		iamRole string
		region  string
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
			if err := setAWSAuth(tt.args.iamRole, tt.args.region); (err != nil) != tt.wantErr {
				t.Errorf("setAWSAuth() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_unSetAWSAuth1(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			unSetAWSAuth()
		})
	}
}
