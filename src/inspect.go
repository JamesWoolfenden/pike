package pike

import (
	"errors"
	"fmt"

	Identity "github.com/jameswoolfenden/identity/src"
	"github.com/rs/zerolog/log"
)

type PolicyDiff struct {
	Over  []string
	Under []string
}

const allow = "Allow"

type identityParseError struct {
	err error
}

func (m *identityParseError) Error() string {
	return fmt.Sprintf("Identity parsing error %v", m.err)
}

type getIAMError struct {
	err error
}

func (m *getIAMError) Error() string {
	return fmt.Sprintf("get IAM error %v", m.err)
}

type compareAllowError struct {
	err error
}

func (m *compareAllowError) Error() string {
	return fmt.Sprintf("compare allow error %v", m.err)
}

func Inspect(directory string, init bool) (PolicyDiff, error) {
	var iacPolicy Identity.Policy

	var Difference PolicyDiff

	rawIACPolicy, err := MakePolicy(directory, nil, init, false, "", "")
	if err != nil {
		if errors.Is(err, &emptyIACError{}) {
			log.Info().Msgf("nothing to do for IAC as %s for directory %s", err, directory)
		} else {
			return Difference, &makePolicyError{err: err}
		}
	}

	iacPolicy, err = Identity.Parse(rawIACPolicy.AWS.JSONOut)
	if err != nil {
		if errors.Is(err, &Identity.EmptyParseError{}) {
			log.Info().Msgf("nothing to do for IAC as parse for %s was empty", directory)
		} else {
			return Difference, &identityParseError{err}
		}
	}

	iamIdentity, err := Identity.GetIam()
	if err != nil {
		log.Info().Msgf("nothing to do for AWS as %s ", err)

		return Difference, &getIAMError{err: err}
	}

	Difference, err = compareAllow(iamIdentity, iacPolicy)
	if err != nil {
		return Difference, &compareAllowError{err: err}
	}

	return Difference, nil
}

type policyDifferenceError struct{}

func (m *policyDifferenceError) Error() string {
	return "invalid input: empty or nil policies/statements"
}

func compareAllow(identity Identity.IAM, policy Identity.Policy) (PolicyDiff, error) {
	// Add at start of function
	if identity.Policies == nil || policy.Statements == nil {
		return PolicyDiff{}, &policyDifferenceError{}
	}

	if len(identity.Policies) == 0 || len(policy.Statements) == 0 {
		return PolicyDiff{}, &policyDifferenceError{}
	}

	identityAllows := make([]string, 0, len(identity.Policies)*2)
	policyAllows := make([]string, 0, len(policy.Statements))

	var difference PolicyDiff

	for _, identityPolicy := range identity.Policies {
		statements := identityPolicy.Statements

		if statements != nil {
			for _, statement := range identityPolicy.Statements {
				if statement.Effect == allow {
					identityAllows = append(identityAllows, statement.Action...)
				}
			}
		}
	}

	for _, policyStatement := range policy.Statements {
		if policyStatement.Effect == allow {
			policyAllows = append(policyAllows, policyStatement.Action...)
		}
	}

	for _, permission := range identityAllows {
		if !contains(policyAllows, permission) {
			difference.Over = append(difference.Over, permission)
		}
	}

	for _, permission := range policyAllows {
		if !contains(identityAllows, permission) {
			difference.Under = append(difference.Under, permission)
		}
	}

	return difference, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}
