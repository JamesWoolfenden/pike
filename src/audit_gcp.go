package pike

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclsyntax"
)

var (
	gcpHighPrivRoles = map[string]bool{
		"roles/editor":                            true,
		"roles/iam.securityAdmin":                 true,
		"roles/resourcemanager.projectIamAdmin":   true,
		"roles/resourcemanager.folderIamAdmin":    true,
		"roles/resourcemanager.organizationAdmin": true,
	}
	gcpImpersonationRoles = map[string]bool{
		"roles/iam.serviceAccountTokenCreator": true,
		"roles/iam.serviceAccountUser":         true,
		"roles/iam.serviceAccountKeyAdmin":     true,
	}
)

func init() {
	for _, scope := range []string{"project", "folder", "organization"} {
		s := scope
		auditHandlers["resource.google_"+s+"_iam_member"] = func(b *hclsyntax.Block, _ string) []Finding {
			return auditGCPBinding(b, true)
		}
		auditHandlers["resource.google_"+s+"_iam_binding"] = func(b *hclsyntax.Block, _ string) []Finding {
			return auditGCPBinding(b, true)
		}
		auditHandlers["resource.google_"+s+"_iam_policy"] = auditGCPAuthoritativePolicy
	}
	auditHandlers["resource.google_service_account_iam_member"] = func(b *hclsyntax.Block, _ string) []Finding {
		return auditGCPBinding(b, false)
	}
	auditHandlers["resource.google_service_account_iam_binding"] = func(b *hclsyntax.Block, _ string) []Finding {
		return auditGCPBinding(b, false)
	}
	auditHandlers["resource.google_project_iam_custom_role"] = auditGCPCustomRole
	auditHandlers["resource.google_organization_iam_custom_role"] = auditGCPCustomRole
}

// auditGCPBinding checks the role and member(s) of an iam_member/iam_binding
// resource. broadScope is true for project/folder/org scope (where granting
// impersonation roles is risky) and false for service-account scope.
func auditGCPBinding(block *hclsyntax.Block, broadScope bool) []Finding {
	var out []Finding

	role := attrString(block, "role")
	switch {
	case role == "roles/owner":
		out = append(out, finding(block, "GCP001", SevCritical,
			"grants primitive role roles/owner",
			"owner is project-admin equivalent; use a least-privilege predefined or custom role"))
	case gcpHighPrivRoles[role]:
		out = append(out, finding(block, "GCP002", SevHigh,
			fmt.Sprintf("grants high-privilege role %s", role),
			"role permits IAM or broad resource changes; scope down if possible"))
	case broadScope && gcpImpersonationRoles[role]:
		out = append(out, finding(block, "GCP004", SevMedium,
			fmt.Sprintf("grants %s at %s scope", role, block.Labels[0]),
			"impersonation roles at broad scope allow lateral movement to any service account"))
	}

	members := attrStringList(block, "members")
	if m := attrString(block, "member"); m != "" {
		members = append(members, m)
	}
	for _, m := range members {
		if m == "allUsers" || m == "allAuthenticatedUsers" {
			out = append(out, finding(block, "GCP006", SevMedium,
				fmt.Sprintf("binds %s", m),
				"public principals should not hold IAM roles on private resources"))
		}
	}

	return out
}

func auditGCPAuthoritativePolicy(block *hclsyntax.Block, _ string) []Finding {
	return []Finding{finding(block, "GCP003", SevHigh,
		"authoritative iam_policy resource replaces all bindings on the target",
		"prefer google_*_iam_member or _iam_binding to avoid removing bindings managed elsewhere")}
}

func auditGCPCustomRole(block *hclsyntax.Block, _ string) []Finding {
	var out []Finding
	for _, p := range attrStringList(block, "permissions") {
		if escalationGCP[p] {
			out = append(out, finding(block, "GCP005", SevHigh,
				fmt.Sprintf("custom role includes escalation permission %s", p),
				"holder can grant themselves additional access"))
		}
	}
	return out
}
