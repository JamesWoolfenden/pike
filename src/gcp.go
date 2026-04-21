package pike

import (
	"fmt"
	"strings"
)

type invalidGCPResourceError struct {
	resource string
}

func (m invalidGCPResourceError) Error() string {
	return fmt.Sprintf("Invalid GCP lookup sourceData type for resource %s", m.resource)
}

type invalidPermissionMapError struct {
	err error
}

func (m invalidPermissionMapError) Error() string {
	return fmt.Sprintf("Invalid Permission Map %v", m.err)
}

// getGCPPermissions for GCP resources.
func getGCPPermissions(result ResourceV2) ([]string, error) {
	if result.TypeName == resource || result.TypeName == "terraform" {
		return getGCPResourcePermissions(result)
	}

	return GetGCPDataPermissions(result)
}

// getGCPRuntimePermissions extracts runtime permissions for GCP resources.
func getGCPRuntimePermissions(result ResourceV2) (RuntimePermission, error) {
	if result.TypeName == resource || result.TypeName == "terraform" {
		return getGCPResourceRuntimePermissions(result)
	}

	return RuntimePermission{}, nil // No runtime permissions for data sources
}

func getGCPResourcePermissions(sourceData ResourceV2) ([]string, error) {
	raw := GCPLookup(sourceData.Name)
	if raw == nil {
		return nil, &notImplementedResourceError{sourceData.Name}
	}

	permissions, err := GetPermissionMap(raw, sourceData.Attributes, sourceData.Name)
	if err != nil {
		return nil, &invalidPermissionMapError{err}
	}

	return permissions, nil
}

func getGCPResourceRuntimePermissions(sourceData ResourceV2) (RuntimePermission, error) {
	raw := GCPLookup(sourceData.Name)
	if raw == nil {
		return RuntimePermission{}, nil
	}

	permissions, err := GetRuntimePermissions(raw, sourceData.Attributes, sourceData.Name)
	if err != nil {
		// Runtime permissions are optional, so don't fail if not found
		return RuntimePermission{}, nil
	}

	if len(permissions) == 0 {
		return RuntimePermission{}, nil
	}

	// Extract service account from attributes
	serviceAccount := extractServiceAccount(sourceData.Attributes)

	return RuntimePermission{
		ResourceType:   sourceData.Name,
		ResourceName:   sourceData.ResourceName,
		ServiceAccount: serviceAccount,
		Permissions:    permissions,
	}, nil
}

// extractServiceAccount finds the service account in the resource attributes
func extractServiceAccount(attributes []string) string {
	// Look for service_account attribute
	for _, attr := range attributes {
		if attr == "service_account" {
			return "custom" // Indicates a custom service account is defined
		}
	}
	return "default" // No service account specified, using default
}

func GCPLookup(name string) []byte {
	if data, ok := gcpResourceMap[name]; ok {
		return data
	}
	// IAM resources come in _binding/_member/_policy variants that share
	// a single JSON file named without the suffix (e.g. google_foo_iam).
	for _, suffix := range []string{"_binding", "_member", "_policy"} {
		if strings.HasSuffix(name, suffix) {
			base := strings.TrimSuffix(name, suffix)
			if data, ok := gcpResourceMap[base]; ok {
				return data
			}
		}
	}
	return nil
}
