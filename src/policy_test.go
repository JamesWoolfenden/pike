package pike

import (
	_ "embed"
	"reflect"
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

//func TestGetPolicy(t *testing.T) {
//	type args struct {
//		actions Sorted
//		output  string
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    string
//		wantErr bool
//	}{
//{"first", args{Sorted{[]string{"s3:DeleteBucket",
//	"s3:CreateBucket",
//	"s3:GetLifecycleConfiguration",
//	"s3:GetBucketTagging",
//	"s3:GetBucketWebsite",
//	"s3:GetBucketLogging",
//	"s3:ListBucket",
//	"s3:GetAccelerateConfiguration",
//	"s3:GetBucketVersioning",
//	"s3:GetBucketAcl",
//	"s3:GetBucketPolicy",
//	"s3:GetReplicationConfiguration",
//	"s3:GetBucketObjectLockConfiguration",
//	"s3:GetObjectAcl",
//	"s3:GetObject",
//	"s3:GetEncryptionConfiguration",
//	"s3:GetBucketRequestPayment",
//	"s3:GetBucketCORS",
//	"s3:DeleteBucket"},
//	nil,
//	nil},
//	""},
//	Policy{
//		Version: "2012-10-17",
//		Statements: []Statement{
//			{"VisualEditor0", "Allow", []string{"s3:CreateBucket",
//				"s3:DeleteBucket",
//				"s3:GetAccelerateConfiguration",
//				"s3:GetBucketAcl",
//				"s3:GetBucketCORS",
//				"s3:GetBucketLogging",
//				"s3:GetBucketObjectLockConfiguration",
//				"s3:GetBucketPolicy",
//				"s3:GetBucketRequestPayment",
//				"s3:GetBucketTagging",
//				"s3:GetBucketVersioning",
//				"s3:GetBucketWebsite",
//				"s3:GetEncryptionConfiguration",
//				"s3:GetLifecycleConfiguration",
//				"s3:GetObject",
//				"s3:GetObjectAcl",
//				"s3:GetReplicationConfiguration",
//				"s3:ListBucket"}, "*"}}},
//	false},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := GetPolicy(tt.args.actions, tt.args.output)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetPolicy() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("GetPolicy() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestAWSPolicy(t *testing.T) {
	type args struct {
		Permissions []string
		output      string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AWSPolicy(tt.args.Permissions, tt.args.output)
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
