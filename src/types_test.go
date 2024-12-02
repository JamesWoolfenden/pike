package pike_test

import (
	"testing"

	pike "github.com/jameswoolfenden/pike/src"
)

func TestOutputPolicy_AsString(t *testing.T) {
	t.Parallel()

	type fields struct {
		AWS   pike.AwsOutput
		GCP   string
		Azure string
	}

	type args struct {
		format string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"json",
			fields{pike.AwsOutput{JSONOut: "[json]"}, "", ""},
			args{"json"},
			"[json]\n",
		},
		{
			"terraform",
			fields{
				pike.AwsOutput{JSONOut: "[json]", Terraform: "terraform"}, "", "",
			},
			args{"terraform"},
			"terraform\n",
		},
		{
			"all",
			fields{pike.AwsOutput{JSONOut: "[json]", Terraform: "terraform"}, "google", "azure"},
			args{"terraform"},
			"terraform\ngoogle\nazure\n",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			Out := pike.OutputPolicy{
				AWS:   tt.fields.AWS,
				GCP:   tt.fields.GCP,
				AZURE: tt.fields.Azure,
			}
			if got := Out.AsString(tt.args.format); got != tt.want {
				t.Errorf("OutputPolicy.AsString() = %v, want %v", got, tt.want)
			}
		})
	}
}
