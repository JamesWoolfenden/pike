//go:build auth

package pike

import (
	"os"
	"reflect"
	"testing"
)

func TestCompare(t *testing.T) {
	t.Parallel()

	type args struct {
		directory string
		arn       string
		init      bool
	}
	os.Setenv("AWS_DEFAULT_PROFILE", "personal")
	os.Setenv("GCP_PROJECT", "pike-412922")

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"fail arn is empty", args{"./testdata/init/nicconf", "", false}, false, true},
		{"fail arn is not policy", args{"./testdata/init/nicconf", "arn:aws:iam::680235478471:user/readonly", false}, false, true},
		{"works but fails", args{"./testdata/init/nicconf", "arn:aws:iam::680235478471:policy/allows3", false}, false, false},
		//code is not aws
		{"gcp-basic-fail", args{"./testdata/gcp/basic", "basic", false}, false, true},
		{"gcp-basic-exist-fail", args{"./testdata/gcp/basic", "projects/pike-412922/roles/terraform_pike", false}, false, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := Compare(tt.args.directory, tt.args.arn, tt.args.init)
			if (err != nil) != tt.wantErr {
				t.Errorf("Compare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Compare() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listEnabledAPIs(t *testing.T) {
	type args struct {
		projectID string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"All", args{"pike-"}, nil, true},
		{"Enabled", args{"488072219970"}, []string{"analyticshub.googleapis.com",
			"artifactregistry.googleapis.com", "autoscaling.googleapis.com", "bigquery.googleapis.com",
			"bigqueryconnection.googleapis.com", "bigquerydatapolicy.googleapis.com", "bigquerymigration.googleapis.com",
			"bigqueryreservation.googleapis.com", "bigquerystorage.googleapis.com", "bigtable.googleapis.com",
			"bigtableadmin.googleapis.com", "cloudapis.googleapis.com", "cloudbuild.googleapis.com",
			"cloudfunctions.googleapis.com", "cloudkms.googleapis.com", "cloudresourcemanager.googleapis.com",
			"cloudtrace.googleapis.com", "composer.googleapis.com", "compute.googleapis.com", "container.googleapis.com",
			"containerfilesystem.googleapis.com", "containerregistry.googleapis.com", "dataform.googleapis.com",
			"dataplex.googleapis.com", "datastore.googleapis.com", "dns.googleapis.com", "gkebackup.googleapis.com",
			"iam.googleapis.com", "iamcredentials.googleapis.com", "logging.googleapis.com", "monitoring.googleapis.com",
			"networkconnectivity.googleapis.com", "oslogin.googleapis.com", "pubsub.googleapis.com", "run.googleapis.com",
			"servicehealth.googleapis.com", "servicemanagement.googleapis.com", "serviceusage.googleapis.com",
			"source.googleapis.com", "sql-component.googleapis.com", "sqladmin.googleapis.com", "storage-api.googleapis.com",
			"storage-component.googleapis.com",
			"storage.googleapis.com"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := listEnabledAPIs(tt.args.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("listEnabledAPIs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listEnabledAPIs() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compareGCPRole(t *testing.T) {
	type args struct {
		directory string
		arn       string
		init      bool
	}

	os.Setenv("GCP_PROJECT", "pike-412922")
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"pass", args{"./testdata/gcp/basic", "projects/pike-412922/roles/terraform_pike", false}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := compareGCPRole(tt.args.directory, tt.args.arn, tt.args.init)
			if (err != nil) != tt.wantErr {
				t.Errorf("compareGCPRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("compareGCPRole() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerifyRole(t *testing.T) {
	type args struct {
		role string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"Fail", args{"projectsmine/duff/roles/mine"}, false, true},
		{"Fail2", args{"projects/duff/noroles/mine"}, false, true},
		{"Fail3", args{"projects/duff/roles"}, false, true},
		{"Fail4", args{"projects/roles/a"}, false, true},
		{"Fail5", args{"mine/duff/roles/mine"}, false, true},

		{"Pass", args{"projects/a/roles/a"}, false, false},
		{"Pass2", args{"projects/duff/roles/mine"}, false, false},
		{"Pass3", args{role: "projects/pike-412922/roles/terraform_pike"}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := verifyGCPRole(tt.args.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
