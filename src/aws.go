package pike

import (
	"encoding/json"
	"fmt"

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
		{
			Permissions, err = GetAWSResourcePermissions(result)
			if err != nil {
				return Permissions, &getAWSResourcePermissionsError{err}
			}
		}
	case data:
		{
			Permissions, err = GetAWSDataPermissions(result)
			if err != nil {
				return Permissions, &getAWSDataPermissionsError{err}
			}
		}
	case module:
		{
			// do nothing this is a module not a base resource type, and
			// we shouldn't really be able to get here unless well bad naming
		}
	default:
		{
			return nil, &unknownPermissionError{result.Name}
		}
	}

	return Permissions, nil
}

// GetAWSResourcePermissions looks up permissions required for resources
//
//goland:noinspection GoLinter
func GetAWSResourcePermissions(result ResourceV2) ([]string, error) {
	var (
		Permissions []string
		err         error
	)

	if temp := AwsLookup(result.Name); temp != nil {
		Permissions, err = GetPermissionMap(temp, result.Attributes, result.Name)
	} else {
		return nil, &notImplementedResourceError{result.Name}
	}

	return Permissions, err
}

func AwsLookup(name string) []byte {
	if name == "" {
		return nil
	}

	return awsResourceMap[name]
}

// Contains looks if slice contains string.
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

// GetPermissionMap Anonymous parsing.
func GetPermissionMap(raw []byte, attributes []string, resource string) ([]string, error) {
	if !json.Valid(raw) || len(raw) == 0 {
		return nil, &invalidJSONError{}
	}

	var mappings []interface{}
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
			for _, entry := range resourceAttributes[attribute].([]interface{}) {
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

func getActionPermissions(permissionMap map[string]interface{}, found []string) ([]string, error) {
	if permissionMap == nil {
		return nil, &parameterNilError{parameter: "permissionMap"}
	}

	for _, action := range []string{apply, plan, modify, destroy} {
		if permissionMap[action] != nil {

			temp, ok := permissionMap[action].([]interface{})

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

	var mappings []interface{}
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
			for _, entry := range runtimeMap[attribute].([]interface{}) {
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

func IsTypeOK(mappings interface{}) (map[string]interface{}, error) {
	temp, ok := mappings.(map[string]interface{})

	if !ok {
		return nil, &assertionError{"mappings to map[string]Interface{}"}
	}

	return temp, nil
}
