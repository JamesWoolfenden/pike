//go:build auth
// +build auth

package pike

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"golang.org/x/net/context"
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

	cfg, _ := config.LoadDefaultConfig(context.TODO())
	client := iam.NewFromConfig(cfg)

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"wait", args{client, "arn:aws:iam::680235478471:policy/allows3", "1.0", 2}, 2, true},
		{"fail", args{client, "arn:aws:iam::680235478471:policy/rub", "1.0", 5}, 5, true},
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
	cfg, _ := config.LoadDefaultConfig(context.TODO())
	client := iam.NewFromConfig(cfg)
	want := "v1"
	tests := []struct {
		name    string
		args    args
		want    *string
		wantErr bool
	}{
		{"fail", args{client, "arn:aws:iam::680235478471:policy/allowsduff"}, nil, true},
		{"current", args{client, "arn:aws:iam::680235478471:policy/allows3"}, &want, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetVersion(tt.args.client, tt.args.PolicyArn)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
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
	cfg, _ := config.LoadDefaultConfig(context.TODO())
	client := iam.NewFromConfig(cfg)
	wantPass := "{\"Statement\":[{\"Action\":\"s3:*\",\"Effect\":\"Allow\",\"Resource\":\"*\"," +
		"\"Sid\":\"VisualEditor0\"}],\"Version\":\"2012-10-17\"}"
	sagemaker := "{\"Statement\":[{\"Action\":[\"s3:DeleteObject\",\"s3:GetObject\",\"s3:ListBucket\"," +
		"\"s3:PutObject\"],\"Effect\":\"Allow\",\"Resource\":[\"arn:aws:s3:::*\"]}],\"Version\":\"2012-10-17\"}"
	tests := []struct {
		name    string
		args    args
		want    *string
		wantErr bool
	}{
		{"pass", args{client, "arn:aws:iam::680235478471:policy/allows3", "v1"}, &wantPass, false},
		{"not found", args{client, "arn:aws:iam::680235478471:policy/basic", "v1"}, nil, true},
		{"not found2", args{client, "arn:aws:iam::680235478471:policy/service-role/AmazonSageMaker-ExecutionPolicy-20211117T143702", "v1"}, &sagemaker, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPolicyVersion(tt.args.client, tt.args.PolicyArn, tt.args.Version)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPolicyVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPolicyVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortActions(t *testing.T) {
	type args struct {
		myPolicy string
	}
	want := "{\"Statement\":[{\"Action\":[\"cognito-idp:DescribeUserPoolClient\",\"cognito-idp:GetSigningCertificate\",\"cognito-idp:ListUserPoolClients\"],\"Effect\":\"Allow\",\"Resource\":\"*\",\"Sid\":\"0\"}],\"Version\":\"2012-10-17\"}"
	tests := []struct {
		name    string
		args    args
		want    *string
		wantErr bool
	}{
		{
			"odd",
			args{
				"{\"Statement\":[{\"Action\":[\"cognito-idp:ListUserPoolClients\",\"cognito-idp:GetSigningCertificate\",\"cognito-idp:DescribeUserPoolClient\"],\"Effect\":\"Allow\",\"Resource\":\"*\",\"Sid\":\"0\"}],\"Version\":\"2012-10-17\"}",
			},
			&want, false,
		},
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
