package pike

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
)

var azureHighPrivBuiltins = map[string]bool{
	"Contributor":                             true,
	"User Access Administrator":               true,
	"Role Based Access Control Administrator": true,
}

func init() {
	auditHandlers["resource.azurerm_role_assignment"] = auditAzureRoleAssignment
	auditHandlers["resource.azurerm_role_definition"] = auditAzureRoleDefinition
}

func auditAzureRoleAssignment(block *hclsyntax.Block, _ string) []Finding {
	name := attrString(block, "role_definition_name")
	switch {
	case strings.EqualFold(name, "Owner"):
		return []Finding{finding(block, "AZURE001", SevCritical,
			"assigns built-in role Owner",
			"Owner grants full control including RBAC; use a least-privilege role")}
	case azureHighPrivBuiltins[name]:
		return []Finding{finding(block, "AZURE002", SevHigh,
			fmt.Sprintf("assigns high-privilege built-in role %s", name),
			"role permits broad resource or RBAC changes; scope down if possible")}
	}
	return nil
}

func auditAzureRoleDefinition(block *hclsyntax.Block, _ string) []Finding {
	var out []Finding
	for _, perms := range block.Body.Blocks {
		if perms.Type != "permissions" {
			continue
		}
		for _, a := range attrStringList(perms, "actions") {
			if a == "*" {
				out = append(out, finding(block, "AZURE003", SevHigh,
					"custom role actions includes \"*\"",
					"wildcard action grants every operation in scope"))
				continue
			}
			if escalationAZURE[a] {
				out = append(out, finding(block, "AZURE004", SevHigh,
					fmt.Sprintf("custom role includes escalation action %s", a),
					"holder can grant themselves additional access"))
			}
		}
	}
	return out
}
