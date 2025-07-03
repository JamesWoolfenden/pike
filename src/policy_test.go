package pike

import (
	_ "embed"
	"reflect"
	"testing"
)

func TestNewAWSPolicy(t *testing.T) {
	t.Parallel()

	type args struct {
		Actions []string
	}

	tests := []struct {
		name    string
		args    args
		want    Policy
		wantErr bool
	}{
		{
			"pass",
			args{[]string{
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
				"s3:ListBucket",
			}},
			Policy{
				Version: "2012-10-17",
				Statements: []Statement{
					{"VisualEditor0", "Allow", []string{
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
						"s3:ListBucket",
					}, []string{"*"}},
				},
			},
			false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewAWSPolicy(tt.args.Actions, false)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewAWSPolicy() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAWSPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPolicy(t *testing.T) {
	t.Parallel()

	type args struct {
		actions Sorted
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"first",
			args{Sorted{
				AWS: []string{},
			}},
			``,
			false,
		},
		{
			"aws",
			args{Sorted{AWS: []string{
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
				"ec2:TerminateInstances",
			}}},
			`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "VisualEditor0",
      "Effect": "Allow",
      "Action": [
        "ec2:DescribeInstanceAttribute",
        "ec2:DescribeInstanceCreditSpecifications",
        "ec2:DescribeInstanceTypes",
        "ec2:DescribeInstances",
        "ec2:DescribeTags",
        "ec2:DescribeVolumes",
        "ec2:ModifyInstanceAttribute",
        "ec2:RunInstances",
        "ec2:StartInstances",
        "ec2:StopInstances",
        "ec2:TerminateInstances"
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
			"short",
			args{Sorted{AWS: []string{"s3:*"}}},
			`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "VisualEditor0",
      "Effect": "Allow",
      "Action": [
        "s3:*"
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

			got, err := GetPolicy(tt.args.actions, false, "")

			if (err != nil) != tt.wantErr {
				t.Errorf("GetPolicy() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			newGot := Minify(got.AWS.JSONOut)
			reallyWant := Minify(tt.want)

			if newGot != reallyWant {
				t.Errorf("GetPolicy() = %v, want %v", got.AWS.JSONOut, tt.want)
			}
		})
	}
}

func TestAWSPolicy(t *testing.T) {
	t.Parallel()

	type args struct {
		Permissions []string
	}

	tests := []struct {
		name    string
		args    args
		want    AwsOutput
		wantErr bool
	}{
		{
			"fail",
			args{[]string{"woof"}},
			AwsOutput{},
			true,
		},
		{"fail2", args{[]string{"woof", "meow:*"}}, AwsOutput{}, true},
		{
			"pass",
			args{[]string{"woof:*"}},
			AwsOutput{JSONOut: `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "VisualEditor0",
      "Effect": "Allow",
      "Action": [
        "woof:*"
      ],
      "Resource": [
        "*"
      ]
    }
  ]
}`},
			false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := AWSPolicy(tt.args.Permissions, false, "")

			if (err != nil) != tt.wantErr {
				t.Errorf("AWSPolicy() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if Minify(got.JSONOut) != Minify(tt.want.JSONOut) {
				t.Errorf("AWSPolicy() = %v, want %v", got.JSONOut, tt.want.JSONOut)
			}
		})
	}
}

func Test_unique(t *testing.T) {
	t.Parallel()

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
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := Unique(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}
