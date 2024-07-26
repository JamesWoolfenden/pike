package pike

import (
	"errors"

	Identity "github.com/jameswoolfenden/identity/src"
	"github.com/rs/zerolog/log"
)

type PolicyDiff struct {
	Over  []string
	Under []string
}

func Inspect(directory string, init bool) (PolicyDiff, error) {
	var iacPolicy Identity.Policy

	var Difference PolicyDiff

	rawIACPolicy, err := MakePolicy(directory, nil, init, false)
	if err != nil {
		if errors.Is(err, &emptyIACError{}) {
			log.Info().Msgf("nothing to do for IAC as %s for directory %s", err, directory)
		} else {
			return Difference, err
		}
	}

	iacPolicy, err = Identity.Parse(rawIACPolicy.AWS.JSONOut)

	if err != nil {
		if errors.Is(err, &Identity.EmptyParseError{}) {
			log.Info().Msgf("nothing to do for IAC as parse for %s was empty", directory)
		} else {
			return Difference, err
		}
	}

	iamIdentity, err := Identity.GetIam()
	if err != nil {
		log.Info().Msgf("nothing to do for AWS as %s ", err)

		return Difference, err
	}

	Difference, err = CompareAllow(iamIdentity, iacPolicy)
	if err != nil {
		return Difference, err
	}

	return Difference, nil
}

func CompareAllow(identity Identity.IAM, policy Identity.Policy) (PolicyDiff, error) {
	var identityAllows []string

	var policyAllows []string

	var difference PolicyDiff

	for _, identityPolicy := range identity.Policies {
		statements := identityPolicy.Statements

		if statements != nil {
			for _, statement := range identityPolicy.Statements {
				if statement.Effect == "Allow" {
					identityAllows = append(identityAllows, statement.Action...)
				}
			}
		}
	}

	for _, policyStatement := range policy.Statements {
		if policyStatement.Effect == "Allow" {
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
