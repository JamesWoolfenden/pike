package pike

import "encoding/json"

// getPlanPermissionMap returns the plan-array permissions from a mapping file.
// Used for all three providers now that plan arrays are populated at source.
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
