package pike_test

import (
	"reflect"
	"testing"

	pike "github.com/jameswoolfenden/pike/src"
)

func TestGetAWSPermissions(t *testing.T) {
	t.Parallel()

	type args struct {
		result pike.ResourceV2
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "found",
			args: args{
				result: pike.ResourceV2{
					TypeName:     "resource",
					Name:         "aws_api_gateway_api_key",
					ResourceName: "MyDemoApiKey",
					Provider:     "aws",
					Attributes:   []string{"name", "tags"},
				},
			},
			want: []string{
				"apigateway:POST",
				"apigateway:PUT",
				"apigateway:PATCH",
				"apigateway:GET",
				"apigateway:DELETE",
				"apigateway:DELETE",
			},
			wantErr: false,
		},
		{
			name: "bogus",
			args: args{
				pike.ResourceV2{
					TypeName:     "resource",
					Name:         "aws_madeup",
					ResourceName: "pike",
					Provider:     "aws",
					Attributes:   []string{"name"},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "found_datasource",
			args: args{
				result: pike.ResourceV2{
					TypeName:     "data",
					Name:         "aws_cloudwatch_log_group",
					ResourceName: "pike",
					Provider:     "aws",
					Attributes:   []string{"name"},
				},
			},
			want: []string{"logs:DescribeLogGroups", "logs:ListTagsLogGroup"},
		},
		{
			name: "bogus_datasource",
			args: args{
				result: pike.ResourceV2{
					TypeName:     "data",
					Name:         "aws_madeup",
					ResourceName: "pike",
					Provider:     "aws",
					Attributes:   []string{"name"},
				},
			},
			wantErr: true,
		},
		{
			name: "module",
			args: args{
				result: pike.ResourceV2{
					TypeName:     "module",
					Name:         "instance",
					ResourceName: "pike",
					Attributes:   []string{"name"},
				},
			},
		},
		{
			name: "duff",
			args: args{
				result: pike.ResourceV2{
					TypeName:     "duff",
					Name:         "instance",
					ResourceName: "pike",
					Attributes:   []string{"name"},
				},
			},
			wantErr: true,
		},
		{
			name: "empty type name",
			args: args{pike.ResourceV2{
				Name:     "aws_s3_bucket",
				TypeName: "",
			}},
			wantErr: true,
		},
		{
			name: "empty resource name",
			args: args{pike.ResourceV2{
				Name:     "",
				TypeName: "resource",
			}},
			wantErr: true,
		},
		{
			name: "valid resource type",
			args: args{pike.ResourceV2{
				Name:     "aws_s3_bucket",
				TypeName: "resource",
			}},
			want: []string{
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
			wantErr: false,
		},
		{
			name: "valid data type",
			args: args{pike.ResourceV2{
				Name:     "aws_s3_bucket",
				TypeName: "data",
			}},
			want:    []string{"s3:ListBucket"},
			wantErr: false,
		},
		{
			name: "module type",
			args: args{pike.ResourceV2{
				Name:     "s3_module",
				TypeName: "module",
			}},
			wantErr: false,
		},
		{
			name: "unknown type",
			args: args{pike.ResourceV2{
				Name:     "aws_s3_bucket",
				TypeName: "unknown",
			}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := pike.GetAWSPermissions(tt.args.result)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetAWSPermissions() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAWSPermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAWSResourcePermissions(t *testing.T) {
	t.Parallel()

	type args struct {
		result pike.ResourceV2
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "found",
			args: args{
				result: pike.ResourceV2{
					TypeName:     "resource",
					Name:         "aws_api_gateway_api_key",
					ResourceName: "MyDemoApiKey",
					Provider:     "aws",
					Attributes:   []string{"name", "tags"},
				},
			},
			want: []string{
				"apigateway:POST",
				"apigateway:PUT",
				"apigateway:PATCH",
				"apigateway:GET",
				"apigateway:DELETE",
				"apigateway:DELETE",
			},
		},
		{
			name: "bogus",
			args: args{
				result: pike.ResourceV2{
					TypeName:     "resource",
					Name:         "aws_madeup",
					ResourceName: "pike",
					Provider:     "aws",
					Attributes:   []string{"name"},
				},
			},
			wantErr: true,
		},
		{
			name: "valid resource",
			args: args{pike.ResourceV2{
				Name:       "aws_s3_bucket",
				Attributes: []string{"bucket"},
			}},
			want: []string{
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
			wantErr: false,
		},
		{
			name: "non-existent resource",
			args: args{pike.ResourceV2{
				Name: "aws_nonexistent_resource",
			}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := pike.GetAWSResourcePermissions(tt.args.result)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetAWSResourcePermissions() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAWSResourcePermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contains(t *testing.T) {
	t.Parallel()

	type args struct {
		s []string
		e string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "found",
			args: args{
				[]string{"dog", "cat"}, "cat",
			},
			want: true,
		},
		{
			name: "missing",
			args: args{
				[]string{"dog", "cat"}, "fox",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := pike.Contains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPermissionMap(t *testing.T) {
	t.Parallel()

	type args struct {
		raw        []byte
		attributes []string
		resource   string
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "found",
			args: args{
				raw:        pike.AWSAcmCertificate,
				attributes: []string{"validation_method", "tags", "domain_name", "lifecycle", "create_before_destroy"},
				resource:   "aws_acm_certificate",
			},
			want: []string{
				"acm:AddTagsToCertificate",
				"acm:RemoveTagsFromCertificate",
				"acm:RequestCertificate",
				"acm:DescribeCertificate",
				"acm:ListTagsForCertificate",
				"acm:DeleteCertificate",
				"acm:DeleteCertificate",
			},
		},
		{
			name: "bogus",
			args: args{
				raw:        []byte("bogus"),
				attributes: []string{},
				resource:   "bogus",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := pike.GetPermissionMap(tt.args.raw, tt.args.attributes, tt.args.resource)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetPermissionMap() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPermissionMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTypeOK(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		args    interface{}
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "valid_map",
			args: map[string]interface{}{
				"key1": "value1",
				"key2": 123,
				"key3": true,
			},
			want: map[string]interface{}{
				"key1": "value1",
				"key2": 123,
				"key3": true,
			},
			wantErr: false,
		},
		{
			name:    "nil_input",
			args:    nil,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "string_input",
			args:    "not a map",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "slice_input",
			args:    []string{"not", "a", "map"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty_map",
			args:    map[string]interface{}{},
			want:    map[string]interface{}{},
			wantErr: false,
		},
		{
			name: "nested_map",
			args: map[string]interface{}{
				"outer": map[string]interface{}{
					"inner": "value",
				},
			},
			want: map[string]interface{}{
				"outer": map[string]interface{}{
					"inner": "value",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := pike.IsTypeOK(tt.args)

			if (err != nil) != tt.wantErr {
				t.Errorf("IsTypeOK() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsTypeOK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAwsLookup(t *testing.T) {
	tests := []struct {
		name         string
		resourceName string
		expectNil    bool
	}{
		{
			name:         "empty resource name",
			resourceName: "",
			expectNil:    true,
		},
		{
			name:         "valid resource",
			resourceName: "aws_s3_bucket",
			expectNil:    false,
		},
		{
			name:         "non-existent resource",
			resourceName: "aws_nonexistent_resource",
			expectNil:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := pike.AwsLookup(tt.resourceName)

			if tt.expectNil && result != nil {
				t.Errorf("expected nil but got %v", result)
			}

			if !tt.expectNil && result == nil {
				t.Errorf("expected non-nil result but got nil")
			}
		})
	}
}
