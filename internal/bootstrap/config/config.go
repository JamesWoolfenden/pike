// Package config parses and validates rampart.yaml, the bootstrap
// configuration file describing the AWS account, source principal,
// deploy role and boundary policies Rampart generates IAM artifacts for.
package config

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

const SupportedVersion = 1

var accountIDPattern = regexp.MustCompile(`^[0-9]{12}$`)

// Config is the root of rampart.yaml.
type Config struct {
	Version            int                      `yaml:"version"`
	AWS                AWSConfig                `yaml:"aws"`
	Principal          PrincipalConfig          `yaml:"principal"`
	DeployRole         DeployRoleConfig         `yaml:"deploy_role"`
	ManagedRoles       ManagedRolesConfig       `yaml:"managed_roles"`
	Policies           PoliciesConfig           `yaml:"policies"`
	PassRole           PassRoleConfig           `yaml:"pass_role"`
	ServiceLinkedRoles ServiceLinkedRolesConfig `yaml:"service_linked_roles"`
	Security           SecurityConfig           `yaml:"security"`
	Output             OutputConfig             `yaml:"output"`
}

type AWSConfig struct {
	AccountID string `yaml:"account_id"`
	Partition string `yaml:"partition"`
	Region    string `yaml:"region"`
}

type MFAConfig struct {
	Required      bool   `yaml:"required"`
	Serial        string `yaml:"serial"`
	MaxAgeSeconds int    `yaml:"max_age_seconds"`
}

type PrincipalConfig struct {
	Type string    `yaml:"type"`
	Name string    `yaml:"name"`
	MFA  MFAConfig `yaml:"mfa"`
}

type DeployRoleConfig struct {
	Name               string `yaml:"name"`
	Path               string `yaml:"path"`
	MaxSessionDuration int    `yaml:"max_session_duration"`
}

type ManagedRolesConfig struct {
	Path string `yaml:"path"`
}

// PolicyConfig identifies a single managed policy by name and IAM path.
type PolicyConfig struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

// WorkloadBoundaryConfig is a PolicyConfig plus the explicit allowlist
// mode described in spec section 13 (the workload boundary is generated
// from configuration, never inferred from the deploy policy).
type WorkloadBoundaryConfig struct {
	Name           string   `yaml:"name"`
	Path           string   `yaml:"path"`
	Mode           string   `yaml:"mode"`
	AllowedActions []string `yaml:"allowed_actions"`
}

type PoliciesConfig struct {
	Deploy           PolicyConfig           `yaml:"deploy"`
	DeployBoundary   PolicyConfig           `yaml:"deploy_boundary"`
	WorkloadBoundary WorkloadBoundaryConfig `yaml:"workload_boundary"`
}

type PassRoleConfig struct {
	AllowedServices []string `yaml:"allowed_services"`
}

type ServiceLinkedRolesConfig struct {
	AllowedServices []string `yaml:"allowed_services"`
}

type SecurityConfig struct {
	FailOnUnknownIAMAction       bool `yaml:"fail_on_unknown_iam_action"`
	DenyBoundaryRemoval          bool `yaml:"deny_boundary_removal"`
	DenyBoundaryPolicyMutation   bool `yaml:"deny_boundary_policy_mutation"`
	DenyHumanPrincipalManagement bool `yaml:"deny_human_principal_management"`
	DenyAccessKeyManagement      bool `yaml:"deny_access_key_management"`
	DenyOrganizationManagement   bool `yaml:"deny_organization_management"`
}

type OutputConfig struct {
	Directory string `yaml:"directory"`
	Overwrite bool   `yaml:"overwrite"`
}

// workloadAllowlistDenylist is the set of actions/wildcards spec section
// 7.1 forbids in policies.workload_boundary.allowed_actions, since any of
// them would let a Terraform-created workload role manage or escalate IAM
// itself.
var workloadAllowlistDenylist = []string{"*", "iam:*", "sts:*", "organizations:*", "account:*"}

