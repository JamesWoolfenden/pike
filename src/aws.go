package pike

import (
	"encoding/json"
	"fmt"
	"slices"

	"github.com/rs/zerolog/log"
)

const (
	terraform string = "terraform"
	module    string = "module"
	resource  string = "resource"
	data      string = "data"
)

// GetAWSPermissions for AWS resources.
func GetAWSPermissions(result ResourceV2) ([]string, error) {
	// Validate the input
	if result.TypeName == "" {
		return nil, &emptyTypeNameError{}
	}

	if result.Name == "" {
		return nil, &emptyNameError{}
	}

	var (
		err         error
		Permissions []string
	)

	switch result.TypeName {
	case resource, terraform:
		Permissions, err = GetAWSResourcePermissions(result)
		if err != nil {
			return Permissions, &getAWSResourcePermissionsError{err}
		}
	case data:
		Permissions, err = GetAWSDataPermissions(result)
		if err != nil {
			return Permissions, &getAWSDataPermissionsError{err}
		}
	case module:
		// module block — not a base resource type, nothing to do
	default:
		return nil, &unknownPermissionError{result.Name}
	}

	return Permissions, nil
}

// GetAWSResourcePermissions looks up permissions required for resources.
func GetAWSResourcePermissions(result ResourceV2) ([]string, error) {
	temp := AwsLookup(result.Name)
	if temp == nil {
		return nil, &notImplementedResourceError{result.Name}
	}
	return GetPermissionMap(temp, result.Attributes, result.AttributeValues, result.Name)
}

func getAWSPlanPermissions(result ResourceV2) ([]string, error) {
	switch result.TypeName {
	case resource, terraform:
		temp := AwsLookup(result.Name)
		if temp == nil {
			return nil, &notImplementedResourceError{result.Name}
		}
		return getPlanPermissionMap(temp, result.Attributes, result.Name, isAWSReadPerm)
	case data:
		temp := AwsDataLookup(result.Name)
		if temp == nil {
			return nil, &notImplementedDatasourceError{result.Name}
		}
		return getPlanPermissionMap(temp, result.Attributes, result.Name, isAWSReadPerm)
	default:
		return nil, nil
	}
}

func AwsLookup(name string) []byte {
	if name == "" {
		return nil
	}

	return awsResourceMap[name]
}

// Contains looks if slice contains string.
func Contains(s []string, e string) bool {
	return slices.Contains(s, e)
}

// GetPermissionMap Anonymous parsing.
func GetPermissionMap(raw []byte, attributes []string, attributeValues map[string][]string, resource string) ([]string, error) {
	if !json.Valid(raw) || len(raw) == 0 {
		return nil, &invalidJSONError{}
	}

	var mappings []any
	err := json.Unmarshal(raw, &mappings)
	if err != nil {
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
				value, ok := entry.(string)

				if !ok {
					log.Error().Msg("failed to cast to string")

					continue
				}

				found = append(
					found,
					value,
				)
			}
		}
	}

	found, err = getActionPermissions(temp, found)
	if err != nil {
		return found, fmt.Errorf("getActionPermissions: %w", err)
	}

	found, err = getAttributeValuePermissions(temp, attributeValues, found)
	if err != nil {
		return found, fmt.Errorf("getAttributeValuePermissions: %w", err)
	}

	return found, nil
}

// getAttributeValuePermissions adds permissions keyed by the literal value
// of an attribute, not merely its presence — e.g. azurerm_private_endpoint's
// subresource_names needs a different permission for "redisCache" than for
// "vault". The mapping JSON's optional "attributeValues" field has the
// shape {"<attribute name>": {"<value>": ["<permission>", ...], ...}, ...}.
// A mapping with no "attributeValues" field, or a value with no matching
// entry, adds nothing here — this never guesses a permission for a value
// the mapping doesn't explicitly know about.
func getAttributeValuePermissions(permissionMap map[string]any, attributeValues map[string][]string, found []string) ([]string, error) {
	raw, ok := permissionMap["attributeValues"]
	if !ok || raw == nil {
		return found, nil
	}

	byAttribute, err := IsTypeOK(raw)
	if err != nil {
		return found, &assertionFailedError{"temp[\"attributeValues\"]", err}
	}

	for attribute, values := range attributeValues {
		rawByValue, ok := byAttribute[attribute]
		if !ok {
			continue
		}

		byValue, err := IsTypeOK(rawByValue)
		if err != nil {
			return found, &assertionFailedError{"attributeValues[\"" + attribute + "\"]", err}
		}

		for _, value := range values {
			entries, ok := byValue[value].([]any)
			if !ok {
				continue
			}

			for _, entry := range entries {
				permission, ok := entry.(string)
				if !ok {
					log.Error().Msg("failed to cast to string")

					continue
				}

				if !Contains(found, permission) {
					found = append(found, permission)
				}
			}
		}
	}

	return found, nil
}

const (
	apply   = "apply"
	plan    = "plan"
	modify  = "modify"
	destroy = "destroy"
)

type parameterNilError struct {
	parameter string
}

func (m *parameterNilError) Error() string {
	return fmt.Sprintf("%s was nil", m.parameter)
}

func getActionPermissions(permissionMap map[string]any, found []string, actions ...string) ([]string, error) {
	if permissionMap == nil {
		return nil, &parameterNilError{parameter: "permissionMap"}
	}

	if len(actions) == 0 {
		actions = []string{apply, plan, modify, destroy}
	}

	for _, action := range actions {
		if permissionMap[action] != nil {

			temp, ok := permissionMap[action].([]any)

			if !ok {
				log.Error().Msg("failed to cast permission map to list")
			}

			for _, entry := range temp {
				value, ok := entry.(string)
				if !ok {
					log.Error().Msg("failed to cast string")

					continue
				}

				found = append(found, value)
			}
		}
	}

	return found, nil
}

// GetRuntimePermissions extracts runtime permissions needed by service accounts from mapping files.
func GetRuntimePermissions(raw []byte, attributes []string, resource string) ([]string, error) {
	if !json.Valid(raw) || len(raw) == 0 {
		return nil, &invalidJSONError{}
	}

	var mappings []any
	err := json.Unmarshal(raw, &mappings)
	if err != nil {
		return nil, &unmarshallJSONError{err, resource}
	}

	if len(mappings) == 0 {
		return nil, nil
	}

	temp, err := IsTypeOK(mappings[0])
	if err != nil {
		return nil, err
	}

	// Check if runtime field exists
	if temp["runtime"] == nil {
		return []string{}, nil // No runtime permissions defined
	}

	runtimeMap, err := IsTypeOK(temp["runtime"])
	if err != nil {
		return nil, &assertionFailedError{"temp[\"runtime\"]", err}
	}

	var found []string

	// Check each attribute to see if it triggers runtime permissions
	for _, attribute := range attributes {
		if runtimeMap[attribute] != nil {
			for _, entry := range runtimeMap[attribute].([]any) {
				value, ok := entry.(string)

				if !ok {
					log.Error().Msg("failed to cast runtime permission to string")
					continue
				}

				found = append(found, value)
			}
		}
	}

	return found, nil
}

func IsTypeOK(mappings any) (map[string]any, error) {
	temp, ok := mappings.(map[string]any)

	if !ok {
		return nil, &assertionError{"mappings to map[string]Interface{}"}
	}

	return temp, nil
}
