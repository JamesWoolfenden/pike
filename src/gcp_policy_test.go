package pike_test

import (
	_ "embed"
	"os"
	"testing"

	pike "github.com/jameswoolfenden/pike/src"
)

func TestGCPPolicy(t *testing.T) {
	t.Parallel()
	os.Setenv("GCP_PROJECT", "pike-412922")
	type args struct {
		permissions []string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"basic",
			args{[]string{"bigquery.datasets.create", "bigquery.jobs.create"}},
			"resource\"google_project_iam_custom_role\"\"terraform_pike\"{project=\"pike-412922\"role_id=\"terraform_pike\"title=\"terraform_pike\"description=\"Auserwithleastprivileges\"permissions=[\"bigquery.datasets.create\",\"bigquery.jobs.create\"]}",
			false,
		},
		{
			"empty",
			args{[]string{}},
			"",
			true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := pike.GCPPolicy(tt.args.permissions, "")

			if (err != nil) != tt.wantErr {
				t.Errorf("GCPPolicy() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			original := Minify(got)
			target := Minify(tt.want)

			if original != target {
				t.Errorf("GCPPolicy() = %v, want %v", got, tt.want)
				t.Errorf("GCPPolicy() = %v, want %v", original, target)
			}
		})
	}
}
