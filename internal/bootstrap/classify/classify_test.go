package classify_test

import (
	"testing"

	"github.com/qj0r9j0vc2/rampart/internal/bootstrap/classify"
)

func TestClassify(t *testing.T) {
	t.Parallel()

	tests := []struct {
		action string
		want   classify.Class
	}{
		// section 10 table examples, verbatim.
		{"iam:GetRole", classify.Read},
		{"iam:ListRolePolicies", classify.Read},
		{"iam:CreateRole", classify.RoleCreate},
		{"iam:PutRolePolicy", classify.RolePolicyWrite},
		{"iam:AttachRolePolicy", classify.RolePolicyWrite},
		{"iam:PassRole", classify.PassRole},
		{"iam:CreateServiceLinkedRole", classify.ServiceLinkedRole},
		{"iam:DeleteRolePermissionsBoundary", classify.BoundaryMutation},
		{"iam:CreatePolicyVersion", classify.PolicyMutation},
		{"iam:CreateUser", classify.HumanPrincipal},
		{"iam:CreateAccessKey", classify.HumanPrincipal},

		// section 12.2 ManageTerraformRoles: every action in that
		// statement must land in role-policy-write.
		{"iam:DeleteRole", classify.RolePolicyWrite},
		{"iam:DeleteRolePolicy", classify.RolePolicyWrite},
		{"iam:DetachRolePolicy", classify.RolePolicyWrite},
		{"iam:PutRolePermissionsBoundary", classify.RolePolicyWrite},
		{"iam:TagRole", classify.RolePolicyWrite},
		{"iam:UntagRole", classify.RolePolicyWrite},

		// section 12.5 DenyBoundaryPolicyMutation: every action in that
		// statement must land in policy-mutation.
		{"iam:DeletePolicyVersion", classify.PolicyMutation},
		{"iam:SetDefaultPolicyVersion", classify.PolicyMutation},
		{"iam:DeletePolicy", classify.PolicyMutation},

		// organizations:*, account:* (section 7.1 / 10) — denied
		// regardless of whether the specific action is a read or a write.
		{"organizations:DescribeOrganization", classify.Organization},
		{"organizations:LeaveOrganization", classify.Organization},
		{"account:GetAccountInformation", classify.Organization},
		{"account:EnableRegion", classify.Organization},

		// generic read-verb rule applies across services, not just iam.
		{"sts:GetCallerIdentity", classify.Read},
		{"iam:ListAttachedRolePolicies", classify.Read},
		{"iam:GenerateCredentialReport", classify.Read},
		{"iam:GetAccountSummary", classify.Read},

		// unmapped IAM/STS writes fail closed.
		{"iam:TagPolicy", classify.Unknown},
		{"sts:AssumeRole", classify.Unknown},
		{"iam:UpdateRoleDescription", classify.Unknown},

		// malformed / non-IAM actions: out of this classifier's scope,
		// so also Unknown rather than guessed at.
		{"ec2:DescribeInstances", classify.Unknown},
		{"no-colon-here", classify.Unknown},
		{"", classify.Unknown},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.action, func(t *testing.T) {
			t.Parallel()

			if got := classify.Classify(tt.action); got != tt.want {
				t.Errorf("Classify(%q) = %q, want %q", tt.action, got, tt.want)
			}
		})
	}
}
