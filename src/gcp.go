package pike

import (
	"log"
)

// GetGCPPermissions for GCP resources
func GetGCPPermissions(result ResourceV2) []string {
	var Permissions []string
	switch result.Name {
	case "googleComputeInstance":
		Permissions = GetPermissionMap(googleComputeInstance, result.Attributes)

	default:
		log.Printf("%s not yet implemented", result.Name)
	}

	return Permissions
}
