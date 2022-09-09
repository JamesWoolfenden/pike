package pike

import (
	"reflect"
	"testing"
)

func TestGetAWSDataPermissions(t *testing.T) {
	type args struct {
		result ResourceV2
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"placeholder",
			args{
				ResourceV2{"data",
					"aws_iam_policy_document",
					"elasticsearch-log-publishing-policy",
					"aws",
					[]string{"statement", "actions", "resources", "principals", "type", "identifiers"}}},
			nil,
			false},
		{"found",
			args{
				ResourceV2{"data",
					"aws_cloudwatch_log_group",
					"pike",
					"aws",
					[]string{"name"}},
			},
			[]string{"logs:DescribeLogGroups", "logs:ListTagsLogGroup"},
			false},
		{"bogus",
			args{
				ResourceV2{"data",
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
			got, err := GetAWSDataPermissions(tt.args.result)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAWSDataPermissions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAWSDataPermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}
