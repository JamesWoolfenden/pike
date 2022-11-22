package pike

import (
	"fmt"
)

// GetGCPDataPermissions gets permissions required for datasources
func GetGCPDataPermissions(result ResourceV2) ([]string, error) {

	TFLookup := map[string]interface{}{
		"google_service_account":    dataGoogleServiceAccount,
		"google_compute_image":      placeholder,
		"google_compute_network":    dataGoogleComputeNetwork,
		"google_compute_subnetwork": dataGoogleComputeSubnetwork,
		"google_compute_zones":      dataGoogleComputeZones,
		"google_project":            dataGoogleProject,
		"google_iam_policy":         placeholder,
		"google_iam_role":           placeholder,
		"google_kms_crypto_key":     dataGoogleKmsCryptoKey,
		"google_kms_key_ring":       dataGoogleKmsKeyRing,
	}

	var Permissions []string
	var err error
	temp := TFLookup[result.Name]
	if temp != nil {
		Permissions, err = GetPermissionMap(TFLookup[result.Name].([]byte), result.Attributes)
	} else {
		return nil, fmt.Errorf("data.%s not implemented", result.Name)
	}

	return Permissions, err
}
