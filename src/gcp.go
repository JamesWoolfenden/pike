package pike

import (
	"log"
)

// GetGCPPermissions for GCP resources
func GetGCPPermissions(result ResourceV2) []string {
	var Permissions []string
	if result.TypeName == "resource" {
		Permissions = GetGCPResourcePermissions(result)
	} else {
		Permissions = GetGCPDataPermissions(result)
	}

	return Permissions
}

// GetGCPResourcePermissions looks up permissions required for resources
func GetGCPResourcePermissions(result ResourceV2) []string {
	TFLookup := map[string]interface{}{
		"google_compute_instance": googleComputeInstance,
		"google_storage_bucket":   googleStorageBucket,
	}

	var Permissions []string

	temp := TFLookup[result.Name]
	if temp != nil {
		Permissions = GetPermissionMap(TFLookup[result.Name].([]byte), result.Attributes)
	} else {
		log.Printf("%s not implemented", result.Name)
	}

	return Permissions
}
