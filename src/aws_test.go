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
			"found",
			args{
				pike.ResourceV2{
					TypeName:     "resource",
					Name:         "aws_api_gateway_api_key",
					ResourceName: "MyDemoApiKey",
					Provider:     "aws",
					Attributes:   []string{"name", "tags"},
				},
			},
			[]string{
				"apigateway:POST",
				"apigateway:PUT",
				"apigateway:PATCH",
				"apigateway:GET",
				"apigateway:DELETE",
				"apigateway:DELETE",
			},
			false,
		},
		{
			"bogus",
			args{
				pike.ResourceV2{
					TypeName:     "resource",
					Name:         "aws_madeup",
					ResourceName: "pike",
					Provider:     "aws",
					Attributes:   []string{"name"},
				},
			},
			nil,
			true,
		},
		{
			"found_datasource",
			args{
				pike.ResourceV2{
					TypeName:     "data",
					Name:         "aws_cloudwatch_log_group",
					ResourceName: "pike",
					Provider:     "aws",
					Attributes:   []string{"name"},
				},
			},
			[]string{"logs:DescribeLogGroups", "logs:ListTagsLogGroup"},
			false,
		},
		{
			"bogus_datasource",
			args{
				pike.ResourceV2{
					TypeName:     "data",
					Name:         "aws_madeup",
					ResourceName: "pike",
					Provider:     "aws",
					Attributes:   []string{"name"},
				},
			},
			nil,
			true,
		},
		{
			"module",
			args{
				pike.ResourceV2{
					TypeName:     "module",
					Name:         "instance",
					ResourceName: "pike",
					Attributes:   []string{"name"},
				},
			},
			nil,
			false,
		},
		{
			"duff",
			args{
				pike.ResourceV2{
					TypeName:     "duff",
					Name:         "instance",
					ResourceName: "pike",
					Attributes:   []string{"name"},
				},
			},
			nil,
			true,
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
			"found",
			args{
				pike.ResourceV2{
					TypeName:     "resource",
					Name:         "aws_api_gateway_api_key",
					ResourceName: "MyDemoApiKey",
					Provider:     "aws",
					Attributes:   []string{"name", "tags"},
				},
			},
			[]string{
				"apigateway:POST",
				"apigateway:PUT",
				"apigateway:PATCH",
				"apigateway:GET",
				"apigateway:DELETE",
				"apigateway:DELETE",
			},
			false,
		},
		{
			"bogus",
			args{
				pike.ResourceV2{
					TypeName:     "resource",
					Name:         "aws_madeup",
					ResourceName: "pike",
					Provider:     "aws",
					Attributes:   []string{"name"},
				},
			},
			nil,
			true,
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
			"found",
			args{
				[]string{"dog", "cat"}, "cat",
			},
			true,
		},
		{
			"missing",
			args{
				[]string{"dog", "cat"}, "fox",
			},
			false,
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
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			"found",
			args{
				pike.AWSAcmCertificate,
				[]string{"validation_method", "tags", "domain_name", "lifecycle", "create_before_destroy"},
			},
			[]string{
				"acm:AddTagsToCertificate",
				"acm:RemoveTagsFromCertificate",
				"acm:RequestCertificate",
				"acm:DescribeCertificate",
				"acm:ListTagsForCertificate",
				"acm:DeleteCertificate",
				"acm:DeleteCertificate",
			},
			false,
		},
		{"bogus", args{[]byte("bogus"), []string{}}, nil, true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := pike.GetPermissionMap(tt.args.raw, tt.args.attributes)
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
