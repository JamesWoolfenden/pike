package pike

import "testing"

func TestOutputPolicy_AsString(t *testing.T) {
	type fields struct {
		AWS   AwsOutput
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
		{"json", fields{AwsOutput{"[json]", ""}, "", ""}, args{"json"}, "[json]\n"},
		{"terraform", fields{AwsOutput{"[json]", "terraform"}, "", ""}, args{"terraform"}, "terraform\n"},
		{"all", fields{AwsOutput{"[json]", "terraform"}, "gcp", "azure"}, args{"terraform"}, "terraform\ngcp\nazure\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Out := OutputPolicy{
				AWS:   tt.fields.AWS,
				GCP:   tt.fields.GCP,
				Azure: tt.fields.Azure,
			}
			if got := Out.AsString(tt.args.format); got != tt.want {
				t.Errorf("OutputPolicy.AsString() = %v, want %v", got, tt.want)
			}
		})
	}
}
