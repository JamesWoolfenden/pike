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
			name: "placeholder",
			args: args{
				result: pike.ResourceV2{
					TypeName:     "data",
					Name:         "aws_iam_policy_document",
					ResourceName: "elasticsearch-log-publishing-policy",
					Provider:     "aws",
					Attributes:   []string{"statement", "actions", "resources", "principals", "type", "identifiers"},
				},
			},
		},
		{
			name: "found",
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
			name: "bogus",
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
