package policy_test

import (
	"testing"

	"github.com/qj0r9j0vc2/rampart/internal/policy"
)

func TestPolicy_Actions_Deduplicates(t *testing.T) {
	t.Parallel()

	p := policy.Policy{
		Statements: []policy.Statement{
			{Actions: []string{"ec2:DescribeInstances", "ec2:RunInstances"}, Resources: []string{"*"}},
			{Actions: []string{"ec2:DescribeInstances", "iam:PassRole"}, Resources: []string{"*"}},
		},
	}

	got := p.Actions()

	want := map[string]bool{
		"ec2:DescribeInstances": false,
		"ec2:RunInstances":      false,
		"iam:PassRole":          false,
	}

	if len(got) != len(want) {
		t.Fatalf("Actions() = %v, want %d unique actions", got, len(want))
	}

	for _, a := range got {
		if _, ok := want[a]; !ok {
			t.Errorf("Actions() contains unexpected action %q", a)
		}

		want[a] = true
	}

	for action, seen := range want {
		if !seen {
			t.Errorf("Actions() is missing %q", action)
		}
	}
}

func TestPolicy_Empty(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		p    policy.Policy
		want bool
	}{
		{"no statements", policy.Policy{}, true},
		{"one statement", policy.Policy{Statements: []policy.Statement{{Actions: []string{"ec2:DescribeInstances"}}}}, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.p.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}
