package pike

import "testing"

// These IAM actions are well-established privilege-escalation primitives
// (Rhino Security Labs' catalog) that escalationAWS previously omitted,
// letting IaC granting exactly one of them slip past pike scan's warning,
// `pike compare --strict`, `pike audit`, and `pike scan --output split`.
func TestEscalationAWS_IncludesKnownPrivescActions(t *testing.T) {
	t.Parallel()

	want := []string{
		"iam:AttachUserPolicy",
		"iam:AttachGroupPolicy",
		"iam:PutUserPolicy",
		"iam:PutGroupPolicy",
		"iam:AddUserToGroup",
		"iam:CreateLoginProfile",
		"iam:UpdateLoginProfile",
	}
	for _, action := range want {
		if !escalationAWS[action] {
			t.Errorf("escalationAWS missing known privesc action %q", action)
		}
	}
}

func TestFindEscalation_DetectsNewAWSActions(t *testing.T) {
	t.Parallel()

	found := findEscalation(Sorted{AWS: []string{"iam:AttachUserPolicy", "s3:GetObject"}})
	if len(found.AWS) != 1 || found.AWS[0] != "iam:AttachUserPolicy" {
		t.Errorf("got %v, want [iam:AttachUserPolicy]", found.AWS)
	}
}

func TestEscalationAWSServiceWildcards_DerivedFromEscalationAWS(t *testing.T) {
	t.Parallel()

	if !escalationAWSServiceWildcards["iam:*"] {
		t.Error(`escalationAWSServiceWildcards["iam:*"] = false, want true (iam: has escalation-class actions)`)
	}
	if !escalationAWSServiceWildcards["sts:*"] {
		t.Error(`escalationAWSServiceWildcards["sts:*"] = false, want true (sts:AssumeRole is escalation-class)`)
	}
	if escalationAWSServiceWildcards["s3:*"] {
		t.Error(`escalationAWSServiceWildcards["s3:*"] = true, want false (s3 has no escalation-class actions)`)
	}
}
