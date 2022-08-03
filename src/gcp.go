package pike

import (
	"log"
)

// GetGCPPermissions for GCP resources
func GetGCPPermissions(result template) []string {
	myAttributes := GetAttributes(result)
	var Permissions []string
	switch result.Resource.name {
	case "google_compute_instance":
		Permissions = GetPermissionMap(google_compute_instance, myAttributes)

	default:
		log.Printf("%s %s not yet implemented", result.Template, result.Resource.name)
	}

	return Permissions
}
