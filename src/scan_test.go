package pike

import (
	"log"
	"path/filepath"
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
		{"aws", args{"../terraform/aws/backup"}, false},
		{"gcp", args{"../terraform/gcp/backup"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Scan(tt.args.dirname, "json", "", false); (err != nil) != tt.wantErr {
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
		want    []string
		wantErr bool
	}{
		{"first", args{"../testdata/scan/examples/simple"}, []string{"../testdata/scan/examples/simple/aws_s3_bucket.pike.tf"}, false},
		{"empty", args{"../testdata/scan"}, nil, false},
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

func TestInit(t *testing.T) {
	type args struct {
		dirName string
	}

	dirName, _ := filepath.Abs("testdata/init/nicconf")

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"remote", args{dirName}, []string{"api_gateway", "dynamodb_table", "lambda_get", "lambda_post"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, modules, err := Init(tt.args.dirName)
			log.Print(modules)
			if (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == "" {
				t.Error("init should return new path to Terraform")
			}
			if !reflect.DeepEqual(modules, tt.want) {
				t.Errorf("Init() got1 = %v, want %v", modules, tt.want)
			}

		})
	}
}

func TestMakePolicy(t *testing.T) {
	type args struct {
		dirName string
		file    string
		init    bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"basic", args{
			"testdata/init/nicconf", "", true},
			"{\"Version\":\"2012-10-17\",\"Statement\":[{\"Sid\":\"VisualEditor0\",\"Effect\":\"Allow\",\"Action\":[\"apigateway:DELETE\",\"apigateway:GET\",\"apigateway:PATCH\",\"apigateway:POST\",\"apigateway:PUT\"],\"Resource\":\"*\"},{\"Sid\":\"VisualEditor1\",\"Effect\":\"Allow\",\"Action\":[\"application-autoscaling:DeleteScalingPolicy\",\"application-autoscaling:DeregisterScalableTarget\",\"application-autoscaling:DescribeScalableTargets\",\"application-autoscaling:DescribeScalingPolicies\",\"application-autoscaling:PutScalingPolicy\",\"application-autoscaling:RegisterScalableTarget\"],\"Resource\":\"*\"},{\"Sid\":\"VisualEditor2\",\"Effect\":\"Allow\",\"Action\":[\"dynamodb:CreateTable\",\"dynamodb:DeleteTable\",\"dynamodb:DescribeContinuousBackups\",\"dynamodb:DescribeTable\",\"dynamodb:DescribeTimeToLive\",\"dynamodb:ListTagsOfResource\",\"dynamodb:TagResource\",\"dynamodb:UntagResource\",\"dynamodb:UpdateTable\",\"dynamodb:UpdateTimeToLive\"],\"Resource\":\"*\"},{\"Sid\":\"VisualEditor3\",\"Effect\":\"Allow\",\"Action\":[\"ec2:DescribeAccountAttributes\"],\"Resource\":\"*\"},{\"Sid\":\"VisualEditor4\",\"Effect\":\"Allow\",\"Action\":[\"iam:AttachRolePolicy\",\"iam:CreatePolicy\",\"iam:CreateRole\",\"iam:CreateServiceLinkedRole\",\"iam:DeletePolicy\",\"iam:DeleteRole\",\"iam:DeleteRolePermissionsBoundary\",\"iam:DetachRolePolicy\",\"iam:GetPolicy\",\"iam:GetPolicyVersion\",\"iam:GetRole\",\"iam:ListAttachedRolePolicies\",\"iam:ListInstanceProfilesForRole\",\"iam:ListPolicies\",\"iam:ListPolicyVersions\",\"iam:ListRolePolicies\",\"iam:PassRole\",\"iam:PutRolePermissionsBoundary\",\"iam:TagPolicy\",\"iam:TagRole\",\"iam:UntagPolicy\",\"iam:UpdateRoleDescription\"],\"Resource\":\"*\"},{\"Sid\":\"VisualEditor5\",\"Effect\":\"Allow\",\"Action\":[\"lambda:AddPermission\",\"lambda:CreateFunction\",\"lambda:DeleteFunction\",\"lambda:GetFunction\",\"lambda:GetFunctionCodeSigningConfig\",\"lambda:GetPolicy\",\"lambda:ListVersionsByFunction\",\"lambda:RemovePermission\",\"lambda:TagResource\",\"lambda:UntagResource\"],\"Resource\":\"*\"},{\"Sid\":\"VisualEditor6\",\"Effect\":\"Allow\",\"Action\":[\"logs:AssociateKmsKey\",\"logs:CreateLogGroup\",\"logs:DeleteLogGroup\",\"logs:DeleteRetentionPolicy\",\"logs:DescribeLogGroups\",\"logs:DisassociateKmsKey\",\"logs:ListTagsLogGroup\",\"logs:PutRetentionPolicy\",\"logs:TagLogGroup\",\"logs:UntagLogGroup\"],\"Resource\":\"*\"},{\"Sid\":\"VisualEditor7\",\"Effect\":\"Allow\",\"Action\":[\"s3:DeleteObject\",\"s3:GetObject\",\"s3:GetObjectTagging\",\"s3:PutObject\"],\"Resource\":\"*\"}]}",
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MakePolicy(tt.args.dirName, tt.args.file, tt.args.init)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakePolicy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(minify(got.AWS.JSONOut), tt.want) {
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
		{"basic", args{"aws_s3_bucket"}, "aws"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHCLType(tt.args.resourceName); got != tt.want {
				t.Errorf("GetHCLType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTFFiles(t *testing.T) {
	type args struct {
		dirName string
	}

	dirName, _ := filepath.Abs("testdata/init/nicconf")

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"basic", args{dirName},
			[]string{
				dirName + "/api_gateway.tf",
				dirName + "/dynamodb.tf",
				dirName + "/lambda_get.tf",
				dirName + "/lambda_post.tf",
				dirName + "/main.tf",
				dirName + "/outputs.tf"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getTFFiles(tt.args.dirName)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTFFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTFFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
