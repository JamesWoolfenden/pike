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
	"iam:AttachUserPolicy":        true,
	"iam:AttachGroupPolicy":       true,
	"iam:PutUserPolicy":           true,
	"iam:PutGroupPolicy":          true,
	"iam:AddUserToGroup":          true,
	"iam:CreateLoginProfile":      true,
	"iam:UpdateLoginProfile":      true,
}

// escalationAWSServiceWildcards is derived from escalationAWS: the set of
// "service:*" wildcards that would grant every escalation-class action in
// that service (e.g. "iam:*" covers iam:PassRole, iam:PutRolePolicy, ...).
var escalationAWSServiceWildcards = func() map[string]bool {
	out := map[string]bool{}
	for action := range escalationAWS {
		if service, _, ok := strings.Cut(action, ":"); ok {
			out[service+":*"] = true
		}
	}
	return out
}()

// escalationAZURE is the set of Azure permissions that grant owner-equivalent
// privilege escalation.
var escalationAZURE = map[string]bool{
	"Microsoft.Authorization/roleAssignments/write":                  true,
	"Microsoft.Authorization/roleDefinitions/write":                  true,
	"Microsoft.ManagedIdentity/userAssignedIdentities/assign/action": true,
}

// escalationComboAWS lists AWS permission combinations where holding every
// permission in the set together opens a privilege-escalation path that no
// single permission in it reveals alone — e.g. iam:PassRole only lets you
// hand a role to a service; paired with lambda:CreateFunction or
// ec2:RunInstances, it lets you choose which compute resource receives that
// role and run arbitrary code as it.
var escalationComboAWS = [][]string{
	{"iam:PassRole", "lambda:CreateFunction"},
	{"iam:PassRole", "ec2:RunInstances"},
	{"iam:CreateAccessKey", "iam:ListUsers"},
}

// findEscalationCombos returns the entries of escalationComboAWS whose every
// permission is present in bag's AWS permissions.
func findEscalationCombos(bag Sorted) [][]string {
	have := Unique(bag.AWS)

	var found [][]string

	for _, combo := range escalationComboAWS {
		allPresent := true

		for _, p := range combo {
			if !Contains(have, p) {
				allPresent = false
				break
			}
		}

		if allPresent {
			found = append(found, combo)
		}
	}

	return found
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

// WarnEscalation prints a warning to stderr when escalation-class permissions,
// or a dangerous combination of otherwise-unremarkable permissions, are
// present. It is always called; callers do not need an opt-in flag.
func WarnEscalation(bag Sorted) {
	found := findEscalation(bag)
	combos := findEscalationCombos(bag)

	if found.empty() && len(combos) == 0 {
		return
	}

	fmt.Fprintln(os.Stderr, "WARNING: escalation-class permissions detected — holder can grant themselves additional access.")
	fmt.Fprintln(os.Stderr, "         Consider the two-role pattern: planner SA (read-only) on all branches,")
	fmt.Fprintln(os.Stderr, "         applier SA (full) on protected branches only.")
	fmt.Fprintln(os.Stderr, "         The escalation-class permissions below should not be granted to either CI")
	fmt.Fprintln(os.Stderr, "         identity — move them to a separate role applied outside CI (e.g. by a human,")
	fmt.Fprintln(os.Stderr, "         or a manually-triggered bootstrap pipeline) instead.")

	if !found.empty() {
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

	if len(combos) > 0 {
		fmt.Fprintln(os.Stderr, "  Dangerous combinations (no single permission above reveals this; together they do):")
		for _, combo := range combos {
			fmt.Fprintf(os.Stderr, "    aws:   %s\n", strings.Join(combo, " + "))
		}
	}
}
