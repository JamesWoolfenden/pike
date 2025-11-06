package pike

import (
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/rs/zerolog/log"
)

func TestScan(t *testing.T) {
	t.Parallel()

	type args struct {
		dirname string
		output  string
		write   bool
	}

	testPath, _ := filepath.Abs("../terraform/aws/backup")

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"aws", args{testPath, "json", false}, false},
		{"aws-out", args{testPath, "terraform", true}, false},
		{"google", args{testPath, "json", false}, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if err := Scan(tt.args.dirname, tt.args.output, nil, false, tt.args.write, false, "", "", ""); (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetTF(t *testing.T) {
	t.Parallel()

	type args struct {
		dirname string
	}

	found := []string{"../testdata/scan/examples/notlocal/module.tf"}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			"first",
			args{"../testdata/scan/examples/simple"},
			[]string{"../testdata/scan/examples/simple/aws_s3_bucket.pike.tf"},
			false,
		},
		{"empty", args{"../testdata/scan"}, nil, false},
		{"notlocal", args{"../testdata/scan/examples/notlocal"}, found, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

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

func TestGetPermissionBag(t *testing.T) {
	t.Parallel()

	type args struct {
		resources []ResourceV2
		provider  string
	}

	//goland:noinspection GoLinter
	tests := []struct {
		name string
		args args
		want Sorted
	}{
		{
			"basic_aws",
			args{
				resources: []ResourceV2{
					{
						TypeName: "terraform",
						Name:     "aws_s3_bucket",
						Provider: "aws",
						Attributes: []string{
							"tags",
						},
					},
				},
			},
			Sorted{
				AWS: []string{
					"s3:PutBucketTagging",
					"s3:DeleteBucket",
					"s3:CreateBucket",
					"s3:GetLifecycleConfiguration",
					"s3:GetBucketTagging",
					"s3:GetBucketWebsite",
					"s3:GetBucketLogging",
					"s3:ListBucket",
					"s3:GetAccelerateConfiguration",
					"s3:GetBucketVersioning",
					"s3:GetBucketAcl",
					"s3:GetBucketPolicy",
					"s3:GetReplicationConfiguration",
					"s3:GetBucketObjectLockConfiguration",
					"s3:GetObjectAcl",
					"s3:GetObject",
					"s3:GetEncryptionConfiguration",
					"s3:GetBucketRequestPayment",
					"s3:GetBucketCORS",
					"s3:DeleteBucket",
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			permissionBag := GetPermissionBag(tt.args.resources, tt.args.provider)

			if !reflect.DeepEqual(permissionBag, tt.want) {
				t.Errorf("MakePolicy() = %v, want %v", permissionBag, tt.want)
			}
		})
	}
}

func Test_stringInSlice(t *testing.T) {
	t.Parallel()

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
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := StringInSlice(tt.args.a, tt.args.list); got != tt.want {
				t.Errorf("StringInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakePolicy(t *testing.T) {
	t.Parallel()

	type args struct {
		dirName string
		file    *string
		init    bool
	}

	bogus := "testdata/scan/examples/simple/bogus.tf"
	actual := "testdata/scan/examples/simple/aws_s3_bucket.pike.tf"
	dynamic := "testdata/scan/examples/dynamic/dynamic.tf"

	//goland:noinspection GoLinter
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"basic",
			args{
				"testdata/init/nicconf", nil, true,
			},
			`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "VisualEditor0",
      "Effect": "Allow",
      "Action": [
        "apigateway:DELETE",
        "apigateway:GET",
        "apigateway:PATCH",
        "apigateway:POST",
        "apigateway:PUT"
      ],
      "Resource": [
        "*"
      ]
    },
    {
      "Sid": "VisualEditor1",
      "Effect": "Allow",
      "Action": [
        "application-autoscaling:DeleteScalingPolicy",
        "application-autoscaling:DeregisterScalableTarget",
        "application-autoscaling:DescribeScalableTargets",
        "application-autoscaling:DescribeScalingPolicies",
        "application-autoscaling:DescribeScheduledActions",
        "application-autoscaling:PutScalingPolicy",
        "application-autoscaling:PutScheduledAction",
        "application-autoscaling:RegisterScalableTarget"
      ],
      "Resource": [
        "*"
      ]
    },
    {
      "Sid": "VisualEditor2",
      "Effect": "Allow",
      "Action": [
        "cloudwatch:DeleteAlarms",
        "cloudwatch:DescribeAlarms",
 		"cloudwatch:GetMetricData",
        "cloudwatch:PutMetricAlarm"
      ],
      "Resource": [
        "*"
      ]
    },
    {
      "Sid": "VisualEditor3",
      "Effect": "Allow",
      "Action": [
        "dynamodb:BatchWriteItem",
        "dynamodb:CreateTable",
        "dynamodb:CreateTableReplica",
        "dynamodb:DeleteItem",
        "dynamodb:DeleteTable",
        "dynamodb:DeleteTableReplica",
        "dynamodb:DescribeContinuousBackups",
        "dynamodb:DescribeKinesisStreamingDestination",
        "dynamodb:DescribeTable",
        "dynamodb:DescribeTimeToLive",
        "dynamodb:DisableKinesisStreamingDestination",
        "dynamodb:EnableKinesisStreamingDestination",
        "dynamodb:GetItem",
        "dynamodb:ListTagsOfResource",
        "dynamodb:PutItem",
        "dynamodb:Query",
        "dynamodb:Scan",
        "dynamodb:TagResource",
        "dynamodb:UntagResource",
        "dynamodb:UpdateContinuousBackups",
        "dynamodb:UpdateItem",
        "dynamodb:UpdateTable",
        "dynamodb:UpdateTimeToLive"
      ],
      "Resource": [
        "*"
      ]
    },
    {
      "Sid": "VisualEditor4",
      "Effect": "Allow",
      "Action": [
        "ec2:DescribeAccountAttributes",
        "ec2:DescribeNetworkInterfaces"
      ],
      "Resource": [
        "*"
      ]
    },
    {
      "Sid": "VisualEditor5",
      "Effect": "Allow",
      "Action": [
        "iam:AttachRolePolicy",
        "iam:CreatePolicy",
        "iam:CreateRole",
        "iam:CreateServiceLinkedRole",
        "iam:DeletePolicy",
        "iam:DeleteRole",
        "iam:DeleteRolePermissionsBoundary",
        "iam:DetachRolePolicy",
        "iam:GetPolicy",
        "iam:GetPolicyVersion",
        "iam:GetRole",
        "iam:ListAttachedRolePolicies",
        "iam:ListInstanceProfilesForRole",
        "iam:ListPolicies",
        "iam:ListPolicyVersions",
        "iam:ListRolePolicies",
        "iam:PassRole",
        "iam:PutRolePermissionsBoundary",
        "iam:TagPolicy",
        "iam:TagRole",
        "iam:UntagPolicy",
        "iam:UntagRole",
        "iam:UpdateRoleDescription"
      ],
      "Resource": [
        "*"
      ]
    },
    {
      "Sid": "VisualEditor6",
      "Effect": "Allow",
      "Action": [
        "kinesis:DescribeStream",
        "kinesis:PutRecords"
      ],
      "Resource": [
        "*"
      ]
    },
    {
      "Sid": "VisualEditor7",
      "Effect": "Allow",
      "Action": [
        "kms:CreateGrant",
        "kms:Decrypt",
        "kms:DescribeKey",
        "kms:Encrypt",
        "kms:GenerateDataKey",
        "kms:ListAliases",
        "kms:RevokeGrant"
      ],
      "Resource": [
        "*"
      ]
    },
    {
      "Sid": "VisualEditor8",
      "Effect": "Allow",
      "Action": [
                        "lambda:AddPermission",
                "lambda:CreateEventSourceMapping",
                "lambda:CreateFunction",
                "lambda:CreateFunctionUrlConfig",
                "lambda:DeleteEventSourceMapping",
                "lambda:DeleteFunction",
                "lambda:DeleteFunctionEventInvokeConfig",
                "lambda:DeleteFunctionUrlConfig",
                "lambda:DeleteLayerVersion",
                "lambda:DeleteProvisionedConcurrencyConfig",
                "lambda:GetCodeSigningConfig",
                "lambda:GetEventSourceMapping",
                "lambda:GetFunction",
                "lambda:GetFunctionCodeSigningConfig",
                "lambda:GetFunctionEventInvokeConfig",
                "lambda:GetFunctionUrlConfig",
                "lambda:GetLayerVersion",
                "lambda:GetPolicy",
                "lambda:GetProvisionedConcurrencyConfig",
                "lambda:ListTags",
                "lambda:ListVersionsByFunction",
                "lambda:PublishLayerVersion",
                "lambda:PutFunctionEventInvokeConfig",
                "lambda:PutProvisionedConcurrencyConfig",
                "lambda:RemovePermission",
                "lambda:TagResource",
                "lambda:UntagResource",
                "lambda:UpdateEventSourceMapping",
                "lambda:UpdateFunctionEventInvokeConfig",
                "lambda:UpdateFunctionUrlConfig"
      ],
      "Resource": [
        "*"
      ]
    },
    {
      "Sid": "VisualEditor9",
      "Effect": "Allow",
      "Action": [
        "logs:AssociateKmsKey",
        "logs:CreateLogGroup",
        "logs:DeleteLogGroup",
        "logs:DeleteRetentionPolicy",
        "logs:DescribeLogGroups",
        "logs:DisassociateKmsKey",
 		"logs:ListTagsForResource",
        "logs:ListTagsLogGroup",
        "logs:PutRetentionPolicy",
        "logs:TagLogGroup",
        "logs:UntagLogGroup"
      ],
      "Resource": [
        "*"
      ]
    },
    {
      "Sid": "VisualEditor10",
      "Effect": "Allow",
      "Action": [
        "s3:DeleteObject",
        "s3:GetObject",
        "s3:GetObjectTagging",
        "s3:GetObjectVersion",
        "s3:PutObject"
      ],
      "Resource": [
        "*"
      ]
    }
  ]
}`,
			false,
		},
		{
			"not a dir",
			args{"bogus", nil, true},
			"",
			true,
		},
		{
			"a file",
			args{"", &actual, false},
			`
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "VisualEditor0",
      "Effect": "Allow",
      "Action": [
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
        "s3:ListBucket"
      ],
      "Resource": [
        "*"
      ]
    }
  ]
}`,
			false,
		},
		{
			"not a file",
			args{"", &bogus, false},
			"",
			true,
		},
		{
			"remote module",
			args{"../testdata/scan/examples/notlocal", &bogus, false},
			"",
			true,
		},
		{
			"dynamic",
			args{"", &dynamic, false},
			`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "VisualEditor0",
      "Effect": "Allow",
      "Action": [
        "autoscaling:CreateAutoScalingGroup",
        "autoscaling:CreateOrUpdateTags",
        "autoscaling:DeleteAutoScalingGroup",
        "autoscaling:DeleteTags",
        "autoscaling:Describe*",
        "autoscaling:DescribeAutoScalingGroups",
        "autoscaling:DescribeScalingActivities",
        "autoscaling:UpdateAutoScalingGroup"
      ],
      "Resource": [
        "*"
      ]
    },
    {
      "Sid": "VisualEditor1",
      "Effect": "Allow",
      "Action": [
        "ec2:Describe*",
        "ec2:Get*",
        "ec2:RunInstances"
      ],
      "Resource": [
        "*"
      ]
    },
    {
      "Sid": "VisualEditor2",
      "Effect": "Allow",
      "Action": [
        "managed-fleets:DeleteAutoScalingGroup",
        "managed-fleets:DeregisterAutoScalingGroup",
        "managed-fleets:Get*",
        "managed-fleets:RegisterAutoScalingGroup",
        "managed-fleets:UpdateAutoScalingGroup"
      ],
      "Resource": [
        "*"
      ]
    },
    {
      "Sid": "VisualEditor3",
      "Effect": "Allow",
      "Action": [
        "ssm:Get*"
      ],
      "Resource": [
        "*"
      ]
    }
  ]
}`,
			false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := MakePolicy(tt.args.dirName, tt.args.file, tt.args.init, false, "", "")

			if (err != nil) != tt.wantErr {
				t.Errorf("MakePolicy() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			JSONOut := Minify(got.AWS.JSONOut)
			want := Minify(tt.want)

			if !reflect.DeepEqual(JSONOut, want) {
				t.Errorf("MakePolicy() = %v, want %v", JSONOut, want)
			}
		})
	}
}

func TestGetHCLType(t *testing.T) {
	t.Parallel()

	type args struct {
		resourceName string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"basic", args{"aws_s3_bucket"}, "aws"},
		{"gcp", args{"google_storage_bucket"}, "google"},
		{"azurerm", args{"azurerm_storage_account"}, "azurerm"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := GetHCLType(tt.args.resourceName); got != tt.want {
				t.Errorf("GetHCLType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTFFiles(t *testing.T) {
	t.Parallel()

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
		{
			"basic",
			args{dirName},
			[]string{
				dirName + "/api_gateway.tf",
				dirName + "/dynamodb.tf",
				dirName + "/lambda_get.tf",
				dirName + "/lambda_post.tf",
				dirName + "/main.tf",
				dirName + "/outputs.tf",
			},
			false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := GetTFFiles(tt.args.dirName)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetTFFiles() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTFFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteOutput(t *testing.T) {
	t.Parallel()

	type args struct {
		OutPolicy OutputPolicy
		output    string
		location  string
	}

	out := OutputPolicy{AWS: AwsOutput{
		JSONOut:   "{\n    \"Version\": \"2012-10-17\",\n    \"Statement\": [\n        {\n            \"Sid\": \"VisualEditor0\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"SNS:CreateTopic\",\n                \"SNS:DeleteTopic\",\n                \"SNS:GetTopicAttributes\",\n                \"SNS:ListTagsForResource\",\n                \"SNS:ListTopics\",\n                \"SNS:SetTopicAttributes\",\n                \"SNS:TagResource\",\n                \"SNS:UnTagResource\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor1\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"acm:AddTagsToCertificate\",\n                \"acm:DeleteCertificate\",\n                \"acm:DescribeCertificate\",\n                \"acm:ListTagsForCertificate\",\n                \"acm:RemoveTagsFromCertificate\",\n                \"acm:RequestCertificate\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor2\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"acm-pa:DescribeCertificateAuthority\",\n                \"acm-pa:ListTags\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor3\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"acm-pca:CreateCertificateAuthority\",\n                \"acm-pca:DeleteCertificateAuthority\",\n                \"acm-pca:GetCertificateAuthorityCertificate\",\n                \"acm-pca:UpdateCertificateAuthority\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor4\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"apigateway:DELETE\",\n                \"apigateway:GET\",\n                \"apigateway:PATCH\",\n                \"apigateway:POST\",\n                \"apigateway:PUT\",\n                \"apigateway:UpdateRestApiPolicy\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor5\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"application-autoscaling:DeleteScalingPolicy\",\n                \"application-autoscaling:DeleteScheduledAction\",\n                \"application-autoscaling:DeregisterScalableTarget\",\n                \"application-autoscaling:DescribeScalableTargets\",\n                \"application-autoscaling:DescribeScalingPolicies\",\n                \"application-autoscaling:DescribeScheduledActions\",\n                \"application-autoscaling:PutScalingPolicy\",\n                \"application-autoscaling:PutScheduledAction\",\n                \"application-autoscaling:RegisterScalableTarget\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor6\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"application-autoscaling:DeleteScalingPolicy\",\n                \"application-autoscaling:DeleteScheduledAction\",\n                \"application-autoscaling:DeregisterScalableTarget\",\n                \"application-autoscaling:DescribeScalableTargets\",\n                \"application-autoscaling:DescribeScalingPolicies\",\n                \"application-autoscaling:DescribeScheduledActions\",\n                \"application-autoscaling:PutScalingPolicy\",\n                \"application-autoscaling:PutScheduledAction\",\n                \"application-autoscaling:RegisterScalableTarget\",\n                \"autoscaling:AttachLoadBalancers\",\n                \"autoscaling:CreateAutoScalingGroup\",\n                \"autoscaling:CreateLaunchConfiguration\",\n                \"autoscaling:DeleteAutoScalingGroup\",\n                \"autoscaling:DeleteLaunchConfiguration\",\n                \"autoscaling:DescribeAutoScalingGroups\",\n                \"autoscaling:DescribeLaunchConfigurations\",\n                \"autoscaling:DescribeScalingActivities\",\n                \"autoscaling:DetachLoadBalancers\",\n                \"autoscaling:UpdateAutoScalingGroup\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor7\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"budgets:ModifyBudget\",\n                \"budgets:ViewBudget\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor8\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"cloudtrail:AddTags\",\n                \"cloudtrail:CreateTrail\",\n                \"cloudtrail:DeleteTrail\",\n                \"cloudtrail:DescribeTrails\",\n                \"cloudtrail:GetEventSelectors\",\n                \"cloudtrail:GetTrailStatus\",\n                \"cloudtrail:ListTags\",\n                \"cloudtrail:PutEventSelectors\",\n                \"cloudtrail:RemoveTags\",\n                \"cloudtrail:StartLogging\",\n                \"cloudtrail:UpdateTrail\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor9\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"cloudwatch:DeleteAlarms\",\n                \"cloudwatch:DescribeAlarms\",\n                \"cloudwatch:ListTagsForResource\",\n                \"cloudwatch:PutMetricAlarm\",\n                \"cloudwatch:TagResource\",\n                \"cloudwatch:UnTagResource\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor10\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"codeartifact:CreateDomain\",\n                \"codeartifact:CreateRepository\",\n                \"codeartifact:DeleteDomain\",\n                \"codeartifact:DeleteDomainPermissionsPolicy\",\n                \"codeartifact:DeleteRepository\",\n                \"codeartifact:DeleteRepositoryPermissionsPolicy\",\n                \"codeartifact:DescribeDomain\",\n                \"codeartifact:DescribeRepository\",\n                \"codeartifact:GetDomainPermissionsPolicy\",\n                \"codeartifact:GetRepositoryPermissionsPolicy\",\n                \"codeartifact:ListTagsForResource\",\n                \"codeartifact:PutDomainPermissionsPolicy\",\n                \"codeartifact:PutRepositoryPermissionsPolicy\",\n                \"codeartifact:TagResource\",\n                \"codeartifact:UntagResource\",\n                \"codeartifact:UpdateRepository\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor11\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"codebuild:BatchGetProjects\",\n                \"codebuild:CreateProject\",\n                \"codebuild:DeleteProject\",\n                \"codebuild:UpdateProject\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor12\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"codecommit:CreateRepository\",\n                \"codecommit:DeleteRepository\",\n                \"codecommit:GetRepository\",\n                \"codecommit:ListBranches\",\n                \"codecommit:ListTagsForResource\",\n                \"codecommit:TagResource\",\n                \"codecommit:UntagResource\",\n                \"codecommit:UpdateRepositoryDescription\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor13\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"codepipeline:CreatePipeline\",\n                \"codepipeline:DeletePipeline\",\n                \"codepipeline:GetPipeline\",\n                \"codepipeline:ListTagsForResource\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor14\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"cognito-idp:AddCustomAttributes\",\n                \"cognito-idp:AdminAddUserToGroup\",\n                \"cognito-idp:AdminCreateUser\",\n                \"cognito-idp:AdminDeleteUser\",\n                \"cognito-idp:AdminGetUser\",\n                \"cognito-idp:AdminListGroupsForUser\",\n                \"cognito-idp:AdminRemoveUserFromGroup\",\n                \"cognito-idp:CreateGroup\",\n                \"cognito-idp:CreateIdentityProvider\",\n                \"cognito-idp:CreateResourceServer\",\n                \"cognito-idp:CreateUserPool\",\n                \"cognito-idp:CreateUserPoolClient\",\n                \"cognito-idp:CreateUserPoolDomain\",\n                \"cognito-idp:DeleteGroup\",\n                \"cognito-idp:DeleteIdentityProvider\",\n                \"cognito-idp:DeleteResourceServer\",\n                \"cognito-idp:DeleteUserPool\",\n                \"cognito-idp:DeleteUserPoolClient\",\n                \"cognito-idp:DeleteUserPoolDomain\",\n                \"cognito-idp:DescribeIdentityProvider\",\n                \"cognito-idp:DescribeResourceServer\",\n                \"cognito-idp:DescribeUserPool\",\n                \"cognito-idp:DescribeUserPoolClient\",\n                \"cognito-idp:DescribeUserPoolDomain\",\n                \"cognito-idp:GetGroup\",\n                \"cognito-idp:GetSigningCertificate\",\n                \"cognito-idp:GetUICustomization\",\n                \"cognito-idp:GetUserPoolMfaConfig\",\n                \"cognito-idp:ListUserPoolClients\",\n                \"cognito-idp:ListUserPools\",\n                \"cognito-idp:SetUICustomization\",\n                \"cognito-idp:SetUserPoolMfaConfig\",\n                \"cognito-idp:TagResource\",\n                \"cognito-idp:UntagResource\",\n                \"cognito-idp:UpdateGroup\",\n                \"cognito-idp:UpdateIdentityProvider\",\n                \"cognito-idp:UpdateResourceServer\",\n                \"cognito-idp:UpdateUserPool\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor15\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"dax:CreateParameterGroup\",\n                \"dax:CreateSubnetGroup\",\n                \"dax:DeleteParameterGroup\",\n                \"dax:DeleteSubnetGroup\",\n                \"dax:DescribeParameterGroups\",\n                \"dax:DescribeParameters\",\n                \"dax:DescribeSubnetGroups\",\n                \"dax:UpdateParameterGroup\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor16\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"ds:AddTagsToResource\",\n                \"ds:CreateDirectory\",\n                \"ds:CreateLogSubscription\",\n                \"ds:DeleteDirectory\",\n                \"ds:DeleteLogSubscription\",\n                \"ds:ListTagsForResource\",\n                \"ds:RemoveTagsFromResource\",\n                \"rds:AddRoleToDBCluster\",\n                \"rds:AddTagsToResource\",\n                \"rds:CreateDBCluster\",\n                \"rds:CreateDBClusterEndpoint\",\n                \"rds:CreateDBClusterParameterGroup\",\n                \"rds:CreateDBClusterSnapshot\",\n                \"rds:CreateDBInstance\",\n                \"rds:CreateDBParameterGroup\",\n                \"rds:CreateDBSubnetGroup\",\n                \"rds:CreateGlobalCluster\",\n                \"rds:CreateOptionGroup\",\n                \"rds:DeleteDBCluster\",\n                \"rds:DeleteDBClusterEndpoint\",\n                \"rds:DeleteDBClusterParameterGroup\",\n                \"rds:DeleteDBClusterSnapshot\",\n                \"rds:DeleteDBParameterGroup\",\n                \"rds:DeleteDBSubnetGroup\",\n                \"rds:DeleteGlobalCluster\",\n                \"rds:DeleteOptionGroup\",\n                \"rds:DescribeCertificates\",\n                \"rds:DescribeDBClusterParameterGroups\",\n                \"rds:DescribeDBClusterParameters\",\n                \"rds:DescribeDBClusterSnapshots\",\n                \"rds:DescribeDBClusters\",\n                \"rds:DescribeDBEngineVersions\",\n                \"rds:DescribeDBInstances\",\n                \"rds:DescribeDBParameterGroups\",\n                \"rds:DescribeDBParameters\",\n                \"rds:DescribeDBSnapshots\",\n                \"rds:DescribeDBSubnetGroups\",\n                \"rds:DescribeEventCategories\",\n                \"rds:DescribeGlobalClusters\",\n                \"rds:DescribeOptionGroups\",\n                \"rds:DescribeOrderableDBInstanceOptions\",\n                \"rds:ListTagsForResource\",\n                \"rds:ModifyDBCluster\",\n                \"rds:ModifyDBClusterEndpoint\",\n                \"rds:ModifyDBClusterParameterGroup\",\n                \"rds:ModifyDBInstance\",\n                \"rds:ModifyDBParameterGroup\",\n                \"rds:ModifyGlobalCluster\",\n                \"rds:ModifyOptionGroup\",\n                \"rds:RemoveRoleFromDBCluster\",\n                \"rds:RemoveTagsFromResource\",\n                \"rds:StartActivityStream\",\n                \"rds:StopActivityStream\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor17\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"dynamodb:BatchWriteItem\",\n                \"dynamodb:CreateTable\",\n                \"dynamodb:CreateTableReplica\",\n                \"dynamodb:DeleteItem\",\n                \"dynamodb:DeleteTable\",\n                \"dynamodb:DeleteTableReplica\",\n                \"dynamodb:DescribeContinuousBackups\",\n                \"dynamodb:DescribeTable\",\n                \"dynamodb:DescribeTimeToLive\",\n                \"dynamodb:GetItem\",\n                \"dynamodb:ListTagsOfResource\",\n                \"dynamodb:PutItem\",\n                \"dynamodb:Query\",\n                \"dynamodb:Scan\",\n                \"dynamodb:TagResource\",\n                \"dynamodb:UntagResource\",\n                \"dynamodb:UpdateContinuousBackups\",\n                \"dynamodb:UpdateItem\",\n                \"dynamodb:UpdateTable\",\n                \"dynamodb:UpdateTimeToLive\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor18\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"ec2:AllocateAddress\",\n                \"ec2:AssociateAddress\",\n                \"ec2:AssociateRouteTable\",\n                \"ec2:AttachInternetGateway\",\n                \"ec2:AttachVolume\",\n                \"ec2:AttachVpnGateway\",\n                \"ec2:AuthorizeSecurityGroupEgress\",\n                \"ec2:AuthorizeSecurityGroupIngress\",\n                \"ec2:CancelCapacityReservation\",\n                \"ec2:CancelSpotInstanceRequests\",\n                \"ec2:CreateCapacityReservation\",\n                \"ec2:CreateDefaultVpc\",\n                \"ec2:CreateFlowLogs\",\n                \"ec2:CreateInternetGateway\",\n                \"ec2:CreateKeyPair\",\n                \"ec2:CreateLaunchTemplate\",\n                \"ec2:CreateLaunchTemplateVersion\",\n                \"ec2:CreateNatGateway\",\n                \"ec2:CreateNetworkAcl\",\n                \"ec2:CreateNetworkAclEntry\",\n                \"ec2:CreateNetworkInterface\",\n                \"ec2:CreateNetworkInterfacePermission\",\n                \"ec2:CreatePlacementGroup\",\n                \"ec2:CreateRoute\",\n                \"ec2:CreateRouteTable\",\n                \"ec2:CreateSecurityGroup\",\n                \"ec2:CreateSubnet\",\n                \"ec2:CreateTags\",\n                \"ec2:CreateVPC\",\n                \"ec2:CreateVolume\",\n                \"ec2:CreateVpcEndpoint\",\n                \"ec2:CreateVpnGateway\",\n                \"ec2:DeleteFlowLogs\",\n                \"ec2:DeleteInternetGateway\",\n                \"ec2:DeleteKeyPair\",\n                \"ec2:DeleteLaunchTemplate\",\n                \"ec2:DeleteNatGateway\",\n                \"ec2:DeleteNetworkAcl\",\n                \"ec2:DeleteNetworkAclEntry\",\n                \"ec2:DeleteNetworkInterface\",\n                \"ec2:DeleteNetworkInterfacePermission\",\n                \"ec2:DeletePlacementGroup\",\n                \"ec2:DeleteRoute\",\n                \"ec2:DeleteRouteTable\",\n                \"ec2:DeleteSecurityGroup\",\n                \"ec2:DeleteSubnet\",\n                \"ec2:DeleteTags\",\n                \"ec2:DeleteVPC\",\n                \"ec2:DeleteVolume\",\n                \"ec2:DeleteVpcEndpoints\",\n                \"ec2:DeleteVpnGateway\",\n                \"ec2:DescribeAccountAttributes\",\n                \"ec2:DescribeAddresses\",\n                \"ec2:DescribeAvailabilityZones\",\n                \"ec2:DescribeCapacityReservations\",\n                \"ec2:DescribeDhcpOptions\",\n                \"ec2:DescribeFlowLogs\",\n                \"ec2:DescribeImages\",\n                \"ec2:DescribeInstanceAttribute\",\n                \"ec2:DescribeInstanceCreditSpecifications\",\n                \"ec2:DescribeInstanceTypes\",\n                \"ec2:DescribeInstances\",\n                \"ec2:DescribeInternetGateways\",\n                \"ec2:DescribeKeyPairs\",\n                \"ec2:DescribeLaunchTemplateVersions\",\n                \"ec2:DescribeLaunchTemplates\",\n                \"ec2:DescribeNatGateways\",\n                \"ec2:DescribeNetworkAcls\",\n                \"ec2:DescribeNetworkInterfaces\",\n                \"ec2:DescribePlacementGroups\",\n                \"ec2:DescribePrefixLists\",\n                \"ec2:DescribeRouteTables\",\n                \"ec2:DescribeSecurityGroups\",\n                \"ec2:DescribeSpotInstanceRequests\",\n                \"ec2:DescribeSubnets\",\n                \"ec2:DescribeTags\",\n                \"ec2:DescribeVolumes\",\n                \"ec2:DescribeVpcAttribute\",\n                \"ec2:DescribeVpcEndpointServices\",\n                \"ec2:DescribeVpcEndpoints\",\n                \"ec2:DescribeVpcs\",\n                \"ec2:DescribeVpnGateways\",\n                \"ec2:DetachInternetGateway\",\n                \"ec2:DetachNetworkInterface\",\n                \"ec2:DetachVolume\",\n                \"ec2:DetachVpnGateway\",\n                \"ec2:DisassociateAddress\",\n                \"ec2:DisassociateRouteTable\",\n                \"ec2:GetEbsDefaultKmsKeyId\",\n                \"ec2:ImportKeyPair\",\n                \"ec2:ModifyCapacityReservation\",\n                \"ec2:ModifyInstanceAttribute\",\n                \"ec2:ModifyVolume\",\n                \"ec2:ModifyVpcEndpoint\",\n                \"ec2:MonitorInstances\",\n                \"ec2:ReleaseAddress\",\n                \"ec2:RequestSpotInstances\",\n                \"ec2:RevokeSecurityGroupEgress\",\n                \"ec2:RevokeSecurityGroupIngress\",\n                \"ec2:RunInstances\",\n                \"ec2:StartInstances\",\n                \"ec2:StopInstances\",\n                \"ec2:TerminateInstances\",\n                \"ec2:UnmonitorInstances\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor19\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"ecr:CreatePullThroughCacheRule\",\n                \"ecr:CreateRepository\",\n                \"ecr:DeleteLifecyclePolicy\",\n                \"ecr:DeletePullThroughCacheRule\",\n                \"ecr:DeleteRepository\",\n                \"ecr:DescribePullThroughCacheRules\",\n                \"ecr:DescribeRepositories\",\n                \"ecr:GetAuthorizationToken\",\n                \"ecr:GetLifecyclePolicy\",\n                \"ecr:ListTagsForResource\",\n                \"ecr:PutImageScanningConfiguration\",\n                \"ecr:PutLifecyclePolicy\",\n                \"ecr:TagResource\",\n                \"ecr:UntagResource\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor20\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"ecs:CreateCluster\",\n                \"ecs:CreateService\",\n                \"ecs:DeleteCluster\",\n                \"ecs:DeleteService\",\n                \"ecs:DeregisterTaskDefinition\",\n                \"ecs:DescribeClusters\",\n                \"ecs:DescribeServices\",\n                \"ecs:DescribeTaskDefinition\",\n                \"ecs:UpdateCluster\",\n                \"ecs:RegisterTaskDefinition\",\n                \"ecs:TagResource\",\n                \"ecs:UntagResource\",\n                \"ecs:UpdateService\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor21\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"eks:DescribeCluster\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor22\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"elasticache:AddTagsToResource\",\n                \"elasticache:CreateCacheParameterGroup\",\n                \"elasticache:CreateCacheSubnetGroup\",\n                \"elasticache:DeleteCacheParameterGroup\",\n                \"elasticache:DeleteCacheSubnetGroup\",\n                \"elasticache:DescribeCacheParameterGroups\",\n                \"elasticache:DescribeCacheParameters\",\n                \"elasticache:DescribeCacheSubnetGroups\",\n                \"elasticache:ListTagsForResource\",\n                \"elasticache:ModifyCacheParameterGroup\",\n                \"elasticache:ModifyCacheSubnetGroup\",\n                \"elasticache:RemoveTagsFromResource\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor23\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"elasticbeanstalk:AddTags\",\n                \"elasticbeanstalk:CreateApplication\",\n                \"elasticbeanstalk:DeleteApplication\",\n                \"elasticbeanstalk:DescribeApplications\",\n                \"elasticbeanstalk:ListAvailableSolutionStacks\",\n                \"elasticbeanstalk:ListTagsForResource\",\n                \"elasticbeanstalk:RemoveTags\",\n                \"elasticbeanstalk:UpdateApplicationResourceLifecycle\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor24\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"elasticfilesystem:CreateAccessPoint\",\n                \"elasticfilesystem:CreateFileSystem\",\n                \"elasticfilesystem:CreateReplicationConfiguration\",\n                \"elasticfilesystem:DeleteAccessPoint\",\n                \"elasticfilesystem:DeleteFileSystem\",\n                \"elasticfilesystem:DeleteFileSystemPolicy\",\n                \"elasticfilesystem:DeleteReplicationConfiguration\",\n                \"elasticfilesystem:DescribeAccessPoints\",\n                \"elasticfilesystem:DescribeBackupPolicy\",\n                \"elasticfilesystem:DescribeFileSystemPolicy\",\n                \"elasticfilesystem:DescribeFileSystems\",\n                \"elasticfilesystem:DescribeLifecycleConfiguration\",\n                \"elasticfilesystem:DescribeMountTargetSecurityGroups\",\n                \"elasticfilesystem:DescribeMountTargets\",\n                \"elasticfilesystem:DescribeReplicationConfigurations\",\n                \"elasticfilesystem:PutBackupPolicy\",\n                \"elasticfilesystem:PutFileSystemPolicy\",\n                \"elasticfilesystem:TagResource\",\n                \"elasticfilesystem:UntagResource\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor25\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"elasticloadbalancing:AddTags\",\n                \"elasticloadbalancing:AttachLoadBalancerToSubnets\",\n                \"elasticloadbalancing:CreateListener\",\n                \"elasticloadbalancing:CreateLoadBalancer\",\n                \"elasticloadbalancing:CreateLoadBalancerListeners\",\n                \"elasticloadbalancing:CreateTargetGroup\",\n                \"elasticloadbalancing:DeleteListener\",\n                \"elasticloadbalancing:DeleteLoadBalancer\",\n                \"elasticloadbalancing:DeleteTargetGroup\",\n                \"elasticloadbalancing:DeregisterTargets\",\n                \"elasticloadbalancing:DescribeListeners\",\n                \"elasticloadbalancing:DescribeLoadBalancerAttributes\",\n                \"elasticloadbalancing:DescribeLoadBalancers\",\n                \"elasticloadbalancing:DescribeTags\",\n                \"elasticloadbalancing:DescribeTargetGroupAttributes\",\n                \"elasticloadbalancing:DescribeTargetGroups\",\n                \"elasticloadbalancing:DescribeTargetHealth\",\n                \"elasticloadbalancing:ModifyListener\",\n                \"elasticloadbalancing:ModifyLoadBalancerAttributes\",\n                \"elasticloadbalancing:ModifyTargetGroupAttributes\",\n                \"elasticloadbalancing:RegisterTargets\",\n                \"elasticloadbalancing:RemoveTags\",\n                \"elasticloadbalancing:SetSecurityGroups\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor26\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"events:DeleteRule\",\n                \"events:DescribeRule\",\n                \"events:ListTagsForResource\",\n                \"events:ListTargetsByRule\",\n                \"events:PutRule\",\n                \"events:PutTargets\",\n                \"events:RemoveTargets\",\n                \"events:TagResource\",\n                \"events:UnTagResource\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor27\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"firehose:CreateDeliveryStream\",\n                \"firehose:DeleteDeliveryStream\",\n                \"firehose:DescribeDeliveryStream\",\n                \"firehose:ListTagsForDeliveryStream\",\n                \"firehose:TagDeliveryStream\",\n                \"firehose:UntagDeliveryStream\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor28\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"glue:CreateClassifier\",\n                \"glue:CreateConnection\",\n                \"glue:CreateCrawler\",\n                \"glue:CreateDatabase\",\n                \"glue:CreateJob\",\n                \"glue:CreateMLTransform\",\n                \"glue:CreateRegistry\",\n                \"glue:CreateSchema\",\n                \"glue:CreateScript\",\n                \"glue:CreateSecurityConfiguration\",\n                \"glue:CreateTable\",\n                \"glue:CreateTrigger\",\n                \"glue:CreateUserDefinedFunction\",\n                \"glue:CreateWorkflow\",\n                \"glue:DeleteClassifier\",\n                \"glue:DeleteConnection\",\n                \"glue:DeleteCrawler\",\n                \"glue:DeleteDatabase\",\n                \"glue:DeleteJob\",\n                \"glue:DeleteMLTransform\",\n                \"glue:DeleteRegistry\",\n                \"glue:DeleteResourcePolicy\",\n                \"glue:DeleteSchema\",\n                \"glue:DeleteSecurityConfiguration\",\n                \"glue:DeleteTable\",\n                \"glue:DeleteTrigger\",\n                \"glue:DeleteUserDefinedFunction\",\n                \"glue:DeleteWorkflow\",\n                \"glue:GetClassifier\",\n                \"glue:GetConnection\",\n                \"glue:GetCrawler\",\n                \"glue:GetDataCatalogEncryptionSettings\",\n                \"glue:GetDatabase\",\n                \"glue:GetJob\",\n                \"glue:GetMLTransform\",\n                \"glue:GetRegistry\",\n                \"glue:GetResourcePolicy\",\n                \"glue:GetSchema\",\n                \"glue:GetSchemaVersion\",\n                \"glue:GetSecurityConfiguration\",\n                \"glue:GetTable\",\n                \"glue:GetTags\",\n                \"glue:GetTrigger\",\n                \"glue:GetUserDefinedFunction\",\n                \"glue:GetWorkflow\",\n                \"glue:PutDataCatalogEncryptionSettings\",\n                \"glue:PutResourcePolicy\",\n                \"glue:TagResource\",\n                \"glue:UntagResource\",\n                \"glue:UpdateClassifier\",\n                \"glue:UpdateConnection\",\n                \"glue:UpdateCrawler\",\n                \"glue:UpdateDatabase\",\n                \"glue:UpdateJob\",\n                \"glue:UpdateMLTransform\",\n                \"glue:UpdateRegistry\",\n                \"glue:UpdateSchema\",\n                \"glue:UpdateTable\",\n                \"glue:UpdateTrigger\",\n                \"glue:UpdateUserDefinedFunction\",\n                \"glue:UpdateWorkflow\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor29\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"iam:AddRoleToInstanceProfile\",\n                \"iam:AddUserToGroup\",\n                \"iam:AttachGroupPolicy\",\n                \"iam:AttachRolePolicy\",\n                \"iam:AttachUserPolicy\",\n                \"iam:CreateAccessKey\",\n                \"iam:CreateGroup\",\n                \"iam:CreateInstanceProfile\",\n                \"iam:CreateLoginProfile\",\n                \"iam:CreatePolicy\",\n                \"iam:CreateRole\",\n                \"iam:CreateServiceLinkedRole\",\n                \"iam:CreateUser\",\n                \"iam:DeleteAccessKey\",\n                \"iam:DeleteGroup\",\n                \"iam:DeleteGroupPolicy\",\n                \"iam:DeleteInstanceProfile\",\n                \"iam:DeleteLoginProfile\",\n                \"iam:DeletePolicy\",\n                \"iam:DeleteRole\",\n                \"iam:DeleteRolePolicy\",\n                \"iam:DeleteServiceLinkedRole\",\n                \"iam:DeleteUser\",\n                \"iam:DeleteUserPolicy\",\n                \"iam:DetachGroupPolicy\",\n                \"iam:DetachRolePolicy\",\n                \"iam:DetachUserPolicy\",\n                \"iam:GetGroup\",\n                \"iam:GetGroupPolicy\",\n                \"iam:GetInstanceProfile\",\n                \"iam:GetLoginProfile\",\n                \"iam:GetPolicy\",\n                \"iam:GetPolicyVersion\",\n                \"iam:GetRole\",\n                \"iam:GetRolePolicy\",\n                \"iam:GetServiceLinkedRoleDeletionStatus\",\n                \"iam:GetUser\",\n                \"iam:GetUserPolicy\",\n                \"iam:ListAccessKeys\",\n                \"iam:ListAttachedGroupPolicies\",\n                \"iam:ListAttachedRolePolicies\",\n                \"iam:ListAttachedUserPolicies\",\n                \"iam:ListEntitiesForPolicy\",\n                \"iam:ListGroupsForUser\",\n                \"iam:ListInstanceProfilesForRole\",\n                \"iam:ListPolicies\",\n                \"iam:ListPolicyVersions\",\n                \"iam:ListRolePolicies\",\n                \"iam:PassRole\",\n                \"iam:PutGroupPolicy\",\n                \"iam:PutRolePolicy\",\n                \"iam:PutUserPolicy\",\n                \"iam:RemoveRoleFromInstanceProfile\",\n                \"iam:RemoveUserFromGroup\",\n                \"iam:TagPolicy\",\n                \"iam:TagRole\",\n                \"iam:TagUser\",\n                \"iam:UnTagRole\",\n                \"iam:UnTagUser\",\n                \"iam:UntagPolicy\",\n                \"iam:UpdateAccessKey\",\n                \"iam:UpdateRoleDescription\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor30\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"inspector:CreateAssessmentTarget\",\n                \"inspector:CreateAssessmentTemplate\",\n                \"inspector:CreateResourceGroup\",\n                \"inspector:DeleteAssessmentTarget\",\n                \"inspector:DeleteAssessmentTemplate\",\n                \"inspector:DescribeAssessmentTargets\",\n                \"inspector:DescribeAssessmentTemplates\",\n                \"inspector:DescribeResourceGroups\",\n                \"inspector:ListEventSubscriptions\",\n                \"inspector:ListRulesPackages\",\n                \"inspector:ListTagsForResource\",\n                \"inspector:SetTagsForResource\",\n                \"inspector:SubscribeToEvent\",\n                \"inspector:UnsubscribeFromEvent\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor31\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"kinesis:AddTagsToStream\",\n                \"kinesis:CreateStream\",\n                \"kinesis:DeleteStream\",\n                \"kinesis:DescribeStreamSummary\",\n                \"kinesis:EnableEnhancedMonitoring\",\n                \"kinesis:IncreaseStreamRetentionPeriod\",\n                \"kinesis:ListTagsForStream\",\n                \"kinesis:RemoveTagsFromStream\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor32\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"kinesisvideo:CreateStream\",\n                \"kinesisvideo:DeleteStream\",\n                \"kinesisvideo:DescribeStream\",\n                \"kinesisvideo:ListTagsForStream\",\n                \"kinesisvideo:TagStream\",\n                \"kinesisvideo:UntagStream\",\n                \"kinesisvideo:UpdateStream\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor33\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"kms:CreateAlias\",\n                \"kms:CreateGrant\",\n                \"kms:CreateKey\",\n                \"kms:Decrypt\",\n                \"kms:DeleteAlias\",\n                \"kms:DescribeKey\",\n                \"kms:DisableKey\",\n                \"kms:EnableKey\",\n                \"kms:EnableKeyRotation\",\n                \"kms:Encrypt\",\n                \"kms:GenerateDataKey*\",\n                \"kms:GetKeyPolicy\",\n                \"kms:GetKeyRotationStatus\",\n                \"kms:ListAliases\",\n                \"kms:ListResourceTags\",\n                \"kms:PutKeyPolicy\",\n                \"kms:ReEncrypt*\",\n                \"kms:ScheduleKeyDeletion\",\n                \"kms:TagResource\",\n                \"kms:UntagResource\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor34\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"lambda:AddPermission\",\n                \"lambda:CreateAlias\",\n                \"lambda:CreateFunction\",\n                \"lambda:DeleteAlias\",\n                \"lambda:DeleteFunction\",\n                \"lambda:GetAlias\",\n                \"lambda:GetFunction\",\n                \"lambda:GetFunctionCodeSigningConfig\",\n                \"lambda:GetPolicy\",\n                \"lambda:ListVersionsByFunction\",\n                \"lambda:RemovePermission\",\n                \"lambda:TagResource\",\n                \"lambda:UntagResource\",\n                \"lambda:UpdateAlias\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor35\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"logs:CreateLogGroup\",\n                \"logs:DeleteLogGroup\",\n                \"logs:DeleteMetricFilter\",\n                \"logs:DeleteResourcePolicy\",\n                \"logs:DeleteRetentionPolicy\",\n                \"logs:DeleteSubscriptionFilter\",\n                \"logs:DescribeLogGroups\",\n                \"logs:DescribeMetricFilters\",\n                \"logs:DescribeResourcePolicies\",\n                \"logs:DescribeSubscriptionFilters\",\n                \"logs:ListTagsLogGroup\",\n                \"logs:PutMetricFilter\",\n                \"logs:PutResourcePolicy\",\n                \"logs:PutRetentionPolicy\",\n                \"logs:PutSubscriptionFilter\",\n                \"logs:TagLogGroup\",\n                \"logs:UntagLogGroup\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor36\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"memorydb:CreateSubnetGroup\",\n                \"memorydb:DeleteSubnetGroup\",\n                \"memorydb:DescribeSubnetGroups\",\n                \"memorydb:ListTags\",\n                \"memorydb:TagResource\",\n                \"memorydb:UntagResource\",\n                \"memorydb:UpdateSubnetGroup\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor37\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"mq:CreateBroker\",\n                \"mq:CreateConfiguration\",\n                \"mq:CreateTags\",\n                \"mq:CreateUser\",\n                \"mq:DeleteBroker\",\n                \"mq:DeleteTags\",\n                \"mq:DeleteUser\",\n                \"mq:DescribeBroker\",\n                \"mq:DescribeConfiguration\",\n                \"mq:DescribeConfigurationRevision\",\n                \"mq:DescribeUser\",\n                \"mq:RebootBroker\",\n                \"mq:UpdateBroker\",\n                \"mq:UpdateConfiguration\",\n                \"mq:UpdateUser\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor38\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"organizations:DescribeOrganization\",\n                \"organizations:ListAWSServiceAccessForOrganization\",\n                \"organizations:ListAccounts\",\n                \"organizations:ListRoots\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor39\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"outposts:ListOutposts\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor40\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"rds:AddRoleToDBCluster\",\n                \"rds:AddTagsToResource\",\n                \"rds:CreateDBCluster\",\n                \"rds:CreateDBClusterEndpoint\",\n                \"rds:CreateDBClusterParameterGroup\",\n                \"rds:CreateDBClusterSnapshot\",\n                \"rds:CreateDBInstance\",\n                \"rds:CreateDBParameterGroup\",\n                \"rds:CreateDBSubnetGroup\",\n                \"rds:CreateGlobalCluster\",\n                \"rds:CreateOptionGroup\",\n                \"rds:DeleteDBCluster\",\n                \"rds:DeleteDBClusterEndpoint\",\n                \"rds:DeleteDBClusterParameterGroup\",\n                \"rds:DeleteDBClusterSnapshot\",\n                \"rds:DeleteDBParameterGroup\",\n                \"rds:DeleteDBSubnetGroup\",\n                \"rds:DeleteGlobalCluster\",\n                \"rds:DeleteOptionGroup\",\n                \"rds:DescribeCertificates\",\n                \"rds:DescribeDBClusterParameterGroups\",\n                \"rds:DescribeDBClusterParameters\",\n                \"rds:DescribeDBClusterSnapshots\",\n                \"rds:DescribeDBClusters\",\n                \"rds:DescribeDBEngineVersions\",\n                \"rds:DescribeDBInstances\",\n                \"rds:DescribeDBParameterGroups\",\n                \"rds:DescribeDBParameters\",\n                \"rds:DescribeDBSnapshots\",\n                \"rds:DescribeDBSubnetGroups\",\n                \"rds:DescribeEventCategories\",\n                \"rds:DescribeGlobalClusters\",\n                \"rds:DescribeOptionGroups\",\n                \"rds:DescribeOrderableDBInstanceOptions\",\n                \"rds:ListTagsForResource\",\n                \"rds:ModifyDBCluster\",\n                \"rds:ModifyDBClusterEndpoint\",\n                \"rds:ModifyDBClusterParameterGroup\",\n                \"rds:ModifyDBInstance\",\n                \"rds:ModifyDBParameterGroup\",\n                \"rds:ModifyGlobalCluster\",\n                \"rds:ModifyOptionGroup\",\n                \"rds:RemoveRoleFromDBCluster\",\n                \"rds:RemoveTagsFromResource\",\n                \"rds:StartActivityStream\",\n                \"rds:StopActivityStream\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor41\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"redshift:CreateAuthenticationProfile\",\n                \"redshift:CreateCluster\",\n                \"redshift:CreateClusterParameterGroup\",\n                \"redshift:CreateClusterSubnetGroup\",\n                \"redshift:CreateEventSubscription\",\n                \"redshift:CreateHsmClientCertificate\",\n                \"redshift:CreateHsmConfiguration\",\n                \"redshift:CreateScheduledAction\",\n                \"redshift:CreateSnapshotCopyGrant\",\n                \"redshift:CreateSnapshotSchedule\",\n                \"redshift:CreateTags\",\n                \"redshift:CreateUsageLimit\",\n                \"redshift:DeleteAuthenticationProfile\",\n                \"redshift:DeleteCluster\",\n                \"redshift:DeleteClusterParameterGroup\",\n                \"redshift:DeleteClusterSubnetGroup\",\n                \"redshift:DeleteEventSubscription\",\n                \"redshift:DeleteHsmClientCertificate\",\n                \"redshift:DeleteHsmConfiguration\",\n                \"redshift:DeleteScheduledAction\",\n                \"redshift:DeleteSnapshotCopyGrant\",\n                \"redshift:DeleteSnapshotSchedule\",\n                \"redshift:DeleteTags\",\n                \"redshift:DeleteUsageLimit\",\n                \"redshift:DescribeAuthenticationProfiles\",\n                \"redshift:DescribeClusterParameterGroups\",\n                \"redshift:DescribeClusterParameters\",\n                \"redshift:DescribeClusterSubnetGroups\",\n                \"redshift:DescribeClusters\",\n                \"redshift:DescribeEventSubscriptions\",\n                \"redshift:DescribeHsmClientCertificates\",\n                \"redshift:DescribeHsmConfigurations\",\n                \"redshift:DescribeLoggingStatus\",\n                \"redshift:DescribeOrderableClusterOptions\",\n                \"redshift:DescribeScheduledActions\",\n                \"redshift:DescribeSnapshotCopyGrants\",\n                \"redshift:DescribeSnapshotSchedules\",\n                \"redshift:DescribeUsageLimits\",\n                \"redshift:DisableLogging\",\n                \"redshift:EnableLogging\",\n                \"redshift:GetClusterCredentials\",\n                \"redshift:ModifyAuthenticationProfile\",\n                \"redshift:ModifyCluster\",\n                \"redshift:ModifyClusterIamRoles\",\n                \"redshift:ModifyClusterParameterGroup\",\n                \"redshift:ModifyClusterSnapshotSchedule\",\n                \"redshift:ModifyClusterSubnetGroup\",\n                \"redshift:ModifyEventSubscription\",\n                \"redshift:ModifyScheduledAction\",\n                \"redshift:ModifySnapshotSchedule\",\n                \"redshift:ModifyUsageLimit\",\n                \"redshift:PauseCluster\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor42\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"route53:AssociateVPCWithHostedZone\",\n                \"route53:ChangeResourceRecordSets\",\n                \"route53:ChangeTagsForResource\",\n                \"route53:CreateHostedZone\",\n                \"route53:DeleteHostedZone\",\n                \"route53:GetChange\",\n                \"route53:GetHostedZone\",\n                \"route53:ListHostedZones\",\n                \"route53:ListResourceRecordSets\",\n                \"route53:ListTagsForResource\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor43\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"s3:CreateBucket\",\n                \"s3:DeleteBucket\",\n                \"s3:DeleteObject\",\n                \"s3:GetAccelerateConfiguration\",\n                \"s3:GetBucketAcl\",\n                \"s3:GetBucketCORS\",\n                \"s3:GetBucketLocation\",\n                \"s3:GetBucketLogging\",\n                \"s3:GetBucketObjectLockConfiguration\",\n                \"s3:GetBucketPolicy\",\n                \"s3:GetBucketPublicAccessBlock\",\n                \"s3:GetBucketRequestPayment\",\n                \"s3:GetBucketTagging\",\n                \"s3:GetBucketVersioning\",\n                \"s3:GetBucketWebsite\",\n                \"s3:GetEncryptionConfiguration\",\n                \"s3:GetLifecycleConfiguration\",\n                \"s3:GetObject\",\n                \"s3:GetObjectAcl\",\n                \"s3:GetObjectTagging\",\n                \"s3:GetReplicationConfiguration\",\n                \"s3:ListAllMyBuckets\",\n                \"s3:ListBucket\",\n                \"s3:PutBucketAcl\",\n                \"s3:PutBucketLogging\",\n                \"s3:PutBucketObjectLockConfiguration\",\n                \"s3:PutBucketPolicy\",\n                \"s3:PutBucketPublicAccessBlock\",\n                \"s3:PutBucketVersioning\",\n                \"s3:PutEncryptionConfiguration\",\n                \"s3:PutLifecycleConfiguration\",\n                \"s3:PutObject\",\n                \"s3:PutObjectLegalHold\",\n                \"s3:PutObjectRetention\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor44\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"secretsmanager:CreateSecret\",\n                \"secretsmanager:DeleteSecret\",\n                \"secretsmanager:DescribeSecret\",\n                \"secretsmanager:GetResourcePolicy\",\n                \"secretsmanager:GetSecretValue\",\n                \"secretsmanager:PutSecretValue\",\n                \"secretsmanager:TagResource\",\n                \"secretsmanager:UntagResource\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor45\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"servicecatalog:CreatePortfolio\",\n                \"servicecatalog:DeletePortfolio\",\n                \"servicecatalog:DescribePortfolio\",\n                \"servicecatalog:TagResource\",\n                \"servicecatalog:UpdatePortfolio\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor46\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"sqs:GetQueueAttributes\",\n                \"sqs:ListQueueTags\",\n                \"sqs:SetQueueAttributes\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor47\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"ssm:AddTagsToResource\",\n                \"ssm:CreateDocument\",\n                \"ssm:CreateMaintenanceWindow\",\n                \"ssm:CreatePatchBaseline\",\n                \"ssm:DeleteDocument\",\n                \"ssm:DeleteMaintenanceWindow\",\n                \"ssm:DeleteParameter\",\n                \"ssm:DeletePatchBaseline\",\n                \"ssm:DeregisterPatchBaselineForPatchGroup\",\n                \"ssm:DeregisterTargetFromMaintenanceWindow\",\n                \"ssm:DeregisterTaskFromMaintenanceWindow\",\n                \"ssm:DescribeDocument\",\n                \"ssm:DescribeDocumentPermission\",\n                \"ssm:DescribeMaintenanceWindowTargets\",\n                \"ssm:DescribeMaintenanceWindowTasks\",\n                \"ssm:DescribeParameters\",\n                \"ssm:DescribePatchGroups\",\n                \"ssm:GetDocument\",\n                \"ssm:GetMaintenanceWindow\",\n                \"ssm:GetParameter\",\n                \"ssm:GetParameters\",\n                \"ssm:GetPatchBaseline\",\n                \"ssm:ListTagsForResource\",\n                \"ssm:PutParameter\",\n                \"ssm:RegisterPatchBaselineForPatchGroup\",\n                \"ssm:RegisterTargetWithMaintenanceWindow\",\n                \"ssm:RegisterTaskWithMaintenanceWindow\",\n                \"ssm:RemoveTagsFromResource\",\n                \"ssm:UpdateDocument\",\n                \"ssm:UpdateMaintenanceWindow\",\n                \"ssm:UpdatePatchBaseline\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor48\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"sso:ListInstances\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor49\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"states:CreateActivity\",\n                \"states:CreateStateMachine\",\n                \"states:DeleteActivity\",\n                \"states:DeleteStateMachine\",\n                \"states:DescribeActivity\",\n                \"states:DescribeStateMachine\",\n                \"states:ListTagsForResource\",\n                \"states:TagResource\",\n                \"states:UntagResource\",\n                \"states:UpdateStateMachine\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor50\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"wafv2:CreateIpSet\",\n                \"wafv2:CreateRegexPatternSet\",\n                \"wafv2:CreateRuleGroup\",\n                \"wafv2:CreateWebACL\",\n                \"wafv2:DeleteIpSet\",\n                \"wafv2:DeleteRegexPatternSet\",\n                \"wafv2:DeleteRuleGroup\",\n                \"wafv2:DeleteWebACL\",\n                \"wafv2:GetIpSet\",\n                \"wafv2:GetRegexPatternSet\",\n                \"wafv2:GetRuleGroup\",\n                \"wafv2:GetWebACL\",\n                \"wafv2:ListIPSets\",\n                \"wafv2:ListRegexPatternSets\",\n                \"wafv2:ListRuleGroups\",\n                \"wafv2:ListTagsForResource\",\n                \"wafv2:ListWebACLs\",\n                \"wafv2:TagResource\",\n                \"wafv2:UntagResource\",\n                \"wafv2:UpdateIpSet\",\n                \"wafv2:UpdateRegexPatternSet\",\n                \"wafv2:UpdateRuleGroup\"\n            ],\n            \"Resource\": \"*\"\n        }\n    ]\n}\n",
		Terraform: "resource \"aws_iam_policy\" \"terraform_pike\" {\n  name_prefix = \"terraform_pike\"\n  path        = \"/\"\n  description = \"Pike Autogenerated policy from IAC\"\n\n  policy = jsonencode({\n    \"Version\": \"2012-10-17\",\n    \"Statement\": [\n        {\n            \"Sid\": \"VisualEditor0\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"SNS:CreateTopic\",\n                \"SNS:DeleteTopic\",\n                \"SNS:GetTopicAttributes\",\n                \"SNS:ListTagsForResource\",\n                \"SNS:ListTopics\",\n                \"SNS:SetTopicAttributes\",\n                \"SNS:TagResource\",\n                \"SNS:UnTagResource\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor1\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"acm:AddTagsToCertificate\",\n                \"acm:DeleteCertificate\",\n                \"acm:DescribeCertificate\",\n                \"acm:ListTagsForCertificate\",\n                \"acm:RemoveTagsFromCertificate\",\n                \"acm:RequestCertificate\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor2\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"acm-pa:DescribeCertificateAuthority\",\n                \"acm-pa:ListTags\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor3\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"acm-pca:CreateCertificateAuthority\",\n                \"acm-pca:DeleteCertificateAuthority\",\n                \"acm-pca:GetCertificateAuthorityCertificate\",\n                \"acm-pca:UpdateCertificateAuthority\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor4\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"apigateway:DELETE\",\n                \"apigateway:GET\",\n                \"apigateway:PATCH\",\n                \"apigateway:POST\",\n                \"apigateway:PUT\",\n                \"apigateway:UpdateRestApiPolicy\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor5\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"application-autoscaling:DeleteScalingPolicy\",\n                \"application-autoscaling:DeleteScheduledAction\",\n                \"application-autoscaling:DeregisterScalableTarget\",\n                \"application-autoscaling:DescribeScalableTargets\",\n                \"application-autoscaling:DescribeScalingPolicies\",\n                \"application-autoscaling:DescribeScheduledActions\",\n                \"application-autoscaling:PutScalingPolicy\",\n                \"application-autoscaling:PutScheduledAction\",\n                \"application-autoscaling:RegisterScalableTarget\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor6\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"application-autoscaling:DeleteScalingPolicy\",\n                \"application-autoscaling:DeleteScheduledAction\",\n                \"application-autoscaling:DeregisterScalableTarget\",\n                \"application-autoscaling:DescribeScalableTargets\",\n                \"application-autoscaling:DescribeScalingPolicies\",\n                \"application-autoscaling:DescribeScheduledActions\",\n                \"application-autoscaling:PutScalingPolicy\",\n                \"application-autoscaling:PutScheduledAction\",\n                \"application-autoscaling:RegisterScalableTarget\",\n                \"autoscaling:AttachLoadBalancers\",\n                \"autoscaling:CreateAutoScalingGroup\",\n                \"autoscaling:CreateLaunchConfiguration\",\n                \"autoscaling:DeleteAutoScalingGroup\",\n                \"autoscaling:DeleteLaunchConfiguration\",\n                \"autoscaling:DescribeAutoScalingGroups\",\n                \"autoscaling:DescribeLaunchConfigurations\",\n                \"autoscaling:DescribeScalingActivities\",\n                \"autoscaling:DetachLoadBalancers\",\n                \"autoscaling:UpdateAutoScalingGroup\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor7\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"budgets:ModifyBudget\",\n                \"budgets:ViewBudget\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor8\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"cloudtrail:AddTags\",\n                \"cloudtrail:CreateTrail\",\n                \"cloudtrail:DeleteTrail\",\n                \"cloudtrail:DescribeTrails\",\n                \"cloudtrail:GetEventSelectors\",\n                \"cloudtrail:GetTrailStatus\",\n                \"cloudtrail:ListTags\",\n                \"cloudtrail:PutEventSelectors\",\n                \"cloudtrail:RemoveTags\",\n                \"cloudtrail:StartLogging\",\n                \"cloudtrail:UpdateTrail\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor9\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"cloudwatch:DeleteAlarms\",\n                \"cloudwatch:DescribeAlarms\",\n                \"cloudwatch:ListTagsForResource\",\n                \"cloudwatch:PutMetricAlarm\",\n                \"cloudwatch:TagResource\",\n                \"cloudwatch:UnTagResource\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor10\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"codeartifact:CreateDomain\",\n                \"codeartifact:CreateRepository\",\n                \"codeartifact:DeleteDomain\",\n                \"codeartifact:DeleteDomainPermissionsPolicy\",\n                \"codeartifact:DeleteRepository\",\n                \"codeartifact:DeleteRepositoryPermissionsPolicy\",\n                \"codeartifact:DescribeDomain\",\n                \"codeartifact:DescribeRepository\",\n                \"codeartifact:GetDomainPermissionsPolicy\",\n                \"codeartifact:GetRepositoryPermissionsPolicy\",\n                \"codeartifact:ListTagsForResource\",\n                \"codeartifact:PutDomainPermissionsPolicy\",\n                \"codeartifact:PutRepositoryPermissionsPolicy\",\n                \"codeartifact:TagResource\",\n                \"codeartifact:UntagResource\",\n                \"codeartifact:UpdateRepository\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor11\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"codebuild:BatchGetProjects\",\n                \"codebuild:CreateProject\",\n                \"codebuild:DeleteProject\",\n                \"codebuild:UpdateProject\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor12\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"codecommit:CreateRepository\",\n                \"codecommit:DeleteRepository\",\n                \"codecommit:GetRepository\",\n                \"codecommit:ListBranches\",\n                \"codecommit:ListTagsForResource\",\n                \"codecommit:TagResource\",\n                \"codecommit:UntagResource\",\n                \"codecommit:UpdateRepositoryDescription\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor13\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"codepipeline:CreatePipeline\",\n                \"codepipeline:DeletePipeline\",\n                \"codepipeline:GetPipeline\",\n                \"codepipeline:ListTagsForResource\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor14\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"cognito-idp:AddCustomAttributes\",\n                \"cognito-idp:AdminAddUserToGroup\",\n                \"cognito-idp:AdminCreateUser\",\n                \"cognito-idp:AdminDeleteUser\",\n                \"cognito-idp:AdminGetUser\",\n                \"cognito-idp:AdminListGroupsForUser\",\n                \"cognito-idp:AdminRemoveUserFromGroup\",\n                \"cognito-idp:CreateGroup\",\n                \"cognito-idp:CreateIdentityProvider\",\n                \"cognito-idp:CreateResourceServer\",\n                \"cognito-idp:CreateUserPool\",\n                \"cognito-idp:CreateUserPoolClient\",\n                \"cognito-idp:CreateUserPoolDomain\",\n                \"cognito-idp:DeleteGroup\",\n                \"cognito-idp:DeleteIdentityProvider\",\n                \"cognito-idp:DeleteResourceServer\",\n                \"cognito-idp:DeleteUserPool\",\n                \"cognito-idp:DeleteUserPoolClient\",\n                \"cognito-idp:DeleteUserPoolDomain\",\n                \"cognito-idp:DescribeIdentityProvider\",\n                \"cognito-idp:DescribeResourceServer\",\n                \"cognito-idp:DescribeUserPool\",\n                \"cognito-idp:DescribeUserPoolClient\",\n                \"cognito-idp:DescribeUserPoolDomain\",\n                \"cognito-idp:GetGroup\",\n                \"cognito-idp:GetSigningCertificate\",\n                \"cognito-idp:GetUICustomization\",\n                \"cognito-idp:GetUserPoolMfaConfig\",\n                \"cognito-idp:ListUserPoolClients\",\n                \"cognito-idp:ListUserPools\",\n                \"cognito-idp:SetUICustomization\",\n                \"cognito-idp:SetUserPoolMfaConfig\",\n                \"cognito-idp:TagResource\",\n                \"cognito-idp:UntagResource\",\n                \"cognito-idp:UpdateGroup\",\n                \"cognito-idp:UpdateIdentityProvider\",\n                \"cognito-idp:UpdateResourceServer\",\n                \"cognito-idp:UpdateUserPool\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor15\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"dax:CreateParameterGroup\",\n                \"dax:CreateSubnetGroup\",\n                \"dax:DeleteParameterGroup\",\n                \"dax:DeleteSubnetGroup\",\n                \"dax:DescribeParameterGroups\",\n                \"dax:DescribeParameters\",\n                \"dax:DescribeSubnetGroups\",\n                \"dax:UpdateParameterGroup\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor16\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"ds:AddTagsToResource\",\n                \"ds:CreateDirectory\",\n                \"ds:CreateLogSubscription\",\n                \"ds:DeleteDirectory\",\n                \"ds:DeleteLogSubscription\",\n                \"ds:ListTagsForResource\",\n                \"ds:RemoveTagsFromResource\",\n                \"rds:AddRoleToDBCluster\",\n                \"rds:AddTagsToResource\",\n                \"rds:CreateDBCluster\",\n                \"rds:CreateDBClusterEndpoint\",\n                \"rds:CreateDBClusterParameterGroup\",\n                \"rds:CreateDBClusterSnapshot\",\n                \"rds:CreateDBInstance\",\n                \"rds:CreateDBParameterGroup\",\n                \"rds:CreateDBSubnetGroup\",\n                \"rds:CreateGlobalCluster\",\n                \"rds:CreateOptionGroup\",\n                \"rds:DeleteDBCluster\",\n                \"rds:DeleteDBClusterEndpoint\",\n                \"rds:DeleteDBClusterParameterGroup\",\n                \"rds:DeleteDBClusterSnapshot\",\n                \"rds:DeleteDBParameterGroup\",\n                \"rds:DeleteDBSubnetGroup\",\n                \"rds:DeleteGlobalCluster\",\n                \"rds:DeleteOptionGroup\",\n                \"rds:DescribeCertificates\",\n                \"rds:DescribeDBClusterParameterGroups\",\n                \"rds:DescribeDBClusterParameters\",\n                \"rds:DescribeDBClusterSnapshots\",\n                \"rds:DescribeDBClusters\",\n                \"rds:DescribeDBEngineVersions\",\n                \"rds:DescribeDBInstances\",\n                \"rds:DescribeDBParameterGroups\",\n                \"rds:DescribeDBParameters\",\n                \"rds:DescribeDBSnapshots\",\n                \"rds:DescribeDBSubnetGroups\",\n                \"rds:DescribeEventCategories\",\n                \"rds:DescribeGlobalClusters\",\n                \"rds:DescribeOptionGroups\",\n                \"rds:DescribeOrderableDBInstanceOptions\",\n                \"rds:ListTagsForResource\",\n                \"rds:ModifyDBCluster\",\n                \"rds:ModifyDBClusterEndpoint\",\n                \"rds:ModifyDBClusterParameterGroup\",\n                \"rds:ModifyDBInstance\",\n                \"rds:ModifyDBParameterGroup\",\n                \"rds:ModifyGlobalCluster\",\n                \"rds:ModifyOptionGroup\",\n                \"rds:RemoveRoleFromDBCluster\",\n                \"rds:RemoveTagsFromResource\",\n                \"rds:StartActivityStream\",\n                \"rds:StopActivityStream\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor17\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"dynamodb:BatchWriteItem\",\n                \"dynamodb:CreateTable\",\n                \"dynamodb:CreateTableReplica\",\n                \"dynamodb:DeleteItem\",\n                \"dynamodb:DeleteTable\",\n                \"dynamodb:DeleteTableReplica\",\n                \"dynamodb:DescribeContinuousBackups\",\n                \"dynamodb:DescribeTable\",\n                \"dynamodb:DescribeTimeToLive\",\n                \"dynamodb:GetItem\",\n                \"dynamodb:ListTagsOfResource\",\n                \"dynamodb:PutItem\",\n                \"dynamodb:Query\",\n                \"dynamodb:Scan\",\n                \"dynamodb:TagResource\",\n                \"dynamodb:UntagResource\",\n                \"dynamodb:UpdateContinuousBackups\",\n                \"dynamodb:UpdateItem\",\n                \"dynamodb:UpdateTable\",\n                \"dynamodb:UpdateTimeToLive\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor18\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"ec2:AllocateAddress\",\n                \"ec2:AssociateAddress\",\n                \"ec2:AssociateRouteTable\",\n                \"ec2:AttachInternetGateway\",\n                \"ec2:AttachVolume\",\n                \"ec2:AttachVpnGateway\",\n                \"ec2:AuthorizeSecurityGroupEgress\",\n                \"ec2:AuthorizeSecurityGroupIngress\",\n                \"ec2:CancelCapacityReservation\",\n                \"ec2:CancelSpotInstanceRequests\",\n                \"ec2:CreateCapacityReservation\",\n                \"ec2:CreateDefaultVpc\",\n                \"ec2:CreateFlowLogs\",\n                \"ec2:CreateInternetGateway\",\n                \"ec2:CreateKeyPair\",\n                \"ec2:CreateLaunchTemplate\",\n                \"ec2:CreateLaunchTemplateVersion\",\n                \"ec2:CreateNatGateway\",\n                \"ec2:CreateNetworkAcl\",\n                \"ec2:CreateNetworkAclEntry\",\n                \"ec2:CreateNetworkInterface\",\n                \"ec2:CreateNetworkInterfacePermission\",\n                \"ec2:CreatePlacementGroup\",\n                \"ec2:CreateRoute\",\n                \"ec2:CreateRouteTable\",\n                \"ec2:CreateSecurityGroup\",\n                \"ec2:CreateSubnet\",\n                \"ec2:CreateTags\",\n                \"ec2:CreateVPC\",\n                \"ec2:CreateVolume\",\n                \"ec2:CreateVpcEndpoint\",\n                \"ec2:CreateVpnGateway\",\n                \"ec2:DeleteFlowLogs\",\n                \"ec2:DeleteInternetGateway\",\n                \"ec2:DeleteKeyPair\",\n                \"ec2:DeleteLaunchTemplate\",\n                \"ec2:DeleteNatGateway\",\n                \"ec2:DeleteNetworkAcl\",\n                \"ec2:DeleteNetworkAclEntry\",\n                \"ec2:DeleteNetworkInterface\",\n                \"ec2:DeleteNetworkInterfacePermission\",\n                \"ec2:DeletePlacementGroup\",\n                \"ec2:DeleteRoute\",\n                \"ec2:DeleteRouteTable\",\n                \"ec2:DeleteSecurityGroup\",\n                \"ec2:DeleteSubnet\",\n                \"ec2:DeleteTags\",\n                \"ec2:DeleteVPC\",\n                \"ec2:DeleteVolume\",\n                \"ec2:DeleteVpcEndpoints\",\n                \"ec2:DeleteVpnGateway\",\n                \"ec2:DescribeAccountAttributes\",\n                \"ec2:DescribeAddresses\",\n                \"ec2:DescribeAvailabilityZones\",\n                \"ec2:DescribeCapacityReservations\",\n                \"ec2:DescribeDhcpOptions\",\n                \"ec2:DescribeFlowLogs\",\n                \"ec2:DescribeImages\",\n                \"ec2:DescribeInstanceAttribute\",\n                \"ec2:DescribeInstanceCreditSpecifications\",\n                \"ec2:DescribeInstanceTypes\",\n                \"ec2:DescribeInstances\",\n                \"ec2:DescribeInternetGateways\",\n                \"ec2:DescribeKeyPairs\",\n                \"ec2:DescribeLaunchTemplateVersions\",\n                \"ec2:DescribeLaunchTemplates\",\n                \"ec2:DescribeNatGateways\",\n                \"ec2:DescribeNetworkAcls\",\n                \"ec2:DescribeNetworkInterfaces\",\n                \"ec2:DescribePlacementGroups\",\n                \"ec2:DescribePrefixLists\",\n                \"ec2:DescribeRouteTables\",\n                \"ec2:DescribeSecurityGroups\",\n                \"ec2:DescribeSpotInstanceRequests\",\n                \"ec2:DescribeSubnets\",\n                \"ec2:DescribeTags\",\n                \"ec2:DescribeVolumes\",\n                \"ec2:DescribeVpcAttribute\",\n                \"ec2:DescribeVpcEndpointServices\",\n                \"ec2:DescribeVpcEndpoints\",\n                \"ec2:DescribeVpcs\",\n                \"ec2:DescribeVpnGateways\",\n                \"ec2:DetachInternetGateway\",\n                \"ec2:DetachNetworkInterface\",\n                \"ec2:DetachVolume\",\n                \"ec2:DetachVpnGateway\",\n                \"ec2:DisassociateAddress\",\n                \"ec2:DisassociateRouteTable\",\n                \"ec2:GetEbsDefaultKmsKeyId\",\n                \"ec2:ImportKeyPair\",\n                \"ec2:ModifyCapacityReservation\",\n                \"ec2:ModifyInstanceAttribute\",\n                \"ec2:ModifyVolume\",\n                \"ec2:ModifyVpcEndpoint\",\n                \"ec2:MonitorInstances\",\n                \"ec2:ReleaseAddress\",\n                \"ec2:RequestSpotInstances\",\n                \"ec2:RevokeSecurityGroupEgress\",\n                \"ec2:RevokeSecurityGroupIngress\",\n                \"ec2:RunInstances\",\n                \"ec2:StartInstances\",\n                \"ec2:StopInstances\",\n                \"ec2:TerminateInstances\",\n                \"ec2:UnmonitorInstances\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor19\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"ecr:CreatePullThroughCacheRule\",\n                \"ecr:CreateRepository\",\n                \"ecr:DeleteLifecyclePolicy\",\n                \"ecr:DeletePullThroughCacheRule\",\n                \"ecr:DeleteRepository\",\n                \"ecr:DescribePullThroughCacheRules\",\n                \"ecr:DescribeRepositories\",\n                \"ecr:GetAuthorizationToken\",\n                \"ecr:GetLifecyclePolicy\",\n                \"ecr:ListTagsForResource\",\n                \"ecr:PutImageScanningConfiguration\",\n                \"ecr:PutLifecyclePolicy\",\n                \"ecr:TagResource\",\n                \"ecr:UntagResource\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor20\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"ecs:CreateCluster\",\n                \"ecs:CreateService\",\n                \"ecs:DeleteCluster\",\n                \"ecs:DeleteService\",\n                \"ecs:DeregisterTaskDefinition\",\n                \"ecs:DescribeClusters\",\n                \"ecs:DescribeServices\",\n                \"ecs:DescribeTaskDefinition\",\n                \"ecs:UpdateCluster\",\n                \"ecs:RegisterTaskDefinition\",\n                \"ecs:TagResource\",\n                \"ecs:UntagResource\",\n                \"ecs:UpdateService\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor21\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"eks:DescribeCluster\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor22\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"elasticache:AddTagsToResource\",\n                \"elasticache:CreateCacheParameterGroup\",\n                \"elasticache:CreateCacheSubnetGroup\",\n                \"elasticache:DeleteCacheParameterGroup\",\n                \"elasticache:DeleteCacheSubnetGroup\",\n                \"elasticache:DescribeCacheParameterGroups\",\n                \"elasticache:DescribeCacheParameters\",\n                \"elasticache:DescribeCacheSubnetGroups\",\n                \"elasticache:ListTagsForResource\",\n                \"elasticache:ModifyCacheParameterGroup\",\n                \"elasticache:ModifyCacheSubnetGroup\",\n                \"elasticache:RemoveTagsFromResource\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor23\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"elasticbeanstalk:AddTags\",\n                \"elasticbeanstalk:CreateApplication\",\n                \"elasticbeanstalk:DeleteApplication\",\n                \"elasticbeanstalk:DescribeApplications\",\n                \"elasticbeanstalk:ListAvailableSolutionStacks\",\n                \"elasticbeanstalk:ListTagsForResource\",\n                \"elasticbeanstalk:RemoveTags\",\n                \"elasticbeanstalk:UpdateApplicationResourceLifecycle\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor24\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"elasticfilesystem:CreateAccessPoint\",\n                \"elasticfilesystem:CreateFileSystem\",\n                \"elasticfilesystem:CreateReplicationConfiguration\",\n                \"elasticfilesystem:DeleteAccessPoint\",\n                \"elasticfilesystem:DeleteFileSystem\",\n                \"elasticfilesystem:DeleteFileSystemPolicy\",\n                \"elasticfilesystem:DeleteReplicationConfiguration\",\n                \"elasticfilesystem:DescribeAccessPoints\",\n                \"elasticfilesystem:DescribeBackupPolicy\",\n                \"elasticfilesystem:DescribeFileSystemPolicy\",\n                \"elasticfilesystem:DescribeFileSystems\",\n                \"elasticfilesystem:DescribeLifecycleConfiguration\",\n                \"elasticfilesystem:DescribeMountTargetSecurityGroups\",\n                \"elasticfilesystem:DescribeMountTargets\",\n                \"elasticfilesystem:DescribeReplicationConfigurations\",\n                \"elasticfilesystem:PutBackupPolicy\",\n                \"elasticfilesystem:PutFileSystemPolicy\",\n                \"elasticfilesystem:TagResource\",\n                \"elasticfilesystem:UntagResource\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor25\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"elasticloadbalancing:AddTags\",\n                \"elasticloadbalancing:AttachLoadBalancerToSubnets\",\n                \"elasticloadbalancing:CreateListener\",\n                \"elasticloadbalancing:CreateLoadBalancer\",\n                \"elasticloadbalancing:CreateLoadBalancerListeners\",\n                \"elasticloadbalancing:CreateTargetGroup\",\n                \"elasticloadbalancing:DeleteListener\",\n                \"elasticloadbalancing:DeleteLoadBalancer\",\n                \"elasticloadbalancing:DeleteTargetGroup\",\n                \"elasticloadbalancing:DeregisterTargets\",\n                \"elasticloadbalancing:DescribeListeners\",\n                \"elasticloadbalancing:DescribeLoadBalancerAttributes\",\n                \"elasticloadbalancing:DescribeLoadBalancers\",\n                \"elasticloadbalancing:DescribeTags\",\n                \"elasticloadbalancing:DescribeTargetGroupAttributes\",\n                \"elasticloadbalancing:DescribeTargetGroups\",\n                \"elasticloadbalancing:DescribeTargetHealth\",\n                \"elasticloadbalancing:ModifyListener\",\n                \"elasticloadbalancing:ModifyLoadBalancerAttributes\",\n                \"elasticloadbalancing:ModifyTargetGroupAttributes\",\n                \"elasticloadbalancing:RegisterTargets\",\n                \"elasticloadbalancing:RemoveTags\",\n                \"elasticloadbalancing:SetSecurityGroups\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor26\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"events:DeleteRule\",\n                \"events:DescribeRule\",\n                \"events:ListTagsForResource\",\n                \"events:ListTargetsByRule\",\n                \"events:PutRule\",\n                \"events:PutTargets\",\n                \"events:RemoveTargets\",\n                \"events:TagResource\",\n                \"events:UnTagResource\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor27\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"firehose:CreateDeliveryStream\",\n                \"firehose:DeleteDeliveryStream\",\n                \"firehose:DescribeDeliveryStream\",\n                \"firehose:ListTagsForDeliveryStream\",\n                \"firehose:TagDeliveryStream\",\n                \"firehose:UntagDeliveryStream\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor28\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"glue:CreateClassifier\",\n                \"glue:CreateConnection\",\n                \"glue:CreateCrawler\",\n                \"glue:CreateDatabase\",\n                \"glue:CreateJob\",\n                \"glue:CreateMLTransform\",\n                \"glue:CreateRegistry\",\n                \"glue:CreateSchema\",\n                \"glue:CreateScript\",\n                \"glue:CreateSecurityConfiguration\",\n                \"glue:CreateTable\",\n                \"glue:CreateTrigger\",\n                \"glue:CreateUserDefinedFunction\",\n                \"glue:CreateWorkflow\",\n                \"glue:DeleteClassifier\",\n                \"glue:DeleteConnection\",\n                \"glue:DeleteCrawler\",\n                \"glue:DeleteDatabase\",\n                \"glue:DeleteJob\",\n                \"glue:DeleteMLTransform\",\n                \"glue:DeleteRegistry\",\n                \"glue:DeleteResourcePolicy\",\n                \"glue:DeleteSchema\",\n                \"glue:DeleteSecurityConfiguration\",\n                \"glue:DeleteTable\",\n                \"glue:DeleteTrigger\",\n                \"glue:DeleteUserDefinedFunction\",\n                \"glue:DeleteWorkflow\",\n                \"glue:GetClassifier\",\n                \"glue:GetConnection\",\n                \"glue:GetCrawler\",\n                \"glue:GetDataCatalogEncryptionSettings\",\n                \"glue:GetDatabase\",\n                \"glue:GetJob\",\n                \"glue:GetMLTransform\",\n                \"glue:GetRegistry\",\n                \"glue:GetResourcePolicy\",\n                \"glue:GetSchema\",\n                \"glue:GetSchemaVersion\",\n                \"glue:GetSecurityConfiguration\",\n                \"glue:GetTable\",\n                \"glue:GetTags\",\n                \"glue:GetTrigger\",\n                \"glue:GetUserDefinedFunction\",\n                \"glue:GetWorkflow\",\n                \"glue:PutDataCatalogEncryptionSettings\",\n                \"glue:PutResourcePolicy\",\n                \"glue:TagResource\",\n                \"glue:UntagResource\",\n                \"glue:UpdateClassifier\",\n                \"glue:UpdateConnection\",\n                \"glue:UpdateCrawler\",\n                \"glue:UpdateDatabase\",\n                \"glue:UpdateJob\",\n                \"glue:UpdateMLTransform\",\n                \"glue:UpdateRegistry\",\n                \"glue:UpdateSchema\",\n                \"glue:UpdateTable\",\n                \"glue:UpdateTrigger\",\n                \"glue:UpdateUserDefinedFunction\",\n                \"glue:UpdateWorkflow\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor29\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"iam:AddRoleToInstanceProfile\",\n                \"iam:AddUserToGroup\",\n                \"iam:AttachGroupPolicy\",\n                \"iam:AttachRolePolicy\",\n                \"iam:AttachUserPolicy\",\n                \"iam:CreateAccessKey\",\n                \"iam:CreateGroup\",\n                \"iam:CreateInstanceProfile\",\n                \"iam:CreateLoginProfile\",\n                \"iam:CreatePolicy\",\n                \"iam:CreateRole\",\n                \"iam:CreateServiceLinkedRole\",\n                \"iam:CreateUser\",\n                \"iam:DeleteAccessKey\",\n                \"iam:DeleteGroup\",\n                \"iam:DeleteGroupPolicy\",\n                \"iam:DeleteInstanceProfile\",\n                \"iam:DeleteLoginProfile\",\n                \"iam:DeletePolicy\",\n                \"iam:DeleteRole\",\n                \"iam:DeleteRolePolicy\",\n                \"iam:DeleteServiceLinkedRole\",\n                \"iam:DeleteUser\",\n                \"iam:DeleteUserPolicy\",\n                \"iam:DetachGroupPolicy\",\n                \"iam:DetachRolePolicy\",\n                \"iam:DetachUserPolicy\",\n                \"iam:GetGroup\",\n                \"iam:GetGroupPolicy\",\n                \"iam:GetInstanceProfile\",\n                \"iam:GetLoginProfile\",\n                \"iam:GetPolicy\",\n                \"iam:GetPolicyVersion\",\n                \"iam:GetRole\",\n                \"iam:GetRolePolicy\",\n                \"iam:GetServiceLinkedRoleDeletionStatus\",\n                \"iam:GetUser\",\n                \"iam:GetUserPolicy\",\n                \"iam:ListAccessKeys\",\n                \"iam:ListAttachedGroupPolicies\",\n                \"iam:ListAttachedRolePolicies\",\n                \"iam:ListAttachedUserPolicies\",\n                \"iam:ListEntitiesForPolicy\",\n                \"iam:ListGroupsForUser\",\n                \"iam:ListInstanceProfilesForRole\",\n                \"iam:ListPolicies\",\n                \"iam:ListPolicyVersions\",\n                \"iam:ListRolePolicies\",\n                \"iam:PassRole\",\n                \"iam:PutGroupPolicy\",\n                \"iam:PutRolePolicy\",\n                \"iam:PutUserPolicy\",\n                \"iam:RemoveRoleFromInstanceProfile\",\n                \"iam:RemoveUserFromGroup\",\n                \"iam:TagPolicy\",\n                \"iam:TagRole\",\n                \"iam:TagUser\",\n                \"iam:UnTagRole\",\n                \"iam:UnTagUser\",\n                \"iam:UntagPolicy\",\n                \"iam:UpdateAccessKey\",\n                \"iam:UpdateRoleDescription\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor30\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"inspector:CreateAssessmentTarget\",\n                \"inspector:CreateAssessmentTemplate\",\n                \"inspector:CreateResourceGroup\",\n                \"inspector:DeleteAssessmentTarget\",\n                \"inspector:DeleteAssessmentTemplate\",\n                \"inspector:DescribeAssessmentTargets\",\n                \"inspector:DescribeAssessmentTemplates\",\n                \"inspector:DescribeResourceGroups\",\n                \"inspector:ListEventSubscriptions\",\n                \"inspector:ListRulesPackages\",\n                \"inspector:ListTagsForResource\",\n                \"inspector:SetTagsForResource\",\n                \"inspector:SubscribeToEvent\",\n                \"inspector:UnsubscribeFromEvent\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor31\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"kinesis:AddTagsToStream\",\n                \"kinesis:CreateStream\",\n                \"kinesis:DeleteStream\",\n                \"kinesis:DescribeStreamSummary\",\n                \"kinesis:EnableEnhancedMonitoring\",\n                \"kinesis:IncreaseStreamRetentionPeriod\",\n                \"kinesis:ListTagsForStream\",\n                \"kinesis:RemoveTagsFromStream\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor32\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"kinesisvideo:CreateStream\",\n                \"kinesisvideo:DeleteStream\",\n                \"kinesisvideo:DescribeStream\",\n                \"kinesisvideo:ListTagsForStream\",\n                \"kinesisvideo:TagStream\",\n                \"kinesisvideo:UntagStream\",\n                \"kinesisvideo:UpdateStream\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor33\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"kms:CreateAlias\",\n                \"kms:CreateGrant\",\n                \"kms:CreateKey\",\n                \"kms:Decrypt\",\n                \"kms:DeleteAlias\",\n                \"kms:DescribeKey\",\n                \"kms:DisableKey\",\n                \"kms:EnableKey\",\n                \"kms:EnableKeyRotation\",\n                \"kms:Encrypt\",\n                \"kms:GenerateDataKey*\",\n                \"kms:GetKeyPolicy\",\n                \"kms:GetKeyRotationStatus\",\n                \"kms:ListAliases\",\n                \"kms:ListResourceTags\",\n                \"kms:PutKeyPolicy\",\n                \"kms:ReEncrypt*\",\n                \"kms:ScheduleKeyDeletion\",\n                \"kms:TagResource\",\n                \"kms:UntagResource\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor34\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"lambda:AddPermission\",\n                \"lambda:CreateAlias\",\n                \"lambda:CreateFunction\",\n                \"lambda:DeleteAlias\",\n                \"lambda:DeleteFunction\",\n                \"lambda:GetAlias\",\n                \"lambda:GetFunction\",\n                \"lambda:GetFunctionCodeSigningConfig\",\n                \"lambda:GetPolicy\",\n                \"lambda:ListVersionsByFunction\",\n                \"lambda:RemovePermission\",\n                \"lambda:TagResource\",\n                \"lambda:UntagResource\",\n                \"lambda:UpdateAlias\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor35\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"logs:CreateLogGroup\",\n                \"logs:DeleteLogGroup\",\n                \"logs:DeleteMetricFilter\",\n                \"logs:DeleteResourcePolicy\",\n                \"logs:DeleteRetentionPolicy\",\n                \"logs:DeleteSubscriptionFilter\",\n                \"logs:DescribeLogGroups\",\n                \"logs:DescribeMetricFilters\",\n                \"logs:DescribeResourcePolicies\",\n                \"logs:DescribeSubscriptionFilters\",\n                \"logs:ListTagsLogGroup\",\n                \"logs:PutMetricFilter\",\n                \"logs:PutResourcePolicy\",\n                \"logs:PutRetentionPolicy\",\n                \"logs:PutSubscriptionFilter\",\n                \"logs:TagLogGroup\",\n                \"logs:UntagLogGroup\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor36\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"memorydb:CreateSubnetGroup\",\n                \"memorydb:DeleteSubnetGroup\",\n                \"memorydb:DescribeSubnetGroups\",\n                \"memorydb:ListTags\",\n                \"memorydb:TagResource\",\n                \"memorydb:UntagResource\",\n                \"memorydb:UpdateSubnetGroup\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor37\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"mq:CreateBroker\",\n                \"mq:CreateConfiguration\",\n                \"mq:CreateTags\",\n                \"mq:CreateUser\",\n                \"mq:DeleteBroker\",\n                \"mq:DeleteTags\",\n                \"mq:DeleteUser\",\n                \"mq:DescribeBroker\",\n                \"mq:DescribeConfiguration\",\n                \"mq:DescribeConfigurationRevision\",\n                \"mq:DescribeUser\",\n                \"mq:RebootBroker\",\n                \"mq:UpdateBroker\",\n                \"mq:UpdateConfiguration\",\n                \"mq:UpdateUser\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor38\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"organizations:DescribeOrganization\",\n                \"organizations:ListAWSServiceAccessForOrganization\",\n                \"organizations:ListAccounts\",\n                \"organizations:ListRoots\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor39\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"outposts:ListOutposts\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor40\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"rds:AddRoleToDBCluster\",\n                \"rds:AddTagsToResource\",\n                \"rds:CreateDBCluster\",\n                \"rds:CreateDBClusterEndpoint\",\n                \"rds:CreateDBClusterParameterGroup\",\n                \"rds:CreateDBClusterSnapshot\",\n                \"rds:CreateDBInstance\",\n                \"rds:CreateDBParameterGroup\",\n                \"rds:CreateDBSubnetGroup\",\n                \"rds:CreateGlobalCluster\",\n                \"rds:CreateOptionGroup\",\n                \"rds:DeleteDBCluster\",\n                \"rds:DeleteDBClusterEndpoint\",\n                \"rds:DeleteDBClusterParameterGroup\",\n                \"rds:DeleteDBClusterSnapshot\",\n                \"rds:DeleteDBParameterGroup\",\n                \"rds:DeleteDBSubnetGroup\",\n                \"rds:DeleteGlobalCluster\",\n                \"rds:DeleteOptionGroup\",\n                \"rds:DescribeCertificates\",\n                \"rds:DescribeDBClusterParameterGroups\",\n                \"rds:DescribeDBClusterParameters\",\n                \"rds:DescribeDBClusterSnapshots\",\n                \"rds:DescribeDBClusters\",\n                \"rds:DescribeDBEngineVersions\",\n                \"rds:DescribeDBInstances\",\n                \"rds:DescribeDBParameterGroups\",\n                \"rds:DescribeDBParameters\",\n                \"rds:DescribeDBSnapshots\",\n                \"rds:DescribeDBSubnetGroups\",\n                \"rds:DescribeEventCategories\",\n                \"rds:DescribeGlobalClusters\",\n                \"rds:DescribeOptionGroups\",\n                \"rds:DescribeOrderableDBInstanceOptions\",\n                \"rds:ListTagsForResource\",\n                \"rds:ModifyDBCluster\",\n                \"rds:ModifyDBClusterEndpoint\",\n                \"rds:ModifyDBClusterParameterGroup\",\n                \"rds:ModifyDBInstance\",\n                \"rds:ModifyDBParameterGroup\",\n                \"rds:ModifyGlobalCluster\",\n                \"rds:ModifyOptionGroup\",\n                \"rds:RemoveRoleFromDBCluster\",\n                \"rds:RemoveTagsFromResource\",\n                \"rds:StartActivityStream\",\n                \"rds:StopActivityStream\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor41\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"redshift:CreateAuthenticationProfile\",\n                \"redshift:CreateCluster\",\n                \"redshift:CreateClusterParameterGroup\",\n                \"redshift:CreateClusterSubnetGroup\",\n                \"redshift:CreateEventSubscription\",\n                \"redshift:CreateHsmClientCertificate\",\n                \"redshift:CreateHsmConfiguration\",\n                \"redshift:CreateScheduledAction\",\n                \"redshift:CreateSnapshotCopyGrant\",\n                \"redshift:CreateSnapshotSchedule\",\n                \"redshift:CreateTags\",\n                \"redshift:CreateUsageLimit\",\n                \"redshift:DeleteAuthenticationProfile\",\n                \"redshift:DeleteCluster\",\n                \"redshift:DeleteClusterParameterGroup\",\n                \"redshift:DeleteClusterSubnetGroup\",\n                \"redshift:DeleteEventSubscription\",\n                \"redshift:DeleteHsmClientCertificate\",\n                \"redshift:DeleteHsmConfiguration\",\n                \"redshift:DeleteScheduledAction\",\n                \"redshift:DeleteSnapshotCopyGrant\",\n                \"redshift:DeleteSnapshotSchedule\",\n                \"redshift:DeleteTags\",\n                \"redshift:DeleteUsageLimit\",\n                \"redshift:DescribeAuthenticationProfiles\",\n                \"redshift:DescribeClusterParameterGroups\",\n                \"redshift:DescribeClusterParameters\",\n                \"redshift:DescribeClusterSubnetGroups\",\n                \"redshift:DescribeClusters\",\n                \"redshift:DescribeEventSubscriptions\",\n                \"redshift:DescribeHsmClientCertificates\",\n                \"redshift:DescribeHsmConfigurations\",\n                \"redshift:DescribeLoggingStatus\",\n                \"redshift:DescribeOrderableClusterOptions\",\n                \"redshift:DescribeScheduledActions\",\n                \"redshift:DescribeSnapshotCopyGrants\",\n                \"redshift:DescribeSnapshotSchedules\",\n                \"redshift:DescribeUsageLimits\",\n                \"redshift:DisableLogging\",\n                \"redshift:EnableLogging\",\n                \"redshift:GetClusterCredentials\",\n                \"redshift:ModifyAuthenticationProfile\",\n                \"redshift:ModifyCluster\",\n                \"redshift:ModifyClusterIamRoles\",\n                \"redshift:ModifyClusterParameterGroup\",\n                \"redshift:ModifyClusterSnapshotSchedule\",\n                \"redshift:ModifyClusterSubnetGroup\",\n                \"redshift:ModifyEventSubscription\",\n                \"redshift:ModifyScheduledAction\",\n                \"redshift:ModifySnapshotSchedule\",\n                \"redshift:ModifyUsageLimit\",\n                \"redshift:PauseCluster\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor42\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"route53:AssociateVPCWithHostedZone\",\n                \"route53:ChangeResourceRecordSets\",\n                \"route53:ChangeTagsForResource\",\n                \"route53:CreateHostedZone\",\n                \"route53:DeleteHostedZone\",\n                \"route53:GetChange\",\n                \"route53:GetHostedZone\",\n                \"route53:ListHostedZones\",\n                \"route53:ListResourceRecordSets\",\n                \"route53:ListTagsForResource\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor43\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"s3:CreateBucket\",\n                \"s3:DeleteBucket\",\n                \"s3:DeleteObject\",\n                \"s3:GetAccelerateConfiguration\",\n                \"s3:GetBucketAcl\",\n                \"s3:GetBucketCORS\",\n                \"s3:GetBucketLocation\",\n                \"s3:GetBucketLogging\",\n                \"s3:GetBucketObjectLockConfiguration\",\n                \"s3:GetBucketPolicy\",\n                \"s3:GetBucketPublicAccessBlock\",\n                \"s3:GetBucketRequestPayment\",\n                \"s3:GetBucketTagging\",\n                \"s3:GetBucketVersioning\",\n                \"s3:GetBucketWebsite\",\n                \"s3:GetEncryptionConfiguration\",\n                \"s3:GetLifecycleConfiguration\",\n                \"s3:GetObject\",\n                \"s3:GetObjectAcl\",\n                \"s3:GetObjectTagging\",\n                \"s3:GetReplicationConfiguration\",\n                \"s3:ListAllMyBuckets\",\n                \"s3:ListBucket\",\n                \"s3:PutBucketAcl\",\n                \"s3:PutBucketLogging\",\n                \"s3:PutBucketObjectLockConfiguration\",\n                \"s3:PutBucketPolicy\",\n                \"s3:PutBucketPublicAccessBlock\",\n                \"s3:PutBucketVersioning\",\n                \"s3:PutEncryptionConfiguration\",\n                \"s3:PutLifecycleConfiguration\",\n                \"s3:PutObject\",\n                \"s3:PutObjectLegalHold\",\n                \"s3:PutObjectRetention\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor44\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"secretsmanager:CreateSecret\",\n                \"secretsmanager:DeleteSecret\",\n                \"secretsmanager:DescribeSecret\",\n                \"secretsmanager:GetResourcePolicy\",\n                \"secretsmanager:GetSecretValue\",\n                \"secretsmanager:PutSecretValue\",\n                \"secretsmanager:TagResource\",\n                \"secretsmanager:UntagResource\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor45\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"servicecatalog:CreatePortfolio\",\n                \"servicecatalog:DeletePortfolio\",\n                \"servicecatalog:DescribePortfolio\",\n                \"servicecatalog:TagResource\",\n                \"servicecatalog:UpdatePortfolio\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor46\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"sqs:GetQueueAttributes\",\n                \"sqs:ListQueueTags\",\n                \"sqs:SetQueueAttributes\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor47\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"ssm:AddTagsToResource\",\n                \"ssm:CreateDocument\",\n                \"ssm:CreateMaintenanceWindow\",\n                \"ssm:CreatePatchBaseline\",\n                \"ssm:DeleteDocument\",\n                \"ssm:DeleteMaintenanceWindow\",\n                \"ssm:DeleteParameter\",\n                \"ssm:DeletePatchBaseline\",\n                \"ssm:DeregisterPatchBaselineForPatchGroup\",\n                \"ssm:DeregisterTargetFromMaintenanceWindow\",\n                \"ssm:DeregisterTaskFromMaintenanceWindow\",\n                \"ssm:DescribeDocument\",\n                \"ssm:DescribeDocumentPermission\",\n                \"ssm:DescribeMaintenanceWindowTargets\",\n                \"ssm:DescribeMaintenanceWindowTasks\",\n                \"ssm:DescribeParameters\",\n                \"ssm:DescribePatchGroups\",\n                \"ssm:GetDocument\",\n                \"ssm:GetMaintenanceWindow\",\n                \"ssm:GetParameter\",\n                \"ssm:GetParameters\",\n                \"ssm:GetPatchBaseline\",\n                \"ssm:ListTagsForResource\",\n                \"ssm:PutParameter\",\n                \"ssm:RegisterPatchBaselineForPatchGroup\",\n                \"ssm:RegisterTargetWithMaintenanceWindow\",\n                \"ssm:RegisterTaskWithMaintenanceWindow\",\n                \"ssm:RemoveTagsFromResource\",\n                \"ssm:UpdateDocument\",\n                \"ssm:UpdateMaintenanceWindow\",\n                \"ssm:UpdatePatchBaseline\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor48\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"sso:ListInstances\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor49\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"states:CreateActivity\",\n                \"states:CreateStateMachine\",\n                \"states:DeleteActivity\",\n                \"states:DeleteStateMachine\",\n                \"states:DescribeActivity\",\n                \"states:DescribeStateMachine\",\n                \"states:ListTagsForResource\",\n                \"states:TagResource\",\n                \"states:UntagResource\",\n                \"states:UpdateStateMachine\"\n            ],\n            \"Resource\": \"*\"\n        },\n        {\n            \"Sid\": \"VisualEditor50\",\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"wafv2:CreateIpSet\",\n                \"wafv2:CreateRegexPatternSet\",\n                \"wafv2:CreateRuleGroup\",\n                \"wafv2:CreateWebACL\",\n                \"wafv2:DeleteIpSet\",\n                \"wafv2:DeleteRegexPatternSet\",\n                \"wafv2:DeleteRuleGroup\",\n                \"wafv2:DeleteWebACL\",\n                \"wafv2:GetIpSet\",\n                \"wafv2:GetRegexPatternSet\",\n                \"wafv2:GetRuleGroup\",\n                \"wafv2:GetWebACL\",\n                \"wafv2:ListIPSets\",\n                \"wafv2:ListRegexPatternSets\",\n                \"wafv2:ListRuleGroups\",\n                \"wafv2:ListTagsForResource\",\n                \"wafv2:ListWebACLs\",\n                \"wafv2:TagResource\",\n                \"wafv2:UntagResource\",\n                \"wafv2:UpdateIpSet\",\n                \"wafv2:UpdateRegexPatternSet\",\n                \"wafv2:UpdateRuleGroup\"\n            ],\n            \"Resource\": \"*\"\n        }\n    ]\n})\n}\n",
	}, GCP: "", AZURE: ""}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "aws", args: args{OutPolicy: out, output: "terraform", location: "."}, wantErr: false},
		{name: "aws-json", args: args{OutPolicy: out, output: "json", location: "."}, wantErr: false},
		{name: "aws-fail", args: args{OutPolicy: out, output: "cdk", location: "."}, wantErr: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if err := WriteOutput(tt.args.OutPolicy, tt.args.output, tt.args.location, ""); (err != nil) != tt.wantErr {
				t.Errorf("WriteOutput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLocateTerraform(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		os      string
		want    string
		wantErr bool
	}{
		{"find", "darwin", "/usr/local/bin/terraform", false},
		{"find", "windows", "C:\\ProgramData\\chocolatey\\bin\\terraform.exe", false},
		{"find", "linux", "/usr/local/bin/terraform", false},
	}

	for _, tt := range tests {
		tt := tt
		if tt.os == runtime.GOOS {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := LocateTerraform()

				if (err != nil) != tt.wantErr {
					t.Errorf("LocateTerraform() error = %v, wantErr %v", err, tt.wantErr)

					return
				}

				if got == "" {
					t.Errorf("LocateTerraform() = %v, expected %v", got, tt.want)
				}

				log.Info().Msgf("terraform is at %s", got)
			})
		}
	}
}

func TestInitWithEmptyDir(t *testing.T) {
	t.Parallel()

	tempDir, err := os.MkdirTemp("", "empty_tf_test")
	if err != nil {
		t.Fatal(err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer os.RemoveAll(tempDir)

	tfPath, modules, err := Init(tempDir)
	if err == nil {
		t.Error("Expected error for empty directory, got nil")
	}

	if tfPath == nil {
		t.Error("TFPath was not set")
	}

	if modules != nil {
		t.Errorf("Expected nil modules for empty directory, got %v", modules)
	}
}

//goland:noinspection GoUnhandledErrorResult
func TestInitWithInvalidTerraformConfig(t *testing.T) {
	t.Parallel()

	tempDir, err := os.MkdirTemp("", "invalid_tf_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create invalid terraform configuration
	invalidConfig := []byte(`
		resource "invalid" {
			bad config
		}
	`)

	err = os.WriteFile(filepath.Join(tempDir, "main.tf"), invalidConfig, 0o644)

	if err != nil {
		t.Fatal(err)
	}

	_, modules, err := Init(tempDir)
	if err == nil {
		t.Error("Expected error for invalid terraform config, got nil")
	}

	if modules != nil {
		t.Errorf("Expected nil modules for invalid config, got %v", modules)
	}
}

//goland:noinspection GoUnhandledErrorResult
func TestInitWithModulesJsonOnly(t *testing.T) {
	t.Parallel()

	tempDir, err := os.MkdirTemp("", "modulesjson_tf_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create .terraform/modules directory with only modules.json
	modulesDir := filepath.Join(tempDir, ".terraform", "modules")
	err = os.MkdirAll(modulesDir, 0o755)

	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile(filepath.Join(modulesDir, "modules.json"), []byte("{}"), 0o644)
	if err != nil {
		t.Fatal(err)
	}

	_, modules, err := Init(tempDir)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(modules) != 0 {
		t.Errorf("Expected empty modules slice, got %v", modules)
	}
}

//goland:noinspection GoUnhandledErrorResult
func TestInitWithDSStoreOnly(t *testing.T) {
	t.Parallel()

	tempDir, err := os.MkdirTemp("", "dsstore_tf_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create .terraform/modules directory with only .DS_Store
	modulesDir := filepath.Join(tempDir, ".terraform", "modules")
	err = os.MkdirAll(modulesDir, 0o755)

	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile(filepath.Join(modulesDir, ".DS_Store"), []byte{}, 0o644)

	if err != nil {
		t.Fatal(err)
	}

	_, modules, err := Init(tempDir)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(modules) != 0 {
		t.Errorf("Expected empty modules slice, got %v", modules)
	}
}

func TestInitWithNonExistentDir(t *testing.T) {
	t.Parallel()

	tfPath, modules, err := Init("/path/that/does/not/exist")

	if err == nil {
		t.Error("Expected error for non-existent directory, got nil")
	}

	if tfPath != nil {
		t.Errorf("Expected nil tfPath for non-existent directory, got %v", *tfPath)
	}

	if modules != nil {
		t.Errorf("Expected nil modules for non-existent directory, got %v", modules)
	}
}
