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
		"google_compute_instance":        googleComputeInstance,
		"google_storage_bucket":          googleStorageBucket,
		"google_storage_bucket_object":   googleStorageBucketObject,
		"google_compute_network":         googleComputeNetwork,
		"google_compute_address":         googleComputeAddress,
		"google_compute_firewall":        googleComputeFirewall,
		"google_project_iam_custom_role": googleProjectIamCustomRole,
		"google_service_account":         googleServiceAccount,
		"google_kms_crypto_key":          googleKmsCryptoKey,
		"google_kms_key_ring":            googleKmsKeyRing,
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
