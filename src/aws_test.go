package pike

import (
	"reflect"
	"testing"
)

func TestGetAWSPermissions(t *testing.T) {
	type args struct {
		result ResourceV2
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"found",
			args{
				ResourceV2{"resource",
					"aws_api_gateway_api_key",
					"MyDemoApiKey",
					"aws",
					[]string{"name", "tags"}}},
			[]string{"apigateway:POST", "apigateway:PUT", "apigateway:PATCH", "apigateway:GET", "apigateway:DELETE", "apigateway:DELETE"},
			false},
		{"bogus",
			args{
				ResourceV2{"resource",
					"aws_madeup",
					"pike",
					"aws",
					[]string{"name"}},
			},
			nil,
			true},
		{"found_datasource",
			args{
				ResourceV2{"data",
					"aws_cloudwatch_log_group",
					"pike",
					"aws",
					[]string{"name"}},
			},
			[]string{"logs:DescribeLogGroups", "logs:ListTagsLogGroup"},
			false},
		{"bogus_datasource",
			args{
				ResourceV2{"data",
					"aws_madeup",
					"pike",
					"aws",
					[]string{"name"}},
			},
			nil,
			true},
		{"module",
			args{
				ResourceV2{"module",
					"instance",
					"pike",
					"",
					[]string{"name"}},
			},
			nil,
			false},
		{"duff",
			args{
				ResourceV2{"duff",
					"instance",
					"pike",
					"",
					[]string{"name"}},
			},
			nil,
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAWSPermissions(tt.args.result)
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
	type args struct {
		result ResourceV2
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"found",
			args{
				ResourceV2{"resource",
					"aws_api_gateway_api_key",
					"MyDemoApiKey",
					"aws",
					[]string{"name", "tags"}}},
			[]string{"apigateway:POST", "apigateway:PUT", "apigateway:PATCH", "apigateway:GET", "apigateway:DELETE", "apigateway:DELETE"},
			false},
		{"bogus",
			args{
				ResourceV2{"resource",
					"aws_madeup",
					"pike",
					"aws",
					[]string{"name"}},
			},
			nil,
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAWSResourcePermissions(tt.args.result)
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
	type args struct {
		s []string
		e string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"found", args{
			[]string{"dog", "cat"}, "cat"},
			true},
		{"missing", args{
			[]string{"dog", "cat"}, "fox"},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPermissionMap(t *testing.T) {
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
		{"found",
			args{
				awsAcmCertificate,
				[]string{"validation_method", "tags", "domain_name", "lifecycle", "create_before_destroy"}},
			[]string{
				"acm:AddTagsToCertificate",
				"acm:RemoveTagsFromCertificate",
				"acm:RequestCertificate",
				"acm:DescribeCertificate",
				"acm:ListTagsForCertificate",
				"acm:DeleteCertificate",
				"acm:DeleteCertificate"},
			false},
		{"bogus", args{[]byte("bogus"), []string{}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPermissionMap(tt.args.raw, tt.args.attributes)
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
