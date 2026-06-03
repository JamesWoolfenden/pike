package pike

import (
	"encoding/json"
	"strings"
)

// gcpReadVerbs is the set of GCP permission verbs that are safe for
// terraform plan — they only read state, never mutate it.
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

// getPlanPermissionMap returns only the plan-array permissions for an AWS
// resource mapping. GCP and Azure derive plan perms by filtering apply.
func getPlanPermissionMap(raw []byte, attributes []string, resource string) ([]string, error) {
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
				if value, ok := entry.(string); ok {
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

// filterPlanPermissionsGCP returns the deduplicated subset of GCP permissions
// safe for terraform plan — those whose verb is in gcpReadVerbs.
func filterPlanPermissionsGCP(perms []string) []string {
	seen := make(map[string]bool)
	var out []string
	for _, p := range perms {
		parts := strings.Split(p, ".")
		if len(parts) == 0 {
			continue
		}
		verb := parts[len(parts)-1]
		if gcpReadVerbs[verb] && !seen[p] {
			seen[p] = true
			out = append(out, p)
		}
	}
	return out
}

// filterPlanPermissionsAzure returns the deduplicated subset of Azure
// permissions safe for terraform plan — those ending in "/read".
func filterPlanPermissionsAzure(perms []string) []string {
	seen := make(map[string]bool)
	var out []string
	for _, p := range perms {
		if strings.HasSuffix(p, "/read") && !seen[p] {
			seen[p] = true
			out = append(out, p)
		}
	}
	return out
}