// Load reads and validates a rampart.yaml file at path. Unknown fields are
// rejected, and IAM paths are normalized (leading and trailing "/") before
// validation runs.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config: %w", err)
	}

	cfg, err := Parse(data)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}

	return cfg, nil
}

// Parse decodes and validates rampart.yaml content already read into memory.
func Parse(data []byte) (*Config, error) {
	var cfg Config

	dec := yaml.NewDecoder(bytes.NewReader(data))
	dec.KnownFields(true)

	if err := dec.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("parsing config: %w", err)
	}

	cfg.normalize()

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// normalize rewrites every configured IAM path to have both a leading and
// trailing "/", so downstream ARN construction never has to special-case
// missing slashes. An empty or root path normalizes to "/".
func (c *Config) normalize() {
	c.DeployRole.Path = normalizeIAMPath(c.DeployRole.Path)
	c.ManagedRoles.Path = normalizeIAMPath(c.ManagedRoles.Path)
	c.Policies.Deploy.Path = normalizeIAMPath(c.Policies.Deploy.Path)
	c.Policies.DeployBoundary.Path = normalizeIAMPath(c.Policies.DeployBoundary.Path)
	c.Policies.WorkloadBoundary.Path = normalizeIAMPath(c.Policies.WorkloadBoundary.Path)
}

func normalizeIAMPath(p string) string {
	trimmed := strings.Trim(p, "/")
	if trimmed == "" {
		return "/"
	}

	return "/" + trimmed + "/"
}

// Validate checks the self-contained rules from spec section 7.1 — the
// ones decidable from rampart.yaml alone. Rules that depend on the
// Terraform scan result (PassRole/CreateServiceLinkedRole services being
// configured when those actions are actually detected) belong to the
// transform/validate stage, once a scanned Policy is available to check
// against.
func (c *Config) Validate() error {
	if c.Version != SupportedVersion {
		return fmt.Errorf("unsupported config version %d (supported: %d)", c.Version, SupportedVersion)
	}

	if !accountIDPattern.MatchString(c.AWS.AccountID) {
		return fmt.Errorf("aws.account_id must be exactly 12 decimal digits, got %q", c.AWS.AccountID)
	}

	if c.ManagedRoles.Path == "/" {
		return fmt.Errorf("managed_roles.path must not be \"/\"")
	}

	if c.Principal.MFA.Required && c.Principal.MFA.Serial == "" {
		return fmt.Errorf("principal.mfa.serial is required when principal.mfa.required is true")
	}

	if err := c.validateDistinctPaths(); err != nil {
		return err
	}

	if err := c.validateWorkloadAllowlist(); err != nil {
		return err
	}

	return nil
}

// validateDistinctPaths enforces "deploy, managed-role and boundary policy
// paths must be distinct" (section 7.1). The reference config in section 7
// deliberately gives deploy_boundary and workload_boundary the *same* path
// ("/boundaries/"), so that pair is not required to differ from each
// other — only the deploy-role path and the managed-role path must each
// stay clear of the other two categories and of each other.
func (c *Config) validateDistinctPaths() error {
	if c.DeployRole.Path == c.ManagedRoles.Path {
		return fmt.Errorf("deploy_role.path and managed_roles.path must be distinct, both are %q", c.DeployRole.Path)
	}

	for _, boundaryPath := range []string{c.Policies.DeployBoundary.Path, c.Policies.WorkloadBoundary.Path} {
		if c.DeployRole.Path == boundaryPath {
			return fmt.Errorf("deploy_role.path must not equal a boundary policy path (%q)", boundaryPath)
		}

		if c.ManagedRoles.Path == boundaryPath {
			return fmt.Errorf("managed_roles.path must not equal a boundary policy path (%q)", boundaryPath)
		}
	}

	return nil
}

func (c *Config) validateWorkloadAllowlist() error {
	for _, action := range c.Policies.WorkloadBoundary.AllowedActions {
		for _, denied := range workloadAllowlistDenylist {
			if action == denied {
				return fmt.Errorf("policies.workload_boundary.allowed_actions must not contain %q", action)
			}
		}
	}

	return nil
}
