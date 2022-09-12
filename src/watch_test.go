package pike

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func TestWatch(t *testing.T) {
	type args struct {
		arn  string
		wait int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		//{"first", args{"arn:aws:iam::680235478471:policy/basic", 5}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Watch(tt.args.arn, tt.args.wait); (err != nil) != tt.wantErr {
				t.Errorf("Watch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWaitForPolicyChange(t *testing.T) {
	type args struct {
		client  *iam.Client
		arn     string
		Version string
		Wait    int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WaitForPolicyChange(tt.args.client, tt.args.arn, tt.args.Version, tt.args.Wait)
			if (err != nil) != tt.wantErr {
				t.Errorf("WaitForPolicyChange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("WaitForPolicyChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetVersion(t *testing.T) {
	type args struct {
		client    *iam.Client
		PolicyArn string
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
			if got := GetVersion(tt.args.client, tt.args.PolicyArn); got != tt.want {
				t.Errorf("GetVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPolicyVersion(t *testing.T) {
	type args struct {
		client    *iam.Client
		PolicyArn string
		Version   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		//{ "",args{"","arn:aws:iam::680235478471:policy/allows3","v1"},"", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPolicyVersion(tt.args.client, tt.args.PolicyArn, tt.args.Version)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPolicyVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetPolicyVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortActions(t *testing.T) {
	type args struct {
		myPolicy string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"odd", args{
			"{\"Statement\":[{\"Action\":[\"cognito-idp:ListUserPoolClients\",\"cognito-idp:GetSigningCertificate\",\"cognito-idp:DescribeUserPoolClient\"],\"Effect\":\"Allow\",\"Resource\":\"*\",\"Sid\":\"0\"}],\"Version\":\"2012-10-17\"}"},
			"{\"Version\":\"2012-10-17\",\"Statement\":[{\"Sid\":\"0\",\"Effect\":\"Allow\",\"Action\":[\"cognito-idp:DescribeUserPoolClient\",\"cognito-idp:GetSigningCertificate\",\"cognito-idp:ListUserPoolClients\"],\"Resource\":\"*\"}]}", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SortActions(tt.args.myPolicy)
			if (err != nil) != tt.wantErr {
				t.Errorf("SortActions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortActions() = %v, want %v", got, tt.want)
			}
		})
	}
}
