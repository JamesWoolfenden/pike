package pike

import (
	"fmt"
	"os"
	"strings"
)

// escalationGCP is the set of GCP permissions that grant owner-equivalent
// privilege escalation — the holder can grant themselves additional access.
var escalationGCP = map[string]bool{
	"resourcemanager.projects.setIamPolicy":      true,
	"resourcemanager.folders.setIamPolicy":       true,
	"resourcemanager.organizations.setIamPolicy": true,
	"iam.roles.create":                           true,
	"iam.roles.update":                           true,
	"iam.roles.delete":                           true,
	"iam.serviceAccounts.setIamPolicy":           true,
	"iam.serviceAccounts.getAccessToken":         true,
	"iam.serviceAccounts.signBlob":               true,
	"iam.serviceAccounts.signJwt":                true,
	"iam.serviceAccounts.implicitDelegation":     true,
	"iam.serviceAccountKeys.create":              true,
}

// escalationAWS is the set of AWS permissions that grant owner-equivalent
// privilege escalation.
var escalationAWS = map[string]bool{
	"iam:PutRolePolicy":           true,
	"iam:AttachRolePolicy":        true,
	"iam:CreatePolicyVersion":     true,
	"iam:SetDefaultPolicyVersion": true,
	"iam:UpdateAssumeRolePolicy":  true,
	"iam:PassRole":                true,
	"iam:CreateAccessKey":         true,
	"sts:AssumeRole":              true,
}

// escalationAZURE is the set of Azure permissions that grant owner-equivalent
// privilege escalation.
var escalationAZURE = map[string]bool{
	"Microsoft.Authorization/roleAssignments/write":                  true,
	"Microsoft.Authorization/roleDefinitions/write":                  true,
	"Microsoft.ManagedIdentity/userAssignedIdentities/assign/action": true,
}

// escalationSet groups escalation-class permissions found per provider.
type escalationSet struct {
	AWS   []string
	GCP   []string
	AZURE []string
}

func (e escalationSet) empty() bool {
	return len(e.AWS) == 0 && len(e.GCP) == 0 && len(e.AZURE) == 0
}

// findEscalation partitions a Sorted permission bag into escalation-class
// and base permissions per provider.
func findEscalation(bag Sorted) escalationSet {
	var found escalationSet
	for _, p := range Unique(bag.AWS) {
		if escalationAWS[p] {
			found.AWS = append(found.AWS, p)
		}
	}
	for _, p := range Unique(bag.GCP) {
		if escalationGCP[p] {
			found.GCP = append(found.GCP, p)
		}
	}
	for _, p := range Unique(bag.AZURE) {
		if escalationAZURE[p] {
			found.AZURE = append(found.AZURE, p)
		}
	}
	return found
}

// WarnEscalation prints a warning to stderr when escalation-class permissions
// are present. It is always called; callers do not need an opt-in flag.
func WarnEscalation(bag Sorted) {
	found := findEscalation(bag)
	if found.empty() {
		return
	}

	fmt.Fprintln(os.Stderr, "WARNING: escalation-class permissions detected — holder can grant themselves additional access.")
	fmt.Fprintln(os.Stderr, "         Consider the two-role pattern: planner SA (read-only) on all branches,")
	fmt.Fprintln(os.Stderr, "         applier SA (full) on protected branches only.")
	fmt.Fprintln(os.Stderr, "  Escalation permissions:")
	if len(found.AWS) > 0 {
		fmt.Fprintf(os.Stderr, "    aws:   %s\n", strings.Join(found.AWS, ", "))
	}
	if len(found.GCP) > 0 {
		fmt.Fprintf(os.Stderr, "    gcp:   %s\n", strings.Join(found.GCP, ", "))
	}
	if len(found.AZURE) > 0 {
		fmt.Fprintf(os.Stderr, "    azure: %s\n", strings.Join(found.AZURE, ", "))
	}
}
