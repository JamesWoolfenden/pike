package config_test

import (
	"strings"
	"testing"

	"github.com/qj0r9j0vc2/rampart/internal/bootstrap/config"
)

func TestLoad_ValidExample(t *testing.T) {
	t.Parallel()

	cfg, err := config.Load("testdata/rampart.yaml")
	if err != nil {
		t.Fatalf("Load() error = %v, want nil", err)
	}

	if cfg.AWS.AccountID != "123456789012" {
		t.Errorf("AWS.AccountID = %q, want %q", cfg.AWS.AccountID, "123456789012")
	}

	if cfg.DeployRole.Name != "TerraformDeployRole" {
		t.Errorf("DeployRole.Name = %q, want %q", cfg.DeployRole.Name, "TerraformDeployRole")
	}

	if len(cfg.Policies.WorkloadBoundary.AllowedActions) != 8 {
		t.Errorf("len(Policies.WorkloadBoundary.AllowedActions) = %d, want 8", len(cfg.Policies.WorkloadBoundary.AllowedActions))
	}
}

func TestParse_PathNormalization(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"already normalized", "/terraform-deploy/", "/terraform-deploy/"},
		{"missing both slashes", "terraform-deploy", "/terraform-deploy/"},
		{"missing leading slash", "terraform-deploy/", "/terraform-deploy/"},
		{"missing trailing slash", "/terraform-deploy", "/terraform-deploy/"},
		{"empty defaults to root", "", "/"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			data := validYAML(func(y *yamlOverrides) {
				y.deployRolePath = tt.input
			})

			cfg, err := config.Parse([]byte(data))
			if err != nil {
				t.Fatalf("Parse() error = %v, want nil", err)
			}

			if cfg.DeployRole.Path != tt.want {
				t.Errorf("DeployRole.Path = %q, want %q", cfg.DeployRole.Path, tt.want)
			}
		})
	}
}

func TestParse_RejectsUnknownField(t *testing.T) {
	t.Parallel()

	data := strings.Replace(baseYAML, "version: 1", "version: 1\nbogus_field: true", 1)

	if _, err := config.Parse([]byte(data)); err == nil {
		t.Fatal("Parse() error = nil, want error for unknown field")
	}
}

func TestValidate_Rules(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		mutate  func(*yamlOverrides)
		wantErr string
	}{
		{
			name:    "unsupported version",
			mutate:  func(y *yamlOverrides) { y.version = "2" },
			wantErr: "unsupported config version",
		},
		{
			name:    "account id too short",
			mutate:  func(y *yamlOverrides) { y.accountID = "12345" },
			wantErr: "12 decimal digits",
		},
		{
			name:    "account id non-numeric",
			mutate:  func(y *yamlOverrides) { y.accountID = "12345678901a" },
			wantErr: "12 decimal digits",
		},
		{
			name:    "managed roles path is root",
			mutate:  func(y *yamlOverrides) { y.managedRolesPath = "/" },
			wantErr: `managed_roles.path must not be "/"`,
		},
		{
			name:    "mfa required without serial",
			mutate:  func(y *yamlOverrides) { y.mfaSerial = "" },
			wantErr: "principal.mfa.serial is required",
		},
		{
			name:    "deploy role path equals managed roles path",
			mutate:  func(y *yamlOverrides) { y.deployRolePath = "/terraform-managed/" },
			wantErr: "must be distinct",
		},
		{
			name:    "deploy role path equals boundary path",
			mutate:  func(y *yamlOverrides) { y.deployRolePath = "/boundaries/" },
			wantErr: "must not equal a boundary policy path",
		},
		{
			name:    "managed roles path equals boundary path",
			mutate:  func(y *yamlOverrides) { y.managedRolesPath = "/boundaries/" },
			wantErr: "must not equal a boundary policy path",
		},
		{
			name:    "workload allowlist contains wildcard",
			mutate:  func(y *yamlOverrides) { y.workloadAction = "*" },
			wantErr: "must not contain",
		},
		{
			name:    "workload allowlist contains iam wildcard",
			mutate:  func(y *yamlOverrides) { y.workloadAction = "iam:*" },
			wantErr: "must not contain",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			data := validYAML(tt.mutate)

			_, err := config.Parse([]byte(data))
			if err == nil {
				t.Fatalf("Parse() error = nil, want error containing %q", tt.wantErr)
			}

			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("Parse() error = %q, want it to contain %q", err.Error(), tt.wantErr)
			}
		})
	}
}

