package pike_test

import (
	"reflect"
	"testing"

	pike "github.com/jameswoolfenden/pike/src"
)

func TestGetAWSDataPermissions(t *testing.T) {
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
			"placeholder",
			args{
				pike.ResourceV2{
					TypeName:     "data",
					Name:         "aws_iam_policy_document",
					ResourceName: "elasticsearch-log-publishing-policy",
					Provider:     "aws",
					Attributes:   []string{"statement", "actions", "resources", "principals", "type", "identifiers"},
				},
			},
			nil,
			false,
		},
		{
			"found",
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
			"bogus",
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
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := pike.GetAWSDataPermissions(tt.args.result)
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
