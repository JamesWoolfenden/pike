// Package scan adapts Pike's existing Terraform/OpenTofu analysis into
// Rampart's provider-neutral internal Policy model (see internal/policy).
// Rampart v0.1 targets AWS only, so the Analyzer here always scans for AWS
// permissions regardless of what other providers appear in the source.
package scan

import (
	"context"

	"github.com/qj0r9j0vc2/rampart/internal/policy"
	pike "github.com/qj0r9j0vc2/rampart/src"
)

// awsProvider selects AWS-only resources out of MakePermissionBag's
// mixed-provider scan; Rampart v0.1 has no use for GCP/Azure permissions.
const awsProvider = "aws"

// Analyzer scans a Terraform/OpenTofu directory and returns the
// permissions it requires as a provider-neutral Policy.
type Analyzer interface {
	Scan(ctx context.Context, directory string) (policy.Policy, error)
}

// PikeAnalyzer is the default Analyzer, backed by Pike's Terraform parser
// and AWS resource-to-permission mappings. It never runs `terraform init`
// and never requires AWS credentials — generation is static analysis only
// (spec section 20.2).
type PikeAnalyzer struct{}

// Scan implements Analyzer.
func (PikeAnalyzer) Scan(ctx context.Context, directory string) (policy.Policy, error) {
	if err := ctx.Err(); err != nil {
		return policy.Policy{}, err
	}

	const (
		init               = false
		suppressDeprecated = false
	)

	bag, err := pike.MakePermissionBag(directory, nil, init, awsProvider, suppressDeprecated)
	if err != nil {
		if pike.IsEmptyIAC(err) {
			return policy.Policy{}, nil
		}

		return policy.Policy{}, err
	}

	if len(bag.AWS) == 0 {
		return policy.Policy{}, nil
	}

	// resources=false: no live AWS account to build real ARNs from, so
	// every statement gets Resource ["*"] grouped by AWS service prefix.
	pikePolicy, err := pike.NewAWSPolicy(bag.AWS, false)
	if err != nil {
		return policy.Policy{}, err
	}

	return fromPikePolicy(pikePolicy), nil
}

func fromPikePolicy(p pike.Policy) policy.Policy {
	statements := make([]policy.Statement, 0, len(p.Statements))

	for _, s := range p.Statements {
		statements = append(statements, policy.Statement{
			Actions:   s.Action,
			Resources: s.Resource,
		})
	}

	return policy.Policy{Statements: statements}
}
