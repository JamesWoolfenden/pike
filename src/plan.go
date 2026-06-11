package pike

import (
	"encoding/json"
	"strings"
)

// gcpReadVerbs is the set of GCP permission verbs that are safe for
// terraform plan — they only read state, never mutate it.
// Keep in sync with tools/populate-plan-perms.py.
var gcpReadVerbs = map[string]bool{
	"get":          true,
	"list":         true,
	"getIamPolicy": true,
	"fetch":        true,
	"search":       true,
	"query":        true,
	"check":        true,
	"access":       true,
	"validate":     true,
	"read":         true,
	"view":         true,
	"lookup":       true,
	"describe":     true,
	"export":       true,
}

// awsReadPrefixes is the set of AWS action-name prefixes considered read-only.
var awsReadPrefixes = []string{
	"Describe", "Get", "List", "BatchGet", "BatchDescribe", "Lookup",
	"Search", "Query", "Scan", "View", "Read", "Head", "Select",
	"Check", "Validate", "Detect", "Estimate", "Preview", "Test",
	"Is", "Verify", "Retrieve",
}

func isGCPReadPerm(p string) bool {
	i := strings.LastIndex(p, ".")
	if i < 0 {
		return false
	}
	return gcpReadVerbs[p[i+1:]]
}

func isAzureReadPerm(p string) bool {
	return strings.HasSuffix(p, "/read")
}

func isAWSReadPerm(p string) bool {
	_, action, ok := strings.Cut(p, ":")
	if !ok {
		action = p
	}
	for _, prefix := range awsReadPrefixes {
		if strings.HasPrefix(action, prefix) {
			return true
		}
	}
	return false
}

// getPlanPermissionMap returns the plan-safe permissions from a mapping file:
// the plan array verbatim, plus any attribute-conditional permissions that
// pass the provider's read-only predicate.
func getPlanPermissionMap(raw []byte, attributes []string, resource string, isReadOnly func(string) bool) ([]string, error) {
	if !json.Valid(raw) || len(raw) == 0 {
		return nil, &invalidJSONError{}
	}

	var mappings []any
	if err := json.Unmarshal(raw, &mappings); err != nil {
		return nil, &unmarshallJSONError{err, resource}
	}

	if len(mappings) == 0 {
		return nil, nil
	}

	temp, err := IsTypeOK(mappings[0])
	if err != nil {
		return nil, err
	}

	if temp["attributes"] == nil {
		return nil, &attributesFieldMissingError{}
	}

	resourceAttributes, err := IsTypeOK(temp["attributes"])
	if err != nil {
		return nil, &assertionFailedError{"temp[\"attributes\"]", err}
	}

	var found []string

	for _, attribute := range attributes {
		if resourceAttributes[attribute] != nil {
			for _, entry := range resourceAttributes[attribute].([]any) {
				if value, ok := entry.(string); ok && isReadOnly(value) {
					found = append(found, value)
				}
			}
		}
	}

	found, err = getActionPermissions(temp, found, plan)
	if err != nil {
		return found, err
	}

	return found, nil
}
