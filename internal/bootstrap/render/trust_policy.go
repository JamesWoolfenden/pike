package render

import (
	"fmt"

	"github.com/qj0r9j0vc2/rampart/internal/bootstrap/config"
)

// UnsupportedPrincipalTypeError is returned for any principal.type other
// than "iam-user" — spec section 2.2 excludes IAM Identity Center source
// principals from v0.1.
type UnsupportedPrincipalTypeError struct {
	Type string
}

func (e *UnsupportedPrincipalTypeError) Error() string {
	return fmt.Sprintf("unsupported principal.type %q (only \"iam-user\" is implemented in v0.1)", e.Type)
}

// TrustDocument is an IAM role trust policy (assume-role policy
// document). Unlike Document, its statements are Principal-based rather
// than Resource-based, and its conditions can use operators Document's
// Condition doesn't need (Bool, NumericLessThanEquals) — so it's a
// distinct type rather than a variant of Statement.
type TrustDocument struct {
	Version   string           `json:"Version"`
	Statement []TrustStatement `json:"Statement"`
}

// TrustStatement is a single statement in a TrustDocument.
type TrustStatement struct {
	Sid       string            `json:"Sid"`
	Effect    string            `json:"Effect"`
	Principal map[string]string `json:"Principal"`
	Action    string            `json:"Action"`
	Condition *TrustCondition   `json:"Condition,omitempty"`
}

// TrustCondition covers the two operators the MFA-enforced trust policy
// needs (spec section 14). encoding/json sorts map keys, so this stays
// canonical without extra work.
type TrustCondition struct {
	Bool                  map[string]string `json:"Bool,omitempty"`
	NumericLessThanEquals map[string]string `json:"NumericLessThanEquals,omitempty"`
}

// TrustPolicy renders trust-policy.json: the MFA-constrained deploy-role
// trust policy that lets the configured source principal assume it
// (spec section 14).
func TrustPolicy(cfg *config.Config) (TrustDocument, error) {
	if cfg.Principal.Type != "iam-user" {
		return TrustDocument{}, &UnsupportedPrincipalTypeError{Type: cfg.Principal.Type}
	}

	statement := TrustStatement{
		Sid:       "AssumeWithMFA",
		Effect:    "Allow",
		Principal: map[string]string{"AWS": userARN(cfg, cfg.Principal.Name)},
		Action:    "sts:AssumeRole",
	}

	if cfg.Principal.MFA.Required {
		condition := &TrustCondition{Bool: map[string]string{"aws:MultiFactorAuthPresent": "true"}}

		if cfg.Principal.MFA.MaxAgeSeconds > 0 {
			condition.NumericLessThanEquals = map[string]string{
				"aws:MultiFactorAuthAge": fmt.Sprintf("%d", cfg.Principal.MFA.MaxAgeSeconds),
			}
		}

		statement.Condition = condition
	}

	return TrustDocument{Version: policyVersion, Statement: []TrustStatement{statement}}, nil
}

// AssumeRolePolicy renders assume-role-policy.json: the minimal
// sts:AssumeRole grant for the source principal (spec section 14) —
// deliberately nothing more, matching the source-principal policy's own
// design constraint ("grants only sts:AssumeRole").
func AssumeRolePolicy(cfg *config.Config) Document {
	return Document{
		Version: policyVersion,
		Statements: []Statement{{
			Sid:      "AssumeTerraformDeployRole",
			Effect:   "Allow",
			Action:   []string{"sts:AssumeRole"},
			Resource: []string{deployRoleARN(cfg)},
		}},
	}
}

func userARN(cfg *config.Config, name string) string {
	return fmt.Sprintf("arn:%s:iam::%s:user/%s", cfg.AWS.Partition, cfg.AWS.AccountID, name)
}

func deployRoleARN(cfg *config.Config) string {
	return fmt.Sprintf("arn:%s:iam::%s:role%s%s", cfg.AWS.Partition, cfg.AWS.AccountID, cfg.DeployRole.Path, cfg.DeployRole.Name)
}
