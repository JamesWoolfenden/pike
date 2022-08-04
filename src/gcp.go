package pike

import (
	"log"
)

// GetGCPPermissions for GCP resources
func GetGCPPermissions(result ResourceV2) []string {
	var Permissions []string
	switch result.Name {
	case "google_compute_instance":
		Permissions = GetPermissionMap(google_compute_instance, result.Attributes)

	default:
		log.Printf("%s not yet implemented", result.Name)
	}

	return Permissions
}
