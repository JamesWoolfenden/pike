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

func TestFindEscalationCombos(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		bag  Sorted
		want int
	}{
		{
			"PassRole alone is not a combo",
			Sorted{AWS: []string{"iam:PassRole", "s3:GetObject"}},
			0,
		},
		{
			"PassRole + lambda:CreateFunction is a combo",
			Sorted{AWS: []string{"iam:PassRole", "lambda:CreateFunction"}},
			1,
		},
		{
			"PassRole + ec2:RunInstances is a combo",
			Sorted{AWS: []string{"iam:PassRole", "ec2:RunInstances"}},
			1,
		},
		{
			"CreateAccessKey + ListUsers is a combo",
			Sorted{AWS: []string{"iam:CreateAccessKey", "iam:ListUsers"}},
			1,
		},
		{
			"two combos present at once",
			Sorted{AWS: []string{"iam:PassRole", "lambda:CreateFunction", "iam:CreateAccessKey", "iam:ListUsers"}},
			2,
		},
		{
			"empty bag",
			Sorted{},
			0,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := findEscalationCombos(tt.bag)
			if len(got) != tt.want {
				t.Errorf("findEscalationCombos() = %v (len %d), want len %d", got, len(got), tt.want)
			}
		})
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
