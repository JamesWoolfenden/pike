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

// gcpPolicyDataTypes are the authoritative *_iam_policy resource types whose
// policy_data may reference a data.google_iam_policy defined in the same
// scan — collectAuditFindings' reference-graph pass looks for exactly these.
var gcpPolicyDataTypes = map[string]bool{
	"google_project_iam_policy":      true,
	"google_folder_iam_policy":       true,
	"google_organization_iam_policy": true,
}

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
	auditHandlers["data.google_iam_policy"] = auditGCPIAMPolicyData
}

// auditGCPRoleBinding checks one role+members pair and returns findings
// attributed to addrBlock — the top-level resource/data block whose address
// names where this binding actually lives, which may differ from the block
// the role/members attributes were read from (e.g. a data.google_iam_policy's
// nested binding block has no address of its own). broadScope is true for
// project/folder/org scope (where granting impersonation roles is risky) and
// false for service-account scope.
func auditGCPRoleBinding(addrBlock *hclsyntax.Block, role string, members []string, broadScope bool) []Finding {
	var out []Finding

	switch {
	case role == "roles/owner":
		out = append(out, finding(addrBlock, "GCP001", SevCritical,
			"grants primitive role roles/owner",
			"owner is project-admin equivalent; use a least-privilege predefined or custom role"))
	case gcpHighPrivRoles[role]:
		out = append(out, finding(addrBlock, "GCP002", SevHigh,
			fmt.Sprintf("grants high-privilege role %s", role),
			"role permits IAM or broad resource changes; scope down if possible"))
	case broadScope && gcpImpersonationRoles[role]:
		out = append(out, finding(addrBlock, "GCP004", SevMedium,
			fmt.Sprintf("grants %s at %s scope", role, addrBlock.Labels[0]),
			"impersonation roles at broad scope allow lateral movement to any service account"))
	}

	for _, m := range members {
		if m == "allUsers" || m == "allAuthenticatedUsers" {
			out = append(out, finding(addrBlock, "GCP006", SevMedium,
				fmt.Sprintf("binds %s", m),
				"public principals should not hold IAM roles on private resources"))
		}
	}

	return out
}

// auditGCPBinding checks the role and member(s) of an iam_member/iam_binding
// resource. broadScope is true for project/folder/org scope (where granting
// impersonation roles is risky) and false for service-account scope.
func auditGCPBinding(block *hclsyntax.Block, broadScope bool) []Finding {
	members := attrStringList(block, "members")
	if m := attrString(block, "member"); m != "" {
		members = append(members, m)
	}
	return auditGCPRoleBinding(block, attrString(block, "role"), members, broadScope)
}

func auditGCPAuthoritativePolicy(block *hclsyntax.Block, _ string) []Finding {
	return []Finding{finding(block, "GCP003", SevHigh,
		"authoritative iam_policy resource replaces all bindings on the target",
		"prefer google_*_iam_member or _iam_binding to avoid removing bindings managed elsewhere")}
}

// auditGCPIAMPolicyData checks each binding block of a data.google_iam_policy
// — the actual role/member content an authoritative google_*_iam_policy
// resource applies via its policy_data reference — using the same
// content-aware checks already applied to _iam_member/_iam_binding
// resources. Findings are attributed to the data source itself, since the
// resource that applies it carries no binding content of its own to check;
// collectAuditFindings' reference-graph pass re-attributes them to whichever
// *_iam_policy resource's policy_data points here. broadScope is assumed
// true: this data source is overwhelmingly used with project/folder/org
// scope authoritative policies, which is also the scope the impersonation-
// role check (GCP004) exists to catch.
func auditGCPIAMPolicyData(block *hclsyntax.Block, _ string) []Finding {
	var out []Finding

	for _, inner := range block.Body.Blocks {
		if inner.Type != "binding" {
			continue
		}
		out = append(out, auditGCPRoleBinding(block, attrString(inner, "role"), attrStringList(inner, "members"), true)...)
	}

	return out
}

// gcpPolicyDataLabels is resolveReferenceGraph's labelsFor for GCP's
// authoritative-policy reference graph: policy_data has no principal of its
// own to name (the data source's bindings already say who), so every
// re-attributed finding gets the same fixed label.
func gcpPolicyDataLabels(*hclsyntax.Block) []string {
	return []string{"authoritative policy"}
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
