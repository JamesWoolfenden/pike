// Package classify implements Rampart's IAM/STS action security
// classification (spec section 10). Every IAM or STS action a deploy
// policy requests gets classified here so deploy-boundary transformation
// knows how to treat it — preserve, scope to the managed-role path,
// explicitly deny, or fail closed. Actions for other services (ec2,
// ecs, ...) never go through this classifier; boundary transformation
// preserves those unchanged.
package classify

import "strings"

// Class is the security classification of a single IAM or STS action.
type Class string

const (
	Read              Class = "read"
	RoleCreate        Class = "role-create"
	RolePolicyWrite   Class = "role-policy-write"
	PassRole          Class = "pass-role"
	ServiceLinkedRole Class = "service-linked-role"
	BoundaryMutation  Class = "boundary-mutation"
	PolicyMutation    Class = "policy-mutation"
	HumanPrincipal    Class = "human-principal"
	Organization      Class = "organization"
	Unknown           Class = "unknown"
)

// exact classifies actions that don't fit a general verb-prefix rule, or
// that would otherwise be misclassified by one (e.g. iam:DeleteRole reads
// like a generic write but must be grouped with the managed-role-path
// writes in section 12.2, not left to fall through to Unknown).
var exact = map[string]Class{
	"iam:CreateRole": RoleCreate,

	// section 12.2 ManageTerraformRoles: the exact write actions the
	// deploy role gets, scoped to the managed-role path.
	"iam:AttachRolePolicy":           RolePolicyWrite,
	"iam:DeleteRole":                 RolePolicyWrite,
	"iam:DeleteRolePolicy":           RolePolicyWrite,
	"iam:DetachRolePolicy":           RolePolicyWrite,
	"iam:PutRolePolicy":              RolePolicyWrite,
	"iam:PutRolePermissionsBoundary": RolePolicyWrite,
	"iam:TagRole":                    RolePolicyWrite,
	"iam:UntagRole":                  RolePolicyWrite,

	"iam:PassRole": PassRole,

	"iam:CreateServiceLinkedRole": ServiceLinkedRole,

	// section 12.5 DenyWorkloadBoundaryRemoval.
	"iam:DeleteRolePermissionsBoundary": BoundaryMutation,

	// section 12.5 DenyBoundaryPolicyMutation.
	"iam:CreatePolicyVersion":     PolicyMutation,
	"iam:DeletePolicyVersion":     PolicyMutation,
	"iam:SetDefaultPolicyVersion": PolicyMutation,
	"iam:DeletePolicy":            PolicyMutation,

	// Human-identity and access-key management (section 10 groups both
	// under the single human-principal class).
	"iam:CreateUser":             HumanPrincipal,
	"iam:DeleteUser":             HumanPrincipal,
	"iam:UpdateUser":             HumanPrincipal,
	"iam:CreateLoginProfile":     HumanPrincipal,
	"iam:UpdateLoginProfile":     HumanPrincipal,
	"iam:DeleteLoginProfile":     HumanPrincipal,
	"iam:CreateAccessKey":        HumanPrincipal,
	"iam:UpdateAccessKey":        HumanPrincipal,
	"iam:DeleteAccessKey":        HumanPrincipal,
	"iam:CreateVirtualMFADevice": HumanPrincipal,
	"iam:DeactivateMFADevice":    HumanPrincipal,
	"iam:DeleteVirtualMFADevice": HumanPrincipal,
	"iam:EnableMFADevice":        HumanPrincipal,
	"iam:ResyncMFADevice":        HumanPrincipal,
}

// readVerbPrefixes are the AWS API verb conventions for read-only calls,
// applied across both iam: and sts: actions.
var readVerbPrefixes = []string{"Get", "List", "Generate", "Describe"}

// Classify returns the security classification for an IAM or STS action
// (e.g. "iam:CreateRole", "sts:GetCallerIdentity"). Unrecognised IAM/STS
// writes return Unknown so callers fail closed (section 4.2 rule 7)
// instead of silently allowing an action nobody has reasoned about.
func Classify(action string) Class {
	if class, ok := exact[action]; ok {
		return class
	}

	service, verb, ok := split(action)
	if !ok {
		return Unknown
	}

	// Organizations/account governance is denied outright regardless of
	// read or write (section 12.5), so it's checked before the read rule
	// and before the iam/sts scope restriction below.
	if service == "organizations" || service == "account" {
		return Organization
	}

	// The read-verb rule only applies within this classifier's actual
	// scope (IAM/STS). Without this, a non-IAM action that happens to
	// start with a read verb (e.g. "ec2:DescribeInstances") would be
	// mislabeled Read instead of Unknown.
	if service != "iam" && service != "sts" {
		return Unknown
	}

	for _, prefix := range readVerbPrefixes {
		if strings.HasPrefix(verb, prefix) {
			return Read
		}
	}

	return Unknown
}

func split(action string) (service, verb string, ok bool) {
	i := strings.IndexByte(action, ':')
	if i < 0 {
		return "", "", false
	}

	return action[:i], action[i+1:], true
}
