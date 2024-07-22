package pike

import (
	Identity "github.com/jameswoolfenden/identity/src"
)

func Inspect(directory string, init bool) ([]string, error) {

	rawIACPolicy, err := MakePolicy(directory, nil, init, false)
	if err != nil {
		return nil, err
	}

	iacPolicy, err := Identity.Parse(rawIACPolicy.AWS.JSONOut)

	if err != nil {
		return nil, err
	}

	iamIdentity, err := Identity.GetIam()
	if err != nil {
		return nil, err
	}

	result, err := CompareAllow(iamIdentity, iacPolicy)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func CompareAllow(identity Identity.IAM, policy Identity.Policy) ([]string, error) {
	var identityAllows []string
	var policyAllows []string
	var over []string

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
			over = append(over, permission)
		}
	}

	return over, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
