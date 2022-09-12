package pike

import (
	_ "embed"
	"testing"
)

func TestGCPPolicy(t *testing.T) {
	type args struct {
		permissions []string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"basic", args{[]string{"bigquery.datasets.create", "bigquery.jobs.create"}},
			"resource \"google_project_iam_custom_role\" \"terraformXVlBzgba\" {\n  project     = \"pike\"\n  role_id     = \"terraform_pike\"\n  title       = \"terraformXVlBzgba\"\n  description = \"A user with least privileges\"\n  permissions = [\n    \"bigquery.datasets.create\",\n    \"bigquery.jobs.create\"\n  ]\n}\n",
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GCPPolicy(tt.args.permissions)
			if (err != nil) != tt.wantErr {
				t.Errorf("GCPPolicy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			original := minify(got)
			target := minify(tt.want)
			if original != target {
				t.Errorf("GCPPolicy() = %v, want %v", got, tt.want)
				t.Errorf("GCPPolicy() = %v, want %v", original, target)
			}
		})
	}
}