func TestValidate_DeployBoundaryAndWorkloadBoundarySharingAPathIsAllowed(t *testing.T) {
	t.Parallel()

	// The reference example intentionally gives deploy_boundary and
	// workload_boundary the same path; only the deploy-role and
	// managed-role paths must stay clear of it.
	if _, err := config.Parse([]byte(baseYAML)); err != nil {
		t.Fatalf("Parse() error = %v, want nil", err)
	}
}

// yamlOverrides lets each Validate test mutate exactly one field of an
// otherwise-valid config, so a failure can be attributed to that field.
type yamlOverrides struct {
	version          string
	accountID        string
	deployRolePath   string
	managedRolesPath string
	mfaSerial        string
	workloadAction   string
}

const baseYAML = `version: 1

aws:
  account_id: "123456789012"
  partition: aws
  region: ap-northeast-2

principal:
  type: iam-user
  name: cli-user
  mfa:
    required: true
    serial: arn:aws:iam::123456789012:mfa/cli-user
    max_age_seconds: 3600

deploy_role:
  name: TerraformDeployRole
  path: /terraform-deploy/
  max_session_duration: 3600

managed_roles:
  path: /terraform-managed/

policies:
  deploy:
    name: TerraformDeployPolicy
    path: /terraform-deploy/
  deploy_boundary:
    name: TerraformDeployBoundary
    path: /boundaries/
  workload_boundary:
    name: TerraformWorkloadBoundary
    path: /boundaries/
    mode: allowlist
    allowed_actions:
      - logs:PutLogEvents

pass_role:
  allowed_services:
    - ecs-tasks.amazonaws.com

service_linked_roles:
  allowed_services:
    - ecs.amazonaws.com

security:
  fail_on_unknown_iam_action: true
  deny_boundary_removal: true
  deny_boundary_policy_mutation: true
  deny_human_principal_management: true
  deny_access_key_management: true
  deny_organization_management: true

output:
  directory: .rampart/bootstrap
  overwrite: false
`

// validYAML applies zero or more field overrides on top of baseYAML via
// targeted string substitution, so each Validate test case only has to
// name the one field it wants to change.
func validYAML(mutate func(*yamlOverrides)) string {
	y := &yamlOverrides{
		version:          "1",
		accountID:        "123456789012",
		deployRolePath:   "/terraform-deploy/",
		managedRolesPath: "/terraform-managed/",
		mfaSerial:        "arn:aws:iam::123456789012:mfa/cli-user",
		workloadAction:   "logs:PutLogEvents",
	}
	mutate(y)

	data := baseYAML
	data = strings.Replace(data, "version: 1", "version: "+y.version, 1)
	data = strings.Replace(data, `account_id: "123456789012"`, `account_id: "`+y.accountID+`"`, 1)
	data = strings.Replace(data, "path: /terraform-deploy/\n  max_session_duration", "path: "+y.deployRolePath+"\n  max_session_duration", 1)
	data = strings.Replace(data, "managed_roles:\n  path: /terraform-managed/", "managed_roles:\n  path: "+y.managedRolesPath, 1)
	data = strings.Replace(data, "serial: arn:aws:iam::123456789012:mfa/cli-user", "serial: "+y.mfaSerial, 1)
	data = strings.Replace(data, "- logs:PutLogEvents", `- "`+y.workloadAction+`"`, 1)

	return data
}
