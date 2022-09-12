package pike

import (
	_ "embed"
	"reflect"
	"strings"
	"testing"
)

func TestNewAWSPolicy(t *testing.T) {
	type args struct {
		Actions []string
	}
	tests := []struct {
		name string
		args args
		want Policy
	}{
		{"pass", args{[]string{
			"s3:CreateBucket",
			"s3:DeleteBucket",
			"s3:GetAccelerateConfiguration",
			"s3:GetBucketAcl",
			"s3:GetBucketCORS",
			"s3:GetBucketLogging",
			"s3:GetBucketObjectLockConfiguration",
			"s3:GetBucketPolicy",
			"s3:GetBucketRequestPayment",
			"s3:GetBucketTagging",
			"s3:GetBucketVersioning",
			"s3:GetBucketWebsite",
			"s3:GetEncryptionConfiguration",
			"s3:GetLifecycleConfiguration",
			"s3:GetObject",
			"s3:GetObjectAcl",
			"s3:GetReplicationConfiguration",
			"s3:ListBucket"}},
			Policy{
				"2012-10-17",
				[]Statement{{"VisualEditor0", "Allow", []string{"s3:CreateBucket",
					"s3:DeleteBucket",
					"s3:GetAccelerateConfiguration",
					"s3:GetBucketAcl",
					"s3:GetBucketCORS",
					"s3:GetBucketLogging",
					"s3:GetBucketObjectLockConfiguration",
					"s3:GetBucketPolicy",
					"s3:GetBucketRequestPayment",
					"s3:GetBucketTagging",
					"s3:GetBucketVersioning",
					"s3:GetBucketWebsite",
					"s3:GetEncryptionConfiguration",
					"s3:GetLifecycleConfiguration",
					"s3:GetObject",
					"s3:GetObjectAcl",
					"s3:GetReplicationConfiguration",
					"s3:ListBucket"}, "*"}},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAWSPolicy(tt.args.Actions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAWSPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPolicy(t *testing.T) {
	type args struct {
		actions Sorted
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"first", args{Sorted{
			[]string{},
			nil,
			nil,
		}},
			"{\"Version\": \"2012-10-17\",\"Statement\": null\n}",
			false,
		},
		{"aws", args{Sorted{[]string{
			"ec2:DescribeInstances",
			"ec2:DescribeTags",
			"ec2:DescribeInstanceAttribute",
			"ec2:DescribeVolumes",
			"ec2:DescribeInstanceTypes",
			"ec2:RunInstances",
			"ec2:DescribeInstanceCreditSpecifications",
			"ec2:StopInstances",
			"ec2:StartInstances",
			"ec2:ModifyInstanceAttribute",
			"ec2:StopInstances",
			"ec2:StopInstances",
			"ec2:TerminateInstances"}, nil, nil}},
			"{\n    \"Version\": \"2012-10-17\",\n    \"Statement\": [\n        {\n            \"Sid\": \"VisualEditor0\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"ec2:DescribeInstanceAttribute\",\n                \"ec2:DescribeInstanceCreditSpecifications\",\n                \"ec2:DescribeInstanceTypes\",\n                \"ec2:DescribeInstances\",\n                \"ec2:DescribeTags\",\n                \"ec2:DescribeVolumes\",\n                \"ec2:ModifyInstanceAttribute\",\n                \"ec2:RunInstances\",\n                \"ec2:StartInstances\",\n                \"ec2:StopInstances\",\n                \"ec2:TerminateInstances\"\n            ],\n            \"Resource\": \"*\"\n        }\n    ]\n}",
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPolicy(tt.args.actions)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPolicy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			newGot := minify(got.AWS.JSONOut)
			reallyWant := minify(tt.want)
			if newGot != reallyWant {
				t.Errorf("GetPolicy() = %v, want %v", got.AWS.JSONOut, tt.want)
			}
		})
	}
}

func minify(JSONOut string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(JSONOut, "\n", ""), "	", ""), " ", ""), "\r\n", "")
}

func TestAWSPolicy(t *testing.T) {
	type args struct {
		Permissions []string
	}
	tests := []struct {
		name    string
		args    args
		want    AwsOutput
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AWSPolicy(tt.args.Permissions)
			if (err != nil) != tt.wantErr {
				t.Errorf("AWSPolicy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AWSPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unique(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"pass", args{[]string{"a", "b", "c", "a", "c"}}, []string{"a", "b", "c"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unique(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unique() = %v, want %v", got, tt.want)
			}
		})
	}
}
